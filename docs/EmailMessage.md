# EmailMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attachments** | Pointer to [**[]Attachment**](Attachment.md) |  | [optional] 
**From** | Pointer to [**From**](From.md) |  | [optional] 
**Groups** | Pointer to **[]string** |  | [optional] 
**HtmlBody** | Pointer to **string** |  | [optional] 
**Ippool** | Pointer to **string** |  | [optional] 
**PreText** | Pointer to **string** |  | [optional] 
**ReplyTo** | Pointer to [**ReplyTo**](ReplyTo.md) |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**Template** | Pointer to **string** |  | [optional] 
**TextBody** | Pointer to **string** |  | [optional] 
**To** | Pointer to [**[]To**](To.md) |  | [optional] 
**TrackClicks** | Pointer to **bool** |  | [optional] 
**TrackOpens** | Pointer to **bool** |  | [optional] 
**Headers** | Pointer to **map[string]interface{}** |  | [optional] 
**AmpBody** | Pointer to **string** |  | [optional] 

## Methods

### NewEmailMessage

`func NewEmailMessage() *EmailMessage`

NewEmailMessage instantiates a new EmailMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailMessageWithDefaults

`func NewEmailMessageWithDefaults() *EmailMessage`

NewEmailMessageWithDefaults instantiates a new EmailMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachments

`func (o *EmailMessage) GetAttachments() []Attachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *EmailMessage) GetAttachmentsOk() (*[]Attachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *EmailMessage) SetAttachments(v []Attachment)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *EmailMessage) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetFrom

`func (o *EmailMessage) GetFrom() From`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *EmailMessage) GetFromOk() (*From, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *EmailMessage) SetFrom(v From)`

SetFrom sets From field to given value.

### HasFrom

`func (o *EmailMessage) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetGroups

`func (o *EmailMessage) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *EmailMessage) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *EmailMessage) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *EmailMessage) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetHtmlBody

`func (o *EmailMessage) GetHtmlBody() string`

GetHtmlBody returns the HtmlBody field if non-nil, zero value otherwise.

### GetHtmlBodyOk

`func (o *EmailMessage) GetHtmlBodyOk() (*string, bool)`

GetHtmlBodyOk returns a tuple with the HtmlBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlBody

`func (o *EmailMessage) SetHtmlBody(v string)`

SetHtmlBody sets HtmlBody field to given value.

### HasHtmlBody

`func (o *EmailMessage) HasHtmlBody() bool`

HasHtmlBody returns a boolean if a field has been set.

### GetIppool

`func (o *EmailMessage) GetIppool() string`

GetIppool returns the Ippool field if non-nil, zero value otherwise.

### GetIppoolOk

`func (o *EmailMessage) GetIppoolOk() (*string, bool)`

GetIppoolOk returns a tuple with the Ippool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIppool

`func (o *EmailMessage) SetIppool(v string)`

SetIppool sets Ippool field to given value.

### HasIppool

`func (o *EmailMessage) HasIppool() bool`

HasIppool returns a boolean if a field has been set.

### GetPreText

`func (o *EmailMessage) GetPreText() string`

GetPreText returns the PreText field if non-nil, zero value otherwise.

### GetPreTextOk

`func (o *EmailMessage) GetPreTextOk() (*string, bool)`

GetPreTextOk returns a tuple with the PreText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreText

`func (o *EmailMessage) SetPreText(v string)`

SetPreText sets PreText field to given value.

### HasPreText

`func (o *EmailMessage) HasPreText() bool`

HasPreText returns a boolean if a field has been set.

### GetReplyTo

`func (o *EmailMessage) GetReplyTo() ReplyTo`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *EmailMessage) GetReplyToOk() (*ReplyTo, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *EmailMessage) SetReplyTo(v ReplyTo)`

SetReplyTo sets ReplyTo field to given value.

### HasReplyTo

`func (o *EmailMessage) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.

### GetSubject

`func (o *EmailMessage) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *EmailMessage) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *EmailMessage) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *EmailMessage) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTemplate

`func (o *EmailMessage) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *EmailMessage) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *EmailMessage) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *EmailMessage) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetTextBody

`func (o *EmailMessage) GetTextBody() string`

GetTextBody returns the TextBody field if non-nil, zero value otherwise.

### GetTextBodyOk

`func (o *EmailMessage) GetTextBodyOk() (*string, bool)`

GetTextBodyOk returns a tuple with the TextBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextBody

`func (o *EmailMessage) SetTextBody(v string)`

SetTextBody sets TextBody field to given value.

### HasTextBody

`func (o *EmailMessage) HasTextBody() bool`

HasTextBody returns a boolean if a field has been set.

### GetTo

`func (o *EmailMessage) GetTo() []To`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *EmailMessage) GetToOk() (*[]To, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *EmailMessage) SetTo(v []To)`

SetTo sets To field to given value.

### HasTo

`func (o *EmailMessage) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetTrackClicks

`func (o *EmailMessage) GetTrackClicks() bool`

GetTrackClicks returns the TrackClicks field if non-nil, zero value otherwise.

### GetTrackClicksOk

`func (o *EmailMessage) GetTrackClicksOk() (*bool, bool)`

GetTrackClicksOk returns a tuple with the TrackClicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackClicks

`func (o *EmailMessage) SetTrackClicks(v bool)`

SetTrackClicks sets TrackClicks field to given value.

### HasTrackClicks

`func (o *EmailMessage) HasTrackClicks() bool`

HasTrackClicks returns a boolean if a field has been set.

### GetTrackOpens

`func (o *EmailMessage) GetTrackOpens() bool`

GetTrackOpens returns the TrackOpens field if non-nil, zero value otherwise.

### GetTrackOpensOk

`func (o *EmailMessage) GetTrackOpensOk() (*bool, bool)`

GetTrackOpensOk returns a tuple with the TrackOpens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackOpens

`func (o *EmailMessage) SetTrackOpens(v bool)`

SetTrackOpens sets TrackOpens field to given value.

### HasTrackOpens

`func (o *EmailMessage) HasTrackOpens() bool`

HasTrackOpens returns a boolean if a field has been set.

### GetHeaders

`func (o *EmailMessage) GetHeaders() map[string]interface{}`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *EmailMessage) GetHeadersOk() (*map[string]interface{}, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *EmailMessage) SetHeaders(v map[string]interface{})`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *EmailMessage) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetAmpBody

`func (o *EmailMessage) GetAmpBody() string`

GetAmpBody returns the AmpBody field if non-nil, zero value otherwise.

### GetAmpBodyOk

`func (o *EmailMessage) GetAmpBodyOk() (*string, bool)`

GetAmpBodyOk returns a tuple with the AmpBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmpBody

`func (o *EmailMessage) SetAmpBody(v string)`

SetAmpBody sets AmpBody field to given value.

### HasAmpBody

`func (o *EmailMessage) HasAmpBody() bool`

HasAmpBody returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


