# Output

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;standard&#x60; - Standard HTTP/HTTPS POST * &#x60;kafka&#x60; - Apache Kafka * &#x60;s3&#x60; - Simple Storage Service (S3) * &#x60;big_query&#x60; - Google BigQuery * &#x60;elasticsearch&#x60; - Elasticsearch * &#x60;splunk&#x60; - Splunk * &#x60;aws_kinesis_firehose&#x60; - AWS Kinesis Data Firehose * &#x60;datadog&#x60; - Datadog * &#x60;qradar&#x60; - IBM QRadar * &#x60;azure_monitor&#x60; - Azure Monitor * &#x60;azure_blob_storage&#x60; - Azure Blob Storage | 
**Attributes** | [**OutputPolymorphic**](OutputPolymorphic.md) |  | 

## Methods

### NewOutput

`func NewOutput(type_ string, attributes OutputPolymorphic, ) *Output`

NewOutput instantiates a new Output object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputWithDefaults

`func NewOutputWithDefaults() *Output`

NewOutputWithDefaults instantiates a new Output object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Output) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Output) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Output) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *Output) GetAttributes() OutputPolymorphic`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Output) GetAttributesOk() (*OutputPolymorphic, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Output) SetAttributes(v OutputPolymorphic)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


