package service

import (
	"giligili/model"
	"giligili/serializer"
)

// 视频投稿服务
type CreateVideoService struct {
	Title string `json:"title" form:"title" binding:"required,min=2,max=30"`
	Info  string `json:"info" form:"info" binding:"required,min=0,max=400"`
}

func (s *CreateVideoService) Create() serializer.Response {
	video := model.Video{
		Title: s.Title,
		Info:  s.Info,
	}
	err := model.DB.Create(&video).Error
	if err != nil {
		return serializer.Response{
			Code:  50001,
			Msg:   "视频保存失败",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildVideo(video),
	}
}
