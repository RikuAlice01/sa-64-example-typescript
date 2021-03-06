package controller
 
import (
	"github.com/virusman1/sa-64-example-typescript/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

// POST /watchvideos
func CreateWatchVideo(c *gin.Context) {
	var watchvideo entity.WatchVideo
	if err := c.ShouldBindJSON(&watchvideo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
 
	if err := entity.DB().Create(&watchvideo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": watchvideo})
}

// GET /watchvideo/:id
func GetWatchVideo(c *gin.Context) {
	var watchvideo entity.WatchVideo
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM watch_videos WHERE id = ?", id).Scan(&watchvideo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Raw("SELECT * FROM resolutions WHERE id = ?", watchvideo.ResolutionID).Scan(&watchvideo.Resolution).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Raw("SELECT * FROM videos WHERE id = ?", watchvideo.PlaylistID).Scan(&watchvideo.Playlist).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Raw("SELECT * FROM playlists WHERE id = ?", watchvideo.VideoID).Scan(&watchvideo.Video).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}	
 
	c.JSON(http.StatusOK, gin.H{"data": watchvideo})
}
 
// GET /watchvideos
func ListWatchVideos(c *gin.Context) {
	var watchvideos []entity.WatchVideo
	if err := entity.DB().Raw("SELECT * FROM watch_videos").Scan(&watchvideos).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for index, watchvideo := range watchvideos {


		if err := entity.DB().Raw("SELECT * FROM resolutions WHERE id = ?", watchvideo.ResolutionID).Scan(&watchvideos[index].Resolution).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	
		if err := entity.DB().Raw("SELECT * FROM videos WHERE id = ?", watchvideo.PlaylistID).Scan(&watchvideos[index].Playlist).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	
		if err := entity.DB().Raw("SELECT * FROM playlists WHERE id = ?", watchvideo.VideoID).Scan(&watchvideos[index].Video).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}	

	}
 
	c.JSON(http.StatusOK, gin.H{"data": watchvideos})
}

// DELETE /watchvideos/:id
func DeleteWatchVideo(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM watch_videos WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "watchvideo not found"})
		return
	}
 
	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /watchvideos
func UpdateWatchVideo(c *gin.Context) {
	var watchvideo entity.WatchVideo
	if err := c.ShouldBindJSON(&watchvideo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
 
	if tx := entity.DB().Where("id = ?", watchvideo.ID).First(&watchvideo); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "watchvideo not found"})
		return
	}
 
	if err := entity.DB().Save(&watchvideo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
 
	c.JSON(http.StatusOK, gin.H{"data": watchvideo})
}

