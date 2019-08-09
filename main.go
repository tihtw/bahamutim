package bahamutim

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	// "strings"
)

var Endpoint = "https://us-central1-hahamut-8888.cloudfunctions.net/messagePush"

func CheckSignature(key []byte, data []byte, signature []byte) bool {
	hash := hmac.New(sha1.New, key)
	hash.Write(data)
	excepted := hash.Sum(nil)
	// fmt.Println("data:", string(data))
	// fmt.Println("signature:", string(signature))
	// fmt.Println("excepted:", fmt.Sprintf("sha1=%x", excepted))
	return string(signature) == fmt.Sprintf("sha1=%x", excepted)
}

func NewClient(accessToken string) *Client {
	return &Client{
		AccessToken: accessToken,
	}

}

type Client struct {
	AccessToken string
}

func (c *Client) SendText(userId string, message string) {

	// 	{
	//  "recipient":{
	//    "id":"<SENDER_ID>"
	//  },
	//  "message":{
	//    "type":"text",
	//    "text":"你都知道了，柚子要失業了，嗚"
	//  }
	// }

	messagePayload := map[string]interface{}{
		"type": "text",
		"text": message,
	}

	payload := map[string]interface{}{
		"recipient": map[string]string{
			"id": userId,
		},
		"message": messagePayload,
	}

	payloadByte, _ := json.Marshal(payload)
	c.send(payloadByte)

	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go
}
func (c *Client) send(payload []byte) error {
	body := bytes.NewReader(payload)
	req, err := http.NewRequest(http.MethodPost, Endpoint+"?access_token="+c.AccessToken, body)
	if err != nil {
		return err
		// handle err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
		// handle err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
		// handle err
	}
	fmt.Println("baha response:", string(data))
	return nil
}
