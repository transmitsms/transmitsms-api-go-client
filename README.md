# Getting started

Fast, secure, and easy integrations

## How to Build


* In order to successfully build and run your SDK files, you are required to have the following setup in your system:
    * **Go**  (Visit [https://golang.org/doc/install](https://golang.org/doc/install) for more details on how to install Go)
    * **Java VM** Version 8 or later
    * **Eclipse 4.6 (Neon)** or later ([http://www.eclipse.org/neon/](http://www.eclipse.org/neon/))
    * **GoClipse** setup within above installed Eclipse (Follow the instructions at [this link](https://github.com/GoClipse/goclipse/blob/latest/documentation/Installation.md#instructions) to setup GoClipse)
* Ensure that ```GOPATH``` environment variable is set in the system variables. If not, set it to your workspace directory where you will be adding your Go projects.
* The generated code uses unirest-go http library. Therefore, you will need internet access to resolve this dependency. If Go is properly installed and configured, run the following command to pull the dependency:

```Go
go get github.com/apimatic/unirest-go
```

This will install unirest-go in the ```GOPATH``` you specified in the system variables.

Now follow the steps mentioned below to build your SDK:

1. Open eclipse in the Go language perspective and click on the ```Import``` option in ```File``` menu.

![Importing SDK into Eclipse - Step 1](https://apidocs.io/illustration/go?step=import0)

2. Select ```General -> Existing Projects into Workspace``` option from the tree list.

![Importing SDK into Eclipse - Step 2](https://apidocs.io/illustration/go?step=import1)

3. In ```Select root directory```, provide path to the unzipped archive for the generated code. Once the path is set and the Project becomes visible under ```Projects``` click ```Finish```

![Importing SDK into Eclipse - Step 3](https://apidocs.io/illustration/go?step=import2&workspaceFolder=TransmitSMS-GoLang&projectName=transmitsms_lib)

4. The Go library will be imported and its files will be visible in the Project Explorer

![Importing SDK into Eclipse - Step 4](https://apidocs.io/illustration/go?step=import3&projectName=transmitsms_lib)

## How to Use

The following section explains how to use the TransmitsmsLib library in a new project.

### 1. Add a new Test Project

Create a new project in Eclipse by ```File``` -> ```New``` -> ```Go Project```

![Add a new project in Eclipse](https://apidocs.io/illustration/go?step=createNewProject0)

Name the Project as ```Test``` and click ```Finish```

![Create a new Maven Project - Step 1](https://apidocs.io/illustration/go?step=createNewProject1)

Create a new directory in the ```src``` directory of this project

![Create a new Maven Project - Step 2](https://apidocs.io/illustration/go?step=createNewProject2&projectName=transmitsms_lib)

Name it ```test.com```

![Create a new Maven Project - Step 3](https://apidocs.io/illustration/go?step=createNewProject3&projectName=transmitsms_lib)

Now create a new file inside ```src/test.com```

![Create a new Maven Project - Step 4](https://apidocs.io/illustration/go?step=createNewProject4&projectName=transmitsms_lib)

Name it ```testsdk.go```

![Create a new Maven Project - Step 5](https://apidocs.io/illustration/go?step=createNewProject5&projectName=transmitsms_lib)

In this Go file, you can start adding code to initialize the client library. Sample code to initialize the client library and using its methods is given in the subsequent sections.

### 2. Configure the Test Project

You need to import your generated library in this project in order to make use of its functions. In order to import the library, you can add its path in the ```GOPATH``` for this project. Follow the below steps:

Right click on the project name and click on ```Properties```

![Adding dependency to the client library - Step 1](https://apidocs.io/illustration/go?step=testProject0&projectName=transmitsms_lib)

Choose ```Go Compiler``` from the side menu. Check ```Use project specific settings``` and uncheck ```Use same value as the GOPATH environment variable.```. By default, the GOPATH value from the environment variables will be visible in ```Eclipse GOPATH```. Do not remove this as this points to the unirest dependency.

![Adding dependency to the client library - Step 2](https://apidocs.io/illustration/go?step=testProject1)

Append the library path to this GOPATH

![Adding dependency to the client library - Step 3](https://apidocs.io/illustration/go?step=testProject2&workspaceFolder=TransmitSMS-GoLang)

Once the path is appended, click on ```OK```

### 3. Build the Test Project

Right click on the project name and click on ```Build Project```

![Build Project](https://apidocs.io/illustration/go?step=buildProject&projectName=transmitsms_lib)

### 4. Run the Test Project

If the build is successful, right click on your Go file and click on ```Run As``` -> ```Go Application```

![Run Project](https://apidocs.io/illustration/go?step=runProject&projectName=transmitsms_lib)

## Initialization

### Authentication
In order to setup authentication of the API client, you need the following information.

| Parameter | Description |
|-----------|-------------|
| username | API Key |
| password | Secret |


To configure these for your generated code, open the file "Configuration.go" and edit it's contents.


# Class Reference

## <a name="list_of_controllers"></a>List of Controllers

* [sms_pkg](#sms_pkg)
* [numbers_pkg](#numbers_pkg)
* [emailsms_pkg](#emailsms_pkg)
* [resellers_pkg](#resellers_pkg)
* [account_pkg](#account_pkg)
* [keywords_pkg](#keywords_pkg)
* [lists_pkg](#lists_pkg)

## <a name="sms_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".sms_pkg") sms_pkg

### Get instance

Factory for the ``` SMS ``` interface can be accessed from the package sms_pkg.

```go
sMS := sms_pkg.NewSMS()
```

### <a name="send_sms"></a>![Method: ](https://apidocs.io/img/method.png ".sms_pkg.SendSMS") SendSMS

> The Send-SMS call is the primary method of sending SMS.
> 
> You can elect to pass us the recipient numbers from your database each time you make a call, or you can elect to store recipient data in a contact list and submit only the list_id to trigger the send. This is best for large databases. To add a list please refer to the add-list call.
> 
> Cost data is returned in the major unit of your account currency, e.g. dollars or pounds
> 
> NOTE: If you do not pass the 'from' parameter the messages will be sent from the shared number pool, unless you have a leased number on your account in which case it will be set as the Caller ID


```go
func (me *SMS_IMPL) SendSMS(
            message string,
            format string,
            to *int64,
            from *string,
            sendAt *string,
            listId *int64,
            dlrCallback *string,
            replyCallback *string,
            validity *string,
            repliesToEmail *string,
            fromShared *bool,
            countrycode *string,
            trackedLinkUrl *string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| message |  ``` Required ```  ``` DefaultValue ```  | Message text |
| format |  ``` Required ```  ``` DefaultValue ```  | Response Format |
| to |  ``` Optional ```  | Number or set of up to 10,000 numbers to send the SMS to. If your number set has some invalid numbers, they won’t cause validation error, but will be returned as ‘fails’ parameter of the response (see example 3). Number must be defined in international format.  Some examples by destination:  AU 61491570156, NZ 64212670129, SG 6598654321, UK 44750017696, US 1213811413 |
| from |  ``` Optional ```  | Set the alphanumeric Caller ID, mobile numbers should be in international format. Maximum 11 characters. No spaces. If not set will use shared number pool. |
| sendAt |  ``` Optional ```  | A time in the future to send the message Note: All returned timestamps are in ISO8601 format e.g. YYYY-MM-DD HH:MM:SS. The zone is always UTC. |
| listId |  ``` Optional ```  | This ID is the numerical reference to one of your recipient lists Note: List ID's are made up of digits and will be returned by the add-list call, or can be found at any time by logging into your account and visiting your contacts page. |
| dlrCallback |  ``` Optional ```  | A URL on your system which we can call to notify you of Delivery Receipts. If required, this Parameter can be different for each message sent and will take precedence over the DLR Callback URL supplied by you in the API Settings. |
| replyCallback |  ``` Optional ```  | A URL on your system which we can call to notify you of incoming messages. If required, this parameter can be different and will take precedence over the Reply Callback URL supplied by you on the API Settings. |
| validity |  ``` Optional ```  | Specify the maximum time to attempt to deliver. In minutes, 0 (zero) implies no limit. |
| repliesToEmail |  ``` Optional ```  | Specify an email address to send responses to this message. NOTE: specified email must be authorised to send messages via add-email or in your account under the 'Email SMS' section. |
| fromShared |  ``` Optional ```  | Forces sending via the shared number when you have virtual numbers |
| countrycode |  ``` Optional ```  | Formats numbers given to international format for this 2 letter country code. i.e. 0422222222 will become 6142222222 when countrycode is AU. Codes available AU Australia, NZ New Zealand SG Singapore GB United Kingdom US United States |
| trackedLinkUrl |  ``` Optional ```  | Converts this URL to unique tapth.is/xxxxxx tracking link for each contact. Inserted into message with variable [tracked-link]. Clicks on this URL will be passed as notifications via 'Link hits callback URL' defined in account settings. |


#### Example Usage

```go
message := "[Enter your message here]"
format := "json"
to,_ := strconv.ParseInt("97", 10, 8)
from := "from"
sendAt := "send_at"
listId,_ := strconv.ParseInt("97", 10, 8)
dlrCallback := "dlr_callback"
replyCallback := "reply_callback"
validity := "validity"
repliesToEmail := "replies_to_email"
fromShared := false
countrycode := "countrycode"
trackedLinkUrl := "tracked_link_url"

var result interface{}
result,_ = sMS.SendSMS(message, format, to, from, sendAt, listId, dlrCallback, replyCallback, validity, repliesToEmail, fromShared, countrycode, trackedLinkUrl)

```


### <a name="format_number"></a>![Method: ](https://apidocs.io/img/method.png ".sms_pkg.FormatNumber") FormatNumber

> Format and validate a given number.


```go
func (me *SMS_IMPL) FormatNumber(
            msisdn string,
            countrycode string,
            format string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| msisdn |  ``` Required ```  | The number to check |
| countrycode |  ``` Required ```  | 2 Letter countrycode to validate number against |
| format |  ``` Required ```  ``` DefaultValue ```  | Response Format |


#### Example Usage

```go
msisdn := "msisdn"
countrycode := "countrycode"
format := "json"

var result interface{}
result,_ = sMS.FormatNumber(msisdn, countrycode, format)

```


### <a name="get_sms"></a>![Method: ](https://apidocs.io/img/method.png ".sms_pkg.GetSMS") GetSMS

> Get information about a message you have sent.


```go
func (me *SMS_IMPL) GetSMS(
            messageId string,
            format string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| messageId |  ``` Required ```  | Message ID |
| format |  ``` Required ```  ``` DefaultValue ```  | Response format either json or xml |


#### Example Usage

```go
messageId := "message_id"
format := "json"

var result interface{}
result,_ = sMS.GetSMS(messageId, format)

```


### <a name="get_sms_stats"></a>![Method: ](https://apidocs.io/img/method.png ".sms_pkg.GetSMSStats") GetSMSStats

> Get the status about a message you have sent.


```go
func (me *SMS_IMPL) GetSMSStats(
            messageId string,
            format string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| messageId |  ``` Required ```  | Message ID |
| format |  ``` Required ```  ``` DefaultValue ```  | Response Format e.g. json or xml |


#### Example Usage

```go
messageId := "message_id"
format := "json"

var result interface{}
result,_ = sMS.GetSMSStats(messageId, format)

```


### <a name="get_sms_responses"></a>![Method: ](https://apidocs.io/img/method.png ".sms_pkg.GetSMSResponses") GetSMSResponses

> Pick up responses to messages you have sent. Filter by keyword or for just one phone number.


```go
func (me *SMS_IMPL) GetSMSResponses(
            format string,
            messageId *string,
            keywordId *string,
            keyword *string,
            number *string,
            msisdn *string,
            page *string,
            max *string,
            includeOriginal *string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| format |  ``` Required ```  ``` DefaultValue ```  | Response format e.g. json or xml |
| messageId |  ``` Optional ```  | Message ID |
| keywordId |  ``` Optional ```  | Keyword ID |
| keyword |  ``` Optional ```  | Keyword |
| number |  ``` Optional ```  | Filter results by response number, If keyword is set |
| msisdn |  ``` Optional ```  | Filter results by a particular mobile number |
| page |  ``` Optional ```  | Page number, for pagination |
| max |  ``` Optional ```  | Maximum results returned per page |
| includeOriginal |  ``` Optional ```  | include text of original message |


#### Example Usage

```go
format := "json"
messageId := "message_id"
keywordId := "keyword_id"
keyword := "keyword"
number := "number"
msisdn := "msisdn"
page := "page"
max := "max"
includeOriginal := "include_original"

var result interface{}
result,_ = sMS.GetSMSResponses(format, messageId, keywordId, keyword, number, msisdn, page, max, includeOriginal)

```


### <a name="get_user_sms_responses"></a>![Method: ](https://apidocs.io/img/method.png ".sms_pkg.GetUserSMSResponses") GetUserSMSResponses

> Pick up responses to messages you have sent. Instead of setting message ID, you should provide a time frame.


```go
func (me *SMS_IMPL) GetUserSMSResponses(
            format string,
            start *string,
            end *string,
            page *string,
            max *string,
            keywords *string,
            includeOriginal *string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| format |  ``` Required ```  ``` DefaultValue ```  | Response format e.g. json or xml |
| start |  ``` Optional ```  | A timestamp to start the report from (Format yyyy-mm-dd, Timezone UTC) |
| end |  ``` Optional ```  | A timestamp to end the report at (Format yyyy-mm-dd, Timezone UTC) |
| page |  ``` Optional ```  | Page number, for pagination |
| max |  ``` Optional ```  | Maximum results returned per page |
| keywords |  ``` Optional ```  | Filter if keyword responses should be included. Can be: ‘only’ - only keyword responses will be included‘omit’ - only regular campaign responses will be included  ‘both’ - both keyword and campaign responses will be included (default) |
| includeOriginal |  ``` Optional ```  | include text of original message |


#### Example Usage

```go
format := "json"
start := "start"
end := "end"
page := "page"
max := "max"
keywords := "keywords"
includeOriginal := "include_original"

var result interface{}
result,_ = sMS.GetUserSMSResponses(format, start, end, page, max, keywords, includeOriginal)

```


### <a name="get_link_hits"></a>![Method: ](https://apidocs.io/img/method.png ".sms_pkg.GetLinkHits") GetLinkHits

> Get the list of recipients who have clicked on your tracked link.


```go
func (me *SMS_IMPL) GetLinkHits(
            messageId string,
            format string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| messageId |  ``` Required ```  | Message ID |
| format |  ``` Required ```  ``` DefaultValue ```  | Response format e.g. json or xml |


#### Example Usage

```go
messageId := "message_id"
format := "json"

var result interface{}
result,_ = sMS.GetLinkHits(messageId, format)

```


### <a name="get_sms_sent"></a>![Method: ](https://apidocs.io/img/method.png ".sms_pkg.GetSMSSent") GetSMSSent

> Get a list of recipients from a message send. Get up to date information such as opt-out status and delivery status.


```go
func (me *SMS_IMPL) GetSMSSent(
            messageId string,
            format string,
            optouts *string,
            page *string,
            max *string,
            delivery *string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| messageId |  ``` Required ```  | Message ID's are made up of digits |
| format |  ``` Required ```  ``` DefaultValue ```  | Response format e.g json or xml |
| optouts |  ``` Optional ```  | Whether to include optouts. Valid options are:    only - only get optouts   omit - do not get optouts   include - get all recipients including optouts (default) |
| page |  ``` Optional ```  | Page number, for pagination |
| max |  ``` Optional ```  | Maximum results returned per page |
| delivery |  ``` Optional ```  | Only show messages with requested delivery status. Valid options are:   delivered - only show delivered messages   failed - only show failed messages   pending - only show pending messages |


#### Example Usage

```go
messageId := "message_id"
format := "json"
optouts := "optouts"
page := "page"
max := "max"
delivery := "delivery"

var result interface{}
result,_ = sMS.GetSMSSent(messageId, format, optouts, page, max, delivery)

```


### <a name="get_sms_delivery_status"></a>![Method: ](https://apidocs.io/img/method.png ".sms_pkg.GetSMSDeliveryStatus") GetSMSDeliveryStatus

> Get the delivery time of a delivered message.


```go
func (me *SMS_IMPL) GetSMSDeliveryStatus(
            messageId string,
            msisdn string,
            format string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| messageId |  ``` Required ```  | Message ID |
| msisdn |  ``` Required ```  | Mobile number |
| format |  ``` Required ```  ``` DefaultValue ```  | Response format e.g. json or xml |


#### Example Usage

```go
messageId := "message_id"
msisdn := "msisdn"
format := "json"

var result interface{}
result,_ = sMS.GetSMSDeliveryStatus(messageId, msisdn, format)

```


### <a name="cancel_sms"></a>![Method: ](https://apidocs.io/img/method.png ".sms_pkg.CancelSMS") CancelSMS

> Cancel a message you have scheduled to be sent in the future.


```go
func (me *SMS_IMPL) CancelSMS(
            id string,
            format string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| id |  ``` Required ```  | Message ID |
| format |  ``` Required ```  ``` DefaultValue ```  | Response format e.g. json or xml |


#### Example Usage

```go
id := "id"
format := "json"

var result interface{}
result,_ = sMS.CancelSMS(id, format)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="numbers_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".numbers_pkg") numbers_pkg

### Get instance

Factory for the ``` NUMBERS ``` interface can be accessed from the package numbers_pkg.

```go
numbers := numbers_pkg.NewNUMBERS()
```

### <a name="lease_number"></a>![Method: ](https://apidocs.io/img/method.png ".numbers_pkg.LeaseNumber") LeaseNumber

> Lease a dedicated virtual number.


```go
func (me *NUMBERS_IMPL) LeaseNumber(
            number string,
            format string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| number |  ``` Required ```  | The virtual number to lease. Omit this field to be given a random number. Use get-numbers to find out which numbers are currently available. |
| format |  ``` Required ```  ``` DefaultValue ```  | Response format e.g. json or xml |


#### Example Usage

```go
number := "number"
format := "json"

var result interface{}
result,_ = numbers.LeaseNumber(number, format)

```


### <a name="edit_number_options"></a>![Method: ](https://apidocs.io/img/method.png ".numbers_pkg.EditNumberOptions") EditNumberOptions

> Edit your dedicated virtual number options.


```go
func (me *NUMBERS_IMPL) EditNumberOptions(
            number string,
            format string,
            forwardEmail *string,
            forwardSms *string,
            forwardUrl *string,
            listId *string,
            welcomeMessage *string,
            membersMessage *string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| number |  ``` Required ```  | The dedicated virtual number. |
| format |  ``` Required ```  ``` DefaultValue ```  | Response format e.g. json or xml |
| forwardEmail |  ``` Optional ```  | Forward incoming messages to a set of email addresses. |
| forwardSms |  ``` Optional ```  | Forward incoming messages to a set of mobile numbers. |
| forwardUrl |  ``` Optional ```  | Forward incoming messages to a URL. |
| listId |  ``` Optional ```  | Add new numbers that message in to this list. |
| welcomeMessage |  ``` Optional ```  | Auto-response for all messages received. |
| membersMessage |  ``` Optional ```  | Auto-response if the number is already on the list. (must be adding the number to a list) |


#### Example Usage

```go
number := "number"
format := "json"
forwardEmail := "forward_email"
forwardSms := "forward_sms"
forwardUrl := "forward_url"
listId := "list_id"
welcomeMessage := "welcome_message"
membersMessage := "members_message"

var result interface{}
result,_ = numbers.EditNumberOptions(number, format, forwardEmail, forwardSms, forwardUrl, listId, welcomeMessage, membersMessage)

```


### <a name="get_numbers"></a>![Method: ](https://apidocs.io/img/method.png ".numbers_pkg.GetNumbers") GetNumbers

> Get a list of numbers either leased by you or available to be leased.


```go
func (me *NUMBERS_IMPL) GetNumbers(
            filter string,
            format string,
            page *string,
            max *string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| filter |  ``` Required ```  ``` DefaultValue ```  | Possible values are owned - retrieve your own response numbers (default) available - retrieve response numbers available for purchase |
| format |  ``` Required ```  ``` DefaultValue ```  | Response format e.g. json or xml |
| page |  ``` Optional ```  | Page number, for pagination |
| max |  ``` Optional ```  | Maximum results returned per page |


#### Example Usage

```go
filter := "owned"
format := "json"
page := "page"
max := "max"

var result interface{}
result,_ = numbers.GetNumbers(filter, format, page, max)

```


### <a name="get_number"></a>![Method: ](https://apidocs.io/img/method.png ".numbers_pkg.GetNumber") GetNumber

> Get detailed information about a response number you have leased.


```go
func (me *NUMBERS_IMPL) GetNumber(
            number string,
            format string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| number |  ``` Required ```  | The virtual number to retrieve |
| format |  ``` Required ```  ``` DefaultValue ```  | Response format e.g. json or xml |


#### Example Usage

```go
number := "number"
format := "json"

var result interface{}
result,_ = numbers.GetNumber(number, format)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="emailsms_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".emailsms_pkg") emailsms_pkg

### Get instance

Factory for the ``` EMAILSMS ``` interface can be accessed from the package emailsms_pkg.

```go
emailSMS := emailsms_pkg.NewEMAILSMS()
```

### <a name="add_email"></a>![Method: ](https://apidocs.io/img/method.png ".emailsms_pkg.AddEmail") AddEmail

> Authorise an email address for sending Email to SMS


```go
func (me *EMAILSMS_IMPL) AddEmail(
            email string,
            number string,
            format string,
            maxSms *string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| email |  ``` Required ```  | Email address to register. You may also register a wild-card email which allows any user on the same domain to use Email to SMS.  Wild-card format: *@example.com |
| number |  ``` Required ```  | Optional dedicated virtual number virtual number |
| format |  ``` Required ```  ``` DefaultValue ```  | Response format e.g. json or xml |
| maxSms |  ``` Optional ```  | The maximum number of SMS messages to send from one email message sent from this email address.  Possible values: 1 - up to 160 characters (default) 2 - up to 306 characters 3 - up to 459 characters 4 - up to 612 characters |


#### Example Usage

```go
email := "email"
number := "number"
format := "json"
maxSms := "max_sms"

var result interface{}
result,_ = emailSMS.AddEmail(email, number, format, maxSms)

```


### <a name="delete_email"></a>![Method: ](https://apidocs.io/img/method.png ".emailsms_pkg.DeleteEmail") DeleteEmail

> Remove an email address from the Email to SMS authorised list.


```go
func (me *EMAILSMS_IMPL) DeleteEmail(
            email string,
            format string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| email |  ``` Required ```  | Email address to remove. You may also use a wild-card email which removes all emails on that domain.  Wild-card format: *@example.com |
| format |  ``` Required ```  ``` DefaultValue ```  | Response format e.g. json or xml |


#### Example Usage

```go
email := "email"
format := "json"

var result interface{}
result,_ = emailSMS.DeleteEmail(email, format)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="resellers_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".resellers_pkg") resellers_pkg

### Get instance

Factory for the ``` RESELLERS ``` interface can be accessed from the package resellers_pkg.

```go
resellers := resellers_pkg.NewRESELLERS()
```

### <a name="get_transaction"></a>![Method: ](https://apidocs.io/img/method.png ".resellers_pkg.GetTransaction") GetTransaction

> Get a list of transactions for an account.


```go
func (me *RESELLERS_IMPL) GetTransaction(
            id string,
            format string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| id |  ``` Required ```  | Transaction ID |
| format |  ``` Required ```  ``` DefaultValue ```  | Response format e.g. json or xml |


#### Example Usage

```go
id := "id"
format := "json"

var result interface{}
result,_ = resellers.GetTransaction(id, format)

```


### <a name="edit_client"></a>![Method: ](https://apidocs.io/img/method.png ".resellers_pkg.EditClient") EditClient

> Edit an existing client


```go
func (me *RESELLERS_IMPL) EditClient(
            clientId string,
            format string,
            name *string,
            contact *string,
            email *string,
            password *string,
            msisdn *string,
            timezone *string,
            clientPays *string,
            smsMargin *string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| clientId |  ``` Required ```  | The ID of the client |
| format |  ``` Required ```  ``` DefaultValue ```  | Response format e.g. json or xml |
| name |  ``` Optional ```  | Client company name. Must be unique |
| contact |  ``` Optional ```  | Contact name |
| email |  ``` Optional ```  | Client email address |
| password |  ``` Optional ```  | Client password |
| msisdn |  ``` Optional ```  | Client phone number |
| timezone |  ``` Optional ```  | A valid timezone, Australia/Sydney. Defaults to your own |
| clientPays |  ``` Optional ```  | Set to true if the client will pay (the default) or false if you will pay |
| smsMargin |  ``` Optional ```  | The number of cents to add to the base SMS price. A decimal value. |


#### Example Usage

```go
clientId := "client_id"
format := "json"
name := "name"
contact := "contact"
email := "email"
password := "password"
msisdn := "msisdn"
timezone := "timezone"
clientPays := "client_pays"
smsMargin := "sms_margin"

var result interface{}
result,_ = resellers.EditClient(clientId, format, name, contact, email, password, msisdn, timezone, clientPays, smsMargin)

```


### <a name="get_transactions"></a>![Method: ](https://apidocs.io/img/method.png ".resellers_pkg.GetTransactions") GetTransactions

> Get a list of transactions for a client.


```go
func (me *RESELLERS_IMPL) GetTransactions(
            clientId string,
            format string,
            start *string,
            end *string,
            page *string,
            max *string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| clientId |  ``` Required ```  | Only retrieve records for a particular client |
| format |  ``` Required ```  ``` DefaultValue ```  | Response format e.g. json or xml |
| start |  ``` Optional ```  | A timestamp to start the report from |
| end |  ``` Optional ```  | A timestamp to end the report at |
| page |  ``` Optional ```  | Page number, for pagination |
| max |  ``` Optional ```  | Maximum results returned per page |


#### Example Usage

```go
clientId := "client_id"
format := "json"
start := "start"
end := "end"
page := "page"
max := "max"

var result interface{}
result,_ = resellers.GetTransactions(clientId, format, start, end, page, max)

```


### <a name="get_client"></a>![Method: ](https://apidocs.io/img/method.png ".resellers_pkg.GetClient") GetClient

> Get detailed information about a client.


```go
func (me *RESELLERS_IMPL) GetClient(
            clientId string,
            format string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| clientId |  ``` Required ```  | The ID of the client |
| format |  ``` Required ```  ``` DefaultValue ```  | Response format e.g. json or xml |


#### Example Usage

```go
clientId := "client_id"
format := "json"

var result interface{}
result,_ = resellers.GetClient(clientId, format)

```


### <a name="get_clients"></a>![Method: ](https://apidocs.io/img/method.png ".resellers_pkg.GetClients") GetClients

> Get a list of all clients.


```go
func (me *RESELLERS_IMPL) GetClients(
            format string,
            page *string,
            max *string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| format |  ``` Required ```  ``` DefaultValue ```  | Response format e.g. json or xml |
| page |  ``` Optional ```  | Page number, for pagination |
| max |  ``` Optional ```  | Maximum results returned per page |


#### Example Usage

```go
format := "json"
page := "page"
max := "max"

var result interface{}
result,_ = resellers.GetClients(format, page, max)

```


### <a name="add_client"></a>![Method: ](https://apidocs.io/img/method.png ".resellers_pkg.AddClient") AddClient

> Add a new client.


```go
func (me *RESELLERS_IMPL) AddClient(
            name string,
            email string,
            password string,
            msisdn string,
            format string,
            contact *string,
            timezone *string,
            clientPays *string,
            smsMargin *string,
            numberMargin *string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| name |  ``` Required ```  | Client company name |
| email |  ``` Required ```  | Client email address |
| password |  ``` Required ```  | Client password |
| msisdn |  ``` Required ```  | Client phone number |
| format |  ``` Required ```  ``` DefaultValue ```  | Response format e.g. json |
| contact |  ``` Optional ```  | Contact name |
| timezone |  ``` Optional ```  | A valid timezone, Australia/Sydney. Defaults to your own |
| clientPays |  ``` Optional ```  | Set to true if the client will pay (the default) or false if you will pay |
| smsMargin |  ``` Optional ```  | The number of cents to add to the base SMS price. A decimal value |
| numberMargin |  ``` Optional ```  | The number of cents to add to the base number price. A decimal value |


#### Example Usage

```go
name := "name"
email := "email"
password := "password"
msisdn := "msisdn"
format := "json"
contact := "contact"
timezone := "timezone"
clientPays := "client_pays"
smsMargin := "sms_margin"
numberMargin := "number_margin"

var result interface{}
result,_ = resellers.AddClient(name, email, password, msisdn, format, contact, timezone, clientPays, smsMargin, numberMargin)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="account_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".account_pkg") account_pkg

### Get instance

Factory for the ``` ACCOUNT ``` interface can be accessed from the package account_pkg.

```go
account := account_pkg.NewACCOUNT()
```

### <a name="get_balance"></a>![Method: ](https://apidocs.io/img/method.png ".account_pkg.GetBalance") GetBalance

> Get a summary of your account balance.


```go
func (me *ACCOUNT_IMPL) GetBalance(format string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| format |  ``` Required ```  ``` DefaultValue ```  | Response format |


#### Example Usage

```go
format := "json"

var result interface{}
result,_ = account.GetBalance(format)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="keywords_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".keywords_pkg") keywords_pkg

### Get instance

Factory for the ``` KEYWORDS ``` interface can be accessed from the package keywords_pkg.

```go
keywords := keywords_pkg.NewKEYWORDS()
```

### <a name="get_keywords"></a>![Method: ](https://apidocs.io/img/method.png ".keywords_pkg.GetKeywords") GetKeywords

> Get a list of existing keywords.


```go
func (me *KEYWORDS_IMPL) GetKeywords(
            format string,
            number *string,
            page *string,
            max *string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| format |  ``` Required ```  ``` DefaultValue ```  | Response format e.g. json or xml |
| number |  ``` Optional ```  | Filter the list by virtual number |
| page |  ``` Optional ```  | Page number, for pagination |
| max |  ``` Optional ```  | Maximum results returned per page |


#### Example Usage

```go
format := "json"
number := "number"
page := "page"
max := "max"

var result interface{}
result,_ = keywords.GetKeywords(format, number, page, max)

```


### <a name="add_keyword"></a>![Method: ](https://apidocs.io/img/method.png ".keywords_pkg.AddKeyword") AddKeyword

> Add a keyword to an existing virtual number.


```go
func (me *KEYWORDS_IMPL) AddKeyword(
            keyword string,
            number string,
            format string,
            reference *string,
            listId *string,
            welcomeMessage *string,
            membersMessage *string,
            activate *string,
            forwardUrl *string,
            forwardEmail *string,
            forwardSms *string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| keyword |  ``` Required ```  | The first word of a text message |
| number |  ``` Required ```  | The dedicated virtual number that the keyword belongs to |
| format |  ``` Required ```  ``` DefaultValue ```  | Response format e.g. json or xml |
| reference |  ``` Optional ```  | Your own reference (up to 100 characters) |
| listId |  ``` Optional ```  | ID of a list to add respondents to, list ID's can be found in the title of a list or in the list page URL |
| welcomeMessage |  ``` Optional ```  | SMS message to send to new members |
| membersMessage |  ``` Optional ```  | SMS message to existing members |
| activate |  ``` Optional ```  | Whether to make the keyword active immediately.  Possible values: true - activate immediately (default) false - create the keyword but do not activate |
| forwardUrl |  ``` Optional ```  | Forward messages to a URL |
| forwardEmail |  ``` Optional ```  | Forward messages to a set of email addresses |
| forwardSms |  ``` Optional ```  | Forward messages to a set of msisdns |


#### Example Usage

```go
keyword := "keyword"
number := "number"
format := "json"
reference := "reference"
listId := "list_id"
welcomeMessage := "welcome_message"
membersMessage := "members_message"
activate := "activate"
forwardUrl := "forward_url"
forwardEmail := "forward_email"
forwardSms := "forward_sms"

var result interface{}
result,_ = keywords.AddKeyword(keyword, number, format, reference, listId, welcomeMessage, membersMessage, activate, forwardUrl, forwardEmail, forwardSms)

```


### <a name="edit_keyword"></a>![Method: ](https://apidocs.io/img/method.png ".keywords_pkg.EditKeyword") EditKeyword

> Edit an existing keyword.


```go
func (me *KEYWORDS_IMPL) EditKeyword(
            keyword string,
            number string,
            format string,
            reference *string,
            status *string,
            listId *string,
            welcomeMessage *string,
            membersMessage *string,
            activate *string,
            forwardUrl *string,
            forwardEmail *string,
            forwardSms *string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| keyword |  ``` Required ```  | The first word of a text message |
| number |  ``` Required ```  | The dedicated virtual number that the keyword belongs to |
| format |  ``` Required ```  ``` DefaultValue ```  | Response format e.g. json or xml |
| reference |  ``` Optional ```  | Your own reference (up to 100 characters) |
| status |  ``` Optional ```  | Your own reference (up to 100 characters) |
| listId |  ``` Optional ```  | ID of a list to add respondents to, list ID's can be found in the title of a list or in the list page URL |
| welcomeMessage |  ``` Optional ```  | SMS message to send to new members |
| membersMessage |  ``` Optional ```  | SMS message to existing members |
| activate |  ``` Optional ```  | Whether to make the keyword active immediately.  Possible values: true - activate immediately (default) false - create the keyword but do not activate |
| forwardUrl |  ``` Optional ```  | Forward messages to a URL |
| forwardEmail |  ``` Optional ```  | Forward messages to a set of email addresses |
| forwardSms |  ``` Optional ```  | Forward messages to a set of msisdns |


#### Example Usage

```go
keyword := "keyword"
number := "number"
format := "json"
reference := "reference"
status := "status"
listId := "list_id"
welcomeMessage := "welcome_message"
membersMessage := "members_message"
activate := "activate"
forwardUrl := "forward_url"
forwardEmail := "forward_email"
forwardSms := "forward_sms"

var result interface{}
result,_ = keywords.EditKeyword(keyword, number, format, reference, status, listId, welcomeMessage, membersMessage, activate, forwardUrl, forwardEmail, forwardSms)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="lists_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".lists_pkg") lists_pkg

### Get instance

Factory for the ``` LISTS ``` interface can be accessed from the package lists_pkg.

```go
lists := lists_pkg.NewLISTS()
```

### <a name="add_to_list"></a>![Method: ](https://apidocs.io/img/method.png ".lists_pkg.AddToList") AddToList

> Add a member to a list.


```go
func (me *LISTS_IMPL) AddToList(
            listId string,
            msisdn string,
            format string,
            firstName *string,
            lastName *string,
            countrycode *string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| listId |  ``` Required ```  | ID of the list to add to |
| msisdn |  ``` Required ```  | Mobile number of the member |
| format |  ``` Required ```  ``` DefaultValue ```  | Response format e.g. json or xml |
| firstName |  ``` Optional ```  | First name of the member |
| lastName |  ``` Optional ```  | Last name of the member |
| countrycode |  ``` Optional ```  | Formats msisdn for the given countrycode |


#### Example Usage

```go
listId := "list_id"
msisdn := "msisdn"
format := "json"
firstName := "first_name"
lastName := "last_name"
countrycode := "countrycode"

var result interface{}
result,_ = lists.AddToList(listId, msisdn, format, firstName, lastName, countrycode)

```


### <a name="add_field_to_list"></a>![Method: ](https://apidocs.io/img/method.png ".lists_pkg.AddFieldToList") AddFieldToList

> Update or add custom fields to a list


```go
func (me *LISTS_IMPL) AddFieldToList(
            listId string,
            field1 string,
            format string,
            field2 *string,
                queryParameters map[string]interface{})(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| listId |  ``` Required ```  | ID of the list to add to |
| field1 |  ``` Required ```  | Custom field value where n is an integer between 1 and 10. You can also use the names of the custom fields you have chosen for your list, e.g. field.birthday. |
| format |  ``` Required ```  ``` DefaultValue ```  | Response format e.g. json or xml |
| field2 |  ``` Optional ```  | Custom field value where n is an integer between 1 and 10. You can also use the names of the custom fields you have chosen for your list, e.g. field.birthday. |
| queryParameters | ``` Optional ``` | Additional optional query parameters are supported by this method |


#### Example Usage

```go
listId := "list_id"
field1 := "field_1"
format := "json"
field2 := "field_2"
// key-value map for optional query parameters
	queryParams := map[string]interface{}{"key" : "value"}


var result interface{}
result,_ = lists.AddFieldToList(listId, field1, format, field2, queryParams)

```


### <a name="add_list"></a>![Method: ](https://apidocs.io/img/method.png ".lists_pkg.AddList") AddList

> Create a new list including the ability to add custom fields.


```go
func (me *LISTS_IMPL) AddList(
            name string,
            format string,
            field1 *string,
                queryParameters map[string]interface{})(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| name |  ``` Required ```  | name |
| format |  ``` Required ```  ``` DefaultValue ```  | Response format e.g. json or xml |
| field1 |  ``` Optional ```  | A custom field name where n is an integer between 1 and 10. Once field names have been set they cannot be changed. |
| queryParameters | ``` Optional ``` | Additional optional query parameters are supported by this method |


#### Example Usage

```go
name := "name"
format := "json"
field1 := "field_1"
// key-value map for optional query parameters
	queryParams := map[string]interface{}{"key" : "value"}


var result interface{}
result,_ = lists.AddList(name, format, field1, queryParams)

```


### <a name="optout_list_member"></a>![Method: ](https://apidocs.io/img/method.png ".lists_pkg.OptoutListMember") OptoutListMember

> Opt a user out of one list or all lists.


```go
func (me *LISTS_IMPL) OptoutListMember(
            listId string,
            msisdn string,
            format string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| listId |  ``` Required ```  | ID of the list to opt the user out of. Set this to 0 (zero) to opt out of all of your lists. |
| msisdn |  ``` Required ```  | Mobile number of the member to opt out |
| format |  ``` Required ```  ``` DefaultValue ```  | Response format e.g json or xml |


#### Example Usage

```go
listId := "list_id"
msisdn := "msisdn"
format := "json"

var result interface{}
result,_ = lists.OptoutListMember(listId, msisdn, format)

```


### <a name="delete_from_list"></a>![Method: ](https://apidocs.io/img/method.png ".lists_pkg.DeleteFromList") DeleteFromList

> Remove a member from one list or all lists.


```go
func (me *LISTS_IMPL) DeleteFromList(
            listId int64,
            msisdn string,
            format string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| listId |  ``` Required ```  | ID of the list to remove from. If set to 0 (zero) the member will be removed from all lists. |
| msisdn |  ``` Required ```  | Mobile number of the member |
| format |  ``` Required ```  ``` DefaultValue ```  | Response format e.g. json or xml |


#### Example Usage

```go
listId,_ := strconv.ParseInt("97", 10, 8)
msisdn := "msisdn"
format := "json"

var result interface{}
result,_ = lists.DeleteFromList(listId, msisdn, format)

```


### <a name="get_list"></a>![Method: ](https://apidocs.io/img/method.png ".lists_pkg.GetList") GetList

> Get information about a list and its members.


```go
func (me *LISTS_IMPL) GetList(
            listId string,
            format string,
            members *string,
            page *string,
            max *string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| listId |  ``` Required ```  | List ID |
| format |  ``` Required ```  ``` DefaultValue ```  | Response format e.g. json or xml |
| members |  ``` Optional ```  | Which types of members to return. Possible values: active - only get active members (default) inactive - only get inactive members all - get active and inactive members none - do not get any members, just metadata |
| page |  ``` Optional ```  | Page number, for pagination |
| max |  ``` Optional ```  | Maximum results returned per page |


#### Example Usage

```go
listId := "list_id"
format := "json"
members := "members"
page := "page"
max := "max"

var result interface{}
result,_ = lists.GetList(listId, format, members, page, max)

```


### <a name="get_lists"></a>![Method: ](https://apidocs.io/img/method.png ".lists_pkg.GetLists") GetLists

> Get the metadata of all your lists.


```go
func (me *LISTS_IMPL) GetLists(
            format string,
            page *string,
            max *string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| format |  ``` Required ```  ``` DefaultValue ```  | Response format e.g. json or xml |
| page |  ``` Optional ```  | Page number, for pagination |
| max |  ``` Optional ```  | Maximum results returned per page |


#### Example Usage

```go
format := "json"
page := "page"
max := "max"

var result interface{}
result,_ = lists.GetLists(format, page, max)

```


### <a name="edit_list_member"></a>![Method: ](https://apidocs.io/img/method.png ".lists_pkg.EditListMember") EditListMember

> Edit a member of a list.


```go
func (me *LISTS_IMPL) EditListMember(
            listId string,
            msisdn string,
            format string,
            firstName *string,
            lastName *string,
            field1 *string,
                queryParameters map[string]interface{})(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| listId |  ``` Required ```  | ID of the list the member belongs to |
| msisdn |  ``` Required ```  | Mobile number of the member to edit |
| format |  ``` Required ```  ``` DefaultValue ```  | Response format e.g. json or xml |
| firstName |  ``` Optional ```  | First name of the member |
| lastName |  ``` Optional ```  | Last name of the member |
| field1 |  ``` Optional ```  | Custom field value where n is an integer between 1 and 10. You can also use the names of the custom fields you have chosen for your list, e.g. field.birthday. To remove a value set it to an empty string. |
| queryParameters | ``` Optional ``` | Additional optional query parameters are supported by this method |


#### Example Usage

```go
listId := "list_id"
msisdn := "msisdn"
format := "json"
firstName := "first_name"
lastName := "last_name"
field1 := "field.1"
// key-value map for optional query parameters
	queryParams := map[string]interface{}{"key" : "value"}


var result interface{}
result,_ = lists.EditListMember(listId, msisdn, format, firstName, lastName, field1, queryParams)

```


### <a name="get_contact"></a>![Method: ](https://apidocs.io/img/method.png ".lists_pkg.GetContact") GetContact

> Get contact information from a list.


```go
func (me *LISTS_IMPL) GetContact(
            listId string,
            msisdn string,
            format string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| listId |  ``` Required ```  | ID of the list the contact is on. |
| msisdn |  ``` Required ```  | Mobile number of the contact. |
| format |  ``` Required ```  ``` DefaultValue ```  | Response format e.g. json or xml |


#### Example Usage

```go
listId := "list_id"
msisdn := "msisdn"
format := "json"

var result interface{}
result,_ = lists.GetContact(listId, msisdn, format)

```


### <a name="add_contacts_bulk"></a>![Method: ](https://apidocs.io/img/method.png ".lists_pkg.AddContactsBulk") AddContactsBulk

> Upload a list of contacts to Burst SMS


```go
func (me *LISTS_IMPL) AddContactsBulk(
            name string,
            fileUrl string,
            format string,
            countrycode *string,
            field1 *string,
                queryParameters map[string]interface{})(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| name |  ``` Required ```  | Name of the list |
| fileUrl |  ``` Required ```  | URL location of the contact list (NB: The list you are uploading requires a column labelled mobile) |
| format |  ``` Required ```  ``` DefaultValue ```  | Response format e.g. json or xml |
| countrycode |  ``` Optional ```  | Specifies which country the numbers are to be formatted in (e.g AU). If uploading numbers for multiple countries, do not define this, you will need to ensure that all the numbers are in correct international format before upload. |
| field1 |  ``` Optional ```  | Adds custom fields to the list. |
| queryParameters | ``` Optional ``` | Additional optional query parameters are supported by this method |


#### Example Usage

```go
name := "name"
fileUrl := "file_url"
format := "json"
countrycode := "countrycode"
field1 := "field_1"
// key-value map for optional query parameters
	queryParams := map[string]interface{}{"key" : "value"}


var result interface{}
result,_ = lists.AddContactsBulk(name, fileUrl, format, countrycode, field1, queryParams)

```


### <a name="remove_list"></a>![Method: ](https://apidocs.io/img/method.png ".lists_pkg.RemoveList") RemoveList

> Delete a list and its members.


```go
func (me *LISTS_IMPL) RemoveList(
            listId string,
            format string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| listId |  ``` Required ```  | List ID |
| format |  ``` Required ```  ``` DefaultValue ```  | Response format e.g. json or xml |


#### Example Usage

```go
listId := "list_id"
format := "json"

var result interface{}
result,_ = lists.RemoveList(listId, format)

```


[Back to List of Controllers](#list_of_controllers)



