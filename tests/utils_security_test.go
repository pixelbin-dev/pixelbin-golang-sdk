package tests

import (
	"errors"
	"net/url"
	"testing"

	"github.com/pixelbin-io/pixelbin-go/v3/sdk/utils/security"
)

type signedUrlInput struct {
	url           string
	expirySeconds int
	accessKey     string
	token         string
}

var SIGN_URL_CASES = []struct {
	scenario string
	input    signedUrlInput
	err      error
}{
	{
		scenario: "Should sign URL",
		input: signedUrlInput{
			url:           "https://cdn.pixelbin.io/v2/dummy-cloudname/original/__playground/playground-default.jpeg",
			expirySeconds: 20,
			accessKey:     "592db51c-3d1d-4ae0-a081-a5a085ee3f1a",
			token:         "dummy-token",
		},
		err: nil,
	},
	{
		scenario: "Should sign Custom domain URL",
		input: signedUrlInput{
			url:           "https://krit.imagebin.io/v2/dummy-cloudname/original/__playground/playground-default.jpeg",
			expirySeconds: 42,
			accessKey:     "f7c5044d-8da8-41cd-aa8c-6955f32a3987",
			token:         "dummy-token",
		},
		err: nil,
	},
	{
		scenario: "Should error if URL already has a signature",
		input: signedUrlInput{
			url:           "https://cdn.pixelbin.io/v2/dummy-cloudname/original/__playground/playground-default.jpeg?pbs=2e7578ba14ef3294a3cc95209fad9801a6abdc917ab8f98e5d1ffb4645a6289e&pbe=1696403372&pbt=2583",
			expirySeconds: 23,
			accessKey:     "1684eed3-9a2b-4507-b6c1-706937f73d25",
			token:         "dummy-token",
		},
		err: errors.New("URL already has a signature"),
	},
}

func Test_ShouldSignURL(t *testing.T) {
	for i, testcase := range SIGN_URL_CASES {
		t.Run(testcase.scenario, func(t *testing.T) {
			signedUrl, err := security.SignURL(testcase.input.url, testcase.input.expirySeconds, testcase.input.accessKey, testcase.input.token)
			if err != nil {
				if err.Error() != testcase.err.Error() {
					t.Errorf("Test_ShouldSignURL %d failed, expected %v, got %v", i, testcase.err, err.Error())
				}
			} else {
				urlParts, err := url.Parse(signedUrl)
				if err != nil {
					t.Errorf("Test_ShouldSignURL %d failed, expected %v, got %v", i, testcase.err, err.Error())
				}
				if !urlParts.Query().Has("pbs") || !urlParts.Query().Has("pbe") || !urlParts.Query().Has("pbt") {
					t.Errorf("Test_ShouldSignURL %d failed, expected %v, got %v", i, testcase.err, err.Error())
				}
			}
		})
	}
}
