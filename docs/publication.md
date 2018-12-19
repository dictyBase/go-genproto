# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [publication.proto](#publication.proto)
    - [Author](#dictybase.publication.Author)
    - [ListPublicationParameters](#dictybase.publication.ListPublicationParameters)
    - [Meta](#dictybase.publication.Meta)
    - [NewPublication](#dictybase.publication.NewPublication)
    - [NewPublication.Data](#dictybase.publication.NewPublication.Data)
    - [Publication](#dictybase.publication.Publication)
    - [Publication.Data](#dictybase.publication.Publication.Data)
    - [PublicationAttributes](#dictybase.publication.PublicationAttributes)
    - [PublicationCollection](#dictybase.publication.PublicationCollection)
    - [PublicationCollection.Data](#dictybase.publication.PublicationCollection.Data)
    - [PublicationId](#dictybase.publication.PublicationId)
    - [PublicationUpdate](#dictybase.publication.PublicationUpdate)
    - [PublicationUpdate.Data](#dictybase.publication.PublicationUpdate.Data)
    - [PublicationUpdateAttributes](#dictybase.publication.PublicationUpdateAttributes)
  
  
  
    - [PublicationService](#dictybase.publication.PublicationService)
  

- [Scalar Value Types](#scalar-value-types)



<a name="publication.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## publication.proto



<a name="dictybase.publication.Author"></a>

### Author
Definition of an individual author


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| last_name | [string](#string) |  | Last name of the author |
| first_name | [string](#string) |  | First name of the author |
| initials | [string](#string) |  | Any initials of the author |
| rank | [int64](#int64) |  | Ranking of the author |






<a name="dictybase.publication.ListPublicationParameters"></a>

### ListPublicationParameters
Defines fields for manipulating output of Publication collection


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cursor | [int64](#int64) |  | A unique pointer to the next set of result in the list |
| limit | [int64](#int64) |  | Maximum number of records that can be fetch per request |
| filter | [string](#string) |  | The `filter` field restricts the data return by the collection. To use it, supply one or multiple allowed fields to filter followed by a filter expression. It uses the following syntax... field_name operator expression

The following fields of `PublicationAttributes` definition are allowed to be used for filtering * journal - Journal where the publication was published (string) * year - Year publication was published (string) * pub_date - Date of publication (string) * pub_type - Type of publication (string) * source - Source of the publication (string) * issue - Issue of the publication (string) * status - Status of the publication (string) * author - Authors of the publication (string)

field_name - Any one of the allowed field_name of the `PublicationAttributes` definition. operator - Defines the type of filter match to use. It could be any of the following four and all of them should be URL-encoded for http request.

 Operators for strings =~ Contains substring !~ Not contains substring === Equals !== Not equals

 Operators for number == Equals &gt; Greater than &lt; Less than =&lt; Less than equal to &gt;= Greater than equal to

expression - The value that will be included or excluded from the result. URL-reserved characters must be URL-encoded for http request.

 filter: &#34;pub_type===journal_article&#34; filter: &#34;source===pubmed&#34;

Filter can be combined using OR or AND boolean logic. * The OR is represented using a comma(,). * The AND is represented using a semi-colon(;). * AND and OR operators can be combined and AND takes precedence over OR.

 filter: &#34;year==2008;journal===Genesis&#34;

Can also accept multiple authors. filter: &#34;author===Vandelay;author===VanNostrand&#34;

The sort field allow to sort the data return by the collection based on fields of `PublicationAttributes. To use it, supply a comma separated one or more allowed field from the definition of `PublicationAttributes`. |






<a name="dictybase.publication.Meta"></a>

### Meta
Metadata definition for traversing the collection


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| next_cursor | [int64](#int64) |  | A unique pointer to the next set of result in the collection. Set the cursor value parameter to the value of next_cursor to retrieve the next set of collection using the same method |
| limit | [int64](#int64) |  | Maximum number of records that can be fetch per request |
| total | [int64](#int64) |  | Total number of records in the collection. |






<a name="dictybase.publication.NewPublication"></a>

### NewPublication
Definition for creating a new publication


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [NewPublication.Data](#dictybase.publication.NewPublication.Data) |  |  |






<a name="dictybase.publication.NewPublication.Data"></a>

### NewPublication.Data



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  | Resource name, by default should be publication |
| attributes | [PublicationAttributes](#dictybase.publication.PublicationAttributes) |  |  |






<a name="dictybase.publication.Publication"></a>

### Publication
Definition of an individual publication


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [Publication.Data](#dictybase.publication.Publication.Data) |  |  |






<a name="dictybase.publication.Publication.Data"></a>

### Publication.Data



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  | Resource name, by default should be publication |
| id | [string](#string) |  | Unique identifier for the publication |
| attributes | [PublicationAttributes](#dictybase.publication.PublicationAttributes) |  |  |






<a name="dictybase.publication.PublicationAttributes"></a>

### PublicationAttributes
Definition of various publication attributes


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| doi | [string](#string) |  | Digital object identifier for publication |
| title | [string](#string) |  | Title of publication |
| abstract | [string](#string) |  | Abstract of publication |
| journal | [string](#string) |  | Journal where the publication was published |
| pub_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | Date publication was published |
| volume | [string](#string) |  | Volume of the publication |
| pages | [string](#string) |  | Pages containing the publication |
| issn | [string](#string) |  | International Standard Serial Number of publication |
| pub_type | [string](#string) |  | Type of publication (i.e. &#34;journal_article&#34;) |
| source | [string](#string) |  | Source of the publication (i.e. &#34;pubmed&#34;) |
| issue | [string](#string) |  | Issue of the publication |
| status | [string](#string) |  | Status of the publication |
| authors | [Author](#dictybase.publication.Author) | repeated | List of authors of the publication |






<a name="dictybase.publication.PublicationCollection"></a>

### PublicationCollection
List of publications


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [PublicationCollection.Data](#dictybase.publication.PublicationCollection.Data) | repeated |  |
| meta | [Meta](#dictybase.publication.Meta) |  |  |






<a name="dictybase.publication.PublicationCollection.Data"></a>

### PublicationCollection.Data



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  | Resource name, by default should be publication |
| id | [string](#string) |  | Unique identifier for the publication |
| attributes | [PublicationAttributes](#dictybase.publication.PublicationAttributes) |  |  |






<a name="dictybase.publication.PublicationId"></a>

### PublicationId



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | Unique identifier for the publication |






<a name="dictybase.publication.PublicationUpdate"></a>

### PublicationUpdate
Definition for updating an existing publication


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [PublicationUpdate.Data](#dictybase.publication.PublicationUpdate.Data) |  |  |






<a name="dictybase.publication.PublicationUpdate.Data"></a>

### PublicationUpdate.Data



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  | Resource name, by default should be publication |
| id | [string](#string) |  | Unique identifier for the publication |
| attributes | [PublicationUpdateAttributes](#dictybase.publication.PublicationUpdateAttributes) |  |  |






<a name="dictybase.publication.PublicationUpdateAttributes"></a>

### PublicationUpdateAttributes
Defines attributes for updating an existing publication


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pub_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | Date of publication |
| volume | [string](#string) |  | Volume of the publication |
| pages | [string](#string) |  | Pages containing the publication |
| pub_type | [string](#string) |  | Type of publication (i.e. &#34;journal_article&#34;) |
| source | [string](#string) |  | Source of the publication (i.e. &#34;pubmed&#34;) |
| status | [string](#string) |  | Status of the publication |
| authors | [Author](#dictybase.publication.Author) | repeated | List of authors of the publication |





 

 

 


<a name="dictybase.publication.PublicationService"></a>

### PublicationService
The publication service specification

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetPublication | [PublicationId](#dictybase.publication.PublicationId) | [Publication](#dictybase.publication.Publication) | Retrieves publication by ID |
| CreatePublication | [NewPublication](#dictybase.publication.NewPublication) | [Publication](#dictybase.publication.Publication) | Create new publication |
| UpdatePublication | [PublicationUpdate](#dictybase.publication.PublicationUpdate) | [Publication](#dictybase.publication.Publication) | Update an existing publication |
| DeletePublication | [PublicationId](#dictybase.publication.PublicationId) | [.google.protobuf.Empty](#google.protobuf.Empty) | Delete an existing publication |
| ListPublications | [ListPublicationParameters](#dictybase.publication.ListPublicationParameters) | [PublicationCollection](#dictybase.publication.PublicationCollection) | List all publications |

 



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

