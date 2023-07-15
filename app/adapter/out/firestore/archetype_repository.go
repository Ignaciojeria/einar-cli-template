package firestore

import (
	"archetype/app/domain/ports/out"
	einar "archetype/app/shared/archetype/firestore"
	"context"
	"fmt"
)

var ArchetypeRepository out.ArchetypeOutBoundPort = func(ctx context.Context, REPLACE_BY_YOUR_DOMAIN map[string]string) error {
	collectionReference := einar.Client.Collection("INSERT_YOUR_COLLECTION_CONSTANTS_HERE")
	fmt.Println(collectionReference)
	//PUT YOUR FIRESTORE OPERATION HERE :
	//....collectionReference.Doc()
	//....collectionReference.Add()
	//....
	return nil
}
