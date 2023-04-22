package mysql

import (
	"fmt"
	"github.com/prajnapras19/otpgeneratorwithrecoveryserver/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Service interface {
	InitDB() *gorm.DB
	GetDSN() string
	GetDB() *gorm.DB
}

type service struct {
	config config.MySQLConfig
	db     *gorm.DB
}

func NewService(mySQLConfig config.MySQLConfig) Service {
	svc := &service{
		config: mySQLConfig,
	}
	svc.db = svc.InitDB()
	return svc
}

func (s *service) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		s.config.Username, s.config.Password, s.config.Hostname, s.config.Port, s.config.DBName,
		s.config.Charset, s.config.ParseTime, s.config.Loc,
	)
}

func (s *service) InitDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(s.GetDSN()), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func (s *service) GetDB() *gorm.DB {
	return s.db
}
