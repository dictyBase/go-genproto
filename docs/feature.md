# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [feature.proto](#feature.proto)
    - [Dbxref](#dictybase.feature.Dbxref)
    - [Feature](#dictybase.feature.Feature)
    - [FeatureAttributes](#dictybase.feature.FeatureAttributes)
    - [FeatureCollection](#dictybase.feature.FeatureCollection)
    - [FeatureConnection](#dictybase.feature.FeatureConnection)
    - [FeatureData](#dictybase.feature.FeatureData)
    - [FeatureDbxref](#dictybase.feature.FeatureDbxref)
    - [FeatureHistory](#dictybase.feature.FeatureHistory)
    - [FeatureId](#dictybase.feature.FeatureId)
    - [FeatureLocation](#dictybase.feature.FeatureLocation)
    - [FeaturePublication](#dictybase.feature.FeaturePublication)
    - [FeatureRelationFilter](#dictybase.feature.FeatureRelationFilter)
    - [FeatureRelationship](#dictybase.feature.FeatureRelationship)
    - [FeatureSynonym](#dictybase.feature.FeatureSynonym)
    - [FeatureUpdate](#dictybase.feature.FeatureUpdate)
    - [FeatureUpdate.Data](#dictybase.feature.FeatureUpdate.Data)
    - [FeatureUpdateAttributes](#dictybase.feature.FeatureUpdateAttributes)
    - [ListParameters](#dictybase.feature.ListParameters)
    - [LocatedFeatureFilter](#dictybase.feature.LocatedFeatureFilter)
    - [NewFeature](#dictybase.feature.NewFeature)
    - [NewFeature.Data](#dictybase.feature.NewFeature.Data)
    - [NewFeatureAttributes](#dictybase.feature.NewFeatureAttributes)
    - [NewFeatureLocation](#dictybase.feature.NewFeatureLocation)
    - [Organism](#dictybase.feature.Organism)
    - [PaginatedFeatureCollection](#dictybase.feature.PaginatedFeatureCollection)
    - [Publication](#dictybase.feature.Publication)
    - [ReferenceFeatureFilter](#dictybase.feature.ReferenceFeatureFilter)
  
  
  
    - [FeatureService](#dictybase.feature.FeatureService)
  

- [Scalar Value Types](#scalar-value-types)



<a name="feature.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## feature.proto



<a name="dictybase.feature.Dbxref"></a>

### Dbxref
An identifier typically from a bioinformatics database(NCBI,Uniprot,Ensembl etc.)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| identifier | [string](#string) |  | Identifier |
| version | [int64](#int64) |  |  |
| database | [string](#string) |  | Source database |






<a name="dictybase.feature.Feature"></a>

### Feature
Definition of a biological entity


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [FeatureData](#dictybase.feature.FeatureData) |  |  |
| links | [dictybase.api.jsonapi.Links](#dictybase.api.jsonapi.Links) |  |  |






<a name="dictybase.feature.FeatureAttributes"></a>

### FeatureAttributes
Definition of various feature fields


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Short human readable textual name |
| id | [string](#string) |  |  |
| created_by | [int64](#int64) |  | Identifier of the user who created the feature |
| updated_by | [int64](#int64) |  | Identifier of the user who updated the feature |
| created_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | Timestamp for creation and update |
| updated_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| dbxref | [FeatureDbxref](#dictybase.feature.FeatureDbxref) | repeated |  |
| organism | [Organism](#dictybase.feature.Organism) |  |  |
| is_obsolete | [bool](#bool) |  | Toggle the obsolete status |
| parents | [FeatureRelationship](#dictybase.feature.FeatureRelationship) | repeated | List of biological parent features |
| children | [FeatureRelationship](#dictybase.feature.FeatureRelationship) | repeated | List of biological children feature |
| location | [FeatureLocation](#dictybase.feature.FeatureLocation) | repeated | List of feature locations in the reference backend |
| synonyms | [string](#string) | repeated | Alternate list of names |
| publications | [Publication](#dictybase.feature.Publication) | repeated |  |
| is_analysis | [bool](#bool) |  | Indicates if the feature is generated(annotated) from the result of an automated analysis |
| version | [int64](#int64) |  | Version of this feature, it increase by 1 for the replacement |
| previous | [string](#string) |  | Earlier instance of the feature that this one has replaced |
| next | [string](#string) |  | Next instance of the feature that this one got replaced with |
| dbxrefs | [Dbxref](#dictybase.feature.Dbxref) | repeated |  |






<a name="dictybase.feature.FeatureCollection"></a>

### FeatureCollection
List of features


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [FeatureData](#dictybase.feature.FeatureData) | repeated |  |
| links | [dictybase.api.jsonapi.Links](#dictybase.api.jsonapi.Links) |  |  |






<a name="dictybase.feature.FeatureConnection"></a>

### FeatureConnection
Definition for connecting two biologically related(parent or
children) features


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| relationship | [FeatureRelationship](#dictybase.feature.FeatureRelationship) |  |  |






<a name="dictybase.feature.FeatureData"></a>

### FeatureData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  | The biological type of feature, generally a sequence ontology term |
| id | [string](#string) |  | Unique id |
| attributes | [FeatureAttributes](#dictybase.feature.FeatureAttributes) |  |  |






<a name="dictybase.feature.FeatureDbxref"></a>

### FeatureDbxref
Definition for managing dbxref with a feature


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| dbxref | [Dbxref](#dictybase.feature.Dbxref) |  |  |






<a name="dictybase.feature.FeatureHistory"></a>

### FeatureHistory
Definition for managing a previous or next feature in the feature
history


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| linked_feature | [string](#string) |  |  |






<a name="dictybase.feature.FeatureId"></a>

### FeatureId



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | Unique feature identifier |






<a name="dictybase.feature.FeatureLocation"></a>

### FeatureLocation
The location of a feature relative to another feature using space-based(interbase) coordinates.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| source_feature | [string](#string) |  | Source or reference feature(identifier) which this location is relative to |
| start | [int64](#int64) |  | start/leftmost boundary of this location in linear range based on interbase coordinates. To convert it to base-oriented system, add 1 to this value. |
| end | [int64](#int64) |  | End/rightmost boundary of this location in linear range based on interbase coordinates. No conversion is needed to convert it to a base oriented system. |
| strand | [int64](#int64) |  | Orientation of this location, should be one of 0,1, or -1 |
| phase | [int64](#int64) |  | Phase of translation with respect to source feature, should be one of 0,1,or 2. |
| residue | [string](#string) |  | Alternate residues which differ from feature residues |
| group | [int64](#int64) |  | Represents derivable extra location of a feature. The default value 0 is used for direct location. Any other derived location gets values greater than 0. For example, the position of an exon on a BAC and in a global chromosomal coordinates. |
| rank | [int64](#int64) |  | Used for feature with multiple locations, 0 is used as default. For example, a blast HSP has two locations, one on the query and one on the subject. |






<a name="dictybase.feature.FeaturePublication"></a>

### FeaturePublication
Definition for managing publication with a feature


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| publication | [FeaturePublication](#dictybase.feature.FeaturePublication) |  |  |






<a name="dictybase.feature.FeatureRelationFilter"></a>

### FeatureRelationFilter



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| relation | [string](#string) |  | Type of relationship, generally a sequence ontology term |






<a name="dictybase.feature.FeatureRelationship"></a>

### FeatureRelationship
Definition for a related feature


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| relation | [string](#string) |  | Type of relationship, generally a sequence ontology term |
| feature | [string](#string) |  | Related feature |






<a name="dictybase.feature.FeatureSynonym"></a>

### FeatureSynonym
Defintion for managing synonym of a feature


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| synonym | [string](#string) |  |  |






<a name="dictybase.feature.FeatureUpdate"></a>

### FeatureUpdate
Definition for updating a feature


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [FeatureUpdate.Data](#dictybase.feature.FeatureUpdate.Data) |  |  |
| update_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  | An optional mask specifying which fields to update. Presence of this field allow partial updates. |
| id | [string](#string) |  |  |






<a name="dictybase.feature.FeatureUpdate.Data"></a>

### FeatureUpdate.Data



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  | The biological type of feature, generally a sequence ontology term |
| attributes | [FeatureUpdateAttributes](#dictybase.feature.FeatureUpdateAttributes) |  |  |
| id | [string](#string) |  |  |






<a name="dictybase.feature.FeatureUpdateAttributes"></a>

### FeatureUpdateAttributes
Definition of feature fields for updating a feature


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Short human readable textual name |
| id | [string](#string) |  |  |
| updated_by | [int64](#int64) |  | Identifier of the user who updated the feature |
| dbxref | [FeatureDbxref](#dictybase.feature.FeatureDbxref) | repeated |  |
| organism | [Organism](#dictybase.feature.Organism) |  |  |
| parents | [FeatureRelationship](#dictybase.feature.FeatureRelationship) | repeated | List of biological parent features |
| children | [FeatureRelationship](#dictybase.feature.FeatureRelationship) | repeated | List of biological children feature |
| location | [FeatureLocation](#dictybase.feature.FeatureLocation) | repeated |  |
| synonyms | [string](#string) | repeated | aLTERNATe list of names |
| publications | [Publication](#dictybase.feature.Publication) | repeated |  |
| is_analysis | [bool](#bool) |  | Indicates if the feature is generated(annotated) from the result of an automated analysis |
| previous | [string](#string) |  | Earlier instance of the feature that this one has replaced |
| next | [string](#string) |  | Next instance of the feature that this one got replaced with |
| dbxrefs | [Dbxref](#dictybase.feature.Dbxref) | repeated |  |






<a name="dictybase.feature.ListParameters"></a>

### ListParameters
ListParameters defines fields for manipulating output of PaginatedFeatureCollection


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pagenum | [int64](#int64) |  | The page number to fetch |
| pagesize | [int64](#int64) |  | Number of records per page |
| filter | [string](#string) |  | The `filter` field restricts the data return by the collection. To use it, supply one or multiple allowed fields to filter followed by a filter expression. It uses the following syntax... field_name operator value

field_name - Any one of the allowed field name that are given below.

 * name - feature name (string). * type - The biological type of feature, generally a sequence ontology term (string) * species - Organism species (string) * genus - Organism genus (string) * strain - Strain name of the organism (string) * taxon_id - Taxon identifier (string) * obsolete - Live or retired feature (bool) * analysis - Feature as a result of analysis (bool)

operator - Defines the type of filter match to use. It could be any of the following four and all of them should be URL-encoded for http request.

 == Equals (URL encoding is %3D%3D) != Not equals ~ Contains/matches substring (case insensitive) !~ Not contains substring (case insensitive) &gt; Greater than &lt; Less than =&lt; Less than equal to &gt;= Greater than equal to

value - The value must be a string, a number or a boolean. URL-reserved characters must be URL-encoded for http request.

 filter: &#34;name==pcA&#34; filter: &#34;name~scr&#34; filter: &#34;obsolete==true&#34;

Filter can be combined using OR or AND boolean logic. * The OR is represented using a comma(,). * The AND is represented using a semi-colon(;). * AND and OR operators can be combined and AND takes precedence over OR.

 filter: &#34;name==pcA;obsolete==true;type=exon&#34; filter: &#34;name==transcription,name==gtx&#34;

The sort field allow to sort the data return by the collection based on allowed fields. To use it, supply a comma separated one or more allowed fields. |






<a name="dictybase.feature.LocatedFeatureFilter"></a>

### LocatedFeatureFilter



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | unique id of the reference feature |
| type | [string](#string) |  | The biological type of feature, generally a sequence ontology term. |
| filter | [string](#string) |  | The `filter` field restricts no of features in the collection by using a combination of custom start and end coordinates of reference feature. To use it, supply the field name followed by the filter expression. It uses the following syntax... field_name operator value

field_name - Any one of the allowed field name that are given below.

 * start - leftmost boundary of location (number). * end - rightmost boundary of location (number).

operator - Defines the type of filter match to use. It could be any of the following and all of them should be URL-encoded for http request.

 &gt; Greater than &lt; Less than =&lt; Less than equal to &gt;= Greater than equal to

value - The value must be a number here. If only one of the value is given, the coordinate of reference feature is for the other by default. URL-reserved characters must be URL-encoded for http request.

 filter: &#34;start&gt;=100&#34; filter: &#34;end&lt;789&#34;

Filter can be combined using AND boolean logic. The AND is represented using a semi-colon(;).

 filter: &#34;start&gt;=89;end&lt;=7430&#34; |






<a name="dictybase.feature.NewFeature"></a>

### NewFeature
Definition for creating new feature


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [NewFeature.Data](#dictybase.feature.NewFeature.Data) |  |  |






<a name="dictybase.feature.NewFeature.Data"></a>

### NewFeature.Data



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  | The biological type of feature, generally a sequence ontology term |
| attributes | [NewFeatureAttributes](#dictybase.feature.NewFeatureAttributes) |  |  |






<a name="dictybase.feature.NewFeatureAttributes"></a>

### NewFeatureAttributes
Definition of various feature fields for creating new feature


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Short human readable textual name |
| created_by | [int64](#int64) |  | Identifier of the user who created the feature |
| updated_by | [int64](#int64) |  | Identifier of the user who updated the feature |
| dbxrefs | [Dbxref](#dictybase.feature.Dbxref) | repeated |  |
| organism | [Organism](#dictybase.feature.Organism) |  |  |
| is_obsolete | [bool](#bool) |  | Toggle the obsolete status |
| parents | [FeatureRelationship](#dictybase.feature.FeatureRelationship) | repeated | List of biological parent features |
| children | [FeatureRelationship](#dictybase.feature.FeatureRelationship) | repeated | List of biological children feature |
| location | [FeatureLocation](#dictybase.feature.FeatureLocation) | repeated |  |
| synonyms | [string](#string) | repeated | Alternate list of names |
| publications | [Publication](#dictybase.feature.Publication) | repeated |  |
| is_analysis | [bool](#bool) |  | Indicates if the feature is generated(annotated) from the result of an automated analysis |
| version | [int64](#int64) |  | Version of this feature, it increase by 1 for the replacement |
| previous | [string](#string) |  | Earlier instance of the feature that this one has replaced |
| next | [string](#string) |  | Next instance of the feature that this one got replaced with |






<a name="dictybase.feature.NewFeatureLocation"></a>

### NewFeatureLocation
Definition for attaching a new location to a feature


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| location | [FeatureLocation](#dictybase.feature.FeatureLocation) |  |  |






<a name="dictybase.feature.Organism"></a>

### Organism
Defintion of an organism entry


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| taxon_id | [string](#string) |  | Taxon identifier |
| genus | [string](#string) |  | Organism genus |
| species | [string](#string) |  | Organism species |
| common_name | [string](#string) |  |  |
| rank | [string](#string) |  | Taxonomic rank |
| scientific_name | [string](#string) |  | Scientific name |
| strain | [string](#string) |  | Strain name |






<a name="dictybase.feature.PaginatedFeatureCollection"></a>

### PaginatedFeatureCollection
List of features with pagination


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [FeatureData](#dictybase.feature.FeatureData) | repeated |  |
| links | [dictybase.api.jsonapi.PaginationLinks](#dictybase.api.jsonapi.PaginationLinks) |  |  |
| pagination | [dictybase.api.jsonapi.Pagination](#dictybase.api.jsonapi.Pagination) |  |  |






<a name="dictybase.feature.Publication"></a>

### Publication
Container for linked publications


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| publication_id | [string](#string) |  | An identifier that refers to a publication |
| source | [string](#string) |  | Source of this publication identifier |






<a name="dictybase.feature.ReferenceFeatureFilter"></a>

### ReferenceFeatureFilter



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| rank | [int64](#int64) |  | Used for feature with multiple locations, 0 is used as default. |





 

 

 


<a name="dictybase.feature.FeatureService"></a>

### FeatureService
The feature service specification

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetFeature | [FeatureId](#dictybase.feature.FeatureId) | [Feature](#dictybase.feature.Feature) | Retrieves the specified feature |
| GetParents | [FeatureRelationFilter](#dictybase.feature.FeatureRelationFilter) | [FeatureCollection](#dictybase.feature.FeatureCollection) | Retrieves all the biological parents features |
| GetChildren | [FeatureRelationFilter](#dictybase.feature.FeatureRelationFilter) | [FeatureCollection](#dictybase.feature.FeatureCollection) | Retrieves all biological children features |
| GetReferenceFeature | [ReferenceFeatureFilter](#dictybase.feature.ReferenceFeatureFilter) | [Feature](#dictybase.feature.Feature) | Retrieves the feature(called reference or source feature) under whose coordinate system the current feature is located. By default rank 0 is assumed |
| GetReferenceFeatures | [FeatureId](#dictybase.feature.FeatureId) | [FeatureCollection](#dictybase.feature.FeatureCollection) | Retrieves all reference features |
| GetLocatedFeatures | [LocatedFeatureFilter](#dictybase.feature.LocatedFeatureFilter) | [PaginatedFeatureCollection](#dictybase.feature.PaginatedFeatureCollection) | Retrieves all features that are located within the bounds of referernce feature |
| ListFeatures | [ListParameters](#dictybase.feature.ListParameters) | [PaginatedFeatureCollection](#dictybase.feature.PaginatedFeatureCollection) | Retrieves all features |
| CreateFeature | [NewFeature](#dictybase.feature.NewFeature) | [Feature](#dictybase.feature.Feature) | Create a new feature |
| UpdateFeature | [FeatureUpdate](#dictybase.feature.FeatureUpdate) | [Feature](#dictybase.feature.Feature) | Update an existing feature |
| AddParentalRelationship | [FeatureConnection](#dictybase.feature.FeatureConnection) | [Feature](#dictybase.feature.Feature) | Add a new parental relation with an existing feature. In case of an existing parent(s), it will be appended to the list. |
| AddChildRelationship | [FeatureConnection](#dictybase.feature.FeatureConnection) | [Feature](#dictybase.feature.Feature) | Add a new child relation with an existing feature. In case of an existing child, it will be appended to the list. |
| AttachLocation | [NewFeatureLocation](#dictybase.feature.NewFeatureLocation) | [Feature](#dictybase.feature.Feature) | Attach a new location entry to connect to a referernce feature. In case of any existing location(s), it will be appended to the list. The reference feature referred in the location needs to exist. |
| AddSynonym | [FeatureSynonym](#dictybase.feature.FeatureSynonym) | [Feature](#dictybase.feature.Feature) | Add a new synonym.In case of existing synonym(s), it will be appended to the list |
| AddPublication | [FeaturePublication](#dictybase.feature.FeaturePublication) | [Feature](#dictybase.feature.Feature) | Add a new publication.In case of existing publication(s), it will append it to the list |
| SetPrevFeatureHistory | [FeatureHistory](#dictybase.feature.FeatureHistory) | [Feature](#dictybase.feature.Feature) | Sets the previous feature in the feature history. |
| SetNextFeatureHistory | [FeatureHistory](#dictybase.feature.FeatureHistory) | [Feature](#dictybase.feature.Feature) | Sets the next feature in the feature history. |
| AddDbxref | [FeatureDbxref](#dictybase.feature.FeatureDbxref) | [Feature](#dictybase.feature.Feature) | Add a new dbxref. In case of existing dbxref it will append it to the list. |
| CreateOrganism | [Organism](#dictybase.feature.Organism) | [Organism](#dictybase.feature.Organism) | Create a new organism entry |
| DeleteFeature | [FeatureId](#dictybase.feature.FeatureId) | [.google.protobuf.Empty](#google.protobuf.Empty) | Delete an existing feature |
| DeleteParentalRelationship | [FeatureConnection](#dictybase.feature.FeatureConnection) | [.google.protobuf.Empty](#google.protobuf.Empty) | Delete a parental relationship with a feature |
| DeleteChildRelationship | [FeatureConnection](#dictybase.feature.FeatureConnection) | [.google.protobuf.Empty](#google.protobuf.Empty) | Delete a child relationship with a feature |
| DetachLocation | [ReferenceFeatureFilter](#dictybase.feature.ReferenceFeatureFilter) | [.google.protobuf.Empty](#google.protobuf.Empty) | Detach an existing location from reference feature |
| DeleteSynonym | [FeatureSynonym](#dictybase.feature.FeatureSynonym) | [.google.protobuf.Empty](#google.protobuf.Empty) | Delete a synonym from the feature. |
| DeletePublication | [FeaturePublication](#dictybase.feature.FeaturePublication) | [.google.protobuf.Empty](#google.protobuf.Empty) | Delete a publication from the feature. |
| RemovePrevFeatureHistory | [FeatureHistory](#dictybase.feature.FeatureHistory) | [.google.protobuf.Empty](#google.protobuf.Empty) | Remove the previous feature from the feature history. |
| RemoveNextFeatureHistory | [FeatureHistory](#dictybase.feature.FeatureHistory) | [.google.protobuf.Empty](#google.protobuf.Empty) | Remove the next feature from the feature history. |
| DeleteDbxref | [FeatureDbxref](#dictybase.feature.FeatureDbxref) | [.google.protobuf.Empty](#google.protobuf.Empty) | Delete dbxref from a feature. |

 



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

