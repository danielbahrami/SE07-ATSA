package database

import (
	"atse/scheduler_system/dto"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	User   string
	Pass   string
	Host   string
	Port   any
	DBName string
	db     *gorm.DB
}

func NewPG(user string, pass string, host string, port any, dbname string) Postgres {
	return Postgres{
		User:   user,
		Pass:   pass,
		Host:   host,
		Port:   port,
		DBName: dbname,
	}
}

func (p *Postgres) Connect() {
	log.Println("Connecting to postgres ...")
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		p.Host, p.Port, p.User, p.Pass, p.DBName)
	/*
		db, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			panic(err)
		}
		p.db = db*/

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	p.db = db
	log.Println("Connection established")
	log.Println("Migrating schemas ...")
	db.AutoMigrate(&dto.Schedule{}, &dto.ProductionLine{})
	log.Println("Done")
}

func (p *Postgres) GetDB() *gorm.DB {
	return p.db
}
