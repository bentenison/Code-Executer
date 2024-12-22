package creatordb

import (
	"strings"

	"github.com/bentenison/microservice/business/domain/creatorbus"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *Store) applyFilter(filter creatorbus.QueryFilter) bson.M {
	data := make(map[string]any)
	if filter.ID != "" {
		data["id"] = *&filter.ID
		// wc = append(wc, "id = :id")
	}

	// if filter.UserID != "" {
	// 	data["user_id"] = *filter.UserID
	// 	wc = append(wc, "user_id = :user_id")
	// }

	if filter.Lang != "" {
		data["language"] = filter.Lang
		// wc = append(wc, "language = :language")
	}
	if filter.Tags != "" {
		var tg []string
		if strings.Contains(filter.Tags, ",") {
			tg = strings.Split(filter.Tags, ",")
		} else {
			tg = append(tg, filter.Tags)
		}
		data["tags"] = tg
		// wc = append(wc, "tags = :tags")
	}
	// if filter.IsQc {
	// tg := strings.Split(filter.Tags, ",")
	data["is_qc"] = filter.IsQc
	// wc = append(wc, "tags = :tags")
	// }

	query := bson.M{}

	// Add filters dynamically
	for key, value := range data {
		switch v := value.(type) {
		case string:
			query[key] = bson.M{"$eq": v} // exact match
		case []string:
			query[key] = bson.M{"$in": v} // match any in the list
		case bool:
			query[key] = bson.M{"$eq": v} // exact match for bool values
		default:
			// Handle other data types as needed (e.g., numbers, dates)
			query[key] = v
		}
	}
	return query
}
