package handler

import (
	"github.com/gin-gonic/gin"
	"otus/internal/controller"
	"otus/internal/model/catalog"
)

type EpisodeItemResponse struct {
	Result struct {
		Item catalog.Episode
	}
}
type EpisodeItemsResponse struct {
	Result struct {
		Items []catalog.Episode
	}
}

// GetEpisode godoc
// @Summary Get episode by ID
// @Description Get detailed information about a TV episode
// @Tags episodes
// @Accept  json
// @Produce  json
// @Param id path int true "Episode ID"
// @Success 200 {object} EpisodeItemResponse "Successfully retrieved episode"
// @Failure 400 {object} ErrorResponse "Not found"
// @Router /otus.episode.get/{id} [get]
func GetEpisode(c *gin.Context) {
	controller.GetAction[*catalog.Episode](c)
}

// GetListEpisode godoc
// @Summary Get episodes
// @Description Get list information about TV episode
// @Tags episodes
// @Accept  json
// @Produce  json
// @Success 200 {object} EpisodeItemsResponse "Successfully retrieved episode"
// @Router /otus.episode.list [get]
func GetListEpisode(c *gin.Context) {
	controller.GetListAction[*catalog.Episode](c)
}

// AddEpisode godoc
// @Summary Create new TV episode
// @Description Add a new episode to the catalog
// @Tags episodes
// @Accept  json
// @Produce  json
// @Param episode body catalog.Episode true "Episode data"
// @Success 200 {object} EpisodeItemResponse
// @Security ApiKeyAuth
// @Router /otus.episode.add [post]
func AddEpisode(c *gin.Context) {
	controller.AddAction[*catalog.Episode](c)
}

// UpdateEpisode godoc
// @Summary Update episode
// @Description Update existing TV episode
// @Tags episodes
// @Accept  json
// @Produce  json
// @Param id path int true "Episode ID"
// @Param episode body catalog.Episode true "Update data"
// @Success 200 {object} EpisodeItemResponse
// @Failure 400 {object} ErrorResponse "Not found"
// @Security ApiKeyAuth
// @Router /otus.episode.update/{id} [put]
func UpdateEpisode(c *gin.Context) {
	controller.UpdateAction[*catalog.Episode](c)
}

// DeleteEpisode godoc
// @Summary Delete episode
// @Description Delete a TV episode from catalog
// @Tags episodes
// @Accept  json
// @Produce  json
// @Param id path int true "Episode ID"
// @Success 200 {object} DeleteResponse
// @Failure 400 {object} ErrorResponse "Not found"
// @Security ApiKeyAuth
// @Router /otus.episode.delete/{id} [delete]
func DeleteEpisode(c *gin.Context) {
	controller.DeleteAction[*catalog.Episode](c)
}
