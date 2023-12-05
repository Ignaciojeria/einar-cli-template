package firestore

import (
	"archetype/app/shared/archetype/container"
	"archetype/app/shared/archetype/slog"
	"archetype/app/shared/config"
	"archetype/app/shared/constants"
	"sync"

	"context"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
)

var (
	Client         *firestore.Client
	collectionRefs sync.Map
)

func init() {
	config.Installations.EnableFirestore = true
	container.InjectInstallation(func() error {
		ctx := context.Background()
		app, err := firebase.NewApp(ctx, &firebase.Config{
			ProjectID: config.GOOGLE_PROJECT_ID.Get(),
		})
		if err != nil {
			slog.Logger.Error("error initializing firebase app", constants.ERROR, err.Error())
			return err
		}
		c, err := app.Firestore(ctx)
		if err != nil {
			slog.Logger.Error("error getting firestore client", constants.ERROR, err.Error())
			return err
		}
		Client = c
		return nil
	})
}

// Collection fetches a *firestore.CollectionRef by name. If the CollectionRef exists in the sync.Map, it's returned, otherwise a new one is created and stored in the map.
func Collection(collectionName string) *firestore.CollectionRef {
	value, ok := collectionRefs.Load(collectionName)
	if ok {
		// If the collection reference was found, return it.
		return value.(*firestore.CollectionRef)
	}

	// If the collection reference was not found, create a new one.
	newCollectionRef := Client.Collection(collectionName)

	// Store the new collection reference in the map.
	collectionRefs.Store(collectionName, newCollectionRef)

	// Return the new collection reference.
	return newCollectionRef
}
