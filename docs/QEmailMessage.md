# QEmailMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountID** | Pointer to **int64** |  | [optional] 
**AmpBody** | Pointer to **string** |  | [optional] 
**Attachments** | Pointer to [**[]Attachment**](Attachment.md) |  | [optional] 
**Attempt** | Pointer to **int64** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**EmailType** | Pointer to **string** |  | [optional] 
**From** | Pointer to [**From**](From.md) |  | [optional] 
**Groups** | Pointer to **[]string** |  | [optional] 
**HeaderBcc** | Pointer to [**[]CopyTo**](CopyTo.md) |  | [optional] 
**HeaderCc** | Pointer to [**[]CopyTo**](CopyTo.md) |  | [optional] 
**HeaderTo** | Pointer to [**To**](To.md) |  | [optional] 
**Headers** | Pointer to **map[string]interface{}** |  | [optional] 
**HtmlBody** | Pointer to **string** |  | [optional] 
**IpID** | Pointer to **int64** |  | [optional] 
**IpPool** | Pointer to **string** |  | [optional] 
**LocalIP** | Pointer to **string** |  | [optional] 
**MessageID** | Pointer to **string** |  | [optional] 
**PreText** | Pointer to **string** |  | [optional] 
**PublicIP** | Pointer to **string** |  | [optional] 
**ReplyTo** | Pointer to [**ReplyTo**](ReplyTo.md) |  | [optional] 
**SubAccountID** | Pointer to **int64** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**SubmittedAt** | Pointer to **int64** |  | [optional] 
**TextBody** | Pointer to **string** |  | [optional] 
**To** | Pointer to [**To**](To.md) |  | [optional] 
**TrackClicks** | Pointer to **bool** |  | [optional] 
**TrackOpens** | Pointer to **bool** |  | [optional] 

## Methods

### NewQEmailMessage

`func NewQEmailMessage() *QEmailMessage`

NewQEmailMessage instantiates a new QEmailMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQEmailMessageWithDefaults

`func NewQEmailMessageWithDefaults() *QEmailMessage`

NewQEmailMessageWithDefaults instantiates a new QEmailMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountID

`func (o *QEmailMessage) GetAccountID() int64`

GetAccountID returns the AccountID field if non-nil, zero value otherwise.

### GetAccountIDOk

`func (o *QEmailMessage) GetAccountIDOk() (*int64, bool)`

GetAccountIDOk returns a tuple with the AccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountID

`func (o *QEmailMessage) SetAccountID(v int64)`

SetAccountID sets AccountID field to given value.

### HasAccountID

`func (o *QEmailMessage) HasAccountID() bool`

HasAccountID returns a boolean if a field has been set.

### GetAmpBody

`func (o *QEmailMessage) GetAmpBody() string`

GetAmpBody returns the AmpBody field if non-nil, zero value otherwise.

### GetAmpBodyOk

`func (o *QEmailMessage) GetAmpBodyOk() (*string, bool)`

GetAmpBodyOk returns a tuple with the AmpBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmpBody

`func (o *QEmailMessage) SetAmpBody(v string)`

SetAmpBody sets AmpBody field to given value.

### HasAmpBody

`func (o *QEmailMessage) HasAmpBody() bool`

HasAmpBody returns a boolean if a field has been set.

### GetAttachments

`func (o *QEmailMessage) GetAttachments() []Attachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *QEmailMessage) GetAttachmentsOk() (*[]Attachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *QEmailMessage) SetAttachments(v []Attachment)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *QEmailMessage) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetAttempt

`func (o *QEmailMessage) GetAttempt() int64`

GetAttempt returns the Attempt field if non-nil, zero value otherwise.

### GetAttemptOk

`func (o *QEmailMessage) GetAttemptOk() (*int64, bool)`

GetAttemptOk returns a tuple with the Attempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempt

`func (o *QEmailMessage) SetAttempt(v int64)`

SetAttempt sets Attempt field to given value.

### HasAttempt

`func (o *QEmailMessage) HasAttempt() bool`

HasAttempt returns a boolean if a field has been set.

### GetCustomFields

`func (o *QEmailMessage) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *QEmailMessage) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *QEmailMessage) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *QEmailMessage) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetEmailType

`func (o *QEmailMessage) GetEmailType() string`

GetEmailType returns the EmailType field if non-nil, zero value otherwise.

### GetEmailTypeOk

`func (o *QEmailMessage) GetEmailTypeOk() (*string, bool)`

GetEmailTypeOk returns a tuple with the EmailType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailType

`func (o *QEmailMessage) SetEmailType(v string)`

SetEmailType sets EmailType field to given value.

### HasEmailType

`func (o *QEmailMessage) HasEmailType() bool`

HasEmailType returns a boolean if a field has been set.

### GetFrom

`func (o *QEmailMessage) GetFrom() From`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *QEmailMessage) GetFromOk() (*From, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *QEmailMessage) SetFrom(v From)`

SetFrom sets From field to given value.

### HasFrom

`func (o *QEmailMessage) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetGroups

`func (o *QEmailMessage) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *QEmailMessage) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *QEmailMessage) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *QEmailMessage) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetHeaderBcc

`func (o *QEmailMessage) GetHeaderBcc() []CopyTo`

GetHeaderBcc returns the HeaderBcc field if non-nil, zero value otherwise.

### GetHeaderBccOk

`func (o *QEmailMessage) GetHeaderBccOk() (*[]CopyTo, bool)`

GetHeaderBccOk returns a tuple with the HeaderBcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderBcc

`func (o *QEmailMessage) SetHeaderBcc(v []CopyTo)`

SetHeaderBcc sets HeaderBcc field to given value.

### HasHeaderBcc

`func (o *QEmailMessage) HasHeaderBcc() bool`

HasHeaderBcc returns a boolean if a field has been set.

### GetHeaderCc

`func (o *QEmailMessage) GetHeaderCc() []CopyTo`

GetHeaderCc returns the HeaderCc field if non-nil, zero value otherwise.

### GetHeaderCcOk

`func (o *QEmailMessage) GetHeaderCcOk() (*[]CopyTo, bool)`

GetHeaderCcOk returns a tuple with the HeaderCc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderCc

`func (o *QEmailMessage) SetHeaderCc(v []CopyTo)`

SetHeaderCc sets HeaderCc field to given value.

### HasHeaderCc

`func (o *QEmailMessage) HasHeaderCc() bool`

HasHeaderCc returns a boolean if a field has been set.

### GetHeaderTo

`func (o *QEmailMessage) GetHeaderTo() To`

GetHeaderTo returns the HeaderTo field if non-nil, zero value otherwise.

### GetHeaderToOk

`func (o *QEmailMessage) GetHeaderToOk() (*To, bool)`

GetHeaderToOk returns a tuple with the HeaderTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderTo

`func (o *QEmailMessage) SetHeaderTo(v To)`

SetHeaderTo sets HeaderTo field to given value.

### HasHeaderTo

`func (o *QEmailMessage) HasHeaderTo() bool`

HasHeaderTo returns a boolean if a field has been set.

### GetHeaders

`func (o *QEmailMessage) GetHeaders() map[string]interface{}`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *QEmailMessage) GetHeadersOk() (*map[string]interface{}, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *QEmailMessage) SetHeaders(v map[string]interface{})`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *QEmailMessage) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetHtmlBody

`func (o *QEmailMessage) GetHtmlBody() string`

GetHtmlBody returns the HtmlBody field if non-nil, zero value otherwise.

### GetHtmlBodyOk

`func (o *QEmailMessage) GetHtmlBodyOk() (*string, bool)`

GetHtmlBodyOk returns a tuple with the HtmlBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlBody

`func (o *QEmailMessage) SetHtmlBody(v string)`

SetHtmlBody sets HtmlBody field to given value.

### HasHtmlBody

`func (o *QEmailMessage) HasHtmlBody() bool`

HasHtmlBody returns a boolean if a field has been set.

### GetIpID

`func (o *QEmailMessage) GetIpID() int64`

GetIpID returns the IpID field if non-nil, zero value otherwise.

### GetIpIDOk

`func (o *QEmailMessage) GetIpIDOk() (*int64, bool)`

GetIpIDOk returns a tuple with the IpID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpID

`func (o *QEmailMessage) SetIpID(v int64)`

SetIpID sets IpID field to given value.

### HasIpID

`func (o *QEmailMessage) HasIpID() bool`

HasIpID returns a boolean if a field has been set.

### GetIpPool

`func (o *QEmailMessage) GetIpPool() string`

GetIpPool returns the IpPool field if non-nil, zero value otherwise.

### GetIpPoolOk

`func (o *QEmailMessage) GetIpPoolOk() (*string, bool)`

GetIpPoolOk returns a tuple with the IpPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPool

`func (o *QEmailMessage) SetIpPool(v string)`

SetIpPool sets IpPool field to given value.

### HasIpPool

`func (o *QEmailMessage) HasIpPool() bool`

HasIpPool returns a boolean if a field has been set.

### GetLocalIP

`func (o *QEmailMessage) GetLocalIP() string`

GetLocalIP returns the LocalIP field if non-nil, zero value otherwise.

### GetLocalIPOk

`func (o *QEmailMessage) GetLocalIPOk() (*string, bool)`

GetLocalIPOk returns a tuple with the LocalIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalIP

`func (o *QEmailMessage) SetLocalIP(v string)`

SetLocalIP sets LocalIP field to given value.

### HasLocalIP

`func (o *QEmailMessage) HasLocalIP() bool`

HasLocalIP returns a boolean if a field has been set.

### GetMessageID

`func (o *QEmailMessage) GetMessageID() string`

GetMessageID returns the MessageID field if non-nil, zero value otherwise.

### GetMessageIDOk

`func (o *QEmailMessage) GetMessageIDOk() (*string, bool)`

GetMessageIDOk returns a tuple with the MessageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageID

`func (o *QEmailMessage) SetMessageID(v string)`

SetMessageID sets MessageID field to given value.

### HasMessageID

`func (o *QEmailMessage) HasMessageID() bool`

HasMessageID returns a boolean if a field has been set.

### GetPreText

`func (o *QEmailMessage) GetPreText() string`

GetPreText returns the PreText field if non-nil, zero value otherwise.

### GetPreTextOk

`func (o *QEmailMessage) GetPreTextOk() (*string, bool)`

GetPreTextOk returns a tuple with the PreText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreText

`func (o *QEmailMessage) SetPreText(v string)`

SetPreText sets PreText field to given value.

### HasPreText

`func (o *QEmailMessage) HasPreText() bool`

HasPreText returns a boolean if a field has been set.

### GetPublicIP

`func (o *QEmailMessage) GetPublicIP() string`

GetPublicIP returns the PublicIP field if non-nil, zero value otherwise.

### GetPublicIPOk

`func (o *QEmailMessage) GetPublicIPOk() (*string, bool)`

GetPublicIPOk returns a tuple with the PublicIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIP

`func (o *QEmailMessage) SetPublicIP(v string)`

SetPublicIP sets PublicIP field to given value.

### HasPublicIP

`func (o *QEmailMessage) HasPublicIP() bool`

HasPublicIP returns a boolean if a field has been set.

### GetReplyTo

`func (o *QEmailMessage) GetReplyTo() ReplyTo`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *QEmailMessage) GetReplyToOk() (*ReplyTo, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *QEmailMessage) SetReplyTo(v ReplyTo)`

SetReplyTo sets ReplyTo field to given value.

### HasReplyTo

`func (o *QEmailMessage) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.

### GetSubAccountID

`func (o *QEmailMessage) GetSubAccountID() int64`

GetSubAccountID returns the SubAccountID field if non-nil, zero value otherwise.

### GetSubAccountIDOk

`func (o *QEmailMessage) GetSubAccountIDOk() (*int64, bool)`

GetSubAccountIDOk returns a tuple with the SubAccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubAccountID

`func (o *QEmailMessage) SetSubAccountID(v int64)`

SetSubAccountID sets SubAccountID field to given value.

### HasSubAccountID

`func (o *QEmailMessage) HasSubAccountID() bool`

HasSubAccountID returns a boolean if a field has been set.

### GetSubject

`func (o *QEmailMessage) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *QEmailMessage) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *QEmailMessage) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *QEmailMessage) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetSubmittedAt

`func (o *QEmailMessage) GetSubmittedAt() int64`

GetSubmittedAt returns the SubmittedAt field if non-nil, zero value otherwise.

### GetSubmittedAtOk

`func (o *QEmailMessage) GetSubmittedAtOk() (*int64, bool)`

GetSubmittedAtOk returns a tuple with the SubmittedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedAt

`func (o *QEmailMessage) SetSubmittedAt(v int64)`

SetSubmittedAt sets SubmittedAt field to given value.

### HasSubmittedAt

`func (o *QEmailMessage) HasSubmittedAt() bool`

HasSubmittedAt returns a boolean if a field has been set.

### GetTextBody

`func (o *QEmailMessage) GetTextBody() string`

GetTextBody returns the TextBody field if non-nil, zero value otherwise.

### GetTextBodyOk

`func (o *QEmailMessage) GetTextBodyOk() (*string, bool)`

GetTextBodyOk returns a tuple with the TextBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextBody

`func (o *QEmailMessage) SetTextBody(v string)`

SetTextBody sets TextBody field to given value.

### HasTextBody

`func (o *QEmailMessage) HasTextBody() bool`

HasTextBody returns a boolean if a field has been set.

### GetTo

`func (o *QEmailMessage) GetTo() To`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *QEmailMessage) GetToOk() (*To, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *QEmailMessage) SetTo(v To)`

SetTo sets To field to given value.

### HasTo

`func (o *QEmailMessage) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetTrackClicks

`func (o *QEmailMessage) GetTrackClicks() bool`

GetTrackClicks returns the TrackClicks field if non-nil, zero value otherwise.

### GetTrackClicksOk

`func (o *QEmailMessage) GetTrackClicksOk() (*bool, bool)`

GetTrackClicksOk returns a tuple with the TrackClicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackClicks

`func (o *QEmailMessage) SetTrackClicks(v bool)`

SetTrackClicks sets TrackClicks field to given value.

### HasTrackClicks

`func (o *QEmailMessage) HasTrackClicks() bool`

HasTrackClicks returns a boolean if a field has been set.

### GetTrackOpens

`func (o *QEmailMessage) GetTrackOpens() bool`

GetTrackOpens returns the TrackOpens field if non-nil, zero value otherwise.

### GetTrackOpensOk

`func (o *QEmailMessage) GetTrackOpensOk() (*bool, bool)`

GetTrackOpensOk returns a tuple with the TrackOpens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackOpens

`func (o *QEmailMessage) SetTrackOpens(v bool)`

SetTrackOpens sets TrackOpens field to given value.

### HasTrackOpens

`func (o *QEmailMessage) HasTrackOpens() bool`

HasTrackOpens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


