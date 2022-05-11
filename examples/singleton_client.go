package examples

import (
	"github.com/rara7777/go-http-client/gohttp"
	"github.com/rara7777/go-http-client/gomime"
	"net/http"
	"time"
)

var (
	httpClient = getHttpClient()
)

func getHttpClient() gohttp.Client {
	headers := make(http.Header)
	headers.Set(gomime.HeaderContentType, gomime.ContentTypeJson)

	client := gohttp.NewBuilder().
		SetHeaders(headers).
		SetConnectionTimeout(2 * time.Second).
		SetResponseTimeout(3 * time.Second).
		SetUserAgent("Fedes-Computer").
		Build()
	return client
}
