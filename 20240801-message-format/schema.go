package yt_msg_format

import (
	"bytes"
	"capnproto.org/go/capnp/v3"
	"compress/gzip"
	"yt-msg-format/capn"
	"yt-msg-format/pbuf"
)

type Employee struct {
	Name     string `json:"name"`
	Position string `json:"position"`
}

type Result struct {
	Employee Employee `json:"employee"`
	Scores   []uint64 `json:"scores"`
}

type Message struct {
	Message string   `json:"message"`
	Result  []Result `json:"result"`
}

var printBenchLen = false

var defaultMessage = Message{
	Message: "success",
	Result: []Result{
		{
			Employee: Employee{
				Name:     "iqbal1",
				Position: "cto",
			},
			Scores: []uint64{6, 8, 99, 1111111},
		},
		{
			Employee: Employee{
				Name:     "iqbal2",
				Position: "cto",
			},
			Scores: []uint64{6, 8, 99, 2222222},
		},
		{
			Employee: Employee{
				Name:     "iqbal3",
				Position: "cto",
			},
			Scores: []uint64{6, 8, 99, 3333333},
		},
		{
			Employee: Employee{
				Name:     "iqbal4",
				Position: "cto",
			},
			Scores: []uint64{6, 8, 99, 4444444},
		},
	},
}

func NewMessageFromProtobuf(m *pbuf.Message) *Message {
	var resList []Result
	for _, res := range m.GetResult() {
		resList = append(resList, Result{
			Employee: Employee{
				Name:     res.GetEmployee().GetName(),
				Position: res.GetEmployee().GetPosition(),
			},
			Scores: res.GetScores(),
		})
	}
	return &Message{
		Message: m.GetMessage(),
		Result:  resList,
	}
}

func NewMessageFromCapnProto(m *capnp.Message) *Message {
	mProto, err := capn.ReadRootMessage(m)
	if err != nil {
		panic(err)
	}
	var resList []Result
	resListProto, err := mProto.Result()
	if err != nil {
		panic(err)
	}
	for i := 0; i < resListProto.Len(); i++ {
		resProto := resListProto.At(i)
		scoresProto, err := resProto.Scores()
		if err != nil {
			panic(err)
		}
		var scores []uint64
		for j := 0; j < scoresProto.Len(); j++ {
			scores = append(scores, scoresProto.At(j))
		}
		employeeProto, err := resProto.Employee()
		if err != nil {
			panic(err)
		}
		employeeName, err := employeeProto.Name()
		if err != nil {
			panic(err)
		}
		employeePosition, err := employeeProto.Position()
		if err != nil {
			panic(err)
		}
		resList = append(resList, Result{
			Employee: Employee{
				Name:     employeeName,
				Position: employeePosition,
			},
			Scores: scores,
		})
	}
	message, err := mProto.Message_()
	if err != nil {
		panic(err)
	}
	return &Message{
		Message: message,
		Result:  resList,
	}
}

func compress(data []byte) ([]byte, error) {
	var plGzip bytes.Buffer
	w := gzip.NewWriter(&plGzip)
	_, err := w.Write(data)
	if err != nil {
		return nil, err
	}
	err = w.Close()
	if err != nil {
		return nil, err
	}
	return plGzip.Bytes(), nil
}

func (m *Message) ToProtoBuf() *pbuf.Message {
	if m == nil {
		return nil
	}
	mProto := pbuf.Message{
		Message: m.Message,
	}
	for _, res := range m.Result {
		mProto.Result = append(mProto.Result, &pbuf.Result{
			Employee: &pbuf.Employee{
				Name:     res.Employee.Name,
				Position: res.Employee.Position,
			},
			Scores: res.Scores,
		})
	}
	return &mProto
}

func (m *Message) ToCapnProto() *capnp.Message {
	if m == nil {
		return nil
	}
	arena := capnp.SingleSegment(nil)
	msg, seg, err := capnp.NewMessage(arena)
	if err != nil {
		panic(err)
	}
	mCapn, err := capn.NewRootMessage(seg)
	if err != nil {
		panic(err)
	}
	if err = mCapn.SetMessage_(m.Message); err != nil {
		panic(err)
	}
	if resList, err := mCapn.NewResult(int32(len(m.Result))); err != nil {
		panic(err)
	} else {
		for i, res := range m.Result {
			resCapn, err := capn.NewResult(resList.Segment())
			if err != nil {
				panic(err)
			}
			emplCapn, err := resCapn.NewEmployee()
			if err != nil {
				panic(err)
			}
			if err = emplCapn.SetPosition(res.Employee.Position); err != nil {
				panic(err)
			}
			if err = emplCapn.SetName(res.Employee.Name); err != nil {
				panic(err)
			}
			if err = resCapn.SetEmployee(emplCapn); err != nil {
				panic(err)
			}
			scoresCapn, err := capnp.NewUInt64List(resList.Segment(), int32(len(res.Scores)))
			if err != nil {
				panic(err)
			}
			for j, score := range res.Scores {
				scoresCapn.Set(j, score)
			}
			if err = resCapn.SetScores(scoresCapn); err != nil {
				panic(err)
			}
			if err = resList.Set(i, resCapn); err != nil {
				panic(err)
			}
		}
		if err = mCapn.SetResult(resList); err != nil {
			panic(err)
		}
	}
	return msg
}
