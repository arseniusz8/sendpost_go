# Go API client for sendpost

Email API and SMTP relay to not just send and measure email sending, but also alert and optimise. We provide you with tools, expertise and support needed to reliably deliver emails to your customers inboxes on time, every time.

## Installation

Install the following dependencies:

```shell
go get github.com/arseniusz8/sendpost_go
```

Put the package under your project folder and add the following in import:

```golang
import "github.com/arseniusz8/sendpost_go"
```

## Getting Started

Please follow the [installation](#installation) instruction and execute the following Go code:

```go
cfg := sendpost.NewConfiguration()
client := sendpost.NewAPIClient(cfg)

emailMessage := sendpost.EmailMessage{}
emailMessage.SetSubject("Hello World")
emailMessage.SetHtmlBody("<strong>it works!</strong>")
emailMessage.SetIppool("PiedPiper")
emailMessage.From = &sendpost.From{}
emailMessage.From.SetEmail("richard@piedpiper.com")
tos := make([]sendpost.To, 0)
to := &sendpost.To{}
to.SetEmail("gavin@hooli.com")
tos = append(tos, *to)
emailMessage.To = tos

emailRequest := sendpost.ApiSendEmailRequest{}
emailRequest = emailRequest.XSubAccountApiKey("your_api_key")
emailRequest = emailRequest.EmailMessage(emailMessage)
res, _, err := client.EmailApi.SendEmailExecute(emailRequest)

## Documentation for API Endpoints

All URIs are relative to *https://api.sendpost.io/api/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*EmailApi* | [**SendEmail**](docs/EmailApi.md#sendemail) | **Post** /subaccount/email/ | 
*EmailApi* | [**SendEmailWithTemplate**](docs/EmailApi.md#sendemailwithtemplate) | **Post** /subaccount/email/template | 


## Documentation For Models

 - [Attachment](docs/Attachment.md)
 - [City](docs/City.md)
 - [CopyTo](docs/CopyTo.md)
 - [Device](docs/Device.md)
 - [EmailMessage](docs/EmailMessage.md)
 - [EmailResponse](docs/EmailResponse.md)
 - [EventMetadata](docs/EventMetadata.md)
 - [From](docs/From.md)
 - [Os](docs/Os.md)
 - [QEmailMessage](docs/QEmailMessage.md)
 - [QEvent](docs/QEvent.md)
 - [ReplyTo](docs/ReplyTo.md)
 - [To](docs/To.md)
 - [UserAgent](docs/UserAgent.md)
 - [WebhookEvent](docs/WebhookEvent.md)


## Documentation For Authorization

Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

hello@sendpost.io

