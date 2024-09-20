package viewer

import (
	"errors"
	"fmt"
	"log"
	"math"
	"net/http"
	"strings"

	"github.com/dreamervulpi/tourneyBot/startgg"
	"github.com/gin-gonic/gin"
)

type Viewer struct {
	Client       *startgg.Client
	PhaseGroupId int64
}

type Set struct {
	RoundName     string
	State         startgg.State
	Player1       string
	Player2       string
	Player1_score string
	Player2_score string
}

func getScore(displayScore, garbage string) string {
	return strings.Split(strings.Split(displayScore, garbage)[1], " ")[1]
}

func (v *Viewer) prepareData() ([]any, error) {
	log.Println(v.PhaseGroupId)
	state, err := v.Client.GetPhaseGroupState(v.PhaseGroupId)
	if err != nil {
		return nil, err
	}

	result := []any{}

	winners := []Set{}
	losers := []Set{}

	// Realise: Change to InProgress
	if state == startgg.IsDone {
		// Realise: Change query to 1
		total, err := v.Client.GetPagesCount(v.PhaseGroupId)
		if err != nil {
			return nil, err
		}
		if total == 0 {
			return nil, fmt.Errorf("prepareData: total count = 0")
		}

		var pages int
		if total <= 60 {
			pages = 1
		} else {
			pages = int(math.Round(float64(total / 60)))
		}

		for i := 0; i < pages; i++ {
			sets, err := v.getBracket(pages, 60)
			if err != nil {
				log.Println(errors.New("error get sets"))
			}
			for _, set := range sets {
				player1_score := getScore(set.DisplayScore, set.Slots[0].Entrant.Name)
				player2_score := getScore(set.DisplayScore, set.Slots[1].Entrant.Name)
				preparedData := Set{
					RoundName:     set.FullRoundText,
					State:         set.State,
					Player1:       set.Slots[0].Entrant.Participants[0].GamerTag,
					Player2:       set.Slots[1].Entrant.Participants[0].GamerTag,
					Player1_score: player1_score,
					Player2_score: player2_score,
				}
				if set.Round > 0 {
					winners = append(winners, preparedData)
				} else {
					losers = append(losers, preparedData)
				}
			}
		}
	}

	result = append(result, winners)
	result = append(result, losers)
	return result, nil
}

func (v *Viewer) Run() error {
	// brackets[0] = slice winners bracket
	// brackets[1] = slice losers bracket
	brackets, err := v.prepareData()
	if err != nil {
		return err
	}

	r := gin.Default()
	r.LoadHTMLGlob("internal/templates/*")
	r.Static("/config", "config")
	r.GET("/settings", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	// elements := []string{"ddddddd", "ssssssss", "aaaaaa"}
	r.GET("/top8", func(c *gin.Context) {
		// Call the HTML method of the Context to render a template
		c.HTML(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			// Use the index.html template
			"top8.html",
			// Pass the data that the page uses (in this case, 'title')
			gin.H{
				"title":   "Top8",
				"winners": brackets[0],
				"losers":  brackets[1],
			},
		)
	})
	return r.Run(":7777")
}
