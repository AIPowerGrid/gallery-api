package endpoints

import (
	"context"
	"fmt"
	redisLib "gallery/db/redis"
	"gallery/models"

	"github.com/go-redis/redis/v8"

	json "github.com/goccy/go-json"
)

var (
	ctx = context.Background()
)

func getColIds(sid, id string) ([]string, error) {

	client := redisLib.GetClient()
	k := fmt.Sprintf("usercollections:%s:%s", sid, id)
	log.Debug(k)
	f, err := client.SMembers(ctx, k).Result()
	return f, err

}
func AddToCollection(sid, id, jobid string) error {
	client := redisLib.GetClient()
	k := fmt.Sprintf("usercollections:%s:%s", sid, id)
	_, err := client.SAdd(ctx, k, jobid).Result()
	return err

}
func RemoveFromCollection(sid, id, jobid string) error {
	client := redisLib.GetClient()
	k := fmt.Sprintf("usercollections:%s:%s", sid, id)
	_, err := client.SRem(ctx, k, jobid).Result()
	return err
}
func GetCollection(sid, col string) ([]models.GalleryJob, error) { // todo fix this to use s3
	var jobs []models.GalleryJob
	client := redisLib.GetClient()
	ids, err := getColIds(sid, col)
	log.Debug(ids, 42)
	if err != nil {
		return nil, err
	}

	p := client.Pipeline()

	var cmders []*redis.StringCmd
	for _, id := range ids {

		k := fmt.Sprintf(`job:%s`, id)
		v := p.Get(ctx, k)
		cmders = append(cmders, v)
	}

	_, err = p.Exec(ctx)
	if err != nil {
		return nil, err
	}
	for _, c := range cmders {
		v := c.Val()
		var job models.GalleryJob
		if err := json.Unmarshal([]byte(v), &job); err != nil {
			return nil, err
		}
		// f := toBase64([]byte(job.ImageData))
		// job.ImageData = f

		jobs = append(jobs, job)
	}
	return jobs, nil

}
