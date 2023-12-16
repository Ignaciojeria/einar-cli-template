package secret_manager

import (
	"archetype/app/shared/archetype/container"
	"archetype/app/shared/config"
	"context"
	"encoding/json"
	"fmt"
	"log"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"google.golang.org/api/iterator"
)

var client *secretmanager.Client

var _ = container.InjectConfiguration(func() error {
	ctx := context.Background()
	c, err := secretmanager.NewClient(ctx)
	if err != nil {
		log.Fatalf("failed to create secretmanager client: %v", err)
		return err
	}
	client = c
	load(map[string]string{string(config.PROJECT_NAME): config.PROJECT_NAME.Get()})
	return nil
})

func listSecretsByLabels(labels map[string]string) ([]string, error) {
	projectId := config.GOOGLE_PROJECT_ID.Get()
	ctx := context.Background()
	filter := ""
	for key, value := range labels {
		if filter != "" {
			filter += " AND "
		}
		filter += fmt.Sprintf("labels.%s=\"%s\"", key, value)
	}

	req := &secretmanagerpb.ListSecretsRequest{
		Parent: "projects/" + projectId,
		Filter: filter,
	}
	it := client.ListSecrets(ctx, req)
	var secrets []string
	for {
		secret, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("failed to list secrets: %v", err)
			return nil, err
		}
		secrets = append(secrets, secret.Name)
		fmt.Println("Found secret:", secret.Name)
	}
	return secrets, nil
}

func load(filters map[string]string) error {
	ctx := context.Background()
	secrets, err := listSecretsByLabels(filters)
	if err != nil {
		return err
	}
	for _, v := range secrets {
		var keyValueJson map[string]string
		accessRequest := &secretmanagerpb.AccessSecretVersionRequest{
			Name: v + "/versions/latest",
		}
		result, err := client.AccessSecretVersion(ctx, accessRequest)
		if err != nil {
			log.Fatalf("failed to access secret version: %v", err)
			return err
		}
		err = json.Unmarshal(result.Payload.Data, &keyValueJson)
		if err != nil {
			log.Fatalf("failed to unmarshaling secret: %v. error %v", v, err)
			return err
		}

		for key, val := range keyValueJson {
			config.Set(key, val)
		}
	}
	if len(secrets) == 0 {
		log.Println("No secrets found")
	}
	return nil
}
