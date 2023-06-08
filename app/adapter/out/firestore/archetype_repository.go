package firestore

import (
	"archetype/app/domain/ports/out"
	"archetype/app/shared/archetype/container"
	einar "archetype/app/shared/archetype/firestore"
	"context"

	"cloud.google.com/go/firestore"
	"github.com/google/uuid"
)

var archetypeCollection *firestore.CollectionRef

func init() {
	container.InjectOutBoundAdapter(func() error {
		archetypeCollection = einar.Client.Collection("INSERT_YOUR_COLLECTION_CONSTANTS_HERE")
		return nil
	}, container.InjectionProps{
		DependencyID: uuid.NewString(),
		Parallel:     false,
	})
}

var ArchetypeRepository out.ArchetypeOutBoundPort = func(ctx context.Context, REPLACE_BY_YOUR_DOMAIN map[string]string) error {
	//PUT YOUR FIRESTORE OPERATION HERE :
	//....archetypeCollection.Doc()
	//....archetypeCollection.Add()
	//....
	return nil
}
