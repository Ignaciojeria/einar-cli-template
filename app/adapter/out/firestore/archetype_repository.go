package firestore

import (
	einar "archetype/app/shared/archetype/firestore"
	"context"

	"cloud.google.com/go/firestore"
)

var ArchetypeRepository = func(ctx context.Context, REPLACE_BY_YOUR_DOMAIN map[string]string) error {
	var _ *firestore.CollectionRef = einar.Collection("INSERT_YOUR_COLLECTION_CONSTANT_HERE")
	return nil
}
