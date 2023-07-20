package helloworld

import (
	"context"

	"cloud.google.com/go/firestore"
)

// Get_documents_count retrieves the total count of existing documents in the collection.
func Get_documents_count(ctx context.Context, client *firestore.Client) (int, error) {
	query := client.Collection("grocery_items_database").Select("ID").Limit(1)

	docs, err := query.Documents(ctx).GetAll()
	if err != nil {
		return 0, err
	}

	return len(docs), nil
}
