package endpoints

import (
	"encoding/base64"
	"fmt"
	"gallery/core"
	"gallery/db/redis"
	"gallery/models"

	json "github.com/goccy/go-json"
)

var (
	log = core.GetLogger()
)

func GetJob(id string) (models.GalleryJob, error) {

	k := fmt.Sprintf("galleryjob:%s", id)
	var job models.GalleryJob
	s, err := redis.Get(k)
	if err != nil {
		return job, err
		// return nil, err
	}
	err = json.Unmarshal([]byte(s), &job)
	return job, err
}

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func GetImage(id string) (models.GalleryJob, error) {

	var j models.GalleryJob
	k := fmt.Sprintf("galleryjob:%s", id)
	s, err := redis.Get(k)
	if err != nil {
		log.Error(err)
		return j, err
	}
	err = json.Unmarshal([]byte(s), &j)
	if err != nil {
		return j, err
	}
	return j, nil
}
