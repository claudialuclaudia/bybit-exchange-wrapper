package util

import (
	"encoding/json"
	"net/http"
)

func GetNGNExchangeRates(symbols string) (map[string]string, error) {
	var rates = map[string]string{}
	var err error
	
	client := http.Client{}
	
	request_url := "https://api.exchangerate.host/latest?base=NGN&places=18&symbols=" + symbols
	request, err := http.NewRequest("GET", request_url, nil)
	if err != nil { return rates, err }

	resp, err := client.Do(request)
	if err != nil { return rates, err }
	
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	rrates := result["rates"]
	for k, v := range rrates.(map[string]interface{}) {
		rates[k] = v.(string)
	}
	return rates, err
}