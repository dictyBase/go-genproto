# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [order.proto](#order.proto)
    - [ListParameters](#dictybase.order.ListParameters)
    - [Meta](#dictybase.order.Meta)
    - [NewOrder](#dictybase.order.NewOrder)
    - [NewOrder.Data](#dictybase.order.NewOrder.Data)
    - [NewOrderAttributes](#dictybase.order.NewOrderAttributes)
    - [Order](#dictybase.order.Order)
    - [Order.Data](#dictybase.order.Order.Data)
    - [OrderAttributes](#dictybase.order.OrderAttributes)
    - [OrderCollection](#dictybase.order.OrderCollection)
    - [OrderCollection.Data](#dictybase.order.OrderCollection.Data)
    - [OrderId](#dictybase.order.OrderId)
    - [OrderUpdate](#dictybase.order.OrderUpdate)
    - [OrderUpdate.Data](#dictybase.order.OrderUpdate.Data)
    - [OrderUpdateAttributes](#dictybase.order.OrderUpdateAttributes)
  
    - [OrderStatus](#dictybase.order.OrderStatus)
  
  
    - [OrderService](#dictybase.order.OrderService)
  

- [Scalar Value Types](#scalar-value-types)



<a name="order.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## order.proto



<a name="dictybase.order.ListParameters"></a>

### ListParameters
ListParameters defines fields for manipulating output of Order collection


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cursor | [int64](#int64) |  | A unique pointer to the next set of result in the list |
| limit | [int64](#int64) |  | Maximum number of records that can be fetch per request |
| filter | [string](#string) |  | The `filter` field restricts the data return by the collection. To use it, supply one or multiple allowed fields to filter followed by a filter expression. It uses the following syntax... field_name operator expression

The following fields of `OrderAttributes` definition are allowed to be used for filtering * item - Items that are part of order (string). * courier - The courier used for delivery (string) * payment - Type of payment being used (string) * status - The status of the order (string) * created_at - Date the items are ordered (number), can be in the following formats: YYYY-MM-DD, YYYY-MM, YYYY

field_name - Any one of the allowed field_name of the `OrderAttributes` definition. operator - Defines the type of filter match to use. It could be any of the following four and all of them should be URL-encoded for http request.

 Operators for strings =~ Contains substring !~ Not contains substring === Equals !== Not equals

 Operators for number == Equals &gt; Greater than &lt; Less than &lt;= Less than equal to &gt;= Greater than equal to

 Operators for dates $== Equals $&gt; Greater than $&lt; Less than $&lt;= Less than equal to $&gt;= Greater than equal to

expression - The value that will be included or excluded from the result. URL-reserved characters must be URL-encoded for http request.

 filter: &#34;status===Shipped&#34; filter: &#34;courier===FedEx&#34;

Filter can be combined using OR or AND boolean logic. * The OR is represented using a comma(,). * The AND is represented using a semi-colon(;). * AND and OR operators can be combined and AND takes precedence over OR.

 filter: &#34;courier===FedEx;payment===Credit&#34; filter: &#34;created_at$&gt;=20181201&#34;

The sort field allow to sort the data return by the collection based on fields of `OrderAttributes. To use it, supply a comma separated one or more allowed field from the definition of `OrderAttributes`. |






<a name="dictybase.order.Meta"></a>

### Meta
Metadata definition for traversing the collection


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| next_cursor | [int64](#int64) |  | A unique pointer to the next set of result in the collection. Set the cursor value parameter to the value of next_cursor to retrieve the next set of collection using the same method |
| limit | [int64](#int64) |  | Maximum number of records that can be fetch per request |
| total | [int64](#int64) |  | Total number of records in the collection. |






<a name="dictybase.order.NewOrder"></a>

### NewOrder
Definition for creating a new order


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [NewOrder.Data](#dictybase.order.NewOrder.Data) |  |  |






<a name="dictybase.order.NewOrder.Data"></a>

### NewOrder.Data



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  | Resource name, by default should be order |
| attributes | [NewOrderAttributes](#dictybase.order.NewOrderAttributes) |  |  |






<a name="dictybase.order.NewOrderAttributes"></a>

### NewOrderAttributes
Defines attributes for creating a new order


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| courier | [string](#string) |  | Name of courier for delivery |
| courier_account | [string](#string) |  | Account identification used for courier service |
| comments | [string](#string) |  | Any comments about the order |
| payment | [string](#string) |  | Type of payment being used |
| purchase_order_num | [string](#string) |  | Order number for purchase |
| status | [OrderStatus](#dictybase.order.OrderStatus) |  | Status of order |
| consumer | [string](#string) |  | Person (user) who is receiving the stocks in mail |
| payer | [string](#string) |  | Person (user) who is paying for the stocks |
| purchaser | [string](#string) |  | Person who is ordering (logged in user) |
| items | [string](#string) | repeated | List of items in the order, in this case it will be mostly biological stocks such as plasmids and strains |






<a name="dictybase.order.Order"></a>

### Order
Definition of an individual order


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [Order.Data](#dictybase.order.Order.Data) |  |  |






<a name="dictybase.order.Order.Data"></a>

### Order.Data



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  | Resource name, by default should be order |
| id | [string](#string) |  | Unique identifier for the order |
| attributes | [OrderAttributes](#dictybase.order.OrderAttributes) |  |  |






<a name="dictybase.order.OrderAttributes"></a>

### OrderAttributes
Definition of various order attributes


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| created_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | Timestamp for creation |
| updated_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | Timestamp for update |
| courier | [string](#string) |  | Name of courier for delivery |
| courier_account | [string](#string) |  | Account identification used for courier service |
| comments | [string](#string) |  | Any comments about the order |
| payment | [string](#string) |  | Type of payment being used |
| purchase_order_num | [string](#string) |  | Order number for purchase |
| status | [OrderStatus](#dictybase.order.OrderStatus) |  | Status of order |
| consumer | [string](#string) |  | Person (user) who is receiving the stocks in mail |
| payer | [string](#string) |  | Person (user) who is paying for the stocks |
| purchaser | [string](#string) |  | Person who is ordering (logged in user) |
| items | [string](#string) | repeated | List of items in the order, in this case it will be mostly biological stocks such as plasmids and strains |






<a name="dictybase.order.OrderCollection"></a>

### OrderCollection
List of orders


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [OrderCollection.Data](#dictybase.order.OrderCollection.Data) | repeated |  |
| meta | [Meta](#dictybase.order.Meta) |  |  |






<a name="dictybase.order.OrderCollection.Data"></a>

### OrderCollection.Data



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  | Resource name, by default should be order |
| id | [string](#string) |  | Unique identifier for the order |
| attributes | [OrderAttributes](#dictybase.order.OrderAttributes) |  |  |






<a name="dictybase.order.OrderId"></a>

### OrderId



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | Unique identifier for the order |






<a name="dictybase.order.OrderUpdate"></a>

### OrderUpdate
Definition for updating an existing order


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [OrderUpdate.Data](#dictybase.order.OrderUpdate.Data) |  |  |






<a name="dictybase.order.OrderUpdate.Data"></a>

### OrderUpdate.Data



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  | Resource name, by default should be order |
| id | [string](#string) |  | Unique identifier for the order |
| attributes | [OrderUpdateAttributes](#dictybase.order.OrderUpdateAttributes) |  |  |






<a name="dictybase.order.OrderUpdateAttributes"></a>

### OrderUpdateAttributes
Defines attributes for updating an existing order


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| courier | [string](#string) |  | Name of courier for delivery |
| courier_account | [string](#string) |  | Account identification used for courier service |
| comments | [string](#string) |  | Any comments about the order |
| payment | [string](#string) |  | Type of payment being used |
| purchase_order_num | [string](#string) |  | Order number for purchase |
| status | [OrderStatus](#dictybase.order.OrderStatus) |  | Status of order |
| items | [string](#string) | repeated | List of items in the order, in this case it will be mostly biological stocks such as plasmids and strains |





 


<a name="dictybase.order.OrderStatus"></a>

### OrderStatus
The pre-defined labels for order status

| Name | Number | Description |
| ---- | ------ | ----------- |
| In_preparation | 0 |  |
| Growing | 1 |  |
| Cancelled | 2 |  |
| Shipped | 3 |  |


 

 


<a name="dictybase.order.OrderService"></a>

### OrderService
The order service specification

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetOrder | [OrderId](#dictybase.order.OrderId) | [Order](#dictybase.order.Order) | Retrieves order by ID |
| CreateOrder | [NewOrder](#dictybase.order.NewOrder) | [Order](#dictybase.order.Order) | Create a new order |
| UpdateOrder | [OrderUpdate](#dictybase.order.OrderUpdate) | [Order](#dictybase.order.Order) | Update an existing order |
| ListOrders | [ListParameters](#dictybase.order.ListParameters) | [OrderCollection](#dictybase.order.OrderCollection) | List all orders |

 



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

