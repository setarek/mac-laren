package mysql

import (
	"mac-laren/pkg/config"
	"mac-laren/pkg/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitConnection(conf *config.Config) (*gorm.DB, error) {

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local", // data source name
		DefaultStringSize:         256,                                                                        // default size for string fields
		DisableDatetimePrecision:  true,                                                                       // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,                                                                       // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,                                                                       // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,                                                                      // auto configure based on currently MySQL version
	}), &gorm.Config{})

	if err != nil {
		logger.Logger.Error().Err(err).Str("appname", conf.GetString("APP_NAME")).Msg("error while openning database connection")
		return nil, err
	}

	return db, nil

}
