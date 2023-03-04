package db

import (
	"backend/model"
	"fmt"
	"log"
	"time"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"os"
)

type Sql struct {
	Db       *gorm.DB
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
}

func (s *Sql) Connect() {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", s.User, s.Password, s.Host, s.Port, s.Dbname)

	var err error

	s.Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)

	}

	//defer db.Close()
	sqlDB, err := s.Db.DB()

	if err != nil {
		log.Fatal("Failed database. \n", err)
		os.Exit(2)

	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)
	fmt.Println("Connect ok")
	//Migratsion(s.Db)

	//defer sqlDB.Close()
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
	s.Migrator().DropTable(model.Role{})
	s.Migrator().DropTable(model.User{})

	s.AutoMigrate(model.Role{})
	s.AutoMigrate(model.User{})

	//s.AutoMigrate(model.Permission{})
	initDataRole(s)
}

func initDataRole(s *gorm.DB) {

	role := model.Role{
		RoleName:    "admin",
		Description: "admin",
	}
	role2 := model.Role{
		RoleName:    "user",
		Description: "user",
	}
	s.Create(&role)
	s.Create(&role2)

}
