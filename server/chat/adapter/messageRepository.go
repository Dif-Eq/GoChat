package adapter

import (
	"context"
	"fmt"
	"os"
	"time"

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

const GET_MESSAGES string = `
select
	m.id,
	u.username,
	m.contents,
	m.created_at
from messages m
join users u on m.user_id = u.id;
`

func GetMessages() []model.Message {
	conn := createConnection()

	defer conn.Close()

	rows, err := conn.Query(context.Background(), GET_MESSAGES)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error occurred querying messages: %v\n", err)
	}

	var data []model.Message

	for rows.Next() {
		values, _ := rows.Values()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error occurred constructing uuid: %v\n", err)
		}

		data = append(data, model.Message{
			Username:  values[1].(string),
			Contents:  values[2].(string),
			CreatedAt: values[3].(time.Time),
		})
	}

	return data
}
