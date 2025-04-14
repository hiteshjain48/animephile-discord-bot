package anime

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type GraphQLRequest struct {
	Query string 						`json:"query"`
	Variables map[string]interface{} 	`json:"variables"`
}

type Media struct {
	Title struct {
		Romaji string `json:"romaji"`
	} `json:"title"`
}

type AiringSchedule struct {
	AiringAt 	int64 	`json:"airingAt`
	Episode 	int 	`json:"episode"`
	Media 		Media 	`json:"media"`
}

type Page struct {
	AiringSchedule []AiringSchedule `json:"airingSchedules"`
}

type Data struct {
	Page Page `json:"page"`
}

type GraphQLResponse struct {
	Data Data `json:"data"`
}

func fetchAiringAll(startOfDay int64, endOfDay int64) ([]AiringSchedule, error) {
	perPage := 50

	allSchedules := []AiringSchedule{}
	page := 1

	for {
		schedules, err := fetchSchedule(page, perPage, startOfDay, endOfDay)
		if err != nil {
			return nil, err
		}

		if len(schedules) == 0 {
			break
		}

		allSchedules = append(allSchedules, schedules...)
		page++
	}
	return allSchedules, nil
}

func fetchSchedule(page int, perPage int, startOfDay int64, endOfDay int64) ([]AiringSchedule, error) {
	query := `
	query ($page: Int, $perPage: Int, $start: Int, $end: Int) {
	  Page(page: $page, perPage: $perPage) {
	    airingSchedules(airingAt_greater: $start, airingAt_lesser: $end) {
	      airingAt
	      episode
	      media {
	        title {
	          romaji
	        }
	      }
	    }
	  }
	}`

	variables := map[string]interface{}{
		"page": 	page,
		"perPage": 	perPage,
		"start": 	startOfDay,
		"end": 		endOfDay,
	}

	reqBody := GraphQLRequest{
		Query: 		query,
		Variables: 	variables,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post("https://graphql.anilist.co", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var response GraphQLResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response.Data.Page.AiringSchedule, nil
}

func GetSchedule() ([]AiringSchedule, error) {
	startOfDay := time.Now().UTC().Truncate(24 * time.Hour).Unix()
	endOfDay := startOfDay + 86399

	schedules, err := fetchAiringAll(startOfDay, endOfDay)
	if err != nil {
		return nil, err
	}
	return schedules, nil
}

