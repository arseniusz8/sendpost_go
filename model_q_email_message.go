/*
SendPost API

Email API and SMTP relay to not just send and measure email sending, but also alert and optimise. We provide you with tools, expertise and support needed to reliably deliver emails to your customers inboxes on time, every time.

API version: 1.0.0
Contact: hello@sendpost.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sendpost

import (
	"encoding/json"
)

// checks if the QEmailMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QEmailMessage{}

// QEmailMessage struct for QEmailMessage
type QEmailMessage struct {
	AccountID *int64 `json:"accountID,omitempty"`
	AmpBody *string `json:"ampBody,omitempty"`
	Attachments []Attachment `json:"attachments,omitempty"`
	Attempt *int64 `json:"attempt,omitempty"`
	CustomFields map[string]interface{} `json:"customFields,omitempty"`
	EmailType *string `json:"emailType,omitempty"`
	From *From `json:"from,omitempty"`
	Groups []string `json:"groups,omitempty"`
	HeaderBcc []CopyTo `json:"headerBcc,omitempty"`
	HeaderCc []CopyTo `json:"headerCc,omitempty"`
	HeaderTo *To `json:"headerTo,omitempty"`
	Headers map[string]interface{} `json:"headers,omitempty"`
	HtmlBody *string `json:"htmlBody,omitempty"`
	IpID *int64 `json:"ipID,omitempty"`
	IpPool *string `json:"ipPool,omitempty"`
	LocalIP *string `json:"localIP,omitempty"`
	MessageID *string `json:"messageID,omitempty"`
	PreText *string `json:"preText,omitempty"`
	PublicIP *string `json:"publicIP,omitempty"`
	ReplyTo *ReplyTo `json:"replyTo,omitempty"`
	SubAccountID *int64 `json:"subAccountID,omitempty"`
	Subject *string `json:"subject,omitempty"`
	SubmittedAt *int64 `json:"submittedAt,omitempty"`
	TextBody *string `json:"textBody,omitempty"`
	To *To `json:"to,omitempty"`
	TrackClicks *bool `json:"trackClicks,omitempty"`
	TrackOpens *bool `json:"trackOpens,omitempty"`
}

// NewQEmailMessage instantiates a new QEmailMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQEmailMessage() *QEmailMessage {
	this := QEmailMessage{}
	return &this
}

// NewQEmailMessageWithDefaults instantiates a new QEmailMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQEmailMessageWithDefaults() *QEmailMessage {
	this := QEmailMessage{}
	return &this
}

// GetAccountID returns the AccountID field value if set, zero value otherwise.
func (o *QEmailMessage) GetAccountID() int64 {
	if o == nil || IsNil(o.AccountID) {
		var ret int64
		return ret
	}
	return *o.AccountID
}

// GetAccountIDOk returns a tuple with the AccountID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QEmailMessage) GetAccountIDOk() (*int64, bool) {
	if o == nil || IsNil(o.AccountID) {
		return nil, false
	}
	return o.AccountID, true
}

// HasAccountID returns a boolean if a field has been set.
func (o *QEmailMessage) HasAccountID() bool {
	if o != nil && !IsNil(o.AccountID) {
		return true
	}

	return false
}

// SetAccountID gets a reference to the given int64 and assigns it to the AccountID field.
func (o *QEmailMessage) SetAccountID(v int64) {
	o.AccountID = &v
}

// GetAmpBody returns the AmpBody field value if set, zero value otherwise.
func (o *QEmailMessage) GetAmpBody() string {
	if o == nil || IsNil(o.AmpBody) {
		var ret string
		return ret
	}
	return *o.AmpBody
}

// GetAmpBodyOk returns a tuple with the AmpBody field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QEmailMessage) GetAmpBodyOk() (*string, bool) {
	if o == nil || IsNil(o.AmpBody) {
		return nil, false
	}
	return o.AmpBody, true
}

// HasAmpBody returns a boolean if a field has been set.
func (o *QEmailMessage) HasAmpBody() bool {
	if o != nil && !IsNil(o.AmpBody) {
		return true
	}

	return false
}

// SetAmpBody gets a reference to the given string and assigns it to the AmpBody field.
func (o *QEmailMessage) SetAmpBody(v string) {
	o.AmpBody = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *QEmailMessage) GetAttachments() []Attachment {
	if o == nil || IsNil(o.Attachments) {
		var ret []Attachment
		return ret
	}
	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QEmailMessage) GetAttachmentsOk() ([]Attachment, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *QEmailMessage) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []Attachment and assigns it to the Attachments field.
func (o *QEmailMessage) SetAttachments(v []Attachment) {
	o.Attachments = v
}

// GetAttempt returns the Attempt field value if set, zero value otherwise.
func (o *QEmailMessage) GetAttempt() int64 {
	if o == nil || IsNil(o.Attempt) {
		var ret int64
		return ret
	}
	return *o.Attempt
}

// GetAttemptOk returns a tuple with the Attempt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QEmailMessage) GetAttemptOk() (*int64, bool) {
	if o == nil || IsNil(o.Attempt) {
		return nil, false
	}
	return o.Attempt, true
}

// HasAttempt returns a boolean if a field has been set.
func (o *QEmailMessage) HasAttempt() bool {
	if o != nil && !IsNil(o.Attempt) {
		return true
	}

	return false
}

// SetAttempt gets a reference to the given int64 and assigns it to the Attempt field.
func (o *QEmailMessage) SetAttempt(v int64) {
	o.Attempt = &v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *QEmailMessage) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QEmailMessage) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *QEmailMessage) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *QEmailMessage) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetEmailType returns the EmailType field value if set, zero value otherwise.
func (o *QEmailMessage) GetEmailType() string {
	if o == nil || IsNil(o.EmailType) {
		var ret string
		return ret
	}
	return *o.EmailType
}

// GetEmailTypeOk returns a tuple with the EmailType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QEmailMessage) GetEmailTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EmailType) {
		return nil, false
	}
	return o.EmailType, true
}

// HasEmailType returns a boolean if a field has been set.
func (o *QEmailMessage) HasEmailType() bool {
	if o != nil && !IsNil(o.EmailType) {
		return true
	}

	return false
}

// SetEmailType gets a reference to the given string and assigns it to the EmailType field.
func (o *QEmailMessage) SetEmailType(v string) {
	o.EmailType = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *QEmailMessage) GetFrom() From {
	if o == nil || IsNil(o.From) {
		var ret From
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QEmailMessage) GetFromOk() (*From, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *QEmailMessage) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given From and assigns it to the From field.
func (o *QEmailMessage) SetFrom(v From) {
	o.From = &v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *QEmailMessage) GetGroups() []string {
	if o == nil || IsNil(o.Groups) {
		var ret []string
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QEmailMessage) GetGroupsOk() ([]string, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *QEmailMessage) HasGroups() bool {
	if o != nil && !IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []string and assigns it to the Groups field.
func (o *QEmailMessage) SetGroups(v []string) {
	o.Groups = v
}

// GetHeaderBcc returns the HeaderBcc field value if set, zero value otherwise.
func (o *QEmailMessage) GetHeaderBcc() []CopyTo {
	if o == nil || IsNil(o.HeaderBcc) {
		var ret []CopyTo
		return ret
	}
	return o.HeaderBcc
}

// GetHeaderBccOk returns a tuple with the HeaderBcc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QEmailMessage) GetHeaderBccOk() ([]CopyTo, bool) {
	if o == nil || IsNil(o.HeaderBcc) {
		return nil, false
	}
	return o.HeaderBcc, true
}

// HasHeaderBcc returns a boolean if a field has been set.
func (o *QEmailMessage) HasHeaderBcc() bool {
	if o != nil && !IsNil(o.HeaderBcc) {
		return true
	}

	return false
}

// SetHeaderBcc gets a reference to the given []CopyTo and assigns it to the HeaderBcc field.
func (o *QEmailMessage) SetHeaderBcc(v []CopyTo) {
	o.HeaderBcc = v
}

// GetHeaderCc returns the HeaderCc field value if set, zero value otherwise.
func (o *QEmailMessage) GetHeaderCc() []CopyTo {
	if o == nil || IsNil(o.HeaderCc) {
		var ret []CopyTo
		return ret
	}
	return o.HeaderCc
}

// GetHeaderCcOk returns a tuple with the HeaderCc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QEmailMessage) GetHeaderCcOk() ([]CopyTo, bool) {
	if o == nil || IsNil(o.HeaderCc) {
		return nil, false
	}
	return o.HeaderCc, true
}

// HasHeaderCc returns a boolean if a field has been set.
func (o *QEmailMessage) HasHeaderCc() bool {
	if o != nil && !IsNil(o.HeaderCc) {
		return true
	}

	return false
}

// SetHeaderCc gets a reference to the given []CopyTo and assigns it to the HeaderCc field.
func (o *QEmailMessage) SetHeaderCc(v []CopyTo) {
	o.HeaderCc = v
}

// GetHeaderTo returns the HeaderTo field value if set, zero value otherwise.
func (o *QEmailMessage) GetHeaderTo() To {
	if o == nil || IsNil(o.HeaderTo) {
		var ret To
		return ret
	}
	return *o.HeaderTo
}

// GetHeaderToOk returns a tuple with the HeaderTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QEmailMessage) GetHeaderToOk() (*To, bool) {
	if o == nil || IsNil(o.HeaderTo) {
		return nil, false
	}
	return o.HeaderTo, true
}

// HasHeaderTo returns a boolean if a field has been set.
func (o *QEmailMessage) HasHeaderTo() bool {
	if o != nil && !IsNil(o.HeaderTo) {
		return true
	}

	return false
}

// SetHeaderTo gets a reference to the given To and assigns it to the HeaderTo field.
func (o *QEmailMessage) SetHeaderTo(v To) {
	o.HeaderTo = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *QEmailMessage) GetHeaders() map[string]interface{} {
	if o == nil || IsNil(o.Headers) {
		var ret map[string]interface{}
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QEmailMessage) GetHeadersOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Headers) {
		return map[string]interface{}{}, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *QEmailMessage) HasHeaders() bool {
	if o != nil && !IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given map[string]interface{} and assigns it to the Headers field.
func (o *QEmailMessage) SetHeaders(v map[string]interface{}) {
	o.Headers = v
}

// GetHtmlBody returns the HtmlBody field value if set, zero value otherwise.
func (o *QEmailMessage) GetHtmlBody() string {
	if o == nil || IsNil(o.HtmlBody) {
		var ret string
		return ret
	}
	return *o.HtmlBody
}

// GetHtmlBodyOk returns a tuple with the HtmlBody field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QEmailMessage) GetHtmlBodyOk() (*string, bool) {
	if o == nil || IsNil(o.HtmlBody) {
		return nil, false
	}
	return o.HtmlBody, true
}

// HasHtmlBody returns a boolean if a field has been set.
func (o *QEmailMessage) HasHtmlBody() bool {
	if o != nil && !IsNil(o.HtmlBody) {
		return true
	}

	return false
}

// SetHtmlBody gets a reference to the given string and assigns it to the HtmlBody field.
func (o *QEmailMessage) SetHtmlBody(v string) {
	o.HtmlBody = &v
}

// GetIpID returns the IpID field value if set, zero value otherwise.
func (o *QEmailMessage) GetIpID() int64 {
	if o == nil || IsNil(o.IpID) {
		var ret int64
		return ret
	}
	return *o.IpID
}

// GetIpIDOk returns a tuple with the IpID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QEmailMessage) GetIpIDOk() (*int64, bool) {
	if o == nil || IsNil(o.IpID) {
		return nil, false
	}
	return o.IpID, true
}

// HasIpID returns a boolean if a field has been set.
func (o *QEmailMessage) HasIpID() bool {
	if o != nil && !IsNil(o.IpID) {
		return true
	}

	return false
}

// SetIpID gets a reference to the given int64 and assigns it to the IpID field.
func (o *QEmailMessage) SetIpID(v int64) {
	o.IpID = &v
}

// GetIpPool returns the IpPool field value if set, zero value otherwise.
func (o *QEmailMessage) GetIpPool() string {
	if o == nil || IsNil(o.IpPool) {
		var ret string
		return ret
	}
	return *o.IpPool
}

// GetIpPoolOk returns a tuple with the IpPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QEmailMessage) GetIpPoolOk() (*string, bool) {
	if o == nil || IsNil(o.IpPool) {
		return nil, false
	}
	return o.IpPool, true
}

// HasIpPool returns a boolean if a field has been set.
func (o *QEmailMessage) HasIpPool() bool {
	if o != nil && !IsNil(o.IpPool) {
		return true
	}

	return false
}

// SetIpPool gets a reference to the given string and assigns it to the IpPool field.
func (o *QEmailMessage) SetIpPool(v string) {
	o.IpPool = &v
}

// GetLocalIP returns the LocalIP field value if set, zero value otherwise.
func (o *QEmailMessage) GetLocalIP() string {
	if o == nil || IsNil(o.LocalIP) {
		var ret string
		return ret
	}
	return *o.LocalIP
}

// GetLocalIPOk returns a tuple with the LocalIP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QEmailMessage) GetLocalIPOk() (*string, bool) {
	if o == nil || IsNil(o.LocalIP) {
		return nil, false
	}
	return o.LocalIP, true
}

// HasLocalIP returns a boolean if a field has been set.
func (o *QEmailMessage) HasLocalIP() bool {
	if o != nil && !IsNil(o.LocalIP) {
		return true
	}

	return false
}

// SetLocalIP gets a reference to the given string and assigns it to the LocalIP field.
func (o *QEmailMessage) SetLocalIP(v string) {
	o.LocalIP = &v
}

// GetMessageID returns the MessageID field value if set, zero value otherwise.
func (o *QEmailMessage) GetMessageID() string {
	if o == nil || IsNil(o.MessageID) {
		var ret string
		return ret
	}
	return *o.MessageID
}

// GetMessageIDOk returns a tuple with the MessageID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QEmailMessage) GetMessageIDOk() (*string, bool) {
	if o == nil || IsNil(o.MessageID) {
		return nil, false
	}
	return o.MessageID, true
}

// HasMessageID returns a boolean if a field has been set.
func (o *QEmailMessage) HasMessageID() bool {
	if o != nil && !IsNil(o.MessageID) {
		return true
	}

	return false
}

// SetMessageID gets a reference to the given string and assigns it to the MessageID field.
func (o *QEmailMessage) SetMessageID(v string) {
	o.MessageID = &v
}

// GetPreText returns the PreText field value if set, zero value otherwise.
func (o *QEmailMessage) GetPreText() string {
	if o == nil || IsNil(o.PreText) {
		var ret string
		return ret
	}
	return *o.PreText
}

// GetPreTextOk returns a tuple with the PreText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QEmailMessage) GetPreTextOk() (*string, bool) {
	if o == nil || IsNil(o.PreText) {
		return nil, false
	}
	return o.PreText, true
}

// HasPreText returns a boolean if a field has been set.
func (o *QEmailMessage) HasPreText() bool {
	if o != nil && !IsNil(o.PreText) {
		return true
	}

	return false
}

// SetPreText gets a reference to the given string and assigns it to the PreText field.
func (o *QEmailMessage) SetPreText(v string) {
	o.PreText = &v
}

// GetPublicIP returns the PublicIP field value if set, zero value otherwise.
func (o *QEmailMessage) GetPublicIP() string {
	if o == nil || IsNil(o.PublicIP) {
		var ret string
		return ret
	}
	return *o.PublicIP
}

// GetPublicIPOk returns a tuple with the PublicIP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QEmailMessage) GetPublicIPOk() (*string, bool) {
	if o == nil || IsNil(o.PublicIP) {
		return nil, false
	}
	return o.PublicIP, true
}

// HasPublicIP returns a boolean if a field has been set.
func (o *QEmailMessage) HasPublicIP() bool {
	if o != nil && !IsNil(o.PublicIP) {
		return true
	}

	return false
}

// SetPublicIP gets a reference to the given string and assigns it to the PublicIP field.
func (o *QEmailMessage) SetPublicIP(v string) {
	o.PublicIP = &v
}

// GetReplyTo returns the ReplyTo field value if set, zero value otherwise.
func (o *QEmailMessage) GetReplyTo() ReplyTo {
	if o == nil || IsNil(o.ReplyTo) {
		var ret ReplyTo
		return ret
	}
	return *o.ReplyTo
}

// GetReplyToOk returns a tuple with the ReplyTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QEmailMessage) GetReplyToOk() (*ReplyTo, bool) {
	if o == nil || IsNil(o.ReplyTo) {
		return nil, false
	}
	return o.ReplyTo, true
}

// HasReplyTo returns a boolean if a field has been set.
func (o *QEmailMessage) HasReplyTo() bool {
	if o != nil && !IsNil(o.ReplyTo) {
		return true
	}

	return false
}

// SetReplyTo gets a reference to the given ReplyTo and assigns it to the ReplyTo field.
func (o *QEmailMessage) SetReplyTo(v ReplyTo) {
	o.ReplyTo = &v
}

// GetSubAccountID returns the SubAccountID field value if set, zero value otherwise.
func (o *QEmailMessage) GetSubAccountID() int64 {
	if o == nil || IsNil(o.SubAccountID) {
		var ret int64
		return ret
	}
	return *o.SubAccountID
}

// GetSubAccountIDOk returns a tuple with the SubAccountID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QEmailMessage) GetSubAccountIDOk() (*int64, bool) {
	if o == nil || IsNil(o.SubAccountID) {
		return nil, false
	}
	return o.SubAccountID, true
}

// HasSubAccountID returns a boolean if a field has been set.
func (o *QEmailMessage) HasSubAccountID() bool {
	if o != nil && !IsNil(o.SubAccountID) {
		return true
	}

	return false
}

// SetSubAccountID gets a reference to the given int64 and assigns it to the SubAccountID field.
func (o *QEmailMessage) SetSubAccountID(v int64) {
	o.SubAccountID = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *QEmailMessage) GetSubject() string {
	if o == nil || IsNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QEmailMessage) GetSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *QEmailMessage) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *QEmailMessage) SetSubject(v string) {
	o.Subject = &v
}

// GetSubmittedAt returns the SubmittedAt field value if set, zero value otherwise.
func (o *QEmailMessage) GetSubmittedAt() int64 {
	if o == nil || IsNil(o.SubmittedAt) {
		var ret int64
		return ret
	}
	return *o.SubmittedAt
}

// GetSubmittedAtOk returns a tuple with the SubmittedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QEmailMessage) GetSubmittedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.SubmittedAt) {
		return nil, false
	}
	return o.SubmittedAt, true
}

// HasSubmittedAt returns a boolean if a field has been set.
func (o *QEmailMessage) HasSubmittedAt() bool {
	if o != nil && !IsNil(o.SubmittedAt) {
		return true
	}

	return false
}

// SetSubmittedAt gets a reference to the given int64 and assigns it to the SubmittedAt field.
func (o *QEmailMessage) SetSubmittedAt(v int64) {
	o.SubmittedAt = &v
}

// GetTextBody returns the TextBody field value if set, zero value otherwise.
func (o *QEmailMessage) GetTextBody() string {
	if o == nil || IsNil(o.TextBody) {
		var ret string
		return ret
	}
	return *o.TextBody
}

// GetTextBodyOk returns a tuple with the TextBody field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QEmailMessage) GetTextBodyOk() (*string, bool) {
	if o == nil || IsNil(o.TextBody) {
		return nil, false
	}
	return o.TextBody, true
}

// HasTextBody returns a boolean if a field has been set.
func (o *QEmailMessage) HasTextBody() bool {
	if o != nil && !IsNil(o.TextBody) {
		return true
	}

	return false
}

// SetTextBody gets a reference to the given string and assigns it to the TextBody field.
func (o *QEmailMessage) SetTextBody(v string) {
	o.TextBody = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *QEmailMessage) GetTo() To {
	if o == nil || IsNil(o.To) {
		var ret To
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QEmailMessage) GetToOk() (*To, bool) {
	if o == nil || IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *QEmailMessage) HasTo() bool {
	if o != nil && !IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given To and assigns it to the To field.
func (o *QEmailMessage) SetTo(v To) {
	o.To = &v
}

// GetTrackClicks returns the TrackClicks field value if set, zero value otherwise.
func (o *QEmailMessage) GetTrackClicks() bool {
	if o == nil || IsNil(o.TrackClicks) {
		var ret bool
		return ret
	}
	return *o.TrackClicks
}

// GetTrackClicksOk returns a tuple with the TrackClicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QEmailMessage) GetTrackClicksOk() (*bool, bool) {
	if o == nil || IsNil(o.TrackClicks) {
		return nil, false
	}
	return o.TrackClicks, true
}

// HasTrackClicks returns a boolean if a field has been set.
func (o *QEmailMessage) HasTrackClicks() bool {
	if o != nil && !IsNil(o.TrackClicks) {
		return true
	}

	return false
}

// SetTrackClicks gets a reference to the given bool and assigns it to the TrackClicks field.
func (o *QEmailMessage) SetTrackClicks(v bool) {
	o.TrackClicks = &v
}

// GetTrackOpens returns the TrackOpens field value if set, zero value otherwise.
func (o *QEmailMessage) GetTrackOpens() bool {
	if o == nil || IsNil(o.TrackOpens) {
		var ret bool
		return ret
	}
	return *o.TrackOpens
}

// GetTrackOpensOk returns a tuple with the TrackOpens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QEmailMessage) GetTrackOpensOk() (*bool, bool) {
	if o == nil || IsNil(o.TrackOpens) {
		return nil, false
	}
	return o.TrackOpens, true
}

// HasTrackOpens returns a boolean if a field has been set.
func (o *QEmailMessage) HasTrackOpens() bool {
	if o != nil && !IsNil(o.TrackOpens) {
		return true
	}

	return false
}

// SetTrackOpens gets a reference to the given bool and assigns it to the TrackOpens field.
func (o *QEmailMessage) SetTrackOpens(v bool) {
	o.TrackOpens = &v
}

func (o QEmailMessage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QEmailMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountID) {
		toSerialize["accountID"] = o.AccountID
	}
	if !IsNil(o.AmpBody) {
		toSerialize["ampBody"] = o.AmpBody
	}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	if !IsNil(o.Attempt) {
		toSerialize["attempt"] = o.Attempt
	}
	if !IsNil(o.CustomFields) {
		toSerialize["customFields"] = o.CustomFields
	}
	if !IsNil(o.EmailType) {
		toSerialize["emailType"] = o.EmailType
	}
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !IsNil(o.Groups) {
		toSerialize["groups"] = o.Groups
	}
	if !IsNil(o.HeaderBcc) {
		toSerialize["headerBcc"] = o.HeaderBcc
	}
	if !IsNil(o.HeaderCc) {
		toSerialize["headerCc"] = o.HeaderCc
	}
	if !IsNil(o.HeaderTo) {
		toSerialize["headerTo"] = o.HeaderTo
	}
	if !IsNil(o.Headers) {
		toSerialize["headers"] = o.Headers
	}
	if !IsNil(o.HtmlBody) {
		toSerialize["htmlBody"] = o.HtmlBody
	}
	if !IsNil(o.IpID) {
		toSerialize["ipID"] = o.IpID
	}
	if !IsNil(o.IpPool) {
		toSerialize["ipPool"] = o.IpPool
	}
	if !IsNil(o.LocalIP) {
		toSerialize["localIP"] = o.LocalIP
	}
	if !IsNil(o.MessageID) {
		toSerialize["messageID"] = o.MessageID
	}
	if !IsNil(o.PreText) {
		toSerialize["preText"] = o.PreText
	}
	if !IsNil(o.PublicIP) {
		toSerialize["publicIP"] = o.PublicIP
	}
	if !IsNil(o.ReplyTo) {
		toSerialize["replyTo"] = o.ReplyTo
	}
	if !IsNil(o.SubAccountID) {
		toSerialize["subAccountID"] = o.SubAccountID
	}
	if !IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	if !IsNil(o.SubmittedAt) {
		toSerialize["submittedAt"] = o.SubmittedAt
	}
	if !IsNil(o.TextBody) {
		toSerialize["textBody"] = o.TextBody
	}
	if !IsNil(o.To) {
		toSerialize["to"] = o.To
	}
	if !IsNil(o.TrackClicks) {
		toSerialize["trackClicks"] = o.TrackClicks
	}
	if !IsNil(o.TrackOpens) {
		toSerialize["trackOpens"] = o.TrackOpens
	}
	return toSerialize, nil
}

type NullableQEmailMessage struct {
	value *QEmailMessage
	isSet bool
}

func (v NullableQEmailMessage) Get() *QEmailMessage {
	return v.value
}

func (v *NullableQEmailMessage) Set(val *QEmailMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableQEmailMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableQEmailMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQEmailMessage(val *QEmailMessage) *NullableQEmailMessage {
	return &NullableQEmailMessage{value: val, isSet: true}
}

func (v NullableQEmailMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQEmailMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


