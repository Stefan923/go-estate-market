package service

import (
	"github.com/Stefan923/go-estate-market/api/dto"
	"github.com/Stefan923/go-estate-market/data/model"
	"github.com/Stefan923/go-estate-market/data/repository"
)

type PostService struct {
	BaseService[model.Post, dto.PostCreationWithIdDto, dto.PostCreationDto, dto.PostDto]
}

func NewPostService() *PostService {
	return &PostService{
		BaseService: BaseService[model.Post, dto.PostCreationWithIdDto, dto.PostCreationDto, dto.PostDto]{
			Repository: repository.NewBaseRepository[model.Post]([]repository.PreloadSetting{}),
		},
	}
}
