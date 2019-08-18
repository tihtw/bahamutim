package bahamutim

// {"botid":"bot@243","time":1565368539146,"messaging":[{"sender_id":"pichubaby","message":{"text":"測試"}}]}

import (
	"encoding/json"
)

type Request struct {
	Botid     string `json:"botid"`
	Time      int64  `json:"time"`
	Messaging []struct {
		SenderID string `json:"sender_id"`
		Message  struct {
			Text string `json:"text"`
		} `json:"message"`
	} `json:"messaging"`
}

func ParseRequest(data []byte) *Request {
	r := Request{}
	json.Unmarshal(data, &r)
	return &r
}

func (r *Request) GetUserID() string {
	return r.Messaging[0].SenderID
}
func (r *Request) GetText() string {
	return r.Messaging[0].Message.Text
}
