package dbtest

import (
	"github.com/bentenison/microservice/business/sdk/mongodb"
	"github.com/bentenison/microservice/foundation/logger"
	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
)

type Database struct {
	DB        DS
	Log       *logger.CustomLogger
	BusDomain BusDomain
}
type DS struct {
	SQL *sqlx.DB
	MGO *mongo.Database
}

func New() (*Database, error) {
	// db, err := sqldb.Open(sqldb.Config{
	// 	User:         "exector",
	// 	Password:     "admin#123",
	// 	Host:         "localhost",
	// 	Name:         "",
	// 	MaxIdleConns: 10,
	// 	MaxOpenConns: 10,
	// })
	// if err != nil {
	// 	return nil, fmt.Errorf("connecting to db: %w", err)
	// }

	// defer db.Close()
	// starting mongio db connection

	mongo, err := mongodb.InitializeMongo(mongodb.Config{
		Username:    "admin",
		Password:    "admin#123",
		AuthDB:      "admin",
		Host:        "localhost",
		Port:        "27017",
		DBName:      "EXECUTOR",
		AllowDirect: false,
	})
	if err != nil {
		return nil, err
	}
	ds := DS{
		MGO: mongo,
		SQL: nil,
	}
	log := logger.NewCustomLogger(map[string]interface{}{
		"service": "Test",
	})
	return &Database{
		DB:        ds,
		Log:       log,
		BusDomain: newBusDomain(log, ds),
	}, nil
}
