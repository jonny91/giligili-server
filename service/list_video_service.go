package service

import (
	"giligili/model"
	"giligili/serializer"
)

//视频列表服务
type ListVideoService struct {
}

func (s *ListVideoService) List() serializer.Response {
	var videos []model.Video
	err := model.DB.Find(&videos).Error
	if err != nil {
		return serializer.Response{
			Code:  50000,
			Msg:   "查询错误",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildVideos(videos),
	}
}
