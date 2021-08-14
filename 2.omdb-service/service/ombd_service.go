package service

import (
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"septanto.github.com/omdb/model"
	"septanto.github.com/omdb/repository"
)

const apiKey = "184f325b"
const url = "http://www.omdbapi.com"

var client = resty.New().SetDebug(true)

func RequestSearchToOmdb(searchword string, pagination string) (*resty.Response, error) {
	queryParam := map[string]string{
		"apikey": apiKey,
		"s":      searchword,
		"page":   pagination,
	}
	resp, err := client.R().SetQueryParams(queryParam).Get(url)
	log := model.SearchLog{}
	log.ID = uuid.New().String()
	log.Request = "s: " + searchword + "; page: " + pagination
	log.Response = string(resp.Body())
	log.TimeCreate = time.Now()
	repository.SaveLogToDb(log)
	return resp, err
}

func RequestDetailToOmdb(imdbid string) (*resty.Response, error) {
	queryParam := map[string]string{
		"apikey": apiKey,
		"i":      imdbid,
	}
	resp, err := client.R().SetQueryParams(queryParam).Get(url)
	log := model.SearchLog{}
	log.ID = uuid.New().String()
	log.Request = "i: " + imdbid
	log.Response = string(resp.Body())
	log.TimeCreate = time.Now()
	repository.SaveLogToDb(log)
	return resp, err
}
