package secrets

import (
	"context"

	"github.com/1password/onepassword-sdk-go"
)

func GetClient() (*onepassword.Client, error) {

	client, err := onepassword.NewClient(
		context.TODO(),
		onepassword.WithServiceAccountToken(token),
		onepassword.WithIntegrationInfo("TCTG API Secrets", "v1.0.0"),
	)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func GetSecret(client *onepassword.Client, secretReference string) (string, error) {

	secret, err := client.Secrets.Resolve(context.TODO(), secretReference)
	if err != nil {
		return "", err
	}

	return secret, nil
}
