package http

import (
	"context"
	"io"
	"net/http"
	"net/url"
)

// Put issues a PUT to the specified URL.
//
// Caller should close resp.Body when done reading from it.
//
// If the provided body is an io.Closer, it is closed after the
// request.
//
// To set custom headers, use NewRequest and DefaultClient.Do.
//
// See the Client.Do method documentation for details on how redirects
// are handled.
func Put(ctx context.Context, url string, contentType string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, "PUT", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return http.DefaultClient.Do(req)
}

func RemoveQueryParam(qurl string) string {
	u, err := url.Parse(qurl)
	if err != nil {
		return qurl
	}
	u.RawQuery = ""
	return u.String()
}
