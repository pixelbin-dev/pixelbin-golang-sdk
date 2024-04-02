package platform

import (
	"fmt"
	"strings"

	"github.com/pixelbin-io/pixelbin-go/v3/sdk/common"
)

// APIClient holds raw data for api execution
type APIClient struct {
	Conf        *PixelbinConfig
	Method      string
	Url         string
	Query       map[string]string
	Body        interface{}
	ContentType string
}

// Execute performs API call
func (c *APIClient) Execute() ([]byte, error) {
	var token string = common.EncodeToBase64(c.Conf.GetAccessToken())
	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", token),
	}
	if c.ContentType != "" && c.ContentType != "multipart/form-data" && c.Body != nil {
		headers["Content-Type"] = c.ContentType
	}
	data := c.Body
	if c.ContentType == "multipart/form-data" {
		data = nil
	}
	queryString := common.MapToUrlString(c.Query)

	model := common.NewSignatureModel(c.Conf.Domain, c.Method, c.Url, queryString, headers, data, []string{"Authorization", "Content-Type"})

	headersWithSign, err := model.AddSignatureToHeaders(false)
	headersWithSign["x-ebg-param"] = common.EncodeToBase64(headersWithSign["x-ebg-param"])
	if err != nil {
		return nil, err
	}
	return common.HttpRequest(strings.ToUpper(c.Method), fmt.Sprintf("%s%s", c.Conf.Domain, c.Url), c.Query, c.Body, headersWithSign)
}
