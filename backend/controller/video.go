package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/virusman1/sa-64-example-typescript/entity"
	"net/http"
)

// POST /videos
func CreateVideo(c *gin.Context) {
	var video entity.Video
	if err := c.ShouldBindJSON(&video); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&video).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": video})
}

// GET /video/:id
func GetVideo(c *gin.Context) {
	var video entity.Video

	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM videos WHERE id = ?", id).Scan(&video).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Raw("SELECT id,first_name,last_name,email,age,birth_day FROM users WHERE id = ?", video.OwnerID).Scan(&video.Owner).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Raw("SELECT * FROM watch_videos WHERE id = ?", video.ID).Scan(&video.WatchVideos).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": video})
}

// GET /videos
func ListVideos(c *gin.Context) {
	var videos []entity.Video
	if err := entity.DB().Raw("SELECT * FROM videos").Scan(&videos).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for index, video := range videos {
		if err := entity.DB().Raw("SELECT id,first_name,last_name,email,age,birth_day FROM users WHERE id = ?", video.OwnerID).Scan(&videos[index].Owner).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := entity.DB().Raw("SELECT * FROM watch_videos WHERE video_id = ?", video.ID).Scan(&videos[index].WatchVideos).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": videos})
}

// DELETE /videos/:id
func DeleteVideo(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM videos WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "video not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /videos
func UpdateVideo(c *gin.Context) {
	var video entity.Video
	if err := c.ShouldBindJSON(&video); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", video.ID).First(&video); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "video not found"})
		return
	}

	if err := entity.DB().Save(&video).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": video})
}
