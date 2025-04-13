package anime

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func getSchedule(startOfDay int64, endOfDay int64) *GraphQLResponse {

}