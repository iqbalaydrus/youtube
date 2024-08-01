package yt_msg_format

import (
	"capnproto.org/go/capnp/v3"
	"testing"
)

func BenchmarkCapnProtoSerializeAll(b *testing.B) {
	var pl []byte
	var err error
	for i := 0; i < b.N; i++ {
		m := defaultMessage.ToCapnProto()
		pl, err = m.Marshal()
		if err != nil {
			b.Fatal(err)
		}
	}
	if printBenchLen {
		b.Logf("payload length: %d\n", len(pl))
	}
}

func BenchmarkCapnProtoSerialize(b *testing.B) {
	var pl []byte
	var err error
	m := defaultMessage.ToCapnProto()
	for i := 0; i < b.N; i++ {
		pl, err = m.Marshal()
		if err != nil {
			b.Fatal(err)
		}
	}
	if printBenchLen {
		b.Logf("payload length: %d\n", len(pl))
	}
}

func BenchmarkCapnProtoSerializeCompact(b *testing.B) {
	var pl []byte
	var err error
	m := defaultMessage.ToCapnProto()
	for i := 0; i < b.N; i++ {
		pl, err = m.MarshalPacked()
		if err != nil {
			b.Fatal(err)
		}
	}
	if printBenchLen {
		b.Logf("payload length: %d\n", len(pl))
	}
}

func BenchmarkCapnProtoSerializeGzip(b *testing.B) {
	var pl []byte
	var err error
	m := defaultMessage.ToCapnProto()
	for i := 0; i < b.N; i++ {
		pl, err = m.Marshal()
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

func BenchmarkCapnProtoSerializeCompactGzip(b *testing.B) {
	var pl []byte
	var err error
	m := defaultMessage.ToCapnProto()
	for i := 0; i < b.N; i++ {
		pl, err = m.MarshalPacked()
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

var capnProtoPayload = []byte{0, 0, 0, 0, 85, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 5, 0, 0, 0, 66, 0, 0, 0, 5, 0, 0, 0, 71, 0, 0, 0, 115, 117, 99, 99, 101, 115, 115, 0, 16, 0, 0, 0, 0, 0, 2, 0, 68, 0, 0, 0, 0, 0, 2, 0, 81, 0, 0, 0, 37, 0, 0, 0, 132, 0, 0, 0, 0, 0, 2, 0, 145, 0, 0, 0, 37, 0, 0, 0, 196, 0, 0, 0, 0, 0, 2, 0, 209, 0, 0, 0, 37, 0, 0, 0, 4, 1, 0, 0, 0, 0, 2, 0, 17, 1, 0, 0, 37, 0, 0, 0, 4, 0, 0, 0, 0, 0, 2, 0, 17, 0, 0, 0, 37, 0, 0, 0, 9, 0, 0, 0, 58, 0, 0, 0, 1, 0, 0, 0, 34, 0, 0, 0, 99, 116, 111, 0, 0, 0, 0, 0, 105, 113, 98, 97, 108, 49, 0, 0, 6, 0, 0, 0, 0, 0, 0, 0, 8, 0, 0, 0, 0, 0, 0, 0, 99, 0, 0, 0, 0, 0, 0, 0, 71, 244, 16, 0, 0, 0, 0, 0, 5, 0, 0, 0, 58, 0, 0, 0, 5, 0, 0, 0, 34, 0, 0, 0, 105, 113, 98, 97, 108, 49, 0, 0, 99, 116, 111, 0, 0, 0, 0, 0, 6, 0, 0, 0, 0, 0, 0, 0, 8, 0, 0, 0, 0, 0, 0, 0, 99, 0, 0, 0, 0, 0, 0, 0, 71, 244, 16, 0, 0, 0, 0, 0, 4, 0, 0, 0, 0, 0, 2, 0, 17, 0, 0, 0, 37, 0, 0, 0, 9, 0, 0, 0, 58, 0, 0, 0, 1, 0, 0, 0, 34, 0, 0, 0, 99, 116, 111, 0, 0, 0, 0, 0, 105, 113, 98, 97, 108, 50, 0, 0, 6, 0, 0, 0, 0, 0, 0, 0, 8, 0, 0, 0, 0, 0, 0, 0, 99, 0, 0, 0, 0, 0, 0, 0, 142, 232, 33, 0, 0, 0, 0, 0, 5, 0, 0, 0, 58, 0, 0, 0, 5, 0, 0, 0, 34, 0, 0, 0, 105, 113, 98, 97, 108, 50, 0, 0, 99, 116, 111, 0, 0, 0, 0, 0, 6, 0, 0, 0, 0, 0, 0, 0, 8, 0, 0, 0, 0, 0, 0, 0, 99, 0, 0, 0, 0, 0, 0, 0, 142, 232, 33, 0, 0, 0, 0, 0, 4, 0, 0, 0, 0, 0, 2, 0, 17, 0, 0, 0, 37, 0, 0, 0, 9, 0, 0, 0, 58, 0, 0, 0, 1, 0, 0, 0, 34, 0, 0, 0, 99, 116, 111, 0, 0, 0, 0, 0, 105, 113, 98, 97, 108, 51, 0, 0, 6, 0, 0, 0, 0, 0, 0, 0, 8, 0, 0, 0, 0, 0, 0, 0, 99, 0, 0, 0, 0, 0, 0, 0, 213, 220, 50, 0, 0, 0, 0, 0, 5, 0, 0, 0, 58, 0, 0, 0, 5, 0, 0, 0, 34, 0, 0, 0, 105, 113, 98, 97, 108, 51, 0, 0, 99, 116, 111, 0, 0, 0, 0, 0, 6, 0, 0, 0, 0, 0, 0, 0, 8, 0, 0, 0, 0, 0, 0, 0, 99, 0, 0, 0, 0, 0, 0, 0, 213, 220, 50, 0, 0, 0, 0, 0, 4, 0, 0, 0, 0, 0, 2, 0, 17, 0, 0, 0, 37, 0, 0, 0, 9, 0, 0, 0, 58, 0, 0, 0, 1, 0, 0, 0, 34, 0, 0, 0, 99, 116, 111, 0, 0, 0, 0, 0, 105, 113, 98, 97, 108, 52, 0, 0, 6, 0, 0, 0, 0, 0, 0, 0, 8, 0, 0, 0, 0, 0, 0, 0, 99, 0, 0, 0, 0, 0, 0, 0, 28, 209, 67, 0, 0, 0, 0, 0, 5, 0, 0, 0, 58, 0, 0, 0, 5, 0, 0, 0, 34, 0, 0, 0, 105, 113, 98, 97, 108, 52, 0, 0, 99, 116, 111, 0, 0, 0, 0, 0, 6, 0, 0, 0, 0, 0, 0, 0, 8, 0, 0, 0, 0, 0, 0, 0, 99, 0, 0, 0, 0, 0, 0, 0, 28, 209, 67, 0, 0, 0, 0, 0}

func BenchmarkCapnProtoDeserializeAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m, err := capnp.Unmarshal(capnProtoPayload)
		if err != nil {
			b.Fatal(err)
		}
		_ = NewMessageFromCapnProto(m)
	}
}

func BenchmarkCapnProtoDeserialize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := capnp.Unmarshal(capnProtoPayload)
		if err != nil {
			b.Fatal(err)
		}
	}
}
