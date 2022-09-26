package repository

import (
	"assignment3/entity"
	"context"
	"database/sql"
	"log"
)

type weatherRepositoryImpl struct {
	DB *sql.DB
}

type WeatherRepository interface {
	UpdateValue(windValue string, waterValue string, report string) error
	GetUpdate(ctx context.Context) (entity.StatusUpdate, error)
}

func ProvideWeatherRepository(DB *sql.DB) *weatherRepositoryImpl {
	return &weatherRepositoryImpl{DB: DB}
}

var (
	UPDATE_DATA     = "UPDATE status SET wind = ?, water = ?, update_data = ? WHERE id = 1"
	GET_REPORT_DATA = "SELECT wind, water, update_data FROM status"
)

func (w weatherRepositoryImpl) UpdateValue(windValue string, waterValue string, report string) error {
	query := UPDATE_DATA

	stmt, err := w.DB.Prepare(query)
	if err != nil {
		log.Printf("[UpdateValue] failed to prepare the statement, err => %v", err)
		return err
	}
	_, err = stmt.Exec(windValue, waterValue, report)
	if err != nil {
		log.Printf("[UpdateValue] failed to execute to the database :%v", err)
		return err
	}
	return nil
}

func (w weatherRepositoryImpl) GetUpdate(ctx context.Context) (entity.StatusUpdate, error) {
	query := GET_REPORT_DATA

	stmt, err := w.DB.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("[GetReport] failed to prepare the statement")
		return entity.StatusUpdate{}, err
	}
	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		log.Printf("[GetReport] failed to query to the database")
		return entity.StatusUpdate{}, err
	}

	update := entity.StatusUpdate{}
	for rows.Next() {
		details := entity.StatusUpdate{}
		err := rows.Scan(
			&details.Weather.Wind,
			&details.Weather.Water,
			&details.Status,
		)
		if err != nil {
			log.Printf("[GetReport] failed to scan the data")
			return entity.StatusUpdate{}, err
		}
		update = details
	}
	return update, nil
}
