package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func ApiCall(url string, method string, data []byte) string {
	request, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	bodyString := string(bodyBytes)

	return bodyString
}
