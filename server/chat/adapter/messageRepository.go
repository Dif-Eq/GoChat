package adapter

import (
	"context"
	"fmt"
	"os"

	"gochat/chat-service/model"

	"github.com/jackc/pgx/v4/pgxpool"
)

func createConnection() *pgxpool.Pool {
	databaseUrl := "postgres://postgres:postgres@localhost:5432/postgres"
	pool, err := pgxpool.Connect(context.Background(), databaseUrl)

	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return pool
}

func GetMessages() []model.Message {
	conn := createConnection()
	defer conn.Close()

	rows, err := conn.Query(context.Background(), "select id, user_id, contents from messages;")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error occurred querying messages: %v\n", err)
	}

	var data []model.Message

	for rows.Next() {
		values, _ := rows.Values()
		contents := values[2].(string)
		data = append(data, model.Message{
			Contents: contents,
		})
	}

	return data
}
