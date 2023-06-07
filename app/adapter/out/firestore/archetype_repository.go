package firestore

import (
	"archetype/app/ports/out"
	fs "archetype/app/shared/archetype/firestore"
	"context"

	"cloud.google.com/go/firestore"
)

var ArchetypeRepository out.ArchetypeOutBoundPort = func(ctx context.Context, REPLACE_BY_YOUR_DOMAIN map[string]string) error {
	var _ *firestore.CollectionRef = fs.Client.Collection("INSERT_YOUR_COLLECTION_PATH_HERE")
	//PUT YOUR FIRESTORE OPERATION HERE :
	//....
	return nil
}
