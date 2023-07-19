# QEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountID** | Pointer to **int64** |  | [optional] 
**EventID** | Pointer to **string** |  | [optional] 
**EventMetadata** | Pointer to [**EventMetadata**](EventMetadata.md) |  | [optional] 
**From** | Pointer to **string** |  | [optional] 
**Groups** | Pointer to **[]string** |  | [optional] 
**IpID** | Pointer to **int64** |  | [optional] 
**MessageID** | Pointer to **string** |  | [optional] 
**MessageType** | Pointer to **string** |  | [optional] 
**SubAccountID** | Pointer to **int64** |  | [optional] 
**SubmittedAt** | Pointer to **int64** |  | [optional] 
**To** | Pointer to **string** |  | [optional] 
**TpspId** | Pointer to **int64** |  | [optional] 
**Type** | Pointer to **int64** |  | [optional] 

## Methods

### NewQEvent

`func NewQEvent() *QEvent`

NewQEvent instantiates a new QEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQEventWithDefaults

`func NewQEventWithDefaults() *QEvent`

NewQEventWithDefaults instantiates a new QEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountID

`func (o *QEvent) GetAccountID() int64`

GetAccountID returns the AccountID field if non-nil, zero value otherwise.

### GetAccountIDOk

`func (o *QEvent) GetAccountIDOk() (*int64, bool)`

GetAccountIDOk returns a tuple with the AccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountID

`func (o *QEvent) SetAccountID(v int64)`

SetAccountID sets AccountID field to given value.

### HasAccountID

`func (o *QEvent) HasAccountID() bool`

HasAccountID returns a boolean if a field has been set.

### GetEventID

`func (o *QEvent) GetEventID() string`

GetEventID returns the EventID field if non-nil, zero value otherwise.

### GetEventIDOk

`func (o *QEvent) GetEventIDOk() (*string, bool)`

GetEventIDOk returns a tuple with the EventID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventID

`func (o *QEvent) SetEventID(v string)`

SetEventID sets EventID field to given value.

### HasEventID

`func (o *QEvent) HasEventID() bool`

HasEventID returns a boolean if a field has been set.

### GetEventMetadata

`func (o *QEvent) GetEventMetadata() EventMetadata`

GetEventMetadata returns the EventMetadata field if non-nil, zero value otherwise.

### GetEventMetadataOk

`func (o *QEvent) GetEventMetadataOk() (*EventMetadata, bool)`

GetEventMetadataOk returns a tuple with the EventMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventMetadata

`func (o *QEvent) SetEventMetadata(v EventMetadata)`

SetEventMetadata sets EventMetadata field to given value.

### HasEventMetadata

`func (o *QEvent) HasEventMetadata() bool`

HasEventMetadata returns a boolean if a field has been set.

### GetFrom

`func (o *QEvent) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *QEvent) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *QEvent) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *QEvent) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetGroups

`func (o *QEvent) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *QEvent) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *QEvent) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *QEvent) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetIpID

`func (o *QEvent) GetIpID() int64`

GetIpID returns the IpID field if non-nil, zero value otherwise.

### GetIpIDOk

`func (o *QEvent) GetIpIDOk() (*int64, bool)`

GetIpIDOk returns a tuple with the IpID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpID

`func (o *QEvent) SetIpID(v int64)`

SetIpID sets IpID field to given value.

### HasIpID

`func (o *QEvent) HasIpID() bool`

HasIpID returns a boolean if a field has been set.

### GetMessageID

`func (o *QEvent) GetMessageID() string`

GetMessageID returns the MessageID field if non-nil, zero value otherwise.

### GetMessageIDOk

`func (o *QEvent) GetMessageIDOk() (*string, bool)`

GetMessageIDOk returns a tuple with the MessageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageID

`func (o *QEvent) SetMessageID(v string)`

SetMessageID sets MessageID field to given value.

### HasMessageID

`func (o *QEvent) HasMessageID() bool`

HasMessageID returns a boolean if a field has been set.

### GetMessageType

`func (o *QEvent) GetMessageType() string`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *QEvent) GetMessageTypeOk() (*string, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *QEvent) SetMessageType(v string)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *QEvent) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.

### GetSubAccountID

`func (o *QEvent) GetSubAccountID() int64`

GetSubAccountID returns the SubAccountID field if non-nil, zero value otherwise.

### GetSubAccountIDOk

`func (o *QEvent) GetSubAccountIDOk() (*int64, bool)`

GetSubAccountIDOk returns a tuple with the SubAccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubAccountID

`func (o *QEvent) SetSubAccountID(v int64)`

SetSubAccountID sets SubAccountID field to given value.

### HasSubAccountID

`func (o *QEvent) HasSubAccountID() bool`

HasSubAccountID returns a boolean if a field has been set.

### GetSubmittedAt

`func (o *QEvent) GetSubmittedAt() int64`

GetSubmittedAt returns the SubmittedAt field if non-nil, zero value otherwise.

### GetSubmittedAtOk

`func (o *QEvent) GetSubmittedAtOk() (*int64, bool)`

GetSubmittedAtOk returns a tuple with the SubmittedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedAt

`func (o *QEvent) SetSubmittedAt(v int64)`

SetSubmittedAt sets SubmittedAt field to given value.

### HasSubmittedAt

`func (o *QEvent) HasSubmittedAt() bool`

HasSubmittedAt returns a boolean if a field has been set.

### GetTo

`func (o *QEvent) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *QEvent) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *QEvent) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *QEvent) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetTpspId

`func (o *QEvent) GetTpspId() int64`

GetTpspId returns the TpspId field if non-nil, zero value otherwise.

### GetTpspIdOk

`func (o *QEvent) GetTpspIdOk() (*int64, bool)`

GetTpspIdOk returns a tuple with the TpspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpspId

`func (o *QEvent) SetTpspId(v int64)`

SetTpspId sets TpspId field to given value.

### HasTpspId

`func (o *QEvent) HasTpspId() bool`

HasTpspId returns a boolean if a field has been set.

### GetType

`func (o *QEvent) GetType() int64`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *QEvent) GetTypeOk() (*int64, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *QEvent) SetType(v int64)`

SetType sets Type field to given value.

### HasType

`func (o *QEvent) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


