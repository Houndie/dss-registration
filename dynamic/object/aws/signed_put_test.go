package aws

import (
	"net/url"
	"strings"
	"testing"
)

func TestSignedPut(t *testing.T) {
	region := "us-east-2"
	bucket := "vaxbucket"

	c, err := NewObjectClient("access_key", "secret_key", region, bucket)
	if err != nil {
		t.Fatal(err)
	}

	registrationID := "registration_id"

	u, err := c.SignedPut(23, registrationID)
	if err != nil {
		t.Fatal(err)
	}

	p, err := url.Parse(u)
	if err != nil {
		t.Fatal(err)
	}

	if strings.Split(p.Host, ".")[0] != bucket {
		t.Fatal(p.Host)
	}

	if p.Path != "/"+registrationID {
		t.Fatal(p.Path)
	}
}
