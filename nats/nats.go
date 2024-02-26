package nats

import (
	"fmt"
	"gallery/core"
	"gallery/models"
	"os"
	"time"

	json "github.com/goccy/go-json"
	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
)

var (
	log = core.GetLogger()
)

var NatsConnection *nats.Conn

func Start() {
	url := os.Getenv("NATS_URL")

	conn, err := nats.Connect(url, nats.UserInfo("admin", "meamadmin321"))
	if err != nil {
		log.Fatal(err)
	}
	NatsConnection = conn

	log.Info("[nats] connected to " + url)

	natsAPI()

}
func ComfySendReq(prompt string, seed int64) (models.Job, error) {
	log.Infof("Sending Comfy Request with prompt %s and seed %d", prompt, seed)
	exJob := models.Job{
		ID:     uuid.NewString(),
		Type:   "imagec",
		Prompt: prompt,
		Seed:   seed,
		Model:  "sdxl",
	}
	b, err := json.Marshal(exJob)
	if err != nil {
		log.Error(err)
		return exJob, err
	}

	_, err = NatsConnection.Request("comfy.request", b, time.Second*3)
	if err != nil {
		log.Error(err)
		return exJob, err
	}
	return exJob, nil

}
func ComfyRequest(prompt string, seed int64) (models.ComfyResponse, error) {
	log.Infof("Sending Comfy Request with prompt %s and seed %d", prompt, seed)
	var resp models.ComfyResponse
	exJob := models.Job{
		ID:     uuid.NewString(),
		Prompt: prompt,
		Seed:   seed,
		Type:   "image",
	}
	b, err := json.Marshal(exJob)
	if err != nil {
		log.Error(err)
		return resp, err
	}

	respChan := fmt.Sprintf("comfy.response.%s", exJob.ID)
	sub, err := NatsConnection.SubscribeSync(respChan)
	if err != nil {
		log.Error(err)
		return resp, err
	}
	_, err = NatsConnection.Request("comfy.request", b, time.Second*3)
	if err != nil {
		log.Error(err)
		return resp, err
	}
	log.Info("Waiting for Reply...")
	next, err := sub.NextMsg(time.Second * 30)
	if err != nil {
		log.Error(err)
		return resp, err
	}
	err = json.Unmarshal(next.Data, &resp)
	if err != nil {
		log.Error(err)
		return resp, err
	}
	go sub.Unsubscribe()
	return resp, nil
}
func natsAPI() {
	log.Info("Starting nats api...")
	// seed := time.Now().UnixMilli()
	go func() {
		for {
			// log.Info("sending request...")
			// go ComfyRequest("Chicken fingers", seed)
			time.Sleep(time.Second * 5)
		}
	}()

}
