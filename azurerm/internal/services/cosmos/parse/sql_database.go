package parse

// NOTE: this file is generated via 'go:generate' - manual changes will be overwritten

import (
	"fmt"

	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/azure"
)

type SqlDatabaseId struct {
	SubscriptionId      string
	ResourceGroup       string
	DatabaseAccountName string
	Name                string
}

func NewSqlDatabaseID(subscriptionId, resourceGroup, databaseAccountName, name string) SqlDatabaseId {
	return SqlDatabaseId{
		SubscriptionId:      subscriptionId,
		ResourceGroup:       resourceGroup,
		DatabaseAccountName: databaseAccountName,
		Name:                name,
	}
}

func (id SqlDatabaseId) ID(_ string) string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DocumentDB/DatabaseAccounts/%s/sqlDatabases/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroup, id.DatabaseAccountName, id.Name)
}

func SqlDatabaseID(input string) (*SqlDatabaseId, error) {
	id, err := azure.ParseAzureResourceID(input)
	if err != nil {
		return nil, err
	}

	resourceId := SqlDatabaseId{
		SubscriptionId: id.SubscriptionID,
		ResourceGroup:  id.ResourceGroup,
	}

	if resourceId.DatabaseAccountName, err = id.PopSegment("DatabaseAccounts"); err != nil {
		return nil, err
	}
	if resourceId.Name, err = id.PopSegment("sqlDatabases"); err != nil {
		return nil, err
	}

	if err := id.ValidateNoEmptySegments(input); err != nil {
		return nil, err
	}

	return &resourceId, nil
}
