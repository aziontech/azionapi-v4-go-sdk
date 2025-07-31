# DelegationSignerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlgorithmType** | [**AlgTypeRequest**](AlgTypeRequest.md) |  | 
**Digest** | **string** |  | 
**DigestType** | [**AlgTypeRequest**](AlgTypeRequest.md) |  | 
**KeyTag** | **int64** |  | 

## Methods

### NewDelegationSignerRequest

`func NewDelegationSignerRequest(algorithmType AlgTypeRequest, digest string, digestType AlgTypeRequest, keyTag int64, ) *DelegationSignerRequest`

NewDelegationSignerRequest instantiates a new DelegationSignerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDelegationSignerRequestWithDefaults

`func NewDelegationSignerRequestWithDefaults() *DelegationSignerRequest`

NewDelegationSignerRequestWithDefaults instantiates a new DelegationSignerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgorithmType

`func (o *DelegationSignerRequest) GetAlgorithmType() AlgTypeRequest`

GetAlgorithmType returns the AlgorithmType field if non-nil, zero value otherwise.

### GetAlgorithmTypeOk

`func (o *DelegationSignerRequest) GetAlgorithmTypeOk() (*AlgTypeRequest, bool)`

GetAlgorithmTypeOk returns a tuple with the AlgorithmType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithmType

`func (o *DelegationSignerRequest) SetAlgorithmType(v AlgTypeRequest)`

SetAlgorithmType sets AlgorithmType field to given value.


### GetDigest

`func (o *DelegationSignerRequest) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *DelegationSignerRequest) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *DelegationSignerRequest) SetDigest(v string)`

SetDigest sets Digest field to given value.


### GetDigestType

`func (o *DelegationSignerRequest) GetDigestType() AlgTypeRequest`

GetDigestType returns the DigestType field if non-nil, zero value otherwise.

### GetDigestTypeOk

`func (o *DelegationSignerRequest) GetDigestTypeOk() (*AlgTypeRequest, bool)`

GetDigestTypeOk returns a tuple with the DigestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigestType

`func (o *DelegationSignerRequest) SetDigestType(v AlgTypeRequest)`

SetDigestType sets DigestType field to given value.


### GetKeyTag

`func (o *DelegationSignerRequest) GetKeyTag() int64`

GetKeyTag returns the KeyTag field if non-nil, zero value otherwise.

### GetKeyTagOk

`func (o *DelegationSignerRequest) GetKeyTagOk() (*int64, bool)`

GetKeyTagOk returns a tuple with the KeyTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyTag

`func (o *DelegationSignerRequest) SetKeyTag(v int64)`

SetKeyTag sets KeyTag field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


