package usecase

import "content-service-v3/services/content-service/domain/repository"

type serviceBanner struct {
	repository repository.RepositoryBanner
	repositoryBannerCategory repository.RepositoryBannerCategory

}

func NewServiceBanner(repository repository.RepositoryBanner, repositoryBannerCategory repository.RepositoryBannerCategory) *serviceBanner {
	return &serviceBanner{repository,repositoryBannerCategory}
}
