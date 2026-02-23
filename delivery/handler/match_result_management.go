package handler

import (
	"ayo/entity"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) CreateMatchResult(c *gin.Context) {

	var req struct {
		MatchID   uuid.UUID `json:"match_id" binding:"required,uuid"`
		HomeScore int       `json:"home_score" binding:"required"`
		AwayScore int       `json:"away_score" binding:"required"`
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

	reqEntity := &entity.MatchResult{
		MatchID:   req.MatchID,
		HomeScore: req.HomeScore,
		AwayScore: req.AwayScore,
	}

	result, err := h.MRM.CreatMatchResult(ctx, reqEntity)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Match process failed",
			"data":    gin.H{},
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Match process successful",
		"data": gin.H{
			"id":          result.MatchID,
			"home_score":  result.HomeScore,
			"away_score":  result.AwayScore,
			"reported_at": result.ReportedAt,
		},
	})
}

func (h *Handler) GetMatchResult(c *gin.Context) {

	ctx := c.Request.Context()

	id := c.Param("id")

	result, err := h.MRM.GetMatchResultByID(ctx, id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Match process failed",
			"data":    gin.H{},
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Match process successful",
		"data": gin.H{
			"id":          result.MatchID,
			"home_score":  result.HomeScore,
			"away_score":  result.AwayScore,
			"reported_at": result.ReportedAt,
		},
	})
}

func (h *Handler) UpdateMatchResult(c *gin.Context) {

	var req struct {
		MatchID   uuid.UUID `json:"match_id" binding:"required,uuid"`
		HomeScore int       `json:"home_score" binding:"required"`
		AwayScore int       `json:"away_score" binding:"required"`
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

	reqEntity := &entity.MatchResult{
		MatchID:   req.MatchID,
		HomeScore: req.HomeScore,
		AwayScore: req.AwayScore,
	}

	result, err := h.MRM.UpdateMatchResult(ctx, id, reqEntity)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Match process failed",
			"data":    gin.H{},
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Match process successful",
		"data": gin.H{
			"id":          result.MatchID,
			"home_score":  result.HomeScore,
			"away_score":  result.AwayScore,
			"reported_at": result.ReportedAt,
		},
	})
}

func (h *Handler) DeleteMatchResult(c *gin.Context) {

	ctx := c.Request.Context()

	id := c.Param("id")

	err := h.MRM.DeleteMatchResult(ctx, id)
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
