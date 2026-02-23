package handler

import (
	"ayo/entity"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) CreateMatch(c *gin.Context) {

	var req struct {
		MatchAt    time.Time `json:"match_at" binding:"required"` // format: YYYY-MM-DD
		HomeTeamID uuid.UUID `json:"home_team_id" binding:"required,uuid"`
		AwayTeamID uuid.UUID `json:"away_team_id" binding:"required,uuid"`
		Status     string    `json:"status"`
	}

	ctx := c.Request.Context()

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Get request failed",
			"data":    gin.H{},
		})
	}

	// parse time

	reqEntity := &entity.Match{
		MatchAt:    req.MatchAt,
		HomeTeamID: req.HomeTeamID,
		AwayTeamID: req.AwayTeamID,
		Status:     req.Status,
	}

	result, err := h.MM.CreatMatch(ctx, reqEntity)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Match process failed",
			"data":    gin.H{},
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Match process successful",
		"data": gin.H{
			"id":           result.ID,
			"match_at":     result.MatchAt,
			"home_team_id": result.HomeTeamID,
			"away_team_id": result.AwayTeamID,
			"status":       result.Status,
		},
	})
}

func (h *Handler) GetMatch(c *gin.Context) {

	ctx := c.Request.Context()

	id := c.Param("id")

	result, err := h.MM.GetMatchByID(ctx, id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Match process failed",
			"data":    gin.H{},
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Match process successful",
		"data": gin.H{
			"id":           result.ID,
			"match_at":     result.MatchAt,
			"home_team_id": result.HomeTeamID,
			"away_team_id": result.AwayTeamID,
			"status":       result.Status,
		},
	})
}

func (h *Handler) UpdateMatch(c *gin.Context) {

	var req struct {
		Status string `json:"status"`
	}

	ctx := c.Request.Context()

	id := c.Param("id")

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Get request failed",
			"data":    gin.H{},
		})
	}

	reqEntity := &entity.Match{
		Status: req.Status,
	}

	result, err := h.MM.UpdateMatch(ctx, id, reqEntity)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Match process failed",
			"data":    gin.H{},
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Match process successful",
		"data": gin.H{
			"id":           result.ID,
			"match_at":     result.MatchAt,
			"home_team_id": result.HomeTeamID,
			"away_team_id": result.AwayTeamID,
			"status":       result.Status,
		},
	})
}

func (h *Handler) DeleteMatch(c *gin.Context) {

	ctx := c.Request.Context()

	id := c.Param("id")

	err := h.MM.DeleteMatch(ctx, id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Match deletion process failed",
			"data":    gin.H{},
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Match deletion process successful",
		"data":    gin.H{},
	})
}
