# azure-sql-failover-go
Utility to failover an Azure SQL database that is part of a failover group using Go. I published this snippet of code as there did not seem to be any sample for Go that showed Azure SQL failover groups and it took some digging to figure it out. The utility is useful when trying to do these operations from a Terraform script.

## Calling the utility

### Environment Variables 

| Variable | Description |
|----------|-------------|
| AZURE_SUBSCRIPTION_ID | Guid of subscription |
| AZURE_TENANT_ID | Azure AD tenand ID |
| AZURE_CLIENT_ID | Service principal [Client ID](https://github.com/cloudfoundry/bosh-azure-cpi-release/blob/master/docs/get-started/create-service-principal.md) |
| AZURE_CLIENT_SECRET | The password/secret associated with the Client ID |

### Command-line parameters

| Parameter position | Description |
|-----------|-------------|
| 1 | Resource Group Name |
| 2 | Database Server Name (the server in the group which will become Primary) |
| 3 | Failover group name |

## Building the code

Change to the src folder and then:
`go get` and then `go build -o sqlfailover`
