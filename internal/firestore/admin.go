package firestore

import (
	"context"
	"errors"

	"cloud.google.com/go/firestore/apiv1/firestorepb"
)

// Get the total number of users
func GetTotalUserCount(ctx *context.Context) (int64, error) {
	return getTotalRecordsCount(ctx, users.String())
}

// Get the total number of notes
func GetTotalNoteCount(ctx *context.Context) (int64, error) {
	return getTotalRecordsCount(ctx, notes.String())
}

// Get the total number of records in a collection
func getTotalRecordsCount(ctx *context.Context, collection string) (int64, error) {
	aggregationQuery := client.Collection(collection).NewAggregationQuery().WithCount("all")

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
