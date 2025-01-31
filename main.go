package main

import (
	"fmt"
	"log"

	"github.com/sxravan/mariadb-metric-collector/collector"
	"github.com/sxravan/mariadb-metric-collector/config"
	"github.com/sxravan/mariadb-metric-collector/db"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Connect to the database
	database, err := db.Connect(cfg)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()

	// Collect metrics
	metrics, err := collector.CollectMetrics(database)
	if err != nil {
		log.Fatalf("Error collecting metrics: %v", err)
	}

	// Print metrics
	for _, metric := range metrics {
		fmt.Printf("Metric: %s, Value: %f\n", metric.Name, metric.Value)
	}

	// Load queries from JSON
	query, err := collector.LoadQuery("queries.json")
	if err != nil {
		log.Fatalf("Error loading query: %v", err)
	}

	// Execute and print query results
	collector.ExecuteQuery(database, query)
}
