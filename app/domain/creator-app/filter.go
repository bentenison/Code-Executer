package creatorapp

import (
	"fmt"
	"strconv"
	"time"

	"github.com/bentenison/microservice/business/domain/creatorbus"
	"github.com/gin-gonic/gin"
)

func parseQueryParams(c *gin.Context) QueryParams {

	filter := QueryParams{
		Page:             c.Query("page"),
		Rows:             c.Query("row"),
		OrderBy:          c.Query("orderBy"),
		ID:               c.Query("question_id"),
		UserID:           c.Query("user_id"),
		Lang:             c.Query("lang"),
		Tags:             c.Query("tags"),
		StartCreatedDate: c.Query("start_created_date"),
		EndCreatedDate:   c.Query("end_created_date"),
	}
	isQc, _ := strconv.ParseBool(c.Query("is_qc"))
	filter.IsQC = isQc
	return filter
}

func parseFilter(qp QueryParams) (creatorbus.QueryFilter, error) {
	var filter creatorbus.QueryFilter

	if qp.ID != "" {
		filter.ID = qp.ID
	}

	if qp.UserID != "" {

		filter.UserID = qp.UserID
	}

	if qp.Lang != "" {
		filter.Lang = qp.Lang
	}
	if qp.Tags != "" {

		filter.Tags = qp.Tags
	}
	filter.IsQc = qp.IsQC
	if qp.StartCreatedDate != "" {
		t, err := time.Parse(time.RFC3339, qp.StartCreatedDate)
		if err != nil {
			return creatorbus.QueryFilter{}, fmt.Errorf("start_created_date invalid format")
		}
		filter.StartCreatedDate = &t
	}

	if qp.EndCreatedDate != "" {
		t, err := time.Parse(time.RFC3339, qp.EndCreatedDate)
		if err != nil {
			return creatorbus.QueryFilter{}, fmt.Errorf("end_created_date invalid format")
		}
		filter.EndCreatedDate = &t
	}

	return filter, nil
}
