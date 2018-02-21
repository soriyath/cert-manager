package powerbiembedded

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// AccessKeyName enumerates the values for access key name.
type AccessKeyName string

const (
	// Key1 ...
	Key1 AccessKeyName = "key1"
	// Key2 ...
	Key2 AccessKeyName = "key2"
)

// CheckNameReason enumerates the values for check name reason.
type CheckNameReason string

const (
	// Invalid ...
	Invalid CheckNameReason = "Invalid"
	// Unavailable ...
	Unavailable CheckNameReason = "Unavailable"
)

// AzureSku ...
type AzureSku struct {
	// Name - SKU name
	Name *string `json:"name,omitempty"`
	// Tier - SKU tier
	Tier *string `json:"tier,omitempty"`
}

// CheckNameRequest ...
type CheckNameRequest struct {
	// Name - Workspace collection name
	Name *string `json:"name,omitempty"`
	// Type - Resource type
	Type *string `json:"type,omitempty"`
}

// CheckNameResponse ...
type CheckNameResponse struct {
	autorest.Response `json:"-"`
	// NameAvailable - Specifies a Boolean value that indicates whether the specified Power BI Workspace Collection name is available to use.
	NameAvailable *bool `json:"nameAvailable,omitempty"`
	// Reason - Reason why the workspace collection name cannot be used. Possible values include: 'Unavailable', 'Invalid'
	Reason CheckNameReason `json:"reason,omitempty"`
	// Message - Message indicating an unavailable name due to a conflict, or a description of the naming rules that are violated.
	Message *string `json:"message,omitempty"`
}

// CreateWorkspaceCollectionRequest ...
type CreateWorkspaceCollectionRequest struct {
	// Location - Azure location
	Location *string             `json:"location,omitempty"`
	Tags     *map[string]*string `json:"tags,omitempty"`
	Sku      *AzureSku           `json:"sku,omitempty"`
}

// Display ...
type Display struct {
	// Provider - The localized friendly form of the resource provider name. This form is also expected to include the publisher/company responsible. Use Title Casing. Begin with “Microsoft” for 1st party services.
	Provider *string `json:"provider,omitempty"`
	// Resource - The localized friendly form of the resource type related to this action/operation. This form should match the public documentation for the resource provider. Use Title Casing. For examples, refer to the “name” section.
	Resource *string `json:"resource,omitempty"`
	// Operation - The localized friendly name for the operation as shown to the user. This name should be concise (to fit in drop downs), but clear (self-documenting). Use Title Casing and include the entity/resource to which it applies.
	Operation *string `json:"operation,omitempty"`
	// Description - The localized friendly description for the operation as shown to the user. This description should be thorough, yet concise. It will be used in tool-tips and detailed views.
	Description *string `json:"description,omitempty"`
	// Origin - The intended executor of the operation; governs the display of the operation in the RBAC UX and the audit logs UX. Default value is 'user,system'
	Origin *string `json:"origin,omitempty"`
}

// Error ...
type Error struct {
	Code    *string        `json:"code,omitempty"`
	Message *string        `json:"message,omitempty"`
	Target  *string        `json:"target,omitempty"`
	Details *[]ErrorDetail `json:"details,omitempty"`
}

// ErrorDetail ...
type ErrorDetail struct {
	Code    *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Target  *string `json:"target,omitempty"`
}

// MigrateWorkspaceCollectionRequest ...
type MigrateWorkspaceCollectionRequest struct {
	// TargetResourceGroup - Name of the resource group the Power BI workspace collections will be migrated to.
	TargetResourceGroup *string   `json:"targetResourceGroup,omitempty"`
	Resources           *[]string `json:"resources,omitempty"`
}

// Operation ...
type Operation struct {
	// Name - The name of the operation being performed on this particular object. This name should match the action name that appears in RBAC / the event service.
	Name    *string  `json:"name,omitempty"`
	Display *Display `json:"display,omitempty"`
}

// OperationList ...
type OperationList struct {
	autorest.Response `json:"-"`
	Value             *[]Operation `json:"value,omitempty"`
}

// UpdateWorkspaceCollectionRequest ...
type UpdateWorkspaceCollectionRequest struct {
	Tags *map[string]*string `json:"tags,omitempty"`
	Sku  *AzureSku           `json:"sku,omitempty"`
}

// Workspace ...
type Workspace struct {
	// ID - Workspace id
	ID *string `json:"id,omitempty"`
	// Name - Workspace name
	Name *string `json:"name,omitempty"`
	// Type - Resource type
	Type *string `json:"type,omitempty"`
	// Properties - Property bag
	Properties *map[string]interface{} `json:"properties,omitempty"`
}

// WorkspaceCollection ...
type WorkspaceCollection struct {
	autorest.Response `json:"-"`
	// ID - Resource id
	ID *string `json:"id,omitempty"`
	// Name - Workspace collection name
	Name *string `json:"name,omitempty"`
	// Type - Resource type
	Type *string `json:"type,omitempty"`
	// Location - Azure location
	Location *string             `json:"location,omitempty"`
	Tags     *map[string]*string `json:"tags,omitempty"`
	Sku      *AzureSku           `json:"sku,omitempty"`
	// Properties - Properties
	Properties *map[string]interface{} `json:"properties,omitempty"`
}

// WorkspaceCollectionAccessKey ...
type WorkspaceCollectionAccessKey struct {
	// KeyName - Key name. Possible values include: 'Key1', 'Key2'
	KeyName AccessKeyName `json:"keyName,omitempty"`
}

// WorkspaceCollectionAccessKeys ...
type WorkspaceCollectionAccessKeys struct {
	autorest.Response `json:"-"`
	// Key1 - Access key 1
	Key1 *string `json:"key1,omitempty"`
	// Key2 - Access key 2
	Key2 *string `json:"key2,omitempty"`
}

// WorkspaceCollectionList ...
type WorkspaceCollectionList struct {
	autorest.Response `json:"-"`
	Value             *[]WorkspaceCollection `json:"value,omitempty"`
}

// WorkspaceCollectionsDeleteFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type WorkspaceCollectionsDeleteFuture struct {
	azure.Future
	req *http.Request
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future WorkspaceCollectionsDeleteFuture) Result(client WorkspaceCollectionsClient) (ar autorest.Response, err error) {
	var done bool
	done, err = future.Done(client)
	if err != nil {
		return
	}
	if !done {
		return ar, autorest.NewError("powerbiembedded.WorkspaceCollectionsDeleteFuture", "Result", "asynchronous operation has not completed")
	}
	if future.PollingMethod() == azure.PollingLocation {
		ar, err = client.DeleteResponder(future.Response())
		return
	}
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, autorest.ChangeToGet(future.req),
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	ar, err = client.DeleteResponder(resp)
	return
}

// WorkspaceList ...
type WorkspaceList struct {
	autorest.Response `json:"-"`
	Value             *[]Workspace `json:"value,omitempty"`
}