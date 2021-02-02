package service

import "github.com/Avepa/shortener/pkg/repository"

type URL interface {
	Add(url, path string) (string, error)
	Get(path string) (string, error)
}

type Service struct {
	URL
}

func NewSrvice(repos *repository.Repository) *Service {
	return &Service{
		URL: NewURLService(repos.URL),
	}
}
