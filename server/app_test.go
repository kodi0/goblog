package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	server     *httptest.Server
	reader     io.Reader
	articleUrl string
)

func init() {
	server = httptest.NewServer(???)

	articleUrl = fmt.Sprintf("%s/article", server.URL)
}

func TestArticleAdd(t *testing.T) {
	articleJson := `{"title": "Test post", "body": "Body of the post"}`

	reader = strings.NewReader(articleJson)

	request, err := http.NewRequest("POST", articleUrl, reader)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 201 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
}
