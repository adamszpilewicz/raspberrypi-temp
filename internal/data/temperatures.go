package data

import (
	"database/sql"
	"log"
	"time"
)

type Temp struct {
	Timestamp   time.Time `json:"timestamp"`
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
}

type TempModel struct {
	DB *sql.DB
}

func (m TempModel) GetAllLimit() ([]*Temp, error) {

	query := `
		select  timestamp, temperature, humidity
		from temperature
		where timestamp > $1
		order by timestamp desc
		limit 20
	`

	rows, err := m.DB.Query(query, time.Now().Add(-24*time.Hour))
	if err != nil {
		log.Println(err)
	}

	defer rows.Close()

	temps := []*Temp{}

	for rows.Next() {

		var temperature Temp

		err := rows.Scan(
			&temperature.Timestamp,
			&temperature.Temperature,
			&temperature.Humidity,
		)
		if err != nil {
			return nil, err
		}
		temps = append(temps, &temperature)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return temps, nil

}
