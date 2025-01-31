package collector

import (
	"database/sql"
	"fmt"
)

type Metric struct {
	Name  string
	Value float64
}

func CollectMetrics(db *sql.DB) ([]Metric, error) {
	var metrics []Metric

	// Example query to fetch server status or any other metric
	rows, err := db.Query("SHOW GLOBAL STATUS")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		var value string
		if err := rows.Scan(&name, &value); err != nil {
			return nil, err
		}

		// Convert value to float64 (as an example)
		metricValue := 0.0
		fmt.Sscanf(value, "%f", &metricValue)

		metrics = append(metrics, Metric{
			Name:  name,
			Value: metricValue,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return metrics, nil
}
