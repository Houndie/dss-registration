package auth0

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/lestrrat-go/jwx/jwk"
	"github.com/sirupsen/logrus"
)

type Authorizer struct {
	jwks     *cachedJWKS
	logger   *logrus.Logger
	audience string
}

type cachedDiscoveryDocument struct {
	dd       *discoveryDocument
	expires  time.Time
	mut      *sync.RWMutex
	client   *http.Client
	logger   *logrus.Logger
	endpoint *url.URL
}

func (c *cachedDiscoveryDocument) get(ctx context.Context) (*discoveryDocument, error) {
	c.mut.RLock()
	if c.dd != nil && c.expires.After(time.Now()) {
		c.logger.Trace("found cached discovery document")
		defer c.mut.RUnlock()
		return c.dd, nil
	}
	c.mut.RUnlock()

	c.mut.Lock()
	defer c.mut.Unlock()
	if c.dd != nil && c.expires.After(time.Now()) {
		c.logger.Trace("discovery document added to cache, returning")
		return c.dd, nil
	}

	req, err := http.NewRequest(http.MethodGet, c.endpoint.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("error generating discovery document request: %w", err)
	}
	req = req.WithContext(ctx)

	c.logger.Trace("making network request for discovery document")
	res, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error fetching discovery document: %w", err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("call to fetch discovery document did not return %d, instead found %d", http.StatusOK, res.StatusCode)
	}
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading bytes from discovery document response: %w", err)
	}

	dd := &discoveryDocument{}
	err = json.Unmarshal(bodyBytes, dd)
	if err != nil {
		return nil, fmt.Errorf("error reading json from discovery document response: %w", err)
	}

	c.dd = dd
	c.expires = time.Now().Add(maxAge(res))
	return dd, nil
}

type cachedJWKS struct {
	jwks    jwk.Set
	dd      *cachedDiscoveryDocument
	expires time.Time
	mut     *sync.RWMutex
	client  *http.Client
	logger  *logrus.Logger
}

func (c *cachedJWKS) get(ctx context.Context) (jwk.Set, error) {
	c.mut.RLock()
	if c.jwks != nil && c.expires.After(time.Now()) {
		c.logger.Trace("found cached json web key set")
		defer c.mut.RUnlock()
		return c.jwks, nil
	}
	c.mut.RUnlock()

	c.mut.Lock()
	defer c.mut.Unlock()
	if c.jwks != nil && c.expires.After(time.Now()) {
		c.logger.Trace("json web key set added to cache, returning")
		return c.jwks, nil
	}

	dd, err := c.dd.get(ctx)
	if err != nil {
		return nil, fmt.Errorf("error fetching discovery document from cache: %w", err)
	}

	req, err := http.NewRequest(http.MethodGet, dd.JWKSURI, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating jwks request: %w", err)
	}
	req = req.WithContext(ctx)

	c.logger.Trace("making network request for json web key set")
	res, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error fetching jwt key source: %w", err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("call to fetch jwt key source did not return %d, instead found %d", http.StatusOK, res.StatusCode)
	}
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading bytes from jwt key source response: %w", err)
	}

	set, err := jwk.Parse(bodyBytes)
	if err != nil {
		return nil, fmt.Errorf("error parsing json web key set: %w", err)
	}

	c.jwks = set
	c.expires = time.Now().Add(maxAge(res))
	return set, nil
}

const (
	discoveryDocumentPath = ".well-known/openid-configuration"
)

type discoveryDocument struct {
	JWKSURI string `json:"jwks_uri"`
}

func maxAge(res *http.Response) time.Duration {
	cacheControl := res.Header.Get("cache-control")
	if cacheControl == "" {
		return 0
	}

	cacheControlItems := strings.Split(cacheControl, ",")
	for _, item := range cacheControlItems {
		if item == "no-cache" || item == "no-store" {
			return 0
		}
	}

	for _, item := range cacheControlItems {
		if !strings.HasPrefix(item, "max-age") {
			continue
		}
		splitItem := strings.Split(item, "=")
		if len(splitItem) != 2 {
			// Weird, malformed header?
			return 0
		}

		maxAge, err := strconv.Atoi(splitItem[1])
		if err != nil {
			return 0
		}

		return time.Duration(maxAge) * time.Second
	}

	return 0
}

func NewAuthorizer(endpoint, audience string, client *http.Client, logger *logrus.Logger) (*Authorizer, error) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, fmt.Errorf("error parsing endpoint: %w", err)
	}
	u.Path = path.Join(u.Path, discoveryDocumentPath)
	dd := &cachedDiscoveryDocument{
		expires:  time.Now(),
		mut:      &sync.RWMutex{},
		client:   client,
		logger:   logger,
		endpoint: u,
	}
	jwks := &cachedJWKS{
		expires: time.Now(),
		mut:     &sync.RWMutex{},
		client:  client,
		logger:  logger,
		dd:      dd,
	}
	return &Authorizer{
		jwks:     jwks,
		logger:   logger,
		audience: audience,
	}, nil
}
