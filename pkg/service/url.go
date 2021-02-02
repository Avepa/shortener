package service

import (
	"github.com/Avepa/shortener/pkg"
	"github.com/Avepa/shortener/pkg/repository"
	"github.com/decred/base58"
	"golang.org/x/crypto/sha3"
)

type URLService struct {
	repo repository.URL
}

func NewURLService(repo repository.URL) *URLService {
	return &URLService{repo: repo}
}

func (s *URLService) Add(url, path string) (string, error) {
	if path == "" {
		path = hash(url, 5)
		err := s.repo.Add(url, path)
		if err != nil {
			if err == pkg.ErrorPathIsBusy {
				u, _ := s.repo.Get(path)

				if u != url {
					for i := int8(6); err == pkg.ErrorPathIsBusy && i < 10; i++ {
						path = hash(url, i)
						err = s.repo.Add(url, path)

						if err == pkg.ErrorPathIsBusy {
							u, _ = s.repo.Get(path)

							if u == url {
								return path, nil
							}
						}
					}
				}
				return path, nil
			}
			return path, err
		}
		return path, nil

	} else {
		return path, s.repo.Add(url, path)
	}
}

func (s *URLService) Get(path string) (string, error) {
	return s.repo.Get(path)
}

func hash(url string, Length int8) string {
	hash := make([]byte, Length)
	sha3.ShakeSum128(hash, []byte(url))
	return base58.Encode(hash)
}
