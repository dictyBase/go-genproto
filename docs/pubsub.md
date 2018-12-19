# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [pubsub.proto](#pubsub.proto)
    - [IdRequest](#dictybase.pubsub.IdRequest)
    - [IdentityReply](#dictybase.pubsub.IdentityReply)
    - [IdentityReq](#dictybase.pubsub.IdentityReq)
    - [Reply](#dictybase.pubsub.Reply)
    - [UserReply](#dictybase.pubsub.UserReply)
  
  
  
  

- [Scalar Value Types](#scalar-value-types)



<a name="pubsub.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## pubsub.proto



<a name="dictybase.pubsub.IdRequest"></a>

### IdRequest
Defintion for communicating an identifer


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="dictybase.pubsub.IdentityReply"></a>

### IdentityReply
Definition for replying various identity definition


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| identity | [dictybase.identity.Identity](#dictybase.identity.Identity) |  | single user |
| status | [google.rpc.Status](#google.rpc.Status) |  | Status error model, look here for detail https://github.com/googleapis/googleapis/blob/master/google/rpc/status.proto |
| exist | [bool](#bool) |  | Flag to indicate the presence of resource |






<a name="dictybase.pubsub.IdentityReq"></a>

### IdentityReq
Definition for requesting identity information


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | Unique id |
| identifier | [string](#string) |  | An unique identifier provided by the third party. Generally it&#39;s an email id, however it could be something else specifically provided by an provider. |
| provider | [string](#string) |  | Name of the provider, for example, orcid, google, facebook etc. |






<a name="dictybase.pubsub.Reply"></a>

### Reply
Definition for communicating the existence of a resource
and if there&#39;s any error during the communication


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exist | [bool](#bool) |  | Flag to indicate the presence of resource |
| status | [google.rpc.Status](#google.rpc.Status) |  | Status error model, look here for detail https://github.com/googleapis/googleapis/blob/master/google/rpc/status.proto |






<a name="dictybase.pubsub.UserReply"></a>

### UserReply
Definition for replying various user definition


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [dictybase.user.User](#dictybase.user.User) |  | single user |
| users | [dictybase.user.UserCollection](#dictybase.user.UserCollection) |  | user collection(list of users) |
| status | [google.rpc.Status](#google.rpc.Status) |  | Status error model, look here for detail https://github.com/googleapis/googleapis/blob/master/google/rpc/status.proto |
| exist | [bool](#bool) |  | Flag to indicate the presence of resource |





 

 

 

 



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

