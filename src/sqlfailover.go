package main

import (
	"context"
	"fmt"
	sqlsdk "github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v3.0/sql"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"log"
	"os"
)

var server string
var database string
var subid string
var resgroup string

// https://godoc.org/github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v3.0/sql#DatabasesClient.Failover
// The only available documentation on calling this API.
func main() {

	var err error
	if len(os.Args) < 4 {
		log.Fatal("Usage: sqlfailover <resource-group> <server> <database> <primary/secondary>")
		os.Exit(1)
	}
	temp := os.Getenv("AZURE_SUBSCRIPTION_ID")
	if len(temp) == 0 {
		log.Fatal("Environment variable AZURE_SUBSCRIPTION_ID not found")
		os.Exit(2)
	}

	subid = temp
	temp = os.Getenv("AZURE_TENANT_ID")
	if len(temp) == 0 {
		log.Fatal("Environment variable AZURE_TENANT_ID not found")
		os.Exit(2)
	}

	temp = os.Getenv("AZURE_CLIENT_ID")
	if len(temp) == 0 {
		log.Fatal("Environment variable AZURE_CLIENT_ID not found")
		os.Exit(2)
	}

	temp = os.Getenv("AZURE_CLIENT_SECRET")
	if len(temp) == 0 {
		log.Fatal("Environment variable AZURE_CLIENT_SECRET not found")
		os.Exit(2)
	}

	resgroup = os.Args[1]
	server = os.Args[2]
	database = os.Args[3]

	// Create auth token from env variables (see here for details https://github.com/Azure/azure-sdk-for-go)
	authorizer, err := auth.NewAuthorizerFromEnvironment()
	if err == nil {
		// Create AzureSQL SDK client
		dbclient := sqlsdk.NewDatabasesClient(subid)
		dbclient.Authorizer = authorizer
		var mode sqlsdk.ReplicaType
		if os.Args[4] == "primary" {
			mode = "Primary"
		} else if os.Args[4] == "secondary" {
			mode = "ReadableSecondary"
		} else {
			log.Fatal("Invalid ReplicaType, must be primary or secondary")
			os.Exit(3)
		}
		ctx := context.Background()
		future, err := dbclient.Failover(ctx, resgroup, server, database, mode)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			err = future.WaitForCompletionRef(ctx, dbclient.Client)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
}
