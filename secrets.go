package secrets

import (
	"context"
	"os"

	"github.com/1password/onepassword-sdk-go"
)

func GetSecret(secretReference string) string {
	token := os.Getenv("OP_SERVICE_ACCOUNT_TOKEN")
	client, err := onepassword.NewClient(
		context.TODO(),
		onepassword.WithServiceAccountToken(token),
		onepassword.WithIntegrationInfo("TCTG API Secrets", "v1.0.0"),
	)
	if err != nil {
		panic(err)
	}
	secret, err := client.Secrets.Resolve(context.TODO(), secretReference)
	if err != nil {
		panic(err)
	}

	return secret
}
