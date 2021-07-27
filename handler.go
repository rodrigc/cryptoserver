package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

// CurrencyHandler calls the /currency endpoint on hitbtc.com
func CurrencyHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	url := fmt.Sprintf("%s/currency", HiBTCBaseEndpoint)
	muxVars := mux.Vars(r)

	if symbol, ok := muxVars["symbol"]; ok {
		// If we specify /currency/<symbol> , we want to call
		//  /api/v2/currency/<symbol> to get data for that symbol.
		// but if we specify /currency/all , we want to call
		// /api/v2/currency to get data for all symbols
		if symbol != "all" {
			url = fmt.Sprintf("%s/%s", url, symbol)
		}

	} else {
		reportError(w, fmt.Errorf("no symbol specified"))
		return
	}

	resp, err := http.Get(url)
	if err != nil {
		reportError(w, err)
		return
	}

	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		reportError(w, err)
		return
	}
	w.WriteHeader(resp.StatusCode)
	w.Write(b)

}
