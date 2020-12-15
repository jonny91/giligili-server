package api

import (
	"giligili/serializer"
	"net/http"

	"github.com/gin-gonic/gin"
)

//CreateVideo 创建视频
func CreateVideo(c *gin.Context) {
	c.JSON(http.StatusOK, serializer.Response{
		Code: 200,
		Msg:  "OK",
	})
}
