package db

import (
	"fmt"
	"time"

	"github.com/J-Rivard/trading-bot/internal/logging"
)

func (d *DB) UpdateEvents() {
	events, err := d.FetchEvents()
	if err != nil {
		d.Log.LogDebug(logging.FormattedLog{
			"action": "db_update_events",
			"error":  err.Error(),
		})
	}
	d.Log.LogDebug(logging.FormattedLog{
		"action":   "db_update_events",
		"metadata": fmt.Sprintf("Events fetched: %d", len(events)),
	})
	d.Events = events

	for {
		ticker := time.NewTicker(30 * time.Second)

		select {
		case <-ticker.C:
			events, err := d.FetchEvents()
			if err != nil {
				d.Log.LogDebug(logging.FormattedLog{
					"action": "db_update_events",
					"error":  err.Error(),
				})
				continue
			}
			d.Log.LogDebug(logging.FormattedLog{
				"action":   "db_update_events",
				"metadata": fmt.Sprintf("Events fetched: %d", len(events)),
			})
			d.Events = events
		}
	}
}

func (d *DB) FetchEvents() ([]*string, error) {
	var events []*string
	queryString := `SELECT event FROM events`

	rows, err := d.Client.Query(queryString)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var event *string
		err = rows.Scan(&event)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}
