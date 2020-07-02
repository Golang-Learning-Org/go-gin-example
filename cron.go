package main

import (
	"log"
	"time"

	"github.com/evanxzj/go-gin-example/models"

	"github.com/robfig/cron/v3"
)

func main() {
	log.Println("Starting...")

	c := cron.New(cron.WithSeconds())
	c.AddFunc("*/10 * * * * *", func() {
		log.Println("Run models.CleanAllTag...")
		models.ClearAllTag()
	})
	c.AddFunc("*/10 * * * * *", func() {
		log.Println("Run models.CleanAllArticle...")
		models.ClearAllArticle()
	})

	c.Start()

	t := time.NewTimer(10 * time.Second)
	for {
		select {
		case <-t.C:
			t.Reset(10 * time.Second)
		}
	}
}
