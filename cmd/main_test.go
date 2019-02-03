package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/esanim/top-coins/pkg/app"
	"github.com/esanim/top-coins/pkg/coins"
	midl "github.com/esanim/top-coins/pkg/middleware"
	"github.com/stretchr/testify/assert"
)

func TestRoutes(t *testing.T) {
	a := assert.New(t)

	app, err := app.NewApp(":8080")
	a.Nil(err)

	// GET /
	coinsHandler := coins.NewCoinsHandler(app)
	app.Handler.HandleFunc("/", midl.WrapMiddleware(coinsHandler.GetCoins))

	ts := httptest.NewServer(app.Handler)
	defer ts.Close()

	tt := []struct {
		name   string
		uri    string
		status int
	}{
		{"coins", "/", 200},
		{"not_found", "/xyz", 404},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			app.Logger.Info().Msg(ts.URL + tc.uri)
			res, err := http.Get(ts.URL + tc.uri)
			a.Nil(err)

			app.Logger.Info().Msg(fmt.Sprintf("%s: %d == %d", tc.name, tc.status, res.StatusCode))
			a.Equal(tc.status, res.StatusCode)
		})
	}
}
