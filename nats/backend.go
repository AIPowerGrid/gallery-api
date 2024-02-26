package nats

import "github.com/nats-io/nats.go"

func GetBackend(m *nats.Msg) {
	defer natsPanic(m.Respond)

}
func SetBackendDB(m *nats.Msg) {
	defer natsPanic(m.Respond)
}
