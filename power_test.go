package solaredge_test

import (
	"context"
	"github.com/clambin/solaredge"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestClient_GetPower(t *testing.T) {
	server := &Server{token: "TESTTOKEN"}
	apiServer := httptest.NewServer(http.HandlerFunc(server.apiHandler))
	defer apiServer.Close()

	client := solaredge.Client{
		Token:      "TESTTOKEN",
		HTTPClient: &http.Client{},
		APIURL:     apiServer.URL,
	}

	siteIDs, err := client.GetSiteIDs(context.Background())
	assert.NoError(t, err)
	if assert.Len(t, siteIDs, 1) {
		assert.Equal(t, 1, siteIDs[0])
	}

	entries, err := client.GetPower(context.Background(), siteIDs[0], time.Now().Add(-1*time.Hour), time.Now())

	assert.NoError(t, err)
	if assert.Len(t, entries, 2) {
		assert.Equal(t, 12.0, entries[0].Value)
		assert.Equal(t, 24.0, entries[1].Value)
	}
}

func TestClient_GetPowerOverview(t *testing.T) {
	server := &Server{token: "TESTTOKEN"}
	apiServer := httptest.NewServer(http.HandlerFunc(server.apiHandler))
	defer apiServer.Close()

	client := solaredge.Client{
		Token:      "TESTTOKEN",
		HTTPClient: &http.Client{},
		APIURL:     apiServer.URL,
	}

	siteIDs, err := client.GetSiteIDs(context.Background())
	assert.NoError(t, err)
	if assert.Len(t, siteIDs, 1) {
		assert.Equal(t, 1, siteIDs[0])
	}

	var lifeTime, lastYear, lastMonth, lastDay, current float64
	lifeTime, lastYear, lastMonth, lastDay, current, err = client.GetPowerOverview(context.Background(), siteIDs[0])

	if assert.NoError(t, err) {
		assert.Equal(t, 10000.0, lifeTime)
		assert.Equal(t, 1000.0, lastYear)
		assert.Equal(t, 100.0, lastMonth)
		assert.Equal(t, 10.0, lastDay)
		assert.Equal(t, 1.0, current)
	}
}
