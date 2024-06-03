package main

import (
	"context"
	"fmt"
	"os"

	"github.com/1password/onepassword-sdk-go"
)

func getSecret(secretReference string) string {
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

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a secret reference as a command line argument.")
		os.Exit(1)
	}
	secretReference := os.Args[1]
	secret := getSecret(secretReference)
	fmt.Print(secret)
}
