package main

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/dreamervulpi/tourneyBot/startgg"
	"github.com/dreamervulpi/tourneyViewer/config"
	"github.com/dreamervulpi/tourneyViewer/internal/viewer"
)

func main() {
	cfg, err := config.LoadConfig("config/config.toml")
	if err != nil {
		log.Println(errors.New("not loaded: ").Error() + err.Error())
	} else {
		v := viewer.Viewer{
			Client: startgg.NewClient(cfg.Token, &http.Client{
				Timeout: time.Second * 10,
			}),
			PhaseGroupId: cfg.PhaseGroupId,
		}
		if err := v.Run(); err != nil {
			log.Println(err.Error())
		}
	}
}
