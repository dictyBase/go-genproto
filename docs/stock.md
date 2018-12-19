# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [stock.proto](#stock.proto)
    - [ListParameters](#dictybase.stock.ListParameters)
    - [Meta](#dictybase.stock.Meta)
    - [NewStock](#dictybase.stock.NewStock)
    - [NewStock.Data](#dictybase.stock.NewStock.Data)
    - [NewStockAttributes](#dictybase.stock.NewStockAttributes)
    - [PlasmidListParameters](#dictybase.stock.PlasmidListParameters)
    - [PlasmidProperties](#dictybase.stock.PlasmidProperties)
    - [Stock](#dictybase.stock.Stock)
    - [Stock.Data](#dictybase.stock.Stock.Data)
    - [StockAttributes](#dictybase.stock.StockAttributes)
    - [StockCollection](#dictybase.stock.StockCollection)
    - [StockCollection.Data](#dictybase.stock.StockCollection.Data)
    - [StockId](#dictybase.stock.StockId)
    - [StockUpdate](#dictybase.stock.StockUpdate)
    - [StockUpdate.Data](#dictybase.stock.StockUpdate.Data)
    - [StockUpdateAttributes](#dictybase.stock.StockUpdateAttributes)
    - [StrainListParameters](#dictybase.stock.StrainListParameters)
    - [StrainProperties](#dictybase.stock.StrainProperties)
    - [StrainUpdateProperties](#dictybase.stock.StrainUpdateProperties)
  
  
  
    - [StockService](#dictybase.stock.StockService)
  

- [Scalar Value Types](#scalar-value-types)



<a name="stock.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## stock.proto



<a name="dictybase.stock.ListParameters"></a>

### ListParameters
ListParameters defines fields for manipulating output of Stock collection


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cursor | [int64](#int64) |  | A unique pointer to the next set of result in the list |
| limit | [int64](#int64) |  | Maximum number of records that can be fetch per request |
| filter | [string](#string) |  | The `filter` field restricts the data return by the collection. To use it, supply one or multiple allowed fields to filter followed by a filter expression. It uses the following syntax... field_name operator expression

The following fields of `StockAttributes` definition are allowed to be used for filtering * depositor - Depositor of the stock (string). * summary - Summary of the stock (string) * created_at - Date the stock was created (number), can be in the following formats: YYYY-MM-DD, YYYY-MM, YYYY * updated_at - Date the stock was updated (number), can be in the following formats: YYYY-MM-DD, YYYY-MM, YYYY

field_name - Any one of the allowed field_name of the `StockAttributes` definition. operator - Defines the type of filter match to use. It could be any of the following four and all of them should be URL-encoded for http request.

 Operators for strings =~ Contains substring !~ Not contains substring === Equals !== Not equals

 Operators for number == Equals &gt; Greater than &lt; Less than &lt;= Less than equal to &gt;= Greater than equal to

 Operators for dates $== Equals $&gt; Greater than $&lt; Less than $&lt;= Less than equal to $&gt;= Greater than equal to

expression - The value that will be included or excluded from the result. URL-reserved characters must be URL-encoded for http request.

 filter: &#34;created_at$&gt;=2018-12-01&#34; filter: &#34;depositor===Costanza&#34;

Filter can be combined using OR or AND boolean logic. * The OR is represented using a comma(,). * The AND is represented using a semi-colon(;). * AND and OR operators can be combined and AND takes precedence over OR.

 filter: &#34;depositor===Benes;created_at$&gt;=2018-12-01&#34;

The sort field allow to sort the data return by the collection based on fields of `StockAttributes. To use it, supply a comma separated one or more allowed field from the definition of `StockAttributes`. |






<a name="dictybase.stock.Meta"></a>

### Meta
Metadata definition for traversing the collection


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| next_cursor | [int64](#int64) |  | A unique pointer to the next set of result in the collection. Set the cursor value parameter to the value of next_cursor to retrieve the next set of collection using the same method |
| limit | [int64](#int64) |  | Maximum number of records that can be fetch per request |
| total | [int64](#int64) |  | Total number of records in the collection. |






<a name="dictybase.stock.NewStock"></a>

### NewStock
Definition for creating a new stock


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [NewStock.Data](#dictybase.stock.NewStock.Data) |  |  |






<a name="dictybase.stock.NewStock.Data"></a>

### NewStock.Data



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  | Resource name, i.e. &#34;strain&#34; |
| id | [string](#string) |  | Unique ID for stock |
| attributes | [NewStockAttributes](#dictybase.stock.NewStockAttributes) |  |  |






<a name="dictybase.stock.NewStockAttributes"></a>

### NewStockAttributes
Defines attributes for creating a new stock


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| created_by | [string](#string) |  | User who created stock entry |
| updated_by | [string](#string) |  | User who updated stock entry |
| summary | [string](#string) |  | Summary of the stock |
| editable_summary | [string](#string) |  | Editable version of the stock summary (Slate JSON format) |
| depositor | [string](#string) |  | Depositor of the stock |
| genes | [string](#string) | repeated | List of associated genes |
| dbxrefs | [string](#string) | repeated | List of database cross references |
| publications | [string](#string) | repeated | List of related publications |
| strain_properties | [StrainProperties](#dictybase.stock.StrainProperties) |  |  |
| plasmid_properties | [PlasmidProperties](#dictybase.stock.PlasmidProperties) |  |  |






<a name="dictybase.stock.PlasmidListParameters"></a>

### PlasmidListParameters
PlasmidListParameters defines fields for manipulating output of Stock plasmid collection


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cursor | [int64](#int64) |  | A unique pointer to the next set of result in the list |
| limit | [int64](#int64) |  | Maximum number of records that can be fetch per request |
| filter | [string](#string) |  | The `filter` field restricts the data return by the collection. To use it, supply one or multiple allowed fields to filter followed by a filter expression. It uses the following syntax... field_name operator expression

The following fields of `PlasmidProperties` definition are allowed to be used for filtering * depositor - Depositor of the plasmid (string). * summary - Summary of the plasmid (string) * keyword - Keywords used for the plasmid (string), searches in the &#34;keywords&#34; attribute * created_at - Date the stock was created (number), can be in the following formats: YYYY-MM-DD, YYYY-MM, YYYY * updated_at - Date the stock was updated (number), can be in the following formats: YYYY-MM-DD, YYYY-MM, YYYY

field_name - Any one of the allowed field_name of the `PlasmidProperties` definition. operator - Defines the type of filter match to use. It could be any of the following four and all of them should be URL-encoded for http request.

 Operators for strings =~ Contains substring !~ Not contains substring === Equals !== Not equals

 Operators for number == Equals &gt; Greater than &lt; Less than &lt;= Less than equal to &gt;= Greater than equal to

 Operators for dates $== Equals $&gt; Greater than $&lt; Less than $&lt;= Less than equal to $&gt;= Greater than equal to

expression - The value that will be included or excluded from the result. URL-reserved characters must be URL-encoded for http request.

 filter: &#34;created_at$&gt;=2018-12-01&#34; filter: &#34;depositor===Costanza&#34;

Filter can be combined using OR or AND boolean logic. * The OR is represented using a comma(,). * The AND is represented using a semi-colon(;). * AND and OR operators can be combined and AND takes precedence over OR.

 filter: &#34;depositor===Benes;created_at$&gt;=2018-12-01&#34;

The sort field allow to sort the data return by the collection based on fields of `PlasmidProperties. To use it, supply a comma separated one or more allowed field from the definition of `PlasmidProperties`. |






<a name="dictybase.stock.PlasmidProperties"></a>

### PlasmidProperties



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| image_map | [string](#string) |  | Image map for the plasmid |
| sequence | [string](#string) |  | Sequence for the plasmid |
| keywords | [string](#string) | repeated | Keywords used for the plasmid |






<a name="dictybase.stock.Stock"></a>

### Stock
Definition of an individual stock


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [Stock.Data](#dictybase.stock.Stock.Data) |  |  |






<a name="dictybase.stock.Stock.Data"></a>

### Stock.Data



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  | Resource name (&#34;strain&#34;, &#34;plasmid&#34;, etc) |
| id | [string](#string) |  | Unique identifier for the stock |
| attributes | [StockAttributes](#dictybase.stock.StockAttributes) |  |  |






<a name="dictybase.stock.StockAttributes"></a>

### StockAttributes
Definition of various stock attributes


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| created_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | Timestamp for creation |
| updated_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | Timestamp for update |
| created_by | [string](#string) |  | User who created stock entry |
| updated_by | [string](#string) |  | User who updated stock entry |
| summary | [string](#string) |  | Summary of the stock |
| editable_summary | [string](#string) |  | Editable version of the stock summary (Slate JSON format) |
| depositor | [string](#string) |  | Depositor of the stock |
| genes | [string](#string) | repeated | List of associated genes |
| dbxrefs | [string](#string) | repeated | List of database cross references |
| publications | [string](#string) | repeated | List of related publications |
| strain_properties | [StrainProperties](#dictybase.stock.StrainProperties) |  |  |
| plasmid_properties | [PlasmidProperties](#dictybase.stock.PlasmidProperties) |  |  |






<a name="dictybase.stock.StockCollection"></a>

### StockCollection
List of stocks


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [StockCollection.Data](#dictybase.stock.StockCollection.Data) | repeated |  |
| meta | [Meta](#dictybase.stock.Meta) |  |  |






<a name="dictybase.stock.StockCollection.Data"></a>

### StockCollection.Data



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  | Resource name |
| id | [string](#string) |  | Unique identifier for the stock |
| attributes | [StockAttributes](#dictybase.stock.StockAttributes) |  |  |






<a name="dictybase.stock.StockId"></a>

### StockId



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | Unique identifier for the stock |






<a name="dictybase.stock.StockUpdate"></a>

### StockUpdate
Definition for creating a new stock


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [StockUpdate.Data](#dictybase.stock.StockUpdate.Data) |  |  |






<a name="dictybase.stock.StockUpdate.Data"></a>

### StockUpdate.Data



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  | Resource name, i.e. &#34;strain&#34; |
| id | [string](#string) |  | Unique ID for stock |
| attributes | [StockUpdateAttributes](#dictybase.stock.StockUpdateAttributes) |  |  |






<a name="dictybase.stock.StockUpdateAttributes"></a>

### StockUpdateAttributes
Defines attributes for updating a stock


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| updated_by | [string](#string) |  | User who updated stock entry |
| summary | [string](#string) |  | Summary of the stock |
| editable_summary | [string](#string) |  | Editable version of the stock summary (Slate JSON format) |
| depositor | [string](#string) |  | Depositor of the stock |
| genes | [string](#string) | repeated | List of associated genes |
| dbxrefs | [string](#string) | repeated | List of database cross references |
| publications | [string](#string) | repeated | List of related publications |
| strain_properties | [StrainUpdateProperties](#dictybase.stock.StrainUpdateProperties) |  |  |
| plasmid_properties | [PlasmidProperties](#dictybase.stock.PlasmidProperties) |  |  |






<a name="dictybase.stock.StrainListParameters"></a>

### StrainListParameters
StrainListParameters defines fields for manipulating output of Stock strain collection


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cursor | [int64](#int64) |  | A unique pointer to the next set of result in the list |
| limit | [int64](#int64) |  | Maximum number of records that can be fetch per request |
| filter | [string](#string) |  | The `filter` field restricts the data return by the collection. To use it, supply one or multiple allowed fields to filter followed by a filter expression. It uses the following syntax... field_name operator expression

The following fields of `StrainProperties` definition are allowed to be used for filtering * depositor - Depositor of the stock (string). * parent - Parental strain (string) * plasmid - Related plasmid for the strain (string) * species - The species of the strain (string) * systematic_name - Unambiguous name for the strain (string) * summary - Summary of the stock (string) * name - Name used for strain (string), searches in the &#34;names&#34; attribute * created_at - Date the stock was created (number), can be in the following formats: YYYY-MM-DD, YYYY-MM, YYYY * updated_at - Date the stock was updated (number), can be in the following formats: YYYY-MM-DD, YYYY-MM, YYYY

field_name - Any one of the allowed field_name of the `StrainProperties` definition. operator - Defines the type of filter match to use. It could be any of the following four and all of them should be URL-encoded for http request.

 Operators for strings =~ Contains substring !~ Not contains substring === Equals !== Not equals

 Operators for number == Equals &gt; Greater than &lt; Less than &lt;= Less than equal to &gt;= Greater than equal to

 Operators for dates $== Equals $&gt; Greater than $&lt; Less than $&lt;= Less than equal to $&gt;= Greater than equal to

expression - The value that will be included or excluded from the result. URL-reserved characters must be URL-encoded for http request.

 filter: &#34;systematic_name===DBS0235459&#34; filter: &#34;plasmid===pDeltaAA,p?AA&#34;

Filter can be combined using OR or AND boolean logic. * The OR is represented using a comma(,). * The AND is represented using a semi-colon(;). * AND and OR operators can be combined and AND takes precedence over OR.

 filter: &#34;systematic_name===DBS0235459;plasmid===pDeltaAA,p?AA&#34;

The sort field allow to sort the data return by the collection based on fields of `StrainProperties. To use it, supply a comma separated one or more allowed field from the definition of `StrainProperties`. |






<a name="dictybase.stock.StrainProperties"></a>

### StrainProperties



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| systematic_name | [string](#string) |  | Unambiguous name for the strain |
| descriptor | [string](#string) |  | Descriptor for the strain, a quick overview of its key genetic modifications |
| species | [string](#string) |  | Species of the strain |
| plasmid | [string](#string) |  | Related plasmid for the strain |
| parents | [string](#string) | repeated | List of parents of the strain |
| names | [string](#string) | repeated | List of names for the strain |






<a name="dictybase.stock.StrainUpdateProperties"></a>

### StrainUpdateProperties



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| systematic_name | [string](#string) |  | Unambiguous name for the strain |
| descriptor | [string](#string) |  | Descriptor for the strain, a quick overview of its key genetic modifications |
| species | [string](#string) |  | Species of the strain |
| plasmid | [string](#string) |  | Related plasmid for the strain |
| parents | [string](#string) | repeated | List of parents of the strain |
| names | [string](#string) | repeated | List of names for the strain |





 

 

 


<a name="dictybase.stock.StockService"></a>

### StockService
The stock service specification

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetStock | [StockId](#dictybase.stock.StockId) | [Stock](#dictybase.stock.Stock) | Retrieves stock by ID |
| CreateStock | [NewStock](#dictybase.stock.NewStock) | [Stock](#dictybase.stock.Stock) | Create a new stock |
| UpdateStock | [StockUpdate](#dictybase.stock.StockUpdate) | [Stock](#dictybase.stock.Stock) | Update an existing stock |
| RemoveStock | [StockId](#dictybase.stock.StockId) | [.google.protobuf.Empty](#google.protobuf.Empty) | Remove an existing stock |
| ListStocks | [ListParameters](#dictybase.stock.ListParameters) | [StockCollection](#dictybase.stock.StockCollection) | List all stocks |
| ListStrains | [StrainListParameters](#dictybase.stock.StrainListParameters) | [StockCollection](#dictybase.stock.StockCollection) | List all strains |
| ListPlasmids | [PlasmidListParameters](#dictybase.stock.PlasmidListParameters) | [StockCollection](#dictybase.stock.StockCollection) | List all plasmids |

 



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

