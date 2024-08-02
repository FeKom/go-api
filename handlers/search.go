package handlers

import (
	"net/http"

	service "github.com/fekom/go-api/services"
	"github.com/gin-gonic/gin"
)

func GetValues(c *gin.Context) {
	SearchTeam := c.Param("search_team")
	if SearchTeam == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Search team parameter is required",
		})
		return
	}

	team := service.SearchByName(SearchTeam)
	// errorMessage := team.errorMessage

	if team.Response.Data != nil {
		// dataId := team.response.Data[0].ID
		// schedule := service.GetScheduleByTeamId(dataId)
		c.JSON(http.StatusOK, team.Response)
		return
	}
	c.JSON(404, gin.H{
		"message": "Search team parameter is required",
	})
}

// func SearchByName(SearchTeam string) {
// 	panic("unimplemented")
// }
