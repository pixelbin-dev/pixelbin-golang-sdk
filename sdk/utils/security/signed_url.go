package security

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func hmacSHA256(key string, message string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}

func generateSignature(urlPath string, expiryTimestamp int64, key string) string {
	urlPath = strings.TrimPrefix(urlPath, "/")
	signature := hmacSHA256(key, fmt.Sprintf("%s%d", urlPath, expiryTimestamp))
	return signature
}

func SignURL(urlString string, expirySeconds int, accessKey string, token string) (string, error) {
	expiryTimestamp := time.Now().Unix() + int64(expirySeconds)

	urlParts, err := url.Parse(urlString)
	if err != nil {
		return "", err
	}

	if urlParts.Query().Get("pbs") != "" {
		return "", errors.New("URL already has a signature")
	}

	signature := generateSignature(urlParts.Path, expiryTimestamp, token)

	urlQuery := urlParts.Query()
	urlQuery.Set("pbs", signature)
	urlQuery.Set("pbe", strconv.FormatInt(expiryTimestamp, 10))
	urlQuery.Set("pbt", accessKey)
	urlParts.RawQuery = urlQuery.Encode()

	return urlParts.String(), nil
}
