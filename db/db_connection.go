package db

import (
	"backend/model"
	"fmt"
	"log"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"gorm.io/gorm"

	"os"
)

type Sql struct {
	Db       *sqlx.DB
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
}

func (s *Sql) Connect() {
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", s.Host, s.Port, s.User, s.Password, s.Dbname)

	var err error

	s.Db, err = sqlx.Connect("postgres", dbinfo)
	/*
		gormDB, _ := gorm.Open(postgres.New(postgres.Config{
			Conn: s.Db,
		}), &gorm.Config{})
		Migratsion(gormDB) */

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)

	}
	s.Db.SetMaxIdleConns(3)
	s.Db.SetMaxOpenConns(3)
	s.Db.SetConnMaxLifetime(0)

	/* 	result, err := db.Exec(
	   		`INSERT INTO users (id , username, email) VALUES ($1 , $2 , $3)`,
	   		"gopher",
	   		"gopher",
	   		"gopher",
	   	)
	   	fmt.Println(result)
	*/
	if err != nil {
		panic(err)
	}

	//defer db.Close()

	fmt.Println("Connect ok")
}

func Migratsion(s *gorm.DB) {

	/* s.Migrator().DropTable(model.Province{})
	s.Migrator().DropTable(model.Schedule{})
	s.Migrator().DropTable(model.Region{}) */
	/* s.Migrator().DropTable(model.User{})
	s.Migrator().DropTable(model.Permission{})
	s.Migrator().DropTable("user_role")
	s.Migrator().DropTable(model.Role{})
	*/
	s.AutoMigrate(model.User{})
	s.AutoMigrate(model.Permission{})

}

func (s *Sql) Close() {
	s.Db.Close()
}
