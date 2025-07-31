# CertificateSigningRequestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Certificate** | Pointer to **NullableString** |  | [optional] 
**PrivateKey** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **string** | The value can&#39;t be changed after the certificate creation.  * &#x60;edge_certificate&#x60; - Edge Certificate * &#x60;trusted_ca_certificate&#x60; - Trusted CA Certificate | [optional] 
**KeyAlgorithm** | Pointer to **string** | * &#x60;rsa_2048&#x60; - 2048-bit RSA * &#x60;rsa_4096&#x60; - 4096-bit RSA * &#x60;ecc_384&#x60; - 384-bit Prime Field Curve | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**CommonName** | **string** |  | 
**AlternativeNames** | Pointer to **[]string** |  | [optional] 
**Country** | **string** |  | 
**State** | **string** |  | 
**Locality** | **string** |  | 
**Organization** | **string** |  | 
**OrganizationUnity** | **string** |  | 
**Email** | **string** |  | 

## Methods

### NewCertificateSigningRequestRequest

`func NewCertificateSigningRequestRequest(name string, commonName string, country string, state string, locality string, organization string, organizationUnity string, email string, ) *CertificateSigningRequestRequest`

NewCertificateSigningRequestRequest instantiates a new CertificateSigningRequestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateSigningRequestRequestWithDefaults

`func NewCertificateSigningRequestRequestWithDefaults() *CertificateSigningRequestRequest`

NewCertificateSigningRequestRequestWithDefaults instantiates a new CertificateSigningRequestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CertificateSigningRequestRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CertificateSigningRequestRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CertificateSigningRequestRequest) SetName(v string)`

SetName sets Name field to given value.


### GetCertificate

`func (o *CertificateSigningRequestRequest) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *CertificateSigningRequestRequest) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *CertificateSigningRequestRequest) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *CertificateSigningRequestRequest) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### SetCertificateNil

`func (o *CertificateSigningRequestRequest) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *CertificateSigningRequestRequest) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
### GetPrivateKey

`func (o *CertificateSigningRequestRequest) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *CertificateSigningRequestRequest) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *CertificateSigningRequestRequest) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *CertificateSigningRequestRequest) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### SetPrivateKeyNil

`func (o *CertificateSigningRequestRequest) SetPrivateKeyNil(b bool)`

 SetPrivateKeyNil sets the value for PrivateKey to be an explicit nil

### UnsetPrivateKey
`func (o *CertificateSigningRequestRequest) UnsetPrivateKey()`

UnsetPrivateKey ensures that no value is present for PrivateKey, not even an explicit nil
### GetType

`func (o *CertificateSigningRequestRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CertificateSigningRequestRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CertificateSigningRequestRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CertificateSigningRequestRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetKeyAlgorithm

`func (o *CertificateSigningRequestRequest) GetKeyAlgorithm() string`

GetKeyAlgorithm returns the KeyAlgorithm field if non-nil, zero value otherwise.

### GetKeyAlgorithmOk

`func (o *CertificateSigningRequestRequest) GetKeyAlgorithmOk() (*string, bool)`

GetKeyAlgorithmOk returns a tuple with the KeyAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyAlgorithm

`func (o *CertificateSigningRequestRequest) SetKeyAlgorithm(v string)`

SetKeyAlgorithm sets KeyAlgorithm field to given value.

### HasKeyAlgorithm

`func (o *CertificateSigningRequestRequest) HasKeyAlgorithm() bool`

HasKeyAlgorithm returns a boolean if a field has been set.

### GetActive

`func (o *CertificateSigningRequestRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CertificateSigningRequestRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CertificateSigningRequestRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CertificateSigningRequestRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCommonName

`func (o *CertificateSigningRequestRequest) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *CertificateSigningRequestRequest) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *CertificateSigningRequestRequest) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.


### GetAlternativeNames

`func (o *CertificateSigningRequestRequest) GetAlternativeNames() []string`

GetAlternativeNames returns the AlternativeNames field if non-nil, zero value otherwise.

### GetAlternativeNamesOk

`func (o *CertificateSigningRequestRequest) GetAlternativeNamesOk() (*[]string, bool)`

GetAlternativeNamesOk returns a tuple with the AlternativeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeNames

`func (o *CertificateSigningRequestRequest) SetAlternativeNames(v []string)`

SetAlternativeNames sets AlternativeNames field to given value.

### HasAlternativeNames

`func (o *CertificateSigningRequestRequest) HasAlternativeNames() bool`

HasAlternativeNames returns a boolean if a field has been set.

### GetCountry

`func (o *CertificateSigningRequestRequest) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CertificateSigningRequestRequest) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CertificateSigningRequestRequest) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetState

`func (o *CertificateSigningRequestRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CertificateSigningRequestRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CertificateSigningRequestRequest) SetState(v string)`

SetState sets State field to given value.


### GetLocality

`func (o *CertificateSigningRequestRequest) GetLocality() string`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *CertificateSigningRequestRequest) GetLocalityOk() (*string, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *CertificateSigningRequestRequest) SetLocality(v string)`

SetLocality sets Locality field to given value.


### GetOrganization

`func (o *CertificateSigningRequestRequest) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *CertificateSigningRequestRequest) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *CertificateSigningRequestRequest) SetOrganization(v string)`

SetOrganization sets Organization field to given value.


### GetOrganizationUnity

`func (o *CertificateSigningRequestRequest) GetOrganizationUnity() string`

GetOrganizationUnity returns the OrganizationUnity field if non-nil, zero value otherwise.

### GetOrganizationUnityOk

`func (o *CertificateSigningRequestRequest) GetOrganizationUnityOk() (*string, bool)`

GetOrganizationUnityOk returns a tuple with the OrganizationUnity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationUnity

`func (o *CertificateSigningRequestRequest) SetOrganizationUnity(v string)`

SetOrganizationUnity sets OrganizationUnity field to given value.


### GetEmail

`func (o *CertificateSigningRequestRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CertificateSigningRequestRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CertificateSigningRequestRequest) SetEmail(v string)`

SetEmail sets Email field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


