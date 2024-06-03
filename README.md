# 1Password Secret Retriever 

## Description

This project is a Go application that retrieves secrets from a 1Password vault via a service account. It can be used as a standalone binary within other applications, such as AWS Lambda functions. This project uses the official [1Password Go SDK](https://github.com/1Password/onepassword-sdk-go), which is currently in beta.

## Prerequisites

- A [1Password Service Account](https://developer.1password.com/docs/service-accounts/) and existing knowledge on how to use it
- Environmental variable: `OP_SERVICE_ACCOUNT_TOKEN=YOUR_TOKEN_HERE`

## Building

To build the binary for this application, clone this repo and run the following command:

macOS: `GOOS=darwin go build -o 1pwsecret main.go`
Linux: `GOOS=linux go build -o 1pwsecret main.go`
Windows: `GOOS=windows go build -o 1pwsecret main.go`