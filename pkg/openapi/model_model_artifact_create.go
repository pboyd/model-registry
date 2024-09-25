/*
Model Registry REST API

REST API for Model Registry to create and manage ML model metadata

API version: v1alpha3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ModelArtifactCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelArtifactCreate{}

// ModelArtifactCreate An ML model artifact.
type ModelArtifactCreate struct {
	ArtifactType string `json:"artifactType"`
	// User provided custom properties which are not defined by its type.
	CustomProperties *map[string]MetadataValue `json:"customProperties,omitempty"`
	// An optional description about the resource.
	Description *string `json:"description,omitempty"`
	// The external id that come from the clients’ system. This field is optional. If set, it must be unique among all resources within a database instance.
	ExternalId *string `json:"externalId,omitempty"`
	// The uniform resource identifier of the physical artifact. May be empty if there is no physical artifact.
	Uri   *string        `json:"uri,omitempty"`
	State *ArtifactState `json:"state,omitempty"`
	// The client provided name of the artifact. This field is optional. If set, it must be unique among all the artifacts of the same artifact type within a database instance and cannot be changed once set.
	Name *string `json:"name,omitempty"`
	// Name of the model format.
	ModelFormatName *string `json:"modelFormatName,omitempty"`
	// Storage secret name.
	StorageKey *string `json:"storageKey,omitempty"`
	// Path for model in storage provided by `storageKey`.
	StoragePath *string `json:"storagePath,omitempty"`
	// Version of the model format.
	ModelFormatVersion *string `json:"modelFormatVersion,omitempty"`
	// Name of the service account with storage secret.
	ServiceAccountName *string `json:"serviceAccountName,omitempty"`
}

// NewModelArtifactCreate instantiates a new ModelArtifactCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelArtifactCreate(artifactType string) *ModelArtifactCreate {
	this := ModelArtifactCreate{}
	var state ArtifactState = ARTIFACTSTATE_UNKNOWN
	this.State = &state
	return &this
}

// NewModelArtifactCreateWithDefaults instantiates a new ModelArtifactCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelArtifactCreateWithDefaults() *ModelArtifactCreate {
	this := ModelArtifactCreate{}
	var artifactType string = "model-artifact"
	this.ArtifactType = artifactType
	var state ArtifactState = ARTIFACTSTATE_UNKNOWN
	this.State = &state
	return &this
}

// GetArtifactType returns the ArtifactType field value
func (o *ModelArtifactCreate) GetArtifactType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ArtifactType
}

// GetArtifactTypeOk returns a tuple with the ArtifactType field value
// and a boolean to check if the value has been set.
func (o *ModelArtifactCreate) GetArtifactTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ArtifactType, true
}

// SetArtifactType sets field value
func (o *ModelArtifactCreate) SetArtifactType(v string) {
	o.ArtifactType = v
}

// GetCustomProperties returns the CustomProperties field value if set, zero value otherwise.
func (o *ModelArtifactCreate) GetCustomProperties() map[string]MetadataValue {
	if o == nil || IsNil(o.CustomProperties) {
		var ret map[string]MetadataValue
		return ret
	}
	return *o.CustomProperties
}

// GetCustomPropertiesOk returns a tuple with the CustomProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelArtifactCreate) GetCustomPropertiesOk() (*map[string]MetadataValue, bool) {
	if o == nil || IsNil(o.CustomProperties) {
		return nil, false
	}
	return o.CustomProperties, true
}

// HasCustomProperties returns a boolean if a field has been set.
func (o *ModelArtifactCreate) HasCustomProperties() bool {
	if o != nil && !IsNil(o.CustomProperties) {
		return true
	}

	return false
}

// SetCustomProperties gets a reference to the given map[string]MetadataValue and assigns it to the CustomProperties field.
func (o *ModelArtifactCreate) SetCustomProperties(v map[string]MetadataValue) {
	o.CustomProperties = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ModelArtifactCreate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelArtifactCreate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ModelArtifactCreate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ModelArtifactCreate) SetDescription(v string) {
	o.Description = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *ModelArtifactCreate) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelArtifactCreate) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *ModelArtifactCreate) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *ModelArtifactCreate) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *ModelArtifactCreate) GetUri() string {
	if o == nil || IsNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelArtifactCreate) GetUriOk() (*string, bool) {
	if o == nil || IsNil(o.Uri) {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *ModelArtifactCreate) HasUri() bool {
	if o != nil && !IsNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *ModelArtifactCreate) SetUri(v string) {
	o.Uri = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ModelArtifactCreate) GetState() ArtifactState {
	if o == nil || IsNil(o.State) {
		var ret ArtifactState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelArtifactCreate) GetStateOk() (*ArtifactState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ModelArtifactCreate) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given ArtifactState and assigns it to the State field.
func (o *ModelArtifactCreate) SetState(v ArtifactState) {
	o.State = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ModelArtifactCreate) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelArtifactCreate) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ModelArtifactCreate) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ModelArtifactCreate) SetName(v string) {
	o.Name = &v
}

// GetModelFormatName returns the ModelFormatName field value if set, zero value otherwise.
func (o *ModelArtifactCreate) GetModelFormatName() string {
	if o == nil || IsNil(o.ModelFormatName) {
		var ret string
		return ret
	}
	return *o.ModelFormatName
}

// GetModelFormatNameOk returns a tuple with the ModelFormatName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelArtifactCreate) GetModelFormatNameOk() (*string, bool) {
	if o == nil || IsNil(o.ModelFormatName) {
		return nil, false
	}
	return o.ModelFormatName, true
}

// HasModelFormatName returns a boolean if a field has been set.
func (o *ModelArtifactCreate) HasModelFormatName() bool {
	if o != nil && !IsNil(o.ModelFormatName) {
		return true
	}

	return false
}

// SetModelFormatName gets a reference to the given string and assigns it to the ModelFormatName field.
func (o *ModelArtifactCreate) SetModelFormatName(v string) {
	o.ModelFormatName = &v
}

// GetStorageKey returns the StorageKey field value if set, zero value otherwise.
func (o *ModelArtifactCreate) GetStorageKey() string {
	if o == nil || IsNil(o.StorageKey) {
		var ret string
		return ret
	}
	return *o.StorageKey
}

// GetStorageKeyOk returns a tuple with the StorageKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelArtifactCreate) GetStorageKeyOk() (*string, bool) {
	if o == nil || IsNil(o.StorageKey) {
		return nil, false
	}
	return o.StorageKey, true
}

// HasStorageKey returns a boolean if a field has been set.
func (o *ModelArtifactCreate) HasStorageKey() bool {
	if o != nil && !IsNil(o.StorageKey) {
		return true
	}

	return false
}

// SetStorageKey gets a reference to the given string and assigns it to the StorageKey field.
func (o *ModelArtifactCreate) SetStorageKey(v string) {
	o.StorageKey = &v
}

// GetStoragePath returns the StoragePath field value if set, zero value otherwise.
func (o *ModelArtifactCreate) GetStoragePath() string {
	if o == nil || IsNil(o.StoragePath) {
		var ret string
		return ret
	}
	return *o.StoragePath
}

// GetStoragePathOk returns a tuple with the StoragePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelArtifactCreate) GetStoragePathOk() (*string, bool) {
	if o == nil || IsNil(o.StoragePath) {
		return nil, false
	}
	return o.StoragePath, true
}

// HasStoragePath returns a boolean if a field has been set.
func (o *ModelArtifactCreate) HasStoragePath() bool {
	if o != nil && !IsNil(o.StoragePath) {
		return true
	}

	return false
}

// SetStoragePath gets a reference to the given string and assigns it to the StoragePath field.
func (o *ModelArtifactCreate) SetStoragePath(v string) {
	o.StoragePath = &v
}

// GetModelFormatVersion returns the ModelFormatVersion field value if set, zero value otherwise.
func (o *ModelArtifactCreate) GetModelFormatVersion() string {
	if o == nil || IsNil(o.ModelFormatVersion) {
		var ret string
		return ret
	}
	return *o.ModelFormatVersion
}

// GetModelFormatVersionOk returns a tuple with the ModelFormatVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelArtifactCreate) GetModelFormatVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ModelFormatVersion) {
		return nil, false
	}
	return o.ModelFormatVersion, true
}

// HasModelFormatVersion returns a boolean if a field has been set.
func (o *ModelArtifactCreate) HasModelFormatVersion() bool {
	if o != nil && !IsNil(o.ModelFormatVersion) {
		return true
	}

	return false
}

// SetModelFormatVersion gets a reference to the given string and assigns it to the ModelFormatVersion field.
func (o *ModelArtifactCreate) SetModelFormatVersion(v string) {
	o.ModelFormatVersion = &v
}

// GetServiceAccountName returns the ServiceAccountName field value if set, zero value otherwise.
func (o *ModelArtifactCreate) GetServiceAccountName() string {
	if o == nil || IsNil(o.ServiceAccountName) {
		var ret string
		return ret
	}
	return *o.ServiceAccountName
}

// GetServiceAccountNameOk returns a tuple with the ServiceAccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelArtifactCreate) GetServiceAccountNameOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceAccountName) {
		return nil, false
	}
	return o.ServiceAccountName, true
}

// HasServiceAccountName returns a boolean if a field has been set.
func (o *ModelArtifactCreate) HasServiceAccountName() bool {
	if o != nil && !IsNil(o.ServiceAccountName) {
		return true
	}

	return false
}

// SetServiceAccountName gets a reference to the given string and assigns it to the ServiceAccountName field.
func (o *ModelArtifactCreate) SetServiceAccountName(v string) {
	o.ServiceAccountName = &v
}

func (o ModelArtifactCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelArtifactCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["artifactType"] = o.ArtifactType
	if !IsNil(o.CustomProperties) {
		toSerialize["customProperties"] = o.CustomProperties
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ExternalId) {
		toSerialize["externalId"] = o.ExternalId
	}
	if !IsNil(o.Uri) {
		toSerialize["uri"] = o.Uri
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ModelFormatName) {
		toSerialize["modelFormatName"] = o.ModelFormatName
	}
	if !IsNil(o.StorageKey) {
		toSerialize["storageKey"] = o.StorageKey
	}
	if !IsNil(o.StoragePath) {
		toSerialize["storagePath"] = o.StoragePath
	}
	if !IsNil(o.ModelFormatVersion) {
		toSerialize["modelFormatVersion"] = o.ModelFormatVersion
	}
	if !IsNil(o.ServiceAccountName) {
		toSerialize["serviceAccountName"] = o.ServiceAccountName
	}
	return toSerialize, nil
}

type NullableModelArtifactCreate struct {
	value *ModelArtifactCreate
	isSet bool
}

func (v NullableModelArtifactCreate) Get() *ModelArtifactCreate {
	return v.value
}

func (v *NullableModelArtifactCreate) Set(val *ModelArtifactCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableModelArtifactCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableModelArtifactCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelArtifactCreate(val *ModelArtifactCreate) *NullableModelArtifactCreate {
	return &NullableModelArtifactCreate{value: val, isSet: true}
}

func (v NullableModelArtifactCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelArtifactCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
