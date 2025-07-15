package handler

import (
	"github.com/gin-gonic/gin"
	"otus/internal/controller"
	"otus/internal/model/catalog"
)

type SeasonItemResponse struct {
	Result struct {
		Item catalog.Season
	}
}
type SeasonItemsResponse struct {
	Result struct {
		Items []catalog.Season
	}
}

// GetSeason godoc
// @Summary Get season by ID
// @Description Get detailed information about a TV season
// @Tags seasons
// @Accept  json
// @Produce  json
// @Param id path int true "Season ID"
// @Success 200 {object} SeasonItemResponse "Successfully retrieved season"
// @Failure 400 {object} ErrorResponse "Not found"
// @Router /otus.season.get/{id} [get]
func GetSeason(c *gin.Context) {
	controller.GetAction[*catalog.Season](c)
}

// GetListSeason godoc
// @Summary Get seasons
// @Description Get list information about TV seasons
// @Tags seasons
// @Accept  json
// @Produce  json
// @Success 200 {object} SeasonItemsResponse "Successfully retrieved season"
// @Router /otus.season.list [get]
func GetListSeason(c *gin.Context) {
	controller.GetListAction[*catalog.Season](c)
}

// AddSeason godoc
// @Summary Create new TV season
// @Description Add a new season to the catalog
// @Tags seasons
// @Accept  json
// @Produce  json
// @Param season body catalog.Season true "Season data"
// @Success 200 {object} SeasonItemResponse
// @Security ApiKeyAuth
// @Router /otus.season.add [post]
func AddSeason(c *gin.Context) {
	controller.AddAction[*catalog.Season](c)
}

// UpdateSeason godoc
// @Summary Update season
// @Description Update existing TV season
// @Tags seasons
// @Accept  json
// @Produce  json
// @Param id path int true "Season ID"
// @Param season body catalog.Season true "Update data"
// @Success 200 {object} SeasonItemResponse
// @Failure 400 {object} ErrorResponse "Not found"
// @Security ApiKeyAuth
// @Router /otus.season.update/{id} [put]
func UpdateSeason(c *gin.Context) {
	controller.UpdateAction[*catalog.Season](c)
}

// DeleteSeason godoc
// @Summary Delete season
// @Description Delete a TV season from catalog
// @Tags seasons
// @Accept  json
// @Produce  json
// @Param id path int true "Season ID"
// @Success 200 {object} DeleteResponse
// @Failure 400 {object} ErrorResponse "Not found"
// @Security ApiKeyAuth
// @Router /otus.season.delete/{id} [delete]
func DeleteSeason(c *gin.Context) {
	controller.DeleteAction[*catalog.Season](c)
}
