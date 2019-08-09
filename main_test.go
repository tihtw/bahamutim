package bahamutim

import (
	"testing"
)

func TestCheckSignature(t *testing.T) {
	cases := []struct {
		Key       string
		Data      string
		Signature string
		Excepted  bool
	}{
		{
			Key:       "00000000000000000000",
			Data:      `{"botid":"bot@243","time":1565363950108,"messaging":[{"sender_id":"pichubaby","message":{"text":"測試"}}]}`,
			Signature: `sha1=d12ba1eb4e095f0ddcad737686b07855b6dbb5ef`,
			Excepted:  true,
		},
		{
			Key:       "00000000000000000000",
			Data:      `{"botid":"bot@243","time":1565363950108,"messaging":[{"sender_id":"pichubaby","message":{"text":"測試"}}]}`,
			Signature: `sha1=d12ba1eb4e095f0ddcad737686b07855b6dbb5e1`,
			Excepted:  false,
		},
		{
			Key:       "00000000000000000001",
			Data:      `{"botid":"bot@243","time":1565363950108,"messaging":[{"sender_id":"pichubaby","message":{"text":"測試"}}]}`,
			Signature: `sha1=d12ba1eb4e095f0ddcad737686b07855b6dbb5ef`,
			Excepted:  false,
		},
		{
			Key:       "00000000000000000000",
			Data:      `{"botid":"bot@243","time":1565363950107,"messaging":[{"sender_id":"pichubaby","message":{"text":"測試"}}]}`,
			Signature: `sha1=d12ba1eb4e095f0ddcad737686b07855b6dbb5ef`,
			Excepted:  false,
		},
	}

	for i, c := range cases {
		actual := CheckSignature([]byte(c.Key), []byte(c.Data), []byte(c.Signature))
		if actual != c.Excepted {
			t.Errorf("excpeted case %d: %v, got %v", i, c.Excepted, actual)
		}
	}

}
