package viewer

import (
	"encoding/json"
	"fmt"

	"github.com/dreamervulpi/tourneyBot/startgg"
)

type RawBracketData struct {
	Data   DataPhaseGroup   `json:"data"`
	Errors []startgg.Errors `json:"errors"`
}

type DataPhaseGroup struct {
	PhaseGroup PhaseGroup `json:"phaseGroup"`
}

// Group(Phase)
type PhaseGroup struct {
	Id   int64 `json:"id"`
	Sets Sets  `json:"sets"`
}

type Sets struct {
	PageInfo PageInfo `json:"pageInfo"`
	Nodes    []Nodes  `json:"nodes"`
}

// Sets counts in Group(Phase)
type PageInfo struct {
	Total int `json:"total"`
}

// Information about Set
type Nodes struct {
	Round         int           `json:"round"`
	FullRoundText string        `json:"fullRoundText"`
	State         startgg.State `json:"state"`
	DisplayScore  string        `json:"displayScore"`
	Slots         []Slots       `json:"slots"`
}

// Slots in set
type Slots struct {
	Entrant Entrant `json:"entrant"`
}

// Player in tournament
type Entrant struct {
	Id           int64          `json:"id"`
	Name         string         `json:"name"`
	Participants []Participants `json:"participants"`
}

type Participants struct {
	GamerTag string `json:"gamerTag"`
}

const getSets = `
	query getSets($phaseGroupId: ID!, $page:Int!, $perPage:Int!){
		phaseGroup(id:$phaseGroupId){
			id
			sets(
				page: $page
				perPage: $perPage
				sortType: STANDARD
				filters: {state: 3}
			){
				pageInfo{
					total
				}
				nodes{
					round
					state
					fullRoundText
					displayScore
					slots{
						entrant{
							id
							name
							participants {
								gamerTag
							}
						}
					}
				}
			}
		}
	}`

func (v *Viewer) getBracket(page int, perPage int) ([]Nodes, error) {
	var variables = map[string]any{
		"phaseGroupId": v.PhaseGroupId,
		"page":         page,
		"perPage":      perPage,
	}

	query, err := json.Marshal(startgg.PrepareQuery(getSets, variables))
	if err != nil {
		return []Nodes{}, fmt.Errorf("JSON Marshal - %w", err)
	}

	data, err := v.Client.RunQuery(query)
	if err != nil {
		return []Nodes{}, err
	}

	results := &RawBracketData{}
	err = json.Unmarshal(data, results)
	if err != nil {
		return nil, fmt.Errorf("JSON Unmarshal - %w", err)
	}
	return results.Data.PhaseGroup.Sets.Nodes, nil
}
