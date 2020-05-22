package google

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Houndie/dss-registration/dynamic/authorizer"
	"github.com/pkg/errors"
)

type userinfo struct {
	Sub string `json:"sub"`
}

func (a *Authorizer) Userinfo(ctx context.Context, accessToken string) (*authorizer.Userinfo, error) {
	dd, err := a.cache.Get(discoveryDocumentKey)
	if err != nil {
		return nil, errors.Wrap(err, "error fetching discovery document")
	}
	ddType, ok := dd.(*discoveryDocument)
	if !ok {
		return nil, errors.New("found incorrect discovery document type")
	}
	req, err := http.NewRequest("GET", ddType.UserinfoEndpoint, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error creating request for userinfo endpoint")
	}

	req.Header.Add("Authorization", "Bearer "+accessToken)
	req = req.WithContext(ctx)

	res, err := a.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "error fetching userinfo from open id endpoint")
	}
	defer res.Body.Close()
	switch res.StatusCode {
	case http.StatusNotFound:
		return nil, authorizer.Unauthenticated
	default:
		return nil, fmt.Errorf("call to fetch user info did not return %d, instead found %d", http.StatusOK, res.StatusCode)
	}

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "error reading user info response body")
	}

	guserinfo := &userinfo{}
	err = json.Unmarshal(bodyBytes, guserinfo)
	if err != nil {
		return nil, errors.Wrap(err, "error reading user info response json")
	}

	tuserinfo := &authorizer.Userinfo{
		UserId: guserinfo.Sub,
	}

	return tuserinfo, nil
}
