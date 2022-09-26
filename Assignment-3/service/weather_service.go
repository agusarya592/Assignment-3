package service

import (
	"assignment3/entity"
	"assignment3/repository"
	"assignment3/utils"
	"context"
	"database/sql"
	"log"
	"strconv"
	"time"
)

type WeatherService interface {
	UpdateValue(ctx context.Context) error
	GetUpdate(ctx context.Context) (*entity.WeatherResponse, error)
}

func ProvideWeatherService(DB *sql.DB) WeatherService {
	repo := repository.ProvideWeatherRepository(DB)
	return ProvideWeather(repo)
}

type weatherServiceImpl struct {
	repo repository.WeatherRepository
}

func ProvideWeather(repo repository.WeatherRepository) *weatherServiceImpl {
	return &weatherServiceImpl{repo: repo}
}

func (w weatherServiceImpl) GetUpdate(ctx context.Context) (*entity.WeatherResponse, error) {
	res, err := w.repo.GetUpdate(ctx)
	if err != nil {
		log.Printf("FAILED! to get the weather update")
		return nil, err
	}
	return entity.CreateWeatherResponse(res), nil
}

func (w weatherServiceImpl) UpdateValue(ctx context.Context) error {
	ticker := time.NewTicker(900 * time.Second)
	quit := make(chan int)

	go func() {
		for {
			select {
			case <-ticker.C:
				windValue := utils.RandWindValue()
				waterValue := utils.RandWaterValue()
				windString := strconv.Itoa(int(windValue)) + " m/s"
				waterString := strconv.Itoa(int(waterValue)) + " m"
				update := utils.UpdateWeather(windValue, waterValue)

				err := w.repo.UpdateValue(windString, waterString, update)
				log.Printf("STATUS UPDATED SUCCESSFULLY!")
				if err != nil {
					log.Printf("STATUS FAILED TO UPDATE!")
				}
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
	return nil
}
