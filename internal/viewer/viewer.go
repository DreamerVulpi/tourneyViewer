package viewer

import (
	"errors"
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
	Winners      []Round
	Losers       []Round
}

type Round struct {
	Name string
	Sets []Set
}

type Set struct {
	State         startgg.State
	Player1       string
	Player2       string
	Player1_score string
	Player2_score string
}

func getScore(displayScore, garbage string) string {
	return strings.Split(strings.Split(displayScore, garbage)[1], " ")[1]
}

func (v *Viewer) prepareData() ([]map[int]Round, error) {
	log.Println(v.PhaseGroupId)

	result := []map[int]Round{}

	winners := map[int]Round{}
	losers := map[int]Round{}

	total, err := v.Client.GetPagesCount(v.PhaseGroupId)
	if err != nil {
		return nil, err
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
		for _, node := range sets {
			// prepare data set
			var p1 string
			var player1_score string
			if node.Slots[0].Entrant != nil {
				p1 = node.Slots[0].Entrant.Participants[0].GamerTag
				player1_score = getScore(node.DisplayScore, node.Slots[0].Entrant.Name)
			} else {
				p1 = "1"
				player1_score = "1"
			}
			var p2 string
			var player2_score string
			if node.Slots[1].Entrant != nil {
				p2 = node.Slots[1].Entrant.Participants[0].GamerTag
				player2_score = getScore(node.DisplayScore, node.Slots[1].Entrant.Name)
			} else {
				p2 = "1"
				player2_score = "1"
			}

			set := Set{
				State:         node.State,
				Player1:       p1,
				Player2:       p2,
				Player1_score: player1_score,
				Player2_score: player2_score,
			}

			if node.Round > 0 {
				// save set to winners bracket
				value, ok := winners[node.Round]
				if !ok {
					round := Round{
						Name: node.FullRoundText,
						Sets: []Set{
							set,
						},
					}
					winners[node.Round] = round
				} else {
					value.Sets = append(value.Sets, set)
					winners[node.Round] = value
				}
			} else {
				// save set to losers bracket
				value, ok := losers[node.Round]
				if !ok {
					round := Round{
						Name: node.FullRoundText,
						Sets: []Set{
							set,
						},
					}
					losers[node.Round] = round
					log.Printf("CREATE | %v !! %v", node.Round, set)
				} else {
					log.Printf("APPEND | %v -> %v", node.Round, set)
					value.Sets = append(value.Sets, set)
					losers[node.Round] = value
					log.Printf("VIEW | %v", value.Sets)
				}
			}
		}
	}
	result = append(result, winners)
	result = append(result, losers)
	return result, nil
}

func (v *Viewer) ProcessUpdate() error {
	brackets, err := v.prepareData()
	if err != nil {
		return err
	}
	// map -> slice for html
	winners := []Round{}
	for _, value := range brackets[0] {
		winners = append(winners, value)
	}
	losers := []Round{}
	for _, value := range brackets[1] {
		losers = append(losers, value)
	}

	v.Winners = winners
	v.Losers = losers
	return nil
}

func (v *Viewer) Run() error {
	brackets, err := v.prepareData()
	if err != nil {
		return err
	}
	// map -> slice for html
	winners := []Round{}
	for _, value := range brackets[0] {
		winners = append(winners, value)
	}
	losers := []Round{}
	for _, value := range brackets[1] {
		losers = append(losers, value)
	}

	v.Winners = winners
	v.Losers = losers

	r := gin.Default()
	r.LoadHTMLGlob("internal/templates/*")
	r.Static("/config", "config")
	r.GET("/settings", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/top8", func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			"top8.html",
			gin.H{
				"title":   "Top8",
				"winners": v.Winners,
				"losers":  v.Losers,
			},
		)
	})
	return r.Run(":7777")
}
