package coins_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/esanim/top-coins/pkg/app"
	"github.com/esanim/top-coins/pkg/handlers/coins"
	"github.com/stretchr/testify/assert"
)

func TestCoinsListHandler(t *testing.T) {
	a := assert.New(t)

	app, err := app.NewApp("8080")
	a.Nil(err)
	subject := coins.NewCoinsHandler(app)

	req, err := http.NewRequest("GET", "localhost:8080/", nil)
	a.Nil(err)
	rec := httptest.NewRecorder()

	subject.GetCoins(rec, req)

	res := rec.Result()
	defer res.Body.Close()
	a.Equal(http.StatusOK, res.StatusCode)

	result, err := ioutil.ReadAll(res.Body)
	a.Nil(err)

	a.Contains(string(result), "\"symbol\":\"BTC\"")
}
