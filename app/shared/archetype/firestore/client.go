package firestore

import (
	"archetype/app/shared/archetype/container"
	"archetype/app/shared/config"
	"archetype/app/shared/utils"
	"sync"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"

	"context"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var (
	Client         *firestore.Client
	collectionRefs sync.Map
)

func init() {
	config.Installations.EnableFirestore = true
	container.InjectInstallation(func() error {
		ctx := context.Background()

		creds, err := utils.DecodeBase64(config.GOOGLE_APPLICATION_CRETENTIALS_B64.Get())
		if err != nil {
			log.Error().Err(err).Msg("error decoding GOOGLE_APPLICATION_CRETENTIALS_B64")
			return err
		}

		sa := option.WithCredentialsJSON(creds)

		app, err := firebase.NewApp(ctx, &firebase.Config{
			ProjectID: config.GOOGLE_PROJECT_ID.Get(),
		}, sa)

		if err != nil {
			log.Error().Err(err).Msg("error initializing firebase app")
			return err
		}

		c, err := app.Firestore(ctx)

		if err != nil {
			log.Error().Err(err).Msg("error getting firestore client")
			return err
		}
		Client = c
		return nil
	}, container.InjectionProps{DependencyID: uuid.NewString()})
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
