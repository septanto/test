package repository

import (
	"github.com/labstack/gommon/log"
	"septanto.github.com/omdb/model"
)

func SaveLogToDb(searchLog model.SearchLog) {
	log.Info("Save Log To Db: ID [" + searchLog.ID + "]")
}
