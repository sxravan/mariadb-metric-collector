package collector

import (
	"encoding/json"
	"os"
)

// Query represents a SQL query with a title
type Query struct {
	Title string `json:"title"`
	Query string `json:"query"`
}

// LoadQuery reads a JSON file and returns a list of queries
func LoadQuery(filename string) ([]Query, error) {
	// Read the file contents using os.ReadFile (new method)
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Parse JSON into a slice of Query structs
	var query []Query
	err = json.Unmarshal(data, &query)
	if err != nil {
		return nil, err
	}

	return query, nil
}
