package api

import (
	"giligili/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

//CreateVideo 创建视频
func CreateVideo(c *gin.Context) {
	s := service.CreateVideoService{}
	if err := c.ShouldBind(&s); err == nil {
		res := s.Create()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

func ShowVideo(c *gin.Context) {
	s := service.ShowVideoService{}
	res := s.Show(c.Param("id"))
	c.JSON(http.StatusOK, res)
}

func ListVideo(c *gin.Context) {
	s := service.ListVideoService{}
	if err := c.ShouldBind(&s); err == nil {
		res := s.List()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

func UpdateVideo(c *gin.Context) {
	s := service.UpdateVideoService{}
	if err := c.ShouldBind(&s); err == nil {
		res := s.UpdateVideo(c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

func DeleteVideo(c *gin.Context) {
	s := service.DeleteVideoService{}
	if err := c.ShouldBind(&s); err == nil {
		res := s.Delete(c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}
