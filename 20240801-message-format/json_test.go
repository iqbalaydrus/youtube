package yt_msg_format

import (
	"encoding/json"
	"testing"
)

func BenchmarkJSONSerialize(b *testing.B) {
	var pl []byte
	var err error
	for i := 0; i < b.N; i++ {
		pl, err = json.Marshal(defaultMessage)
		if err != nil {
			b.Fatal(err)
		}
	}
	if printBenchLen {
		b.Logf("payload length: %d\n", len(pl))
	}
}

func BenchmarkJSONSerializeGzip(b *testing.B) {
	var pl []byte
	var err error
	for i := 0; i < b.N; i++ {
		pl, err = json.Marshal(defaultMessage)
		if err != nil {
			b.Fatal(err)
		}
		pl, err = compress(pl)
		if err != nil {
			b.Fatal(err)
		}
	}
	if printBenchLen {
		b.Logf("payload length: %d\n", len(pl))
	}
}

var jsonPayload = []byte(`{
  "message": "success",
  "result": [
    {
      "employee": {
        "name": "iqbal1",
        "position": "cto"
      },
      "scores": [6, 8, 99, 1111111]
    },
    {
      "employee": {
        "name": "iqbal2",
        "position": "cto"
      },
      "scores": [6, 8, 99, 2222222]
    },
    {
      "employee": {
        "name": "iqbal3",
        "position": "cto"
      },
      "scores": [6, 8, 99, 3333333]
    },
    {
      "employee": {
        "name": "iqbal4",
        "position": "cto"
      },
      "scores": [6, 8, 99, 4444444]
    }
  ]
}
`)

func BenchmarkJSONDeserialize(b *testing.B) {
	var m Message
	for i := 0; i < b.N; i++ {
		err := json.Unmarshal(jsonPayload, &m)
		if err != nil {
			b.Fatal(err)
		}
	}
}
