package yt_msg_format

import (
	"github.com/vmihailenco/msgpack/v5"
	"testing"
)

func BenchmarkMsgPackSerialize(b *testing.B) {
	var pl []byte
	var err error
	for i := 0; i < b.N; i++ {
		pl, err = msgpack.Marshal(defaultMessage)
		if err != nil {
			b.Fatal(err)
		}
	}
	if printBenchLen {
		b.Logf("payload length: %d\n", len(pl))
	}
}

func BenchmarkMsgPackSerializeGzip(b *testing.B) {
	var pl []byte
	var err error
	for i := 0; i < b.N; i++ {
		pl, err = msgpack.Marshal(defaultMessage)
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

var msgPackPayload = []byte{130, 167, 77, 101, 115, 115, 97, 103, 101, 167, 115, 117, 99, 99, 101, 115, 115, 166, 82, 101, 115, 117, 108, 116, 148, 130, 168, 69, 109, 112, 108, 111, 121, 101, 101, 130, 164, 78, 97, 109, 101, 166, 105, 113, 98, 97, 108, 49, 168, 80, 111, 115, 105, 116, 105, 111, 110, 163, 99, 116, 111, 166, 83, 99, 111, 114, 101, 115, 148, 207, 0, 0, 0, 0, 0, 0, 0, 6, 207, 0, 0, 0, 0, 0, 0, 0, 8, 207, 0, 0, 0, 0, 0, 0, 0, 99, 207, 0, 0, 0, 0, 0, 16, 244, 71, 130, 168, 69, 109, 112, 108, 111, 121, 101, 101, 130, 164, 78, 97, 109, 101, 166, 105, 113, 98, 97, 108, 50, 168, 80, 111, 115, 105, 116, 105, 111, 110, 163, 99, 116, 111, 166, 83, 99, 111, 114, 101, 115, 148, 207, 0, 0, 0, 0, 0, 0, 0, 6, 207, 0, 0, 0, 0, 0, 0, 0, 8, 207, 0, 0, 0, 0, 0, 0, 0, 99, 207, 0, 0, 0, 0, 0, 33, 232, 142, 130, 168, 69, 109, 112, 108, 111, 121, 101, 101, 130, 164, 78, 97, 109, 101, 166, 105, 113, 98, 97, 108, 51, 168, 80, 111, 115, 105, 116, 105, 111, 110, 163, 99, 116, 111, 166, 83, 99, 111, 114, 101, 115, 148, 207, 0, 0, 0, 0, 0, 0, 0, 6, 207, 0, 0, 0, 0, 0, 0, 0, 8, 207, 0, 0, 0, 0, 0, 0, 0, 99, 207, 0, 0, 0, 0, 0, 50, 220, 213, 130, 168, 69, 109, 112, 108, 111, 121, 101, 101, 130, 164, 78, 97, 109, 101, 166, 105, 113, 98, 97, 108, 52, 168, 80, 111, 115, 105, 116, 105, 111, 110, 163, 99, 116, 111, 166, 83, 99, 111, 114, 101, 115, 148, 207, 0, 0, 0, 0, 0, 0, 0, 6, 207, 0, 0, 0, 0, 0, 0, 0, 8, 207, 0, 0, 0, 0, 0, 0, 0, 99, 207, 0, 0, 0, 0, 0, 67, 209, 28}

func BenchmarkMsgPackDeserialize(b *testing.B) {
	var m Message
	for i := 0; i < b.N; i++ {
		err := msgpack.Unmarshal(msgPackPayload, &m)
		if err != nil {
			b.Fatal(err)
		}
	}
}
