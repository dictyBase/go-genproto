# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [payload.proto](#payload.proto)
    - [Data](#dictybase.api.jsonapi.Data)
    - [DataCollection](#dictybase.api.jsonapi.DataCollection)
    - [Links](#dictybase.api.jsonapi.Links)
    - [Meta](#dictybase.api.jsonapi.Meta)
    - [Pagination](#dictybase.api.jsonapi.Pagination)
    - [PaginationLinks](#dictybase.api.jsonapi.PaginationLinks)
  
  
  
  

- [request.proto](#request.proto)
    - [DeleteRequest](#dictybase.api.jsonapi.DeleteRequest)
    - [ExistResponse](#dictybase.api.jsonapi.ExistResponse)
    - [GetEmailRequest](#dictybase.api.jsonapi.GetEmailRequest)
    - [GetRequest](#dictybase.api.jsonapi.GetRequest)
    - [GetRequestWithFields](#dictybase.api.jsonapi.GetRequestWithFields)
    - [HealthzIdRequest](#dictybase.api.jsonapi.HealthzIdRequest)
    - [IdRequest](#dictybase.api.jsonapi.IdRequest)
    - [IdResponse](#dictybase.api.jsonapi.IdResponse)
    - [ListRequest](#dictybase.api.jsonapi.ListRequest)
    - [RelationshipRequest](#dictybase.api.jsonapi.RelationshipRequest)
    - [RelationshipRequestWithPagination](#dictybase.api.jsonapi.RelationshipRequestWithPagination)
    - [SimpleListRequest](#dictybase.api.jsonapi.SimpleListRequest)
  
  
  
  

- [Scalar Value Types](#scalar-value-types)



<a name="payload.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## payload.proto



<a name="dictybase.api.jsonapi.Data"></a>

### Data
A [resource identifier object](http://jsonapi.org/format/#document-resource-identifier-objects).


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  | The resource name. |
| id | [int64](#int64) |  | Unique id. |






<a name="dictybase.api.jsonapi.DataCollection"></a>

### DataCollection
Definition for resource identifier collection objects.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| data | [Data](#dictybase.api.jsonapi.Data) | repeated |  |






<a name="dictybase.api.jsonapi.Links"></a>

### Links
A container for http links.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| self | [string](#string) |  | A http link. It points to the resource itself. |
| related | [string](#string) |  | A http link. It points to a related resource. |






<a name="dictybase.api.jsonapi.Meta"></a>

### Meta
Top level meta container.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pagination | [Pagination](#dictybase.api.jsonapi.Pagination) |  |  |






<a name="dictybase.api.jsonapi.Pagination"></a>

### Pagination
A container for various pagination properties


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| records | [int64](#int64) |  | Total number of entries, regardless of pages. |
| total | [int64](#int64) |  | Total number of pages. |
| size | [int64](#int64) |  | Number of entries per page. |
| number | [int64](#int64) |  | Current page number. |






<a name="dictybase.api.jsonapi.PaginationLinks"></a>

### PaginationLinks
A container for pagination links.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| self | [string](#string) |  | A http link to the resource itself. |
| next | [string](#string) |  | A http link to the next page of data. |
| prev | [string](#string) |  | A http link to the previous page of data. |
| last | [string](#string) |  | A http link to the last page of data. |
| first | [string](#string) |  | A http link to the first page of data. |





 

 

 

 



<a name="request.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## request.proto



<a name="dictybase.api.jsonapi.DeleteRequest"></a>

### DeleteRequest
A `DeleteRequest` defines the url parameter that must be passed
to remove a singular resource.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | An unique identifier, for example: &#34;/users/34&#34; |






<a name="dictybase.api.jsonapi.ExistResponse"></a>

### ExistResponse
ExistResponse wraps a boolean response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exist | [bool](#bool) |  | exist or non-existant |






<a name="dictybase.api.jsonapi.GetEmailRequest"></a>

### GetEmailRequest
A `GetEmailRequest` is identical to GetRequest except `email` id used as unique identifier.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| email | [string](#string) |  | Email id &#34;/users/newman@seinfeld.org&#34; |
| include | [string](#string) |  | include query parameter to retrieve any particular or particular combination of relationships. Multiple include values are delimited by comma(,).

For example, /{resource_name}/13?include=baz /{resource_name}/13?include=baz,bot |
| fields | [string](#string) |  | fields query parameter to retrieve any particular or any particular combination of attributes. Multiple fields values are delimited by comma(,).

For example, /{resource_name}/29?fields=foo /{resource_name}/?fields=foo,bar |






<a name="dictybase.api.jsonapi.GetRequest"></a>

### GetRequest
A `GetRequest` defines various url and query parameters that could be passed
in a HTTP **GET** request to a singular resource. Majority of the request
parameters are identical or similar to [jsonapi](http://jsonapi.org).


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | An unique identifier, for example: &#34;/users/34&#34; |
| include | [string](#string) |  | include query parameter to retrieve any particular or particular combination of relationships. Multiple include values are delimited by comma(,).

For example, /{resource_name}/13?include=baz /{resource_name}/13?include=baz,bot |
| fields | [string](#string) |  | fields query parameter to retrieve any particular or any particular combination of attributes. Multiple fields values are delimited by comma(,).

For example /{resource_name}/29?fields=foo /{resource_name}/?fields=foo,bar |






<a name="dictybase.api.jsonapi.GetRequestWithFields"></a>

### GetRequestWithFields
A `GetRequestWithFields` is a subset of GetRequest which only allow the fields parameter.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | An unique identifier, for example: &#34;/users/34&#34; |
| fields | [string](#string) |  | fields query parameter to retrieve any particular or any particular combination of attributes. Multiple fields values are delimited by comma(,).

For example /{resource_name}/29?fields=foo /{resource_name}/?fields=foo,bar |






<a name="dictybase.api.jsonapi.HealthzIdRequest"></a>

### HealthzIdRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="dictybase.api.jsonapi.IdRequest"></a>

### IdRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | An unique identifier |






<a name="dictybase.api.jsonapi.IdResponse"></a>

### IdResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | An unique identifier |






<a name="dictybase.api.jsonapi.ListRequest"></a>

### ListRequest
A `ListRequest` defines various url and query parameters that could be
passed in a HTTP **GET** request to a collection resource. All collection
resources are expected to support pagination. Majority of the request
parameters are identical or similar to [jsonapi](http://jsonapi.org).


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| include | [string](#string) |  | include query parameter to retrieve any particular or particular combination of relationships. Multiple include values are delimited by comma(,).

For example, /{resource_name}/13?include=baz /{resource_name}/13?include=baz,bot |
| fields | [string](#string) |  | fields query parameter to retrieve any particular or any particular combination of attributes. Multiple fields values are delimited by comma(,).

For example /{resource_name}/29?fields=foo /{resource_name}/?fields=foo,bar |
| pagenum | [int64](#int64) |  | The page number to fetch |
| pagesize | [int64](#int64) |  | Number of records per page |
| filter | [string](#string) |  | The `filter` query parameter restricts the data return by the collection. To use it, supply an attribute to filter, followed by a filter expression. It uses the following syntax... attribute operator expression attribute - Any one of the valid attribute of the resource. operator - Defines the type of filter match to use. It could be any of the following four and all of them should be URL-encoded.

 == Equals (URL encoding is %3D%3D) != Not equals =@ Contains substring !@ Not contains substring

expression - The value that will be included or excluded from the result. URL-reserved characters must be URL-encoded. For example, the following filter returns all users with last name `Gag`. /users?filter=last_name%3D%3Dgag

Filter can be combined using OR or AND boolean logic. * The OR is represented using a comma(,). * The AND is represented using a semi-colon(;). * AND and OR operators can be combined and AND takes precedence over OR. |






<a name="dictybase.api.jsonapi.RelationshipRequest"></a>

### RelationshipRequest
A `RelationshipRequest` defines the url parameter for relationship resources
that are given in the links field of relationship object


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | An unique identifier, for example: &#34;/users/45/roles&#34; or &#34;/users/45/relationships/roles&#34; |






<a name="dictybase.api.jsonapi.RelationshipRequestWithPagination"></a>

### RelationshipRequestWithPagination
A `RelationshipRequestWithPagination` is a `RelationshipRequest` with pagination


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| pagenum | [int64](#int64) |  |  |
| pagesize | [int64](#int64) |  |  |






<a name="dictybase.api.jsonapi.SimpleListRequest"></a>

### SimpleListRequest
A `SimpleListRequest` is identical to `ListRequest` except it does not support
pagination. The rest of the parameters are identical to `ListRequest` definition.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| include | [string](#string) |  | include query parameter to retrieve any particular or particular combination of relationships. Multiple include values are delimited by comma(,).

For example, /{resource_name}/13?include=baz /{resource_name}/13?include=baz,bot |
| fields | [string](#string) |  | fields query parameter to retrieve any particular or any particular combination of attributes. Multiple fields values are delimited by comma(,).

For example /{resource_name}/29?fields=foo /{resource_name}/?fields=foo,bar |
| filter | [string](#string) |  | The `filter` query parameter restricts the data return by the collection. To use it, supply an attribute to filter, followed by a filter expression. It uses the following syntax... attribute operator expression attribute - Any one of the valid attribute of the resource. operator - Defines the type of filter match to use. It could be any of the following four and all of them should be URL-encoded.

 == Equals (URL encoding is %3D%3D) != Not equals =@ Contains substring !@ Not contains substring

expression - The value that will be included or excluded from the result. URL-reserved characters must be URL-encoded. For example, the following filter returns all users with last name `Gag`. /users?filter=last_name%3D%3Dgag

Filter can be combined using OR or AND boolean logic. * The OR is represented using a comma(,). * The AND is represented using a semi-colon(;). * AND and OR operators can be combined and AND takes precedence over OR. |





 

 

 

 



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

