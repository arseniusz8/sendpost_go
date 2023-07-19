# EmailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | Pointer to **int32** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**MessageId** | Pointer to **string** |  | [optional] 
**SubmittedAt** | Pointer to **int64** |  | [optional] 
**To** | Pointer to **string** |  | [optional] 

## Methods

### NewEmailResponse

`func NewEmailResponse() *EmailResponse`

NewEmailResponse instantiates a new EmailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailResponseWithDefaults

`func NewEmailResponseWithDefaults() *EmailResponse`

NewEmailResponseWithDefaults instantiates a new EmailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *EmailResponse) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *EmailResponse) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *EmailResponse) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *EmailResponse) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetMessage

`func (o *EmailResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *EmailResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *EmailResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *EmailResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMessageId

`func (o *EmailResponse) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *EmailResponse) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *EmailResponse) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *EmailResponse) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetSubmittedAt

`func (o *EmailResponse) GetSubmittedAt() int64`

GetSubmittedAt returns the SubmittedAt field if non-nil, zero value otherwise.

### GetSubmittedAtOk

`func (o *EmailResponse) GetSubmittedAtOk() (*int64, bool)`

GetSubmittedAtOk returns a tuple with the SubmittedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedAt

`func (o *EmailResponse) SetSubmittedAt(v int64)`

SetSubmittedAt sets SubmittedAt field to given value.

### HasSubmittedAt

`func (o *EmailResponse) HasSubmittedAt() bool`

HasSubmittedAt returns a boolean if a field has been set.

### GetTo

`func (o *EmailResponse) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *EmailResponse) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *EmailResponse) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *EmailResponse) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


