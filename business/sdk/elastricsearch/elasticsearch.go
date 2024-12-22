package elastricsearch

import (
	"fmt"

	"github.com/elastic/go-elasticsearch/v8"
)

func InitElasticsearch() (*elasticsearch.Client, error) {
	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200", // Replace with your Elasticsearch address
		},
	})
	if err != nil {
		return nil, fmt.Errorf("error initializing Elasticsearch client: %v", err)
	}

	return client, nil
}
