package elastricsearch

import (
	"fmt"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
)

func InitElasticsearch() (*elasticsearch.Client, error) {
	maxRetries := 5
	baseDelay := 1 * time.Second
	var client *elasticsearch.Client
	var err error

	for i := 0; i < maxRetries; i++ {
		client, err = elasticsearch.NewClient(elasticsearch.Config{
			Addresses: []string{
				"http://localhost:9200", // Replace with your Elasticsearch address
			},
		})
		if err != nil {
			return nil, fmt.Errorf("error initializing Elasticsearch client: %v", err)
		}

		// Check if the Elasticsearch server is available
		res, err := client.Info()
		if err == nil && res.StatusCode == 200 {
			fmt.Println("Elasticsearch client initialized successfully.")
			return client, nil
		}

		// Log the error and retry
		fmt.Printf("Attempt %d failed: %v. Retrying in %v...\n", i+1, err, baseDelay)
		time.Sleep(baseDelay * (1 << i)) // Exponential backoff
	}

	return nil, fmt.Errorf("failed to connect to Elasticsearch after %d attempts", maxRetries)
}
