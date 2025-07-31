# ResponseBadRequestCertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **[]string** |  | [optional] 
**Certificate** | Pointer to **[]string** |  | [optional] 
**PrivateKey** | Pointer to **[]string** |  | [optional] 
**Issuer** | Pointer to **[]string** |  | [optional] 
**Validity** | Pointer to **[]string** |  | [optional] 
**SubjectName** | Pointer to [**ResponseBadRequestCertificateSubjectName**](ResponseBadRequestCertificateSubjectName.md) |  | [optional] 
**Type** | Pointer to **[]string** |  | [optional] 
**Managed** | Pointer to **[]string** |  | [optional] 
**Status** | Pointer to **[]string** |  | [optional] 
**StatusDetail** | Pointer to **[]string** |  | [optional] 
**Csr** | Pointer to **[]string** |  | [optional] 
**KeyAlgorithm** | Pointer to **[]string** |  | [optional] 
**Challenge** | Pointer to **[]string** |  | [optional] 
**Authority** | Pointer to **[]string** |  | [optional] 
**Active** | Pointer to **[]string** |  | [optional] 
**ProductVersion** | Pointer to **[]string** |  | [optional] 
**LastEditor** | Pointer to **[]string** |  | [optional] 
**LastModified** | Pointer to **[]string** |  | [optional] 
**RenewedAt** | Pointer to **[]string** |  | [optional] 
**Detail** | Pointer to **string** |  | [optional] 

## Methods

### NewResponseBadRequestCertificate

`func NewResponseBadRequestCertificate() *ResponseBadRequestCertificate`

NewResponseBadRequestCertificate instantiates a new ResponseBadRequestCertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseBadRequestCertificateWithDefaults

`func NewResponseBadRequestCertificateWithDefaults() *ResponseBadRequestCertificate`

NewResponseBadRequestCertificateWithDefaults instantiates a new ResponseBadRequestCertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResponseBadRequestCertificate) GetId() []string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponseBadRequestCertificate) GetIdOk() (*[]string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponseBadRequestCertificate) SetId(v []string)`

SetId sets Id field to given value.

### HasId

`func (o *ResponseBadRequestCertificate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ResponseBadRequestCertificate) GetName() []string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponseBadRequestCertificate) GetNameOk() (*[]string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponseBadRequestCertificate) SetName(v []string)`

SetName sets Name field to given value.

### HasName

`func (o *ResponseBadRequestCertificate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCertificate

`func (o *ResponseBadRequestCertificate) GetCertificate() []string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *ResponseBadRequestCertificate) GetCertificateOk() (*[]string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *ResponseBadRequestCertificate) SetCertificate(v []string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *ResponseBadRequestCertificate) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetPrivateKey

`func (o *ResponseBadRequestCertificate) GetPrivateKey() []string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *ResponseBadRequestCertificate) GetPrivateKeyOk() (*[]string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *ResponseBadRequestCertificate) SetPrivateKey(v []string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *ResponseBadRequestCertificate) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetIssuer

`func (o *ResponseBadRequestCertificate) GetIssuer() []string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *ResponseBadRequestCertificate) GetIssuerOk() (*[]string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *ResponseBadRequestCertificate) SetIssuer(v []string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *ResponseBadRequestCertificate) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetValidity

`func (o *ResponseBadRequestCertificate) GetValidity() []string`

GetValidity returns the Validity field if non-nil, zero value otherwise.

### GetValidityOk

`func (o *ResponseBadRequestCertificate) GetValidityOk() (*[]string, bool)`

GetValidityOk returns a tuple with the Validity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidity

`func (o *ResponseBadRequestCertificate) SetValidity(v []string)`

SetValidity sets Validity field to given value.

### HasValidity

`func (o *ResponseBadRequestCertificate) HasValidity() bool`

HasValidity returns a boolean if a field has been set.

### GetSubjectName

`func (o *ResponseBadRequestCertificate) GetSubjectName() ResponseBadRequestCertificateSubjectName`

GetSubjectName returns the SubjectName field if non-nil, zero value otherwise.

### GetSubjectNameOk

`func (o *ResponseBadRequestCertificate) GetSubjectNameOk() (*ResponseBadRequestCertificateSubjectName, bool)`

GetSubjectNameOk returns a tuple with the SubjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectName

`func (o *ResponseBadRequestCertificate) SetSubjectName(v ResponseBadRequestCertificateSubjectName)`

SetSubjectName sets SubjectName field to given value.

### HasSubjectName

`func (o *ResponseBadRequestCertificate) HasSubjectName() bool`

HasSubjectName returns a boolean if a field has been set.

### GetType

`func (o *ResponseBadRequestCertificate) GetType() []string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponseBadRequestCertificate) GetTypeOk() (*[]string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponseBadRequestCertificate) SetType(v []string)`

SetType sets Type field to given value.

### HasType

`func (o *ResponseBadRequestCertificate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetManaged

`func (o *ResponseBadRequestCertificate) GetManaged() []string`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *ResponseBadRequestCertificate) GetManagedOk() (*[]string, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *ResponseBadRequestCertificate) SetManaged(v []string)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *ResponseBadRequestCertificate) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### GetStatus

`func (o *ResponseBadRequestCertificate) GetStatus() []string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResponseBadRequestCertificate) GetStatusOk() (*[]string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResponseBadRequestCertificate) SetStatus(v []string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResponseBadRequestCertificate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDetail

`func (o *ResponseBadRequestCertificate) GetStatusDetail() []string`

GetStatusDetail returns the StatusDetail field if non-nil, zero value otherwise.

### GetStatusDetailOk

`func (o *ResponseBadRequestCertificate) GetStatusDetailOk() (*[]string, bool)`

GetStatusDetailOk returns a tuple with the StatusDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDetail

`func (o *ResponseBadRequestCertificate) SetStatusDetail(v []string)`

SetStatusDetail sets StatusDetail field to given value.

### HasStatusDetail

`func (o *ResponseBadRequestCertificate) HasStatusDetail() bool`

HasStatusDetail returns a boolean if a field has been set.

### GetCsr

`func (o *ResponseBadRequestCertificate) GetCsr() []string`

GetCsr returns the Csr field if non-nil, zero value otherwise.

### GetCsrOk

`func (o *ResponseBadRequestCertificate) GetCsrOk() (*[]string, bool)`

GetCsrOk returns a tuple with the Csr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsr

`func (o *ResponseBadRequestCertificate) SetCsr(v []string)`

SetCsr sets Csr field to given value.

### HasCsr

`func (o *ResponseBadRequestCertificate) HasCsr() bool`

HasCsr returns a boolean if a field has been set.

### GetKeyAlgorithm

`func (o *ResponseBadRequestCertificate) GetKeyAlgorithm() []string`

GetKeyAlgorithm returns the KeyAlgorithm field if non-nil, zero value otherwise.

### GetKeyAlgorithmOk

`func (o *ResponseBadRequestCertificate) GetKeyAlgorithmOk() (*[]string, bool)`

GetKeyAlgorithmOk returns a tuple with the KeyAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyAlgorithm

`func (o *ResponseBadRequestCertificate) SetKeyAlgorithm(v []string)`

SetKeyAlgorithm sets KeyAlgorithm field to given value.

### HasKeyAlgorithm

`func (o *ResponseBadRequestCertificate) HasKeyAlgorithm() bool`

HasKeyAlgorithm returns a boolean if a field has been set.

### GetChallenge

`func (o *ResponseBadRequestCertificate) GetChallenge() []string`

GetChallenge returns the Challenge field if non-nil, zero value otherwise.

### GetChallengeOk

`func (o *ResponseBadRequestCertificate) GetChallengeOk() (*[]string, bool)`

GetChallengeOk returns a tuple with the Challenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallenge

`func (o *ResponseBadRequestCertificate) SetChallenge(v []string)`

SetChallenge sets Challenge field to given value.

### HasChallenge

`func (o *ResponseBadRequestCertificate) HasChallenge() bool`

HasChallenge returns a boolean if a field has been set.

### GetAuthority

`func (o *ResponseBadRequestCertificate) GetAuthority() []string`

GetAuthority returns the Authority field if non-nil, zero value otherwise.

### GetAuthorityOk

`func (o *ResponseBadRequestCertificate) GetAuthorityOk() (*[]string, bool)`

GetAuthorityOk returns a tuple with the Authority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthority

`func (o *ResponseBadRequestCertificate) SetAuthority(v []string)`

SetAuthority sets Authority field to given value.

### HasAuthority

`func (o *ResponseBadRequestCertificate) HasAuthority() bool`

HasAuthority returns a boolean if a field has been set.

### GetActive

`func (o *ResponseBadRequestCertificate) GetActive() []string`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ResponseBadRequestCertificate) GetActiveOk() (*[]string, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ResponseBadRequestCertificate) SetActive(v []string)`

SetActive sets Active field to given value.

### HasActive

`func (o *ResponseBadRequestCertificate) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetProductVersion

`func (o *ResponseBadRequestCertificate) GetProductVersion() []string`

GetProductVersion returns the ProductVersion field if non-nil, zero value otherwise.

### GetProductVersionOk

`func (o *ResponseBadRequestCertificate) GetProductVersionOk() (*[]string, bool)`

GetProductVersionOk returns a tuple with the ProductVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVersion

`func (o *ResponseBadRequestCertificate) SetProductVersion(v []string)`

SetProductVersion sets ProductVersion field to given value.

### HasProductVersion

`func (o *ResponseBadRequestCertificate) HasProductVersion() bool`

HasProductVersion returns a boolean if a field has been set.

### GetLastEditor

`func (o *ResponseBadRequestCertificate) GetLastEditor() []string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *ResponseBadRequestCertificate) GetLastEditorOk() (*[]string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *ResponseBadRequestCertificate) SetLastEditor(v []string)`

SetLastEditor sets LastEditor field to given value.

### HasLastEditor

`func (o *ResponseBadRequestCertificate) HasLastEditor() bool`

HasLastEditor returns a boolean if a field has been set.

### GetLastModified

`func (o *ResponseBadRequestCertificate) GetLastModified() []string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *ResponseBadRequestCertificate) GetLastModifiedOk() (*[]string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *ResponseBadRequestCertificate) SetLastModified(v []string)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *ResponseBadRequestCertificate) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetRenewedAt

`func (o *ResponseBadRequestCertificate) GetRenewedAt() []string`

GetRenewedAt returns the RenewedAt field if non-nil, zero value otherwise.

### GetRenewedAtOk

`func (o *ResponseBadRequestCertificate) GetRenewedAtOk() (*[]string, bool)`

GetRenewedAtOk returns a tuple with the RenewedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewedAt

`func (o *ResponseBadRequestCertificate) SetRenewedAt(v []string)`

SetRenewedAt sets RenewedAt field to given value.

### HasRenewedAt

`func (o *ResponseBadRequestCertificate) HasRenewedAt() bool`

HasRenewedAt returns a boolean if a field has been set.

### GetDetail

`func (o *ResponseBadRequestCertificate) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ResponseBadRequestCertificate) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ResponseBadRequestCertificate) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *ResponseBadRequestCertificate) HasDetail() bool`

HasDetail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


