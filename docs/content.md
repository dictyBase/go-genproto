# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [content.proto](#content.proto)
    - [Content](#dictybase.content.Content)
    - [ContentAttributes](#dictybase.content.ContentAttributes)
    - [ContentData](#dictybase.content.ContentData)
    - [ContentIdRequest](#dictybase.content.ContentIdRequest)
    - [ContentRequest](#dictybase.content.ContentRequest)
    - [ExistingContentAttributes](#dictybase.content.ExistingContentAttributes)
    - [NewContentAttributes](#dictybase.content.NewContentAttributes)
    - [StoreContentRequest](#dictybase.content.StoreContentRequest)
    - [StoreContentRequest.Data](#dictybase.content.StoreContentRequest.Data)
    - [UpdateContentRequest](#dictybase.content.UpdateContentRequest)
    - [UpdateContentRequest.Data](#dictybase.content.UpdateContentRequest.Data)
  
  
  
    - [ContentService](#dictybase.content.ContentService)
  

- [Scalar Value Types](#scalar-value-types)



<a name="content.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## content.proto



<a name="dictybase.content.Content"></a>

### Content



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [ContentData](#dictybase.content.ContentData) |  |  |
| links | [dictybase.api.jsonapi.Links](#dictybase.api.jsonapi.Links) |  |  |






<a name="dictybase.content.ContentAttributes"></a>

### ContentAttributes
Definition of various content fields


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | page name |
| slug | [string](#string) |  | page slug. Look here https://en.wikipedia.org/wiki/Semantic_URL#Slug to know about slug |
| created_by | [int64](#int64) |  | id of the user who created the content |
| updated_by | [int64](#int64) |  | id of the user who updated the content |
| created_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | Timestamp for creation and update |
| updated_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| content | [string](#string) |  | serialized page content(for example serialized draft js object) |
| namespace | [string](#string) |  | namespace for the page |






<a name="dictybase.content.ContentData"></a>

### ContentData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  | The resource name |
| id | [int64](#int64) |  | Unique id |
| attributes | [ContentAttributes](#dictybase.content.ContentAttributes) |  |  |
| links | [dictybase.api.jsonapi.Links](#dictybase.api.jsonapi.Links) |  |  |






<a name="dictybase.content.ContentIdRequest"></a>

### ContentIdRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | Unique id to identify content |






<a name="dictybase.content.ContentRequest"></a>

### ContentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| slug | [string](#string) |  | Url slug Look here https://en.wikipedia.org/wiki/Semantic_URL#Slug to know about slug The slug name should be unique |






<a name="dictybase.content.ExistingContentAttributes"></a>

### ExistingContentAttributes
Fields that can be updated
Changing either or both of name and namespace
attributes alter the slug for the page


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| updated_by | [int64](#int64) |  | user id who is updating this content The existence of user will be verified(not implemented yet) using the `user` microservice backend. |
| content | [string](#string) |  | serialized page content(for example serialized draft js object) |






<a name="dictybase.content.NewContentAttributes"></a>

### NewContentAttributes
Definition for fields that are needed for storing the content


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | page name |
| created_by | [int64](#int64) |  | user id who is creating this content The existence of user will be verified(not implemented yet) using the `user` microservice backend. |
| content | [string](#string) |  | page content, expected to be serialized `JSON` string. |
| namespace | [string](#string) |  | namespace for the page, it is prepended to the name to generate an unique slug. |






<a name="dictybase.content.StoreContentRequest"></a>

### StoreContentRequest
Definition for storing new content


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [StoreContentRequest.Data](#dictybase.content.StoreContentRequest.Data) |  |  |






<a name="dictybase.content.StoreContentRequest.Data"></a>

### StoreContentRequest.Data



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  | resource name |
| attributes | [NewContentAttributes](#dictybase.content.NewContentAttributes) |  |  |






<a name="dictybase.content.UpdateContentRequest"></a>

### UpdateContentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [UpdateContentRequest.Data](#dictybase.content.UpdateContentRequest.Data) |  |  |
| id | [int64](#int64) |  |  |
| update_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  | An optional mask specifying which fields to update. Presence of this field allow partial updates. |






<a name="dictybase.content.UpdateContentRequest.Data"></a>

### UpdateContentRequest.Data



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  | resource name |
| id | [int64](#int64) |  | unique id |
| attributes | [ExistingContentAttributes](#dictybase.content.ExistingContentAttributes) |  |  |





 

 

 


<a name="dictybase.content.ContentService"></a>

### ContentService
The content service definition

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetContentBySlug | [ContentRequest](#dictybase.content.ContentRequest) | [Content](#dictybase.content.Content) | Gets the content of specified page(slug) |
| GetContent | [ContentIdRequest](#dictybase.content.ContentIdRequest) | [Content](#dictybase.content.Content) |  |
| StoreContent | [StoreContentRequest](#dictybase.content.StoreContentRequest) | [Content](#dictybase.content.Content) | Store the content of a new page(slug) |
| UpdateContent | [UpdateContentRequest](#dictybase.content.UpdateContentRequest) | [Content](#dictybase.content.Content) | Update the content of an existing page |
| DeleteContent | [ContentIdRequest](#dictybase.content.ContentIdRequest) | [.google.protobuf.Empty](#google.protobuf.Empty) | Delete an existing page along with its content |
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

