package qywchat

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"webhooker/src/utils"
)

// Response qywchat response
type Response struct {
	Code int    `json:"errcode,omitempty"`
	Msg  string `json:"errmsg,omitempty"`
}

// HTTPRequest make http request
func HTTPRequest(method string, path string, body string) (b []byte, err error) {
	resq, err := utils.DoRequest(method, path, body)
	if err != nil {
		log.Printf("doRequest err: %v\n", err)
		return nil, err
	}

	resp := Response{}
	err = json.Unmarshal(resq, &resp)
	if err != nil {
		log.Println("Error unmarshalling response.")
		return nil, err
	}

	if resp.Code == 0 {
		return resq, nil
	}

	if resp.Code == 40014 || resp.Code == 42001 {
		return utils.DoRequest(method, refreshURL(path), body)
	}
	// log.Printf("HTTP request got error code: %d, of msg: %s", resp.Code, resp.Msg)
	return nil, fmt.Errorf("HTTP request got error code: %d, of msg: %s", resp.Code, resp.Msg)
}

func refreshURL(u string) string {
	getAccessToken()
	path, err := url.Parse(u)
	if err != nil {
		log.Panic(err)
	}
	query := path.Query()
	query.Set("access_token", Token)
	path.RawQuery = query.Encode()
	return path.String()
}
