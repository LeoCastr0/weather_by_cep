package service

import (
	"github.com/leocastr0/weather_by_cep/internal/infra/repository"

	zipcode "github.com/leocastr0/weather_by_cep/internal/domain"
)

//go:generate mockery --name ZipCodeService --outpkg mock --output mock --filename zipcode.go --with-expecter=true

type ZipCodeService interface {
	GetLocationByZipCode(zipCode string) (*zipcode.Location, error)
}

type zipCodeService struct {
	repository repository.ZipCodeRepository
}

func NewZipCodeService(repo repository.ZipCodeRepository) ZipCodeService {
	return &zipCodeService{
		repository: repo,
	}
}

func (s *zipCodeService) GetLocationByZipCode(zipCode string) (*zipcode.Location, error) {
	return s.repository.GetLocationByZipCode(zipCode)
}
