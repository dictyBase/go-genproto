# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [identity.proto](#identity.proto)
    - [CreateIdentityReq](#dictybase.identity.CreateIdentityReq)
    - [CreateIdentityReq.Data](#dictybase.identity.CreateIdentityReq.Data)
    - [Identity](#dictybase.identity.Identity)
    - [IdentityAttributes](#dictybase.identity.IdentityAttributes)
    - [IdentityData](#dictybase.identity.IdentityData)
    - [IdentityProviderReq](#dictybase.identity.IdentityProviderReq)
    - [NewIdentityAttributes](#dictybase.identity.NewIdentityAttributes)
  
  
  
    - [IdentityService](#dictybase.identity.IdentityService)
  

- [Scalar Value Types](#scalar-value-types)



<a name="identity.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## identity.proto



<a name="dictybase.identity.CreateIdentityReq"></a>

### CreateIdentityReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [CreateIdentityReq.Data](#dictybase.identity.CreateIdentityReq.Data) |  |  |






<a name="dictybase.identity.CreateIdentityReq.Data"></a>

### CreateIdentityReq.Data



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  | resource name |
| attributes | [NewIdentityAttributes](#dictybase.identity.NewIdentityAttributes) |  |  |






<a name="dictybase.identity.Identity"></a>

### Identity



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [IdentityData](#dictybase.identity.IdentityData) |  |  |
| links | [dictybase.api.jsonapi.Links](#dictybase.api.jsonapi.Links) |  |  |






<a name="dictybase.identity.IdentityAttributes"></a>

### IdentityAttributes
Definition for various fields


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| identifier | [string](#string) |  | An unique identifier provided by the third party. Generally it&#39;s an email id, however it could be something else specifically provided by an provider. |
| provider | [string](#string) |  | Name of the provider, for example, orcid, google, facebook etc. |
| user_id | [int64](#int64) |  | The id of the user to which this identity is connected. This id could be used to fetch a complete user response from the user service |
| created_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | Timestamp for creation and update |
| updated_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="dictybase.identity.IdentityData"></a>

### IdentityData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  | The resource name |
| id | [int64](#int64) |  | Unique id |
| attributes | [IdentityAttributes](#dictybase.identity.IdentityAttributes) |  |  |
| links | [dictybase.api.jsonapi.Links](#dictybase.api.jsonapi.Links) |  |  |






<a name="dictybase.identity.IdentityProviderReq"></a>

### IdentityProviderReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| identifier | [string](#string) |  | An unique identifier provided by the third party. Generally it&#39;s an email id, however it could be something else specifically provided by an provider. |
| provider | [string](#string) |  | Name of the provider, for example, orcid, google, facebook etc. |






<a name="dictybase.identity.NewIdentityAttributes"></a>

### NewIdentityAttributes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| identifier | [string](#string) |  | An unique identifier provided by the third party. Generally it&#39;s an email id, however it could be something else specifically provided by an provider. |
| provider | [string](#string) |  | Name of the provider, for example, orcid, google, facebook etc. |
| user_id | [int64](#int64) |  | The id of the user to which this identity is connected. This id could be used to fetch a complete user response from the user service |





 

 

 


<a name="dictybase.identity.IdentityService"></a>

### IdentityService
The content service definition

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetIdentityFromProvider | [IdentityProviderReq](#dictybase.identity.IdentityProviderReq) | [Identity](#dictybase.identity.Identity) | Gets the specified identity |
| GetIdentity | [.dictybase.api.jsonapi.IdRequest](#dictybase.api.jsonapi.IdRequest) | [Identity](#dictybase.identity.Identity) |  |
| ExistProviderIdentity | [IdentityProviderReq](#dictybase.identity.IdentityProviderReq) | [.dictybase.api.jsonapi.ExistResponse](#dictybase.api.jsonapi.ExistResponse) |  |
| CreateIdentity | [CreateIdentityReq](#dictybase.identity.CreateIdentityReq) | [Identity](#dictybase.identity.Identity) | Create a new identity |
| DeleteIdentity | [.dictybase.api.jsonapi.IdRequest](#dictybase.api.jsonapi.IdRequest) | [.google.protobuf.Empty](#google.protobuf.Empty) | Delete an existing identity |
| Healthz | [.dictybase.api.jsonapi.HealthzIdRequest](#dictybase.api.jsonapi.HealthzIdRequest) | [.google.protobuf.Empty](#google.protobuf.Empty) | Basic health check that always return success |

 



## Scalar Value Types

| .proto Type | Notes | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | --------- | ----------- |
| <a name="double" /> double |  | double | double | float |
| <a name="float" /> float |  | float | float | float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long |
| <a name="bool" /> bool |  | bool | boolean | boolean |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str |

