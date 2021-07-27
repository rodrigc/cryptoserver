package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testURL = "http://localhost:8081/currency"

// TestCurrencyAll calls the /currency/all endpoint on cryptoserver
func TestCurrencyAll(t *testing.T) {
	url := fmt.Sprintf("%s/all", testURL)
	resp, err := http.Get(url)
	assert.Nil(t, err)
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)

	assert.Equal(t, resp.StatusCode, 200)

	var dat []map[string]interface{}
	err = json.Unmarshal(b, &dat)
	assert.Nil(t, err)
}

// TestCurrencyBTC calls the /currency/btc endpoint on cryptoserver
func TestCurrencyBTC(t *testing.T) {
	url := fmt.Sprintf("%s/btc", testURL)
	resp, err := http.Get(url)
	assert.Nil(t, err)
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)

	assert.Equal(t, resp.StatusCode, 200)

	var dat map[string]interface{}
	err = json.Unmarshal(b, &dat)
	assert.Nil(t, err)

	val, ok := dat["fullName"]
	assert.True(t, ok)
	assert.Equal(t, val, "Bitcoin")
}

// TestCurrencyNonExistent calls /currency with a nonexistent currency
func TestCurrencyNonExistent(t *testing.T) {
	url := fmt.Sprintf("%s/nonexistent", testURL)
	resp, err := http.Get(url)
	assert.Nil(t, err)
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)

	assert.Equal(t, resp.StatusCode, 400)

	var dat map[string]interface{}
	err = json.Unmarshal(b, &dat)
	assert.Nil(t, err)

	val, ok := dat["error"]
	assert.True(t, ok)
	assert.NotEmpty(t, val)
}
