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

func GetValues(c *gin.Context) {
	SearchTeam := c.Param("search_team")
	if SearchTeam == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Search team parameter is required",
		})
		return
	}

	ApiURL := ApiBaseUrl + SearchTeam + "?api_token=" + ApiToken

	Resp, err := http.Get(ApiURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to fetch data from external API",
		})
		return
	}
	defer Resp.Body.Close()

	if Resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "Received non-200 response from external API",
		})
		return
	}

	data, err := io.ReadAll(Resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to read response body",
		})
		return
	}
	var Target map[string]interface{}
	err = json.Unmarshal(data, &Target)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
	}
	c.JSON(http.StatusOK, Target)

}
