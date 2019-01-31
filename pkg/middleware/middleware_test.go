package midl_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	mid "github.com/esanim/top-coins/pkg/middleware"
	"github.com/stretchr/testify/assert"
)

func TestJSONHeader(t *testing.T) {
	a := assert.New(t)

	ts := httptest.NewServer(http.HandlerFunc(mid.JSONHeaderMiddleware(func(w http.ResponseWriter, r *http.Request) {})))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	a.NotNil(t, err)

	a.Equal(200, res.StatusCode)
	a.Contains(res.Header["Content-Type"], "application/json")
}
