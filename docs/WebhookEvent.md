# WebhookEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | Pointer to [**QEvent**](QEvent.md) |  | [optional] 
**EmailMessage** | Pointer to [**QEmailMessage**](QEmailMessage.md) |  | [optional] 

## Methods

### NewWebhookEvent

`func NewWebhookEvent() *WebhookEvent`

NewWebhookEvent instantiates a new WebhookEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookEventWithDefaults

`func NewWebhookEventWithDefaults() *WebhookEvent`

NewWebhookEventWithDefaults instantiates a new WebhookEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *WebhookEvent) GetEvent() QEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *WebhookEvent) GetEventOk() (*QEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *WebhookEvent) SetEvent(v QEvent)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *WebhookEvent) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetEmailMessage

`func (o *WebhookEvent) GetEmailMessage() QEmailMessage`

GetEmailMessage returns the EmailMessage field if non-nil, zero value otherwise.

### GetEmailMessageOk

`func (o *WebhookEvent) GetEmailMessageOk() (*QEmailMessage, bool)`

GetEmailMessageOk returns a tuple with the EmailMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailMessage

`func (o *WebhookEvent) SetEmailMessage(v QEmailMessage)`

SetEmailMessage sets EmailMessage field to given value.

### HasEmailMessage

`func (o *WebhookEvent) HasEmailMessage() bool`

HasEmailMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


