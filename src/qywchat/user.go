package qywchat

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
)

// User user info
type User struct {
	Mobile       string   `json:"mobile,omitempty"`
	UID          string   `json:"userid,omitempty"`
	UserName     string   `json:"name,omitempty"`
	Department   []int    `json:"department,omitempty"`
	Position     string   `json:"position,omitempty"`
	Email        string   `json:"email,omitempty"`
	DirectLeader []string `json:"direct_leader,omitempty"`
}

// GetUIDResp get user id response
type GetUIDResp struct {
	ErrCode int    `json:"errorcode,omitempty"`
	ErrMsg  string `json:"errmsg"`
	UserID  string `json:"userid"`
}

// GetUserIDByMobile get user id by mobile
func GetUserIDByMobile(mobile string) (string, error) {
	query := url.Values{}
	query.Set("access_token", Token)

	u := fmt.Sprintf(URLGetUID + "?" + query.Encode())
	body := fmt.Sprintf("{\"mobile\":\"%s\"}", mobile)

	b, err := HTTPRequest("POST", u, body)
	if err != nil {
		return "", err
	}

	res := GetUIDResp{}
	err = json.Unmarshal(b, &res)
	if err != nil {
		return "", err
	}

	return string(res.UserID), nil
}

func GetUserInfo(uid string) (*User, error) {
	query := url.Values{}
	query.Set("access_token", Token)
	query.Set("userid", uid)
	u := fmt.Sprintf(URLGetUserInfo + "?" + query.Encode())
	b, err := HTTPRequest("GET", u, "")
	if err != nil {
		log.Printf("make request Error: %s\n", err)
	}

	user := &User{}
	err = json.Unmarshal(b, user)
	if err != nil {
		log.Printf("json Unmarshal error")
		return nil, err
	}

	return user, nil
}
