package qywchat

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"webhooker/src/utils"
)

// Token accessToken
var Token string

// AccessToken accesstoken response
type AccessToken struct {
	ErrCode     int    `json:"errorcode,omitempty"`
	ErrMsg      string `json:"errmsg"`
	AccessToken string `json:"access_token,omitempty"`
	ExpireIn    int    `json:"expire_in"`
}

func getAccessToken() {
	query := url.Values{}
	query.Add("corpid", CorpID)
	query.Add("corpsecret", Corpsecret)

	u := fmt.Sprintf(URLGetToken + "?" + query.Encode())
	body, err := utils.DoRequest("GET", u, "")
	if err != nil {
		log.Fatalf("make http request error: %v\n", err)
		log.Panic("get access token error")
	}

	accessToken := AccessToken{}
	err = json.Unmarshal(body, &accessToken)
	if err != nil {
		panic("get access token error")
	}
	Token = accessToken.AccessToken
}
