# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [annotation.proto](#annotation.proto)
    - [AnnotationId](#dictybase.annotation.AnnotationId)
    - [EntryAnnotationRequest](#dictybase.annotation.EntryAnnotationRequest)
    - [ListParameters](#dictybase.annotation.ListParameters)
    - [Meta](#dictybase.annotation.Meta)
    - [NewTaggedAnnotation](#dictybase.annotation.NewTaggedAnnotation)
    - [NewTaggedAnnotation.Data](#dictybase.annotation.NewTaggedAnnotation.Data)
    - [NewTaggedAnnotationAttributes](#dictybase.annotation.NewTaggedAnnotationAttributes)
    - [TaggedAnnotation](#dictybase.annotation.TaggedAnnotation)
    - [TaggedAnnotation.Data](#dictybase.annotation.TaggedAnnotation.Data)
    - [TaggedAnnotationAttributes](#dictybase.annotation.TaggedAnnotationAttributes)
    - [TaggedAnnotationCollection](#dictybase.annotation.TaggedAnnotationCollection)
    - [TaggedAnnotationCollection.Data](#dictybase.annotation.TaggedAnnotationCollection.Data)
    - [TaggedAnnotationUpdate](#dictybase.annotation.TaggedAnnotationUpdate)
    - [TaggedAnnotationUpdate.Data](#dictybase.annotation.TaggedAnnotationUpdate.Data)
    - [TaggedAnnotationUpdateAttributes](#dictybase.annotation.TaggedAnnotationUpdateAttributes)
  
  
  
    - [TaggedAnnotationService](#dictybase.annotation.TaggedAnnotationService)
  

- [Scalar Value Types](#scalar-value-types)



<a name="annotation.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## annotation.proto



<a name="dictybase.annotation.AnnotationId"></a>

### AnnotationId



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | unique identifier for the annotation |






<a name="dictybase.annotation.EntryAnnotationRequest"></a>

### EntryAnnotationRequest
Definition for various fields that are needed for fetching annotation for an
entry. The tag, ontology and entry_id must be provided, version and rank are
optional and their default values are used.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tag | [string](#string) |  | An identifiable tagname for the annotation, primarily a structured tag, generally an ontology term. |
| entry_id | [string](#string) |  | unique identifier of a biological entity that is being annotated |
| ontology | [string](#string) |  | Name of ontology in which the tag name is taken |
| rank | [int64](#int64) |  | Ordering of annotation when an entry has multiple annotations with identical tag from the same ontology. By default, rank 0 is assumed. |
| is_obsolete | [bool](#bool) |  | Status for active or retired annotation. Active annotation is chosen by default. |






<a name="dictybase.annotation.ListParameters"></a>

### ListParameters
ListParameters defines fields for manipulating output of TaggedAnnotation collection


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cursor | [int64](#int64) |  | A unique pointer to the next set of result in the list |
| limit | [int64](#int64) |  | Maximum number of records that can be fetch per request |
| filter | [string](#string) |  | The `filter` field restricts the data return by the collection. To use it, supply one or multiple allowed fields to filter followed by a filter expression. It uses the following syntax... field_name operator expression

The following fields of `AnnotationAttributes` definition are allowed to be used for filtering * entry_id - The entry that is being annotated (string) * value - The annotation in plain text format (string) * created_by - Email id of the user (string) * tag - Tag name, a term from an ontology (string). * ontology - Ontology that provides the tag names (string). * version - Version no (number).

field_name - Any one of the allowed field_name of the `AnnotationAttributes` definition. operator - Defines the type of filter match to use. It could be any of the following four and all of them should be URL-encoded for http request.

 Operators for strings =~ Contains substring !~ Not contains substring === Equals !== Not equals

 Operators for numbers == Equals != Not equals &gt; Greater than &lt; Less than =&lt; Less than equal to &gt;= Greater than equal to

expression - The value that will be included or excluded from the result. URL-reserved characters must be URL-encoded for http request.

 filter: &#34;created_by==caboose@abc.com&#34; filter: &#34;entry_id==DDB_G4839483&#34; filter: &#34;value~actin&#34;

Filter can be combined using OR or AND boolean logic. * The OR is represented using a comma(,). * The AND is represented using a semi-colon(;). * AND and OR operators can be combined and AND takes precedence over OR.

 filter: &#34;value~cytoskeletion;tag==cell membrane;ontology==cellular&#34;

The sort field allow to sort the data return by the collection based on fields of `AnnotationAttributes. To use it, supply a comma separated one or more allowed field from the definition of `AnnotationAttributes`. |






<a name="dictybase.annotation.Meta"></a>

### Meta
Metadata definition for traversing the collection


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| next_cursor | [int64](#int64) |  | A unique pointer to the next set of result in the collection. Set the cursor value parameter to the value of next_cursor to retrieve the next set of collection using the same method |
| limit | [int64](#int64) |  | Maximum number of records that can be fetch per request |
| total | [int64](#int64) |  | Total number of records in the collection. |






<a name="dictybase.annotation.NewTaggedAnnotation"></a>

### NewTaggedAnnotation
Definition for creating a new tagged annotation


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [NewTaggedAnnotation.Data](#dictybase.annotation.NewTaggedAnnotation.Data) |  |  |






<a name="dictybase.annotation.NewTaggedAnnotation.Data"></a>

### NewTaggedAnnotation.Data



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  | resource name, by default should be annotation |
| attributes | [NewTaggedAnnotationAttributes](#dictybase.annotation.NewTaggedAnnotationAttributes) |  |  |






<a name="dictybase.annotation.NewTaggedAnnotationAttributes"></a>

### NewTaggedAnnotationAttributes
NewTaggedAnnotation defines attributes for creating a new annotation


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [string](#string) |  | annotation in plain text format |
| editable_value | [string](#string) |  | serialized text content in a format recognized by frontend rich text editor |
| created_by | [string](#string) |  | Unique identifier(generally email) of the user who created the annotation |
| tag | [string](#string) |  | An identifiable tagname for the annotation, primarily a structured tag, generally an ontology term. |
| entry_id | [string](#string) |  | unique identifier of a biological entity that is being annotated |
| ontology | [string](#string) |  | Name of ontology from which the tag name is taken |
| rank | [int64](#int64) |  | Ordering of annotation when an entry has multiple annotations with identical tag from the same ontology. By default, rank 0 is used. |






<a name="dictybase.annotation.TaggedAnnotation"></a>

### TaggedAnnotation
Definition of an tag value based biological annotation where the tag always
represents a term from ontology.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [TaggedAnnotation.Data](#dictybase.annotation.TaggedAnnotation.Data) |  |  |






<a name="dictybase.annotation.TaggedAnnotation.Data"></a>

### TaggedAnnotation.Data



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  | resource name, by default should be annotation |
| id | [string](#string) |  | unique identifier for the annotation |
| attributes | [TaggedAnnotationAttributes](#dictybase.annotation.TaggedAnnotationAttributes) |  |  |






<a name="dictybase.annotation.TaggedAnnotationAttributes"></a>

### TaggedAnnotationAttributes
Definition of various tagged annotation attributes


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [string](#string) |  | annotation in plain text format |
| editable_value | [string](#string) |  | serialized text content in a format recognized by frontend rich text editor |
| created_by | [string](#string) |  | Unique identifier(generally email) of the user who created the annotation |
| created_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | Timestamp for creation |
| tag | [string](#string) |  | An identifiable tagname for the annotation, primarily a structured tag, generally an ontology term. |
| version | [int64](#int64) |  | version refers to the current version no |
| entry_id | [string](#string) |  | unique identifier of a biological entity that is being annotated |
| ontology | [string](#string) |  | Name of ontology in which the tag name is taken |
| rank | [int64](#int64) |  | Ordering of annotation when an entry has multiple annotations with identical tag from the same ontology. |
| is_obsolete | [bool](#bool) |  | Status for active or retired annotation |






<a name="dictybase.annotation.TaggedAnnotationCollection"></a>

### TaggedAnnotationCollection
List of paginated tagged annotations


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [TaggedAnnotationCollection.Data](#dictybase.annotation.TaggedAnnotationCollection.Data) | repeated |  |
| meta | [Meta](#dictybase.annotation.Meta) |  |  |






<a name="dictybase.annotation.TaggedAnnotationCollection.Data"></a>

### TaggedAnnotationCollection.Data



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  | resource name, by default should be annotation |
| id | [string](#string) |  | unique identifier for the annotation |
| attributes | [TaggedAnnotationAttributes](#dictybase.annotation.TaggedAnnotationAttributes) |  |  |






<a name="dictybase.annotation.TaggedAnnotationUpdate"></a>

### TaggedAnnotationUpdate
Definition for updating an existing annotation


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [TaggedAnnotationUpdate.Data](#dictybase.annotation.TaggedAnnotationUpdate.Data) |  |  |






<a name="dictybase.annotation.TaggedAnnotationUpdate.Data"></a>

### TaggedAnnotationUpdate.Data



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  | resource name, by default should be annotation |
| id | [string](#string) |  | unique identifier for the annotation |
| attributes | [TaggedAnnotationUpdateAttributes](#dictybase.annotation.TaggedAnnotationUpdateAttributes) |  |  |






<a name="dictybase.annotation.TaggedAnnotationUpdateAttributes"></a>

### TaggedAnnotationUpdateAttributes
TaggedUpdateAnnotation defines attributes for updating an existing annotation


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [string](#string) |  | annotation in plain text format |
| editable_value | [string](#string) |  | serialized text content in a format recognized by frontend rich text editor |
| created_by | [string](#string) |  | Unique identifier(generally email) of the user who created the annotation |





 

 

 


<a name="dictybase.annotation.TaggedAnnotationService"></a>

### TaggedAnnotationService
The tagged annotation service specification

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetAnnotation | [AnnotationId](#dictybase.annotation.AnnotationId) | [TaggedAnnotation](#dictybase.annotation.TaggedAnnotation) | Retrieves the specified tagged annotation |
| GetEntryAnnotation | [EntryAnnotationRequest](#dictybase.annotation.EntryAnnotationRequest) | [TaggedAnnotation](#dictybase.annotation.TaggedAnnotation) | Retrieves a single tagged annotation associated with an entry |
| ListAnnotations | [ListParameters](#dictybase.annotation.ListParameters) | [TaggedAnnotationCollection](#dictybase.annotation.TaggedAnnotationCollection) | Retrieves all tagged annotations |
| CreateAnnotation | [NewTaggedAnnotation](#dictybase.annotation.NewTaggedAnnotation) | [TaggedAnnotation](#dictybase.annotation.TaggedAnnotation) | Create a tagged annotation |
| UpdateAnnotation | [TaggedAnnotationUpdate](#dictybase.annotation.TaggedAnnotationUpdate) | [TaggedAnnotation](#dictybase.annotation.TaggedAnnotation) | Update an existing annotation, in this case a new annotation entry is created with a link to the previous annotation(copy on write). |
| DeleteAnnotation | [AnnotationId](#dictybase.annotation.AnnotationId) | [.google.protobuf.Empty](#google.protobuf.Empty) | Delete an existing annotation |

 



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

