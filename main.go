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

func (c *Client) SendSticker(userId string, group string, id string) {

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
		"type":          "sticker",
		"sticker_group": group,
		"sticker_id":    id,
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

func (c *Client) SendImage(userId string, image []byte) {

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
		"type": "sticker",
		// "sticker_group": group,
		// "sticker_id":    id,
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

func (c *Client) SendBotStart(userId string, text string, id string) []byte {

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
		"type":      "botStart",
		"start_img": "5f5c1a23a913f7857c762da3476d1e67.JPG",
		"init": map[string]interface{}{
			// "image": "5f5c1a23a913f7857c762da3476d1e67.JPG",
			"hp": map[string]interface{}{
				"max":     100,
				"current": 50,
				"color":   "#6cdee0",
			},
			// "text": map[string]interface{}{
			// 	"message": "test",
			// 	"color":   "#6cdee0",
			// },
			// "button": map[string]interface{}{
			// 	"style": 2,
			// 	"setting": []map[string]interface{}{
			// 		{
			// 			"disabled": false,
			// 			"order":    1,
			// 			"text":     "開始？",
			// 			"command":  "/maidwhite start",
			// 		},
			// 	},
			// },
		},
	}

	payload := map[string]interface{}{
		"recipient": map[string]string{
			"id": userId,
		},
		"message": messagePayload,
	}

	payloadByte, _ := json.Marshal(payload)
	response, _ := c.send(payloadByte)
	return response

	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go
}

func (c *Client) SendBotEvent(userId string, eventId string, currentHp int) {

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
		"type":     "botEvent",
		"event_id": eventId,
		// "image": "5f5c1a23a913f7857c762da3476d1e67.JPG",
		"hp": map[string]interface{}{
			"max":     100,
			"current": currentHp,
			"color":   "#6cdee0",
		},
		// "text": map[string]interface{}{
		// 	"message": "test",
		// 	"color":   "#6cdee0",
		// },
		// "button": map[string]interface{}{
		// 	"style": 2,
		// 	"setting": []map[string]interface{}{
		// 		{
		// 			"disabled": false,
		// 			"order":    1,
		// 			"text":     "開始？",
		// 			"command":  "/maidwhite start",
		// 		},
		// 	},
		// },

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

func (c *Client) send(payload []byte) ([]byte, error) {
	body := bytes.NewReader(payload)
	req, err := http.NewRequest(http.MethodPost, Endpoint+"?access_token="+c.AccessToken, body)
	if err != nil {
		return nil, err
		// handle err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
		// handle err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
		// handle err
	}
	fmt.Println("baha response:", string(data))
	return data, nil
}

func (c *Client) unloadImage(payload []byte) error {
	body := bytes.NewReader(payload)
	req, err := http.NewRequest(http.MethodPost, Endpoint+"?access_token="+c.AccessToken, body)
	if err != nil {
		return err
		// handle err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

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
