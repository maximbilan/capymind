//coverage:ignore file

package firestore

import (
	"context"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

// Delete all documents for the given query
func deleteAllDocs(ctx *context.Context, query firestore.Query, batchSize int) error {
	bulkwriter := client.BulkWriter(*ctx)

	for {
		// Get a batch of documents
		iter := query.Limit(batchSize).Documents(*ctx)
		numDeleted := 0

		// Iterate through the documents, adding
		// a delete operation for each one to the BulkWriter.
		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				return err
			}

			bulkwriter.Delete(doc.Ref)
			numDeleted++
		}

		// If there are no documents to delete,
		// the process is over.
		if numDeleted == 0 {
			bulkwriter.End()
			break
		}

		bulkwriter.Flush()
	}
	return nil
}
