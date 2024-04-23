package repository

import (
	"Stat4Market/internal/domain"
	"context"
	"crypto/tls"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"log"
	"time"
)

type Repository interface {
	CreateEvent(ctx context.Context, event domain.Event) error
}
type Event struct {
	db driver.Conn
}

func NewEvent(db driver.Conn) Repository {
	return &Event{
		db: db,
	}
}

func Connect() (driver.Conn, error) {
	var (
		ctx       = context.Background()
		conn, err = clickhouse.Open(&clickhouse.Options{
			Addr: []string{"localhost:8123"},
			Auth: clickhouse.Auth{
				Database: "default",
				Username: "default",
				Password: "password",
			},
			ClientInfo: clickhouse.ClientInfo{
				Products: []struct {
					Name    string
					Version string
				}{
					{Name: "an-example-go-client", Version: "0.1"},
				},
			},

			Debugf: func(format string, v ...interface{}) {
				fmt.Printf(format, v)
			},
			TLS: &tls.Config{
				InsecureSkipVerify: true,
			},
		})
	)

	if err != nil {
		return nil, err
	}

	if err := conn.Ping(ctx); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("Exception [%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		}
		return nil, err
	}
	return conn, nil
}

func (e *Event) CreateEvent(ctx context.Context, event domain.Event) error {
	querySql := `INSERT INTO events
    (eventType,
     userID, 
     eventTime,
     payload,
     ) VALUES (?, ?, ?, ?)`
	err := e.db.Exec(ctx, querySql,
		event.EventType,
		event.UserId,
		event.EventTime,
		event.Payload,
	)
	if err != nil {
		return err
	}
	return nil
}

//Задача 2
//Реализовать на GO:
//
//Вставку тестовых данных в таблицу events.
//Вывод событий по заданному eventType и временному диапазону.

func (e *Event) InsertTest(ctx context.Context) error {
	querySql := `INSERT INTO events
    (eventId,
     eventType,
     userID, 
     eventTime,
     payload,
     ) VALUES (?, ?, ?, ?, ?)`
	err := e.db.Exec(ctx, querySql,
		1,
		"login",
		1,
		time.Now(),
		`{"key":"value"}`,
	)
	if err != nil {
		return err
	}
	return nil
}

func (e *Event) ListEvents(ctx context.Context, eventType string, start, end time.Time) error {
	querySql := `SELECT * FROM events WHERE eventType = ? AND eventTime BETWEEN ? AND ?`
	rows, err := e.db.Query(ctx, querySql, eventType, start, end)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var eventID int64
		var eventType string
		var userID int64
		var eventTime time.Time
		var payload string
		if err := rows.Scan(&eventID, &eventType, &userID, &eventTime, &payload); err != nil {
			return err
		}
		log.Printf("Event ID: %d, Type: %s, User ID: %d, Time: %v, Payload: %s\n",
			eventID, eventType, userID, eventTime, payload)
	}
	return nil
}
