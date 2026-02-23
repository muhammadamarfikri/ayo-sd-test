package handler

import (
	"ayo/entity"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) CreatePlayer(c *gin.Context) {

	var req struct {
		TeamID      uuid.UUID `json:"team_id" db:"team_id"`
		Name        string    `json:"name" db:"name"`
		HeightCM    int       `json:"height_cm" db:"height_cm"`
		WeightKG    int       `json:"weight_kg" db:"weight_kg"`
		Position    string    `json:"position" db:"position"`
		ShirtNumber int       `json:"shirt_number" db:"shirt_number"`
	}

	ctx := c.Request.Context()

	fmt.Println("kocak")

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Get request failed",
			"data":    gin.H{},
		})
	}

	reqEntity := &entity.Player{
		Name:        req.Name,
		TeamID:      req.TeamID,
		HeightCM:    req.HeightCM,
		WeightKG:    req.WeightKG,
		Position:    req.Position,
		ShirtNumber: req.ShirtNumber,
	}

	result, err := h.PM.CreatPlayer(ctx, reqEntity)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Player process failed",
			"data": gin.H{
				"id":   result.ID,
				"name": result.Name,
			},
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Player process successful",
		"data": gin.H{
			"id":   result.ID,
			"name": result.Name,
		},
	})
}

func (h *Handler) GetPlayer(c *gin.Context) {

	ctx := c.Request.Context()

	id := c.Param("id")

	result, err := h.PM.GetPlayer(ctx, id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Player process failed",
			"data": gin.H{
				"id":   result.ID,
				"name": result.Name,
			},
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Player process successful",
		"data": gin.H{
			"name":         result.Name,
			"team_id":      result.TeamID,
			"height_cm":    result.HeightCM,
			"weight_kg":    result.WeightKG,
			"position":     result.Position,
			"shirt_number": result.ShirtNumber,
		},
	})
}

func (h *Handler) UpdatePlayer(c *gin.Context) {

	var req struct {
		TeamID      uuid.UUID `json:"team_id" db:"team_id"`
		Name        string    `json:"name" db:"name"`
		HeightCM    int       `json:"height_cm" db:"height_cm"`
		WeightKG    int       `json:"weight_kg" db:"weight_kg"`
		Position    string    `json:"position" db:"position"`
		ShirtNumber int       `json:"shirt_number" db:"shirt_number"`
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

	reqEntity := &entity.Player{
		Name:        req.Name,
		TeamID:      req.TeamID,
		HeightCM:    req.HeightCM,
		WeightKG:    req.WeightKG,
		Position:    req.Position,
		ShirtNumber: req.ShirtNumber,
	}

	result, err := h.PM.UpdatePlayer(ctx, id, reqEntity)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Player process failed",
			"data": gin.H{
				"id":   result.ID,
				"name": result.Name,
			},
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Player process successful",
		"data": gin.H{
			"name":         result.Name,
			"team_id":      result.TeamID,
			"height_cm":    result.HeightCM,
			"weight_kg":    result.WeightKG,
			"position":     result.Position,
			"shirt_number": result.ShirtNumber,
		},
	})
}

func (h *Handler) DeletePlayer(c *gin.Context) {

	ctx := c.Request.Context()

	id := c.Param("id")

	err := h.PM.DeletePlayer(ctx, id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Player deletion process failed",
			"data":    gin.H{},
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Player deletion process successful",
		"data":    gin.H{},
	})
}
