package mysql

import (
	"fmt"
	"mac-laren/pkg/config"
	"mac-laren/pkg/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitConnection(conf *config.Config) (*gorm.DB, error) {

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", conf.GetString("MYSQL_USER"), conf.GetString("MYSQL_PASSWORD"), conf.GetString("MYSQL_HOST"), conf.GetString("MYSQL_PORT"), conf.GetString("MYSQL_DATABASE")),
		DefaultStringSize: 256,
	}), &gorm.Config{})

	if err != nil {
		logger.Logger.Error().Err(err).Str("appname", conf.GetString("APP_NAME")).Msg("error while openning database connection")
		return nil, err
	}

	return db, nil

}
