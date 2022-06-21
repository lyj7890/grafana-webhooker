package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func DoRequest(method string, uri string, body string) (b []byte, err error) {
	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest(method, uri, strings.NewReader(body))
	// log.Printf("do http request, %s\t%s\t%s\n", method, uri, strings.NewReader(body))
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf(fmt.Sprintf("Response http code is not 200: %d", resp.StatusCode))
	}
	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}
