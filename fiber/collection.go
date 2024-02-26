package fiber

import (
	"gallery/core"
	"gallery/endpoints"

	"github.com/gofiber/fiber/v2"
)

func getCollection(c *fiber.Ctx) error {
	id := c.Params("id")
	sid := c.Get("sessionId")
	if sid == "" {
		return core.SendErrString(c, "no session id")
	}
	cols, err := endpoints.GetCollection(sid, id)
	if err != nil {
		log.Error(err)
		return core.SendErr(c, err, false)
	}
	return c.JSON(cols)
}
func AddToCollection(c *fiber.Ctx) error {
	id := c.Params("id")

	sid := c.Get("sessionId")
	if sid == "" {
		return core.SendErrString(c, "no session id")
	}
	jobId := c.Query("id")
	if jobId == "" {
		return core.SendErrString(c, "no id passed")
	}
	job, err := endpoints.GetJob(jobId)
	if err != nil {
		log.Error(err)
		return core.SendErr(c, err, false)
	}
	if job.Prompt == "" || job.Seed == "" {
		return core.SendErrString(c, "job id is not valid job")
	}
	err = endpoints.AddToCollection(sid, id, jobId)
	if err != nil {
		log.Error(err)
		return core.SendErr(c, err, false)
	}
	return c.JSON(fiber.Map{"success": true})

}

type rmPayload struct {
	ID    string `json:"_id"`
	JobID string `json:"jobid"`
}

func rmFromCollection(c *fiber.Ctx) error {
	sid := c.Get("sessionId")
	if sid == "" {
		return core.SendErrString(c, "no session id")
	}
	colId := c.Params("id")
	jobId := c.Params("job")
	err := endpoints.RemoveFromCollection(sid, colId, jobId)
	if err != nil {
		return core.SendErr(c, err, true)
	}
	items, err := endpoints.GetCollection(sid, colId)
	if err != nil {
		log.Error(err)
		return core.SendErr(c, err, false)
	}
	return c.JSON(items)
}
