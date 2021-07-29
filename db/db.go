package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
)

var GlobalPoolConnection *pgxpool.Pool

func CreatePoolConn() error {
	var err error
	var host = os.Getenv("HOST")
	var port = os.Getenv("PORT")
	var db = os.Getenv("DB")
	var user = os.Getenv("USER")
	var pass = os.Getenv("PASS")
	var connectionString = fmt.Sprintf("postgresql://%s:%s/%s?user=%s&password=%s", host, port, db, user, pass)
	GlobalPoolConnection, err = pgxpool.Connect(context.Background(), connectionString)
	PingDB()
	if err != nil {
		return errors.New("connection not made: " + err.Error())
	}

	return nil
}

func PingDB() {
	var err = GlobalPoolConnection.Ping(context.Background())
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Connection to DB made")
	}
}