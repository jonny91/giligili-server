package service

import (
	"giligili/model"
	"giligili/serializer"
	"net/http"
)

type UpdateVideoService struct {
	CreateVideoService
}

func (s *UpdateVideoService) UpdateVideo(id string) serializer.Response {
	var video model.Video
	err := model.DB.First(&video, id).Error
	if err != nil {
		return serializer.Response{
			Code:  http.StatusNotFound,
			Msg:   "视频不存在",
			Error: err.Error(),
		}
	}

	video.Title = s.Title
	video.Info = s.Info
	err = model.DB.Save(&video).Error
	if err != nil {
		return serializer.Response{
			Code:  50002,
			Msg:   "更新视频失败",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildVideo(video),
	}
}
