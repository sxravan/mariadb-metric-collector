package collector

import (
	"database/sql"
	"fmt"
	"log"
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

func ExecuteQuery(db *sql.DB, query []Query) {
	for _, q := range query {
		// Execute the query and get a result set
		rows, err := db.Query(q.Query)
		if err != nil {
			log.Printf("Error executing query [%s]: %v", q.Title, err)
			continue
		}
		defer rows.Close()

		// Iterate over the rows returned by the query
		for rows.Next() {
			var result string
			// Scan the values into the result variable (change this based on the expected type of your result)
			err := rows.Scan(&result)
			if err != nil {
				log.Printf("Error scanning row for query [%s]: %v", q.Title, err)
				continue
			}

			// Print the result for each row
			fmt.Printf("%s: %s\n", q.Title, result)
		}

		// Check if there was an error iterating over rows
		if err := rows.Err(); err != nil {
			log.Printf("Error iterating over rows for query [%s]: %v", q.Title, err)
		}
	}
}
