package external

import (
	"fmt"
	"os"

	"github.com/ITEBARPLKelompok3/peminjaman-ruangan/backend/config"
	"github.com/getsentry/sentry-go"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Postgres is a struct to represent postgresql connection requirement
//
//	type Postgres struct {
//		AppName, Host, Port, Database, Username, Password string
//	}
type Postgres struct {
	AppName          string `mapstructure:"APP_NAME"`
	PostgresHost     string `mapstructure:"POSTGRES_HOST"`
	PostgresPort     string `mapstructure:"POSTGRES_PORT"`
	PostgresUser     string `mapstructure:"POSTGRES_USER"`
	PostgresPassword string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresDatabase string `mapstructure:"POSTGRES_DATABASE"`
}

// GetConn is a function to establish connection to postgres database
func (p *Postgres) GetConn() (db *gorm.DB, err error) {
	// load app.env file data to struct
	// config, err := config.LoadConfig(".")
	// // handle errors
	// if err != nil {
	// 	log.Fatalf("can't load environment app.env: %v", err)
	// }

	connString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable timezone=Asia/Jakarta",
		p.PostgresHost,
		p.PostgresPort,
		p.PostgresUser,
		p.PostgresDatabase,
		p.PostgresPassword)
	// if p.SslCert != "" {
	// 	connString = fmt.Sprintf("%s sslmode=verify-ca sslrootcert=%s", connString, p.SslCert)
	// }
	if p.AppName != "" {
		connString = fmt.Sprintf("%s application_name=%s", connString, p.AppName)
	}
	db, err = gorm.Open(postgres.Open(connString))
	if err != nil {
		sentry.CaptureException(err)
		log.Error(err)
		return
	}
	// db.(config.GORMLog)
	// db.DB().SetMaxIdleConns(config.PostgresMinConn)
	// db.DB().SetMaxOpenConns(config.PostgresMaxConn)
	return
}

// GetPostgresClient is an extended function to establish connection to postgres database
func GetPostgresClient() (db *gorm.DB, err error) {
	// load app.env file data to struct
	config, err := config.LoadConfig(".")
	// handle errors
	if err != nil {
		log.Fatalf("can't load environment app.env: %v", err)
	}

	appName := os.Getenv("APP_NAME")
	if appName == "" {
		appName = "unknown"
	}
	connection := &Postgres{
		AppName:          appName,
		PostgresHost:     config.PostgresHost,
		PostgresPort:     config.PostgresPort,
		PostgresDatabase: config.PostgresDatabase,
		PostgresUser:     config.PostgresUser,
		PostgresPassword: config.PostgresPassword,
	}
	db, err = connection.GetConn()
	if err != nil {
		sentry.CaptureException(err)
		log.Error(err)
		return
	}
	if db == nil {
		err := fmt.Errorf("postgres connection to %s does not exist", config.PostgresHost)
		sentry.CaptureException(err)
		return nil, err
	}
	return db, nil
}
