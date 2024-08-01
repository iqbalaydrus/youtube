package yt_msg_format

import (
	"google.golang.org/protobuf/proto"
	"testing"
	"yt-msg-format/pbuf"
)

func BenchmarkProtobufSerializeAll(b *testing.B) {
	var err error
	for i := 0; i < b.N; i++ {
		m := defaultMessage.ToProtoBuf()
		_, err = proto.Marshal(m)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkProtobufSerialize(b *testing.B) {
	var pl []byte
	var err error
	m := defaultMessage.ToProtoBuf()
	for i := 0; i < b.N; i++ {
		pl, err = proto.Marshal(m)
		if err != nil {
			b.Fatal(err)
		}
	}
	if printBenchLen {
		b.Logf("payload length: %d\n", len(pl))
	}
}

func BenchmarkProtobufSerializeGzip(b *testing.B) {
	var pl []byte
	var err error
	m := defaultMessage.ToProtoBuf()
	for i := 0; i < b.N; i++ {
		pl, err = proto.Marshal(m)
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

var protobufPayload = []byte{10, 7, 115, 117, 99, 99, 101, 115, 115, 18, 23, 10, 13, 10, 6, 105, 113, 98, 97, 108, 49, 18, 3, 99, 116, 111, 18, 6, 6, 8, 99, 199, 232, 67, 18, 24, 10, 13, 10, 6, 105, 113, 98, 97, 108, 50, 18, 3, 99, 116, 111, 18, 7, 6, 8, 99, 142, 209, 135, 1, 18, 24, 10, 13, 10, 6, 105, 113, 98, 97, 108, 51, 18, 3, 99, 116, 111, 18, 7, 6, 8, 99, 213, 185, 203, 1, 18, 24, 10, 13, 10, 6, 105, 113, 98, 97, 108, 52, 18, 3, 99, 116, 111, 18, 7, 6, 8, 99, 156, 162, 143, 2}

func BenchmarkProtobufDeserializeAll(b *testing.B) {
	var m pbuf.Message
	for i := 0; i < b.N; i++ {
		err := proto.Unmarshal(protobufPayload, &m)
		if err != nil {
			b.Fatal(err)
		}
		_ = NewMessageFromProtobuf(&m)
	}
}

func BenchmarkProtobufDeserialize(b *testing.B) {
	var m pbuf.Message
	for i := 0; i < b.N; i++ {
		err := proto.Unmarshal(protobufPayload, &m)
		if err != nil {
			b.Fatal(err)
		}
	}
}
