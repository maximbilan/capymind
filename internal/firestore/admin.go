package firestore

import (
	"context"
	"errors"
	"time"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/firestore/apiv1/firestorepb"
)

// Get the total number of users
func GetTotalUserCount(ctx *context.Context) (int64, error) {
	return getTotalRecordsCount(ctx, users.String(), nil)
}

// Get the number of active users (User.timestamp less than 7 days)
func GetActiveUserCount(ctx *context.Context) (int64, error) {
	sevenDaysAgo := time.Now().AddDate(0, 0, -7)
	filter := func(query firestore.Query) firestore.Query {
		return query.Where("timestamp", ">", sevenDaysAgo)
	}
	return getTotalRecordsCount(ctx, users.String(), filter)
}

// Get the total number of notes
func GetTotalNoteCount(ctx *context.Context) (int64, error) {
	return getTotalRecordsCount(ctx, notes.String(), nil)
}

// Get the total number of records in a collection with an optional filter
func getTotalRecordsCount(ctx *context.Context, collection string, filter func(firestore.Query) firestore.Query) (int64, error) {
	query := client.Collection(collection).Query

	if filter != nil {
		query = filter(query)
	}

	aggregationQuery := query.NewAggregationQuery().WithCount("all")

	results, err := aggregationQuery.Get(*ctx)
	if err != nil {
		return 0, err
	}

	count, ok := results["all"]
	if !ok {
		return 0, errors.New("[Firestore]: Couldn't get alias for COUNT from results")
	}

	countValue := count.(*firestorepb.Value)
	return countValue.GetIntegerValue(), nil
}
