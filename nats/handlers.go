package nats

import (
	"fmt"
	"gallery/models"

	json "github.com/goccy/go-json"
	"github.com/nats-io/nats.go"
)

type urlQuery struct {
	Id string `json:"id"`
}

func _returnErr(m *nats.Msg, err error, msg string) {
	log.Error(err)
	var r JSResponse
	if msg != "" {
		r = JSResponse{Success: false, Message: msg}
	} else {
		r = JSResponse{Success: false, Message: err.Error()}
	}
	b, _ := json.Marshal(r)
	m.Respond(b)
}

type saveNodeData struct {
	Node      *models.Node `json:"node"`
	ProjectID string       `json:"projectId"`
}

func natsPanic(cb func([]byte) error) {
	if r := recover(); r != nil {
		fmt.Println("Recovering from panic:", r)
		s := "Recovered Panic: "
		s += fmt.Sprint(r)
		f := JSResponse{Success: false, Message: s}
		b, err := json.Marshal(f)
		if err == nil {
			cb(b)
			// m.Respond(b)
		}
	}
}
