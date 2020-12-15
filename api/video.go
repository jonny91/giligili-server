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
	if err := c.ShouldBind(&s); err == nil {
		res := s.Show(c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}
