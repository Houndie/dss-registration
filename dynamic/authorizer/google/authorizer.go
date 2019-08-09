package google

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/bluele/gcache"
	"github.com/pkg/errors"
)

type discoveryDocument struct {
	UserinfoEndpoint string `json:"userinfo_endpoint"`
}

const (
	discoveryDocumentKey = "discoveryDocumentKey"
	discoveryDocumentURI = "https://accounts.google.com/.well-known/openid-configuration"
)

type Authorizer struct {
	cache  gcache.Cache
	client *http.Client
}

func NewAuthorizer(client *http.Client) *Authorizer {
	return &Authorizer{
		client: client,
		cache: gcache.New(10).LFU().LoaderExpireFunc(func(interface{}) (interface{}, *time.Duration, error) {
			res, err := client.Get(discoveryDocumentURI)
			if err != nil {
				return nil, nil, errors.Wrap(err, "error fetching discovery document")
			}
			defer res.Body.Close()
			if res.StatusCode != http.StatusOK {
				return nil, nil, fmt.Errorf("call to fetch discovery document did not return %d, instead found %d", http.StatusOK, res.StatusCode)
			}
			bodyBytes, err := ioutil.ReadAll(res.Body)
			if err != nil {
				return nil, nil, errors.Wrap(err, "error reading bytes from discovery document response")
			}

			dd := &discoveryDocument{}
			err = json.Unmarshal(bodyBytes, dd)
			if err != nil {
				return nil, nil, errors.Wrap(err, "error reading json from discovery document response")
			}

			cacheControl := res.Header.Get("cache-control")
			nilDuration := time.Duration(0)
			if cacheControl == "" {
				return dd, &nilDuration, nil
			}

			cacheControlItems := strings.Split(cacheControl, ",")
			for _, item := range cacheControlItems {
				if item == "no-cache" || item == "no-store" {
					return dd, &nilDuration, nil
				}
			}

			for _, item := range cacheControlItems {
				if strings.HasPrefix(item, "max-age") {
					splitItem := strings.Split(item, "=")
					if len(splitItem) != 2 {
						// Weird, malformed header?
						return dd, &nilDuration, nil
					}

					maxAge, err := strconv.Atoi(splitItem[1])
					if err != nil {
						return dd, &nilDuration, nil
					}

					maxAgeDuration := time.Duration(maxAge) * time.Second

					return dd, &maxAgeDuration, nil
				}
			}

			return dd, &nilDuration, nil
		}).Build(),
	}
}
