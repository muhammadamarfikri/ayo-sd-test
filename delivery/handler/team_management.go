package handler

import (
	"ayo/entity"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateTeam(c *gin.Context) {

	var req struct {
		Name                string `json:"name"`
		Logo                string `json:"logo"`
		FoundedYear         int    `json:"founded_year"`
		HeadquartersAddress string `json:"headquarter_address"`
		HeadquartersCity    string `json:"headquarter_city"`
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

	reqEntity := &entity.Team{
		Name:                req.Name,
		LogoURL:             &req.Logo,
		FoundedYear:         req.FoundedYear,
		HeadquartersAddress: req.HeadquartersAddress,
		HeadquartersCity:    req.HeadquartersCity,
	}

	result, err := h.TM.CreatTeam(ctx, reqEntity)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Team process failed",
			"data": gin.H{
				"id":   result.ID,
				"name": result.Name,
			},
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Team process successful",
		"data": gin.H{
			"id":   result.ID,
			"name": result.Name,
		},
	})
}

func (h *Handler) GetTeam(c *gin.Context) {

	ctx := c.Request.Context()

	id := c.Param("id")

	result, err := h.TM.GetTeam(ctx, id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Team process failed",
			"data": gin.H{
				"id":   result.ID,
				"name": result.Name,
			},
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Team process successful",
		"data": gin.H{
			"id":                  result.ID,
			"name":                result.Name,
			"founded_year":        result.FoundedYear,
			"headquarter_address": result.HeadquartersAddress,
			"headquarter_city":    result.HeadquartersCity,
		},
	})
}

func (h *Handler) UpdateTeam(c *gin.Context) {

	var req struct {
		Name                string `json:"name"`
		Logo                string `json:"logo"`
		FoundedYear         int    `json:"founded_year"`
		HeadquartersAddress string `json:"headquarter_address"`
		HeadquartersCity    string `json:"headquarter_city"`
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

	reqEntity := &entity.Team{
		Name:                req.Name,
		LogoURL:             &req.Logo,
		FoundedYear:         req.FoundedYear,
		HeadquartersAddress: req.HeadquartersAddress,
		HeadquartersCity:    req.HeadquartersCity,
	}

	result, err := h.TM.UpdateTeam(ctx, id, reqEntity)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Team process failed",
			"data": gin.H{
				"id":   result.ID,
				"name": result.Name,
			},
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Team process successful",
		"data": gin.H{
			"id":                  result.ID,
			"name":                result.Name,
			"founded_year":        result.FoundedYear,
			"headquarter_address": result.HeadquartersAddress,
			"headquarter_city":    result.HeadquartersCity,
		},
	})
}

func (h *Handler) DeleteTeam(c *gin.Context) {

	ctx := c.Request.Context()

	id := c.Param("id")

	err := h.TM.DeleteTeam(ctx, id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Team deletion process failed",
			"data":    gin.H{},
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Team deletion process successful",
		"data":    gin.H{},
	})
}
