package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	ApiBaseUrl = "https://api.sportmonks.com/v3/football/teams/search/"
	ApiToken   = "05E5yYwY1vYsAuQgt6ZX2Oz3AVnhA8gB8EstLL8fsYVuwCIbvNertUhnjp8m"
)

type SearchApiResponse struct {
	Data []struct {
		ID           int    `json:"id"`
		SportID      int    `json:"sport_id"`
		CountryID    int    `json:"country_id"`
		VenueID      int    `json:"venue_id"`
		Gender       string `json:"gender"`
		Name         string `json:"name"`
		ShortCode    string `json:"short_code"`
		ImagePath    string `json:"image_path"`
		Founded      int    `json:"founded"`
		Type         string `json:"type"`
		Placeholder  bool   `json:"placeholder"`
		LastPlayedAt string `json:"last_played_at"`
	} `json:"data"`
	Pagination struct {
		Count       int  `json:"count"`
		PerPage     int  `json:"per_page"`
		CurrentPage int  `json:"current_page"`
		NextPage    any  `json:"next_page"`
		HasMore     bool `json:"has_more"`
	} `json:"pagination"`
	Subscription []struct {
		Meta  []any `json:"meta"`
		Plans []struct {
			Plan     string `json:"plan"`
			Sport    string `json:"sport"`
			Category string `json:"category"`
		} `json:"plans"`
		AddOns  []any `json:"add_ons"`
		Widgets []any `json:"widgets"`
	} `json:"subscription"`
	RateLimit struct {
		ResetsInSeconds int    `json:"resets_in_seconds"`
		Remaining       int    `json:"remaining"`
		RequestedEntity string `json:"requested_entity"`
	} `json:"rate_limit"`
	Timezone string `json:"timezone"`
}
type SearchResponse struct {
	errorMessage string
	response     SearchApiResponse
}

func searchByName(name string) SearchResponse {
	ApiURL := ApiBaseUrl + name + "?api_token=" + ApiToken
	Resp, err := http.Get(ApiURL)
	if err != nil {
		return SearchResponse{errorMessage: "failed to fetch data from external API", response: SearchApiResponse{}}
	}
	defer Resp.Body.Close()
	if Resp.StatusCode != http.StatusOK {
		return SearchResponse{errorMessage: "Received non-200 response from external API", response: SearchApiResponse{}}
	}
	data, err := io.ReadAll(Resp.Body)
	if err != nil {
		return SearchResponse{errorMessage: "Failed to read response body", response: SearchApiResponse{}}
	}
	var Target SearchApiResponse
	err = json.Unmarshal(data, &Target)
	if err != nil {
		return SearchResponse{errorMessage: "Failed to parse response", response: SearchApiResponse{}}
	}
	return SearchResponse{errorMessage: "", response: Target}
}

func GetValues(c *gin.Context) {
	SearchTeam := c.Param("search_team")
	if SearchTeam == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Search team parameter is required",
		})
		return
	}

	team := searchByName(SearchTeam)
	// errorMessage := team.errorMessage

	if team.response.Data != nil {
		// dataId := team.response.Data[0].ID
		// schedule := service.GetScheduleByTeamId(dataId)
		c.JSON(http.StatusOK, team.response)
		return
	}
	c.JSON(404, gin.H{
		"message": "Search team parameter is required",
	})
}
