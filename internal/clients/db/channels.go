package db

type Subscription struct {
	Channel *string
}

func (d *DB) AddSubscription(id string) error {
	insertString := `INSERT INTO channels(channelid) VALUES($1) ON CONFLICT DO NOTHING`

	d.Client.QueryRow(insertString, id)

	return nil
}

func (d *DB) RemoveSubscription(id string) error {
	deleteString := `DELETE FROM channels WHERE channelid=$1`

	d.Client.QueryRow(deleteString, id)

	return nil
}

func (d *DB) FetchSubscriptions() ([]*Subscription, error) {
	var subs []*Subscription
	queryString := `SELECT channelid FROM channels`

	rows, err := d.Client.Query(queryString)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		sub := &Subscription{}
		err = rows.Scan(&sub.Channel)
		if err != nil {
			return nil, err
		}

		subs = append(subs, sub)
	}

	return subs, nil
}
