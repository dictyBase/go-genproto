# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [user.proto](#user.proto)
    - [CreatePermissionRequest](#dictybase.user.CreatePermissionRequest)
    - [CreatePermissionRequest.Data](#dictybase.user.CreatePermissionRequest.Data)
    - [CreateRoleRequest](#dictybase.user.CreateRoleRequest)
    - [CreateRoleRequest.Data](#dictybase.user.CreateRoleRequest.Data)
    - [CreateUserRequest](#dictybase.user.CreateUserRequest)
    - [CreateUserRequest.Data](#dictybase.user.CreateUserRequest.Data)
    - [ExistingRoleRelationships](#dictybase.user.ExistingRoleRelationships)
    - [ExistingRoleRelationships.Permissions](#dictybase.user.ExistingRoleRelationships.Permissions)
    - [ExistingRoleRelationships.Users](#dictybase.user.ExistingRoleRelationships.Users)
    - [ExistingUserRelationships](#dictybase.user.ExistingUserRelationships)
    - [ExistingUserRelationships.Roles](#dictybase.user.ExistingUserRelationships.Roles)
    - [NewRoleRelationships](#dictybase.user.NewRoleRelationships)
    - [NewRoleRelationships.Permissions](#dictybase.user.NewRoleRelationships.Permissions)
    - [NewRoleRelationships.Users](#dictybase.user.NewRoleRelationships.Users)
    - [NewUserRelationships](#dictybase.user.NewUserRelationships)
    - [NewUserRelationships.Roles](#dictybase.user.NewUserRelationships.Roles)
    - [Permission](#dictybase.user.Permission)
    - [PermissionAttributes](#dictybase.user.PermissionAttributes)
    - [PermissionCollection](#dictybase.user.PermissionCollection)
    - [PermissionData](#dictybase.user.PermissionData)
    - [Role](#dictybase.user.Role)
    - [RoleAttributes](#dictybase.user.RoleAttributes)
    - [RoleCollection](#dictybase.user.RoleCollection)
    - [RoleData](#dictybase.user.RoleData)
    - [UpdatePermissionRequest](#dictybase.user.UpdatePermissionRequest)
    - [UpdatePermissionRequest.Data](#dictybase.user.UpdatePermissionRequest.Data)
    - [UpdateRoleRequest](#dictybase.user.UpdateRoleRequest)
    - [UpdateRoleRequest.Data](#dictybase.user.UpdateRoleRequest.Data)
    - [UpdateUserRequest](#dictybase.user.UpdateUserRequest)
    - [UpdateUserRequest.Data](#dictybase.user.UpdateUserRequest.Data)
    - [User](#dictybase.user.User)
    - [UserAttributes](#dictybase.user.UserAttributes)
    - [UserCollection](#dictybase.user.UserCollection)
    - [UserData](#dictybase.user.UserData)
  
  
  
    - [PermissionService](#dictybase.user.PermissionService)
    - [RoleService](#dictybase.user.RoleService)
    - [UserService](#dictybase.user.UserService)
  

- [Scalar Value Types](#scalar-value-types)



<a name="user.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## user.proto



<a name="dictybase.user.CreatePermissionRequest"></a>

### CreatePermissionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [CreatePermissionRequest.Data](#dictybase.user.CreatePermissionRequest.Data) |  |  |






<a name="dictybase.user.CreatePermissionRequest.Data"></a>

### CreatePermissionRequest.Data
The payload for new user


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  |  |
| attributes | [PermissionAttributes](#dictybase.user.PermissionAttributes) |  |  |






<a name="dictybase.user.CreateRoleRequest"></a>

### CreateRoleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [CreateRoleRequest.Data](#dictybase.user.CreateRoleRequest.Data) |  |  |






<a name="dictybase.user.CreateRoleRequest.Data"></a>

### CreateRoleRequest.Data
The payload for new role


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  |  |
| attributes | [RoleAttributes](#dictybase.user.RoleAttributes) |  |  |
| relationships | [NewRoleRelationships](#dictybase.user.NewRoleRelationships) |  |  |






<a name="dictybase.user.CreateUserRequest"></a>

### CreateUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [CreateUserRequest.Data](#dictybase.user.CreateUserRequest.Data) |  |  |






<a name="dictybase.user.CreateUserRequest.Data"></a>

### CreateUserRequest.Data
The payload for new user


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  |  |
| attributes | [UserAttributes](#dictybase.user.UserAttributes) |  |  |
| relationships | [NewUserRelationships](#dictybase.user.NewUserRelationships) |  |  |






<a name="dictybase.user.ExistingRoleRelationships"></a>

### ExistingRoleRelationships
The relationship definition for existing roles.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| permissions | [ExistingRoleRelationships.Permissions](#dictybase.user.ExistingRoleRelationships.Permissions) |  |  |
| users | [ExistingRoleRelationships.Users](#dictybase.user.ExistingRoleRelationships.Users) |  |  |






<a name="dictybase.user.ExistingRoleRelationships.Permissions"></a>

### ExistingRoleRelationships.Permissions
Relationships with permission definition.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| links | [dictybase.api.jsonapi.Links](#dictybase.api.jsonapi.Links) |  |  |
| data | [dictybase.api.jsonapi.Data](#dictybase.api.jsonapi.Data) | repeated |  |






<a name="dictybase.user.ExistingRoleRelationships.Users"></a>

### ExistingRoleRelationships.Users
Relationships with user definition.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| links | [dictybase.api.jsonapi.Links](#dictybase.api.jsonapi.Links) |  |  |
| data | [dictybase.api.jsonapi.Data](#dictybase.api.jsonapi.Data) | repeated |  |






<a name="dictybase.user.ExistingUserRelationships"></a>

### ExistingUserRelationships
The relationship definition for existing users.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| roles | [ExistingUserRelationships.Roles](#dictybase.user.ExistingUserRelationships.Roles) |  |  |






<a name="dictybase.user.ExistingUserRelationships.Roles"></a>

### ExistingUserRelationships.Roles
Relationships with role resource.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| links | [dictybase.api.jsonapi.Links](#dictybase.api.jsonapi.Links) |  | Http links with role resource. |
| data | [dictybase.api.jsonapi.Data](#dictybase.api.jsonapi.Data) | repeated | A role [resource identifier object](http://jsonapi.org/format/#document-resource-identifier-objects). |






<a name="dictybase.user.NewRoleRelationships"></a>

### NewRoleRelationships
The relationship definition for creating new roles.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| permissions | [NewRoleRelationships.Permissions](#dictybase.user.NewRoleRelationships.Permissions) |  |  |
| users | [NewRoleRelationships.Users](#dictybase.user.NewRoleRelationships.Users) |  |  |






<a name="dictybase.user.NewRoleRelationships.Permissions"></a>

### NewRoleRelationships.Permissions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [dictybase.api.jsonapi.Data](#dictybase.api.jsonapi.Data) | repeated |  |






<a name="dictybase.user.NewRoleRelationships.Users"></a>

### NewRoleRelationships.Users



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [dictybase.api.jsonapi.Data](#dictybase.api.jsonapi.Data) | repeated |  |






<a name="dictybase.user.NewUserRelationships"></a>

### NewUserRelationships
The relationship definition for creating new users.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| roles | [NewUserRelationships.Roles](#dictybase.user.NewUserRelationships.Roles) |  |  |






<a name="dictybase.user.NewUserRelationships.Roles"></a>

### NewUserRelationships.Roles
Relationships with role resource.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [dictybase.api.jsonapi.Data](#dictybase.api.jsonapi.Data) | repeated | A role [resource identifier object](http://jsonapi.org/format/#document-resource-identifier-objects). |






<a name="dictybase.user.Permission"></a>

### Permission
A resource for managing user permission.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [PermissionData](#dictybase.user.PermissionData) |  |  |
| links | [dictybase.api.jsonapi.Links](#dictybase.api.jsonapi.Links) |  |  |






<a name="dictybase.user.PermissionAttributes"></a>

### PermissionAttributes
A container for permission fields.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| permission | [string](#string) |  | Kind of permission, for example read, write, admin etc. |
| description | [string](#string) |  | Brief description of the type of permission. |
| created_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | Timestamp for creation and update |
| updated_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| resource | [string](#string) |  | Resource(object) on which this permission is granted |






<a name="dictybase.user.PermissionCollection"></a>

### PermissionCollection
A permission collection resource.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [PermissionData](#dictybase.user.PermissionData) | repeated |  |
| links | [dictybase.api.jsonapi.Links](#dictybase.api.jsonapi.Links) |  |  |






<a name="dictybase.user.PermissionData"></a>

### PermissionData
A top level container for permission data.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  | The resource name. |
| id | [int64](#int64) |  | Unique id. |
| attributes | [PermissionAttributes](#dictybase.user.PermissionAttributes) |  |  |
| links | [dictybase.api.jsonapi.Links](#dictybase.api.jsonapi.Links) |  |  |






<a name="dictybase.user.Role"></a>

### Role
A definition for managing user roles.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [RoleData](#dictybase.user.RoleData) |  |  |
| links | [dictybase.api.jsonapi.Links](#dictybase.api.jsonapi.Links) |  |  |
| included | [google.protobuf.Any](#google.protobuf.Any) | repeated |  |






<a name="dictybase.user.RoleAttributes"></a>

### RoleAttributes
A container for role fields.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| role | [string](#string) |  |  |
| description | [string](#string) |  |  |
| created_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | Timestamp for creation and update |
| updated_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="dictybase.user.RoleCollection"></a>

### RoleCollection
A role collection definition.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [RoleData](#dictybase.user.RoleData) | repeated |  |
| links | [dictybase.api.jsonapi.Links](#dictybase.api.jsonapi.Links) |  |  |
| included | [google.protobuf.Any](#google.protobuf.Any) | repeated |  |






<a name="dictybase.user.RoleData"></a>

### RoleData
A top level container for role data.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  |  |
| id | [int64](#int64) |  |  |
| attributes | [RoleAttributes](#dictybase.user.RoleAttributes) |  |  |
| links | [dictybase.api.jsonapi.Links](#dictybase.api.jsonapi.Links) |  |  |
| relationships | [ExistingRoleRelationships](#dictybase.user.ExistingRoleRelationships) |  |  |






<a name="dictybase.user.UpdatePermissionRequest"></a>

### UpdatePermissionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [UpdatePermissionRequest.Data](#dictybase.user.UpdatePermissionRequest.Data) |  |  |
| id | [int64](#int64) |  | Unique id, required |
| update_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  | An optional mask specifying which fields to update. Presence of this field allow partial updates. |






<a name="dictybase.user.UpdatePermissionRequest.Data"></a>

### UpdatePermissionRequest.Data



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  |  |
| id | [int64](#int64) |  |  |
| attributes | [PermissionAttributes](#dictybase.user.PermissionAttributes) |  |  |






<a name="dictybase.user.UpdateRoleRequest"></a>

### UpdateRoleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [UpdateRoleRequest.Data](#dictybase.user.UpdateRoleRequest.Data) |  |  |
| id | [int64](#int64) |  | Unique id, required |
| update_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  | An optional mask specifying which fields to update. Presence of this field allow partial updates. |






<a name="dictybase.user.UpdateRoleRequest.Data"></a>

### UpdateRoleRequest.Data



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  |  |
| id | [int64](#int64) |  |  |
| attributes | [RoleAttributes](#dictybase.user.RoleAttributes) |  |  |
| relationships | [ExistingRoleRelationships](#dictybase.user.ExistingRoleRelationships) |  |  |






<a name="dictybase.user.UpdateUserRequest"></a>

### UpdateUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [UpdateUserRequest.Data](#dictybase.user.UpdateUserRequest.Data) |  |  |
| id | [int64](#int64) |  | Unique id, required |
| update_mask | [google.protobuf.FieldMask](#google.protobuf.FieldMask) |  | An optional mask specifying which fields to update. Presence of this field allow partial updates. |






<a name="dictybase.user.UpdateUserRequest.Data"></a>

### UpdateUserRequest.Data



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  |  |
| id | [int64](#int64) |  |  |
| attributes | [UserAttributes](#dictybase.user.UserAttributes) |  |  |
| relationships | [ExistingUserRelationships](#dictybase.user.ExistingUserRelationships) |  |  |






<a name="dictybase.user.User"></a>

### User
A user resource.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [UserData](#dictybase.user.UserData) |  |  |
| links | [dictybase.api.jsonapi.Links](#dictybase.api.jsonapi.Links) |  |  |
| included | [google.protobuf.Any](#google.protobuf.Any) | repeated |  |






<a name="dictybase.user.UserAttributes"></a>

### UserAttributes
A container for user fields.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| first_name | [string](#string) |  | First name. |
| last_name | [string](#string) |  | Last name. |
| email | [string](#string) |  | Email. |
| organization | [string](#string) |  | Organization in which the user belong. |
| group_name | [string](#string) |  | Group in which the user belong. |
| first_address | [string](#string) |  | Address. |
| second_address | [string](#string) |  | More address. |
| city | [string](#string) |  | City. |
| state | [string](#string) |  | State. |
| zipcode | [string](#string) |  | Zipcode. |
| country | [string](#string) |  | Country. |
| phone | [string](#string) |  | Phone no. |
| is_active | [bool](#bool) |  | Current status of user. |
| created_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | Timestamp for creation and update |
| updated_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="dictybase.user.UserCollection"></a>

### UserCollection
A user collection resource.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [UserData](#dictybase.user.UserData) | repeated |  |
| links | [dictybase.api.jsonapi.PaginationLinks](#dictybase.api.jsonapi.PaginationLinks) |  |  |
| meta | [dictybase.api.jsonapi.Meta](#dictybase.api.jsonapi.Meta) |  |  |
| included | [google.protobuf.Any](#google.protobuf.Any) | repeated |  |






<a name="dictybase.user.UserData"></a>

### UserData
A top level container for user data.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  | The resource name. |
| id | [int64](#int64) |  | Unique id. |
| attributes | [UserAttributes](#dictybase.user.UserAttributes) |  |  |
| links | [dictybase.api.jsonapi.Links](#dictybase.api.jsonapi.Links) |  |  |
| relationships | [ExistingUserRelationships](#dictybase.user.ExistingUserRelationships) |  |  |





 

 

 


<a name="dictybase.user.PermissionService"></a>

### PermissionService
The permission service definition

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetPermission | [.dictybase.api.jsonapi.GetRequestWithFields](#dictybase.api.jsonapi.GetRequestWithFields) | [Permission](#dictybase.user.Permission) | Gets the specified permission |
| ListPermissions | [.dictybase.api.jsonapi.SimpleListRequest](#dictybase.api.jsonapi.SimpleListRequest) | [PermissionCollection](#dictybase.user.PermissionCollection) | List all permissions |
| CreatePermission | [CreatePermissionRequest](#dictybase.user.CreatePermissionRequest) | [Permission](#dictybase.user.Permission) | Create an permission |
| UpdatePermission | [UpdatePermissionRequest](#dictybase.user.UpdatePermissionRequest) | [Permission](#dictybase.user.Permission) | Update an permission |
| DeletePermission | [.dictybase.api.jsonapi.DeleteRequest](#dictybase.api.jsonapi.DeleteRequest) | [.google.protobuf.Empty](#google.protobuf.Empty) | Delete an permission |


<a name="dictybase.user.RoleService"></a>

### RoleService
The role service definition

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetRole | [.dictybase.api.jsonapi.GetRequest](#dictybase.api.jsonapi.GetRequest) | [Role](#dictybase.user.Role) | Gets the specified role |
| GetRelatedUsers | [.dictybase.api.jsonapi.RelationshipRequestWithPagination](#dictybase.api.jsonapi.RelationshipRequestWithPagination) | [UserCollection](#dictybase.user.UserCollection) | Gets all related users |
| GetRelatedPermissions | [.dictybase.api.jsonapi.RelationshipRequest](#dictybase.api.jsonapi.RelationshipRequest) | [PermissionCollection](#dictybase.user.PermissionCollection) | Gets all related permissions |
| ListRoles | [.dictybase.api.jsonapi.SimpleListRequest](#dictybase.api.jsonapi.SimpleListRequest) | [RoleCollection](#dictybase.user.RoleCollection) | List all roles. Both *users* and *permissions* relationships are allowed in the include parameter. |
| CreateRole | [CreateRoleRequest](#dictybase.user.CreateRoleRequest) | [Role](#dictybase.user.Role) | Create an role |
| CreateUserRelationship | [.dictybase.api.jsonapi.DataCollection](#dictybase.api.jsonapi.DataCollection) | [.google.protobuf.Empty](#google.protobuf.Empty) | Create user relationship with role |
| CreatePermissionRelationship | [.dictybase.api.jsonapi.DataCollection](#dictybase.api.jsonapi.DataCollection) | [.google.protobuf.Empty](#google.protobuf.Empty) | Create permission relationship with role |
| UpdateRole | [UpdateRoleRequest](#dictybase.user.UpdateRoleRequest) | [Role](#dictybase.user.Role) | Update an role |
| UpdateUserRelationship | [.dictybase.api.jsonapi.DataCollection](#dictybase.api.jsonapi.DataCollection) | [.google.protobuf.Empty](#google.protobuf.Empty) | Update existing user relationship with role |
| UpdatePermissionRelationship | [.dictybase.api.jsonapi.DataCollection](#dictybase.api.jsonapi.DataCollection) | [.google.protobuf.Empty](#google.protobuf.Empty) | Update existing permission relationship with role |
| DeleteRole | [.dictybase.api.jsonapi.DeleteRequest](#dictybase.api.jsonapi.DeleteRequest) | [.google.protobuf.Empty](#google.protobuf.Empty) | Delete an role |
| DeleteUserRelationship | [.dictybase.api.jsonapi.DataCollection](#dictybase.api.jsonapi.DataCollection) | [.google.protobuf.Empty](#google.protobuf.Empty) | Delete existing user relationship with role |
| DeletePermissionRelationship | [.dictybase.api.jsonapi.DataCollection](#dictybase.api.jsonapi.DataCollection) | [.google.protobuf.Empty](#google.protobuf.Empty) | Delete existing permission relationship with role |


<a name="dictybase.user.UserService"></a>

### UserService
The user service definition

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| ExistUser | [.dictybase.api.jsonapi.IdRequest](#dictybase.api.jsonapi.IdRequest) | [.dictybase.api.jsonapi.ExistResponse](#dictybase.api.jsonapi.ExistResponse) | Check the existence of user |
| GetUser | [.dictybase.api.jsonapi.GetRequest](#dictybase.api.jsonapi.GetRequest) | [User](#dictybase.user.User) | Gets the specified user |
| GetUserByEmail | [.dictybase.api.jsonapi.GetEmailRequest](#dictybase.api.jsonapi.GetEmailRequest) | [User](#dictybase.user.User) | Gets the specified user by their email id |
| GetRelatedRoles | [.dictybase.api.jsonapi.RelationshipRequest](#dictybase.api.jsonapi.RelationshipRequest) | [RoleCollection](#dictybase.user.RoleCollection) | Gets all related roles |
| ListUsers | [.dictybase.api.jsonapi.ListRequest](#dictybase.api.jsonapi.ListRequest) | [UserCollection](#dictybase.user.UserCollection) | List all users. Only *roles* relationship is allowed for inclusion. |
| CreateUser | [CreateUserRequest](#dictybase.user.CreateUserRequest) | [User](#dictybase.user.User) | Create an user |
| CreateRoleRelationship | [.dictybase.api.jsonapi.DataCollection](#dictybase.api.jsonapi.DataCollection) | [.google.protobuf.Empty](#google.protobuf.Empty) | Create relationship links with roles |
| UpdateUser | [UpdateUserRequest](#dictybase.user.UpdateUserRequest) | [User](#dictybase.user.User) | Update an user |
| UpdateRoleRelationship | [.dictybase.api.jsonapi.DataCollection](#dictybase.api.jsonapi.DataCollection) | [.google.protobuf.Empty](#google.protobuf.Empty) | Update relationship links with roles |
| DeleteUser | [.dictybase.api.jsonapi.DeleteRequest](#dictybase.api.jsonapi.DeleteRequest) | [.google.protobuf.Empty](#google.protobuf.Empty) | Delete an user |
| DeleteRoleRelationship | [.dictybase.api.jsonapi.DataCollection](#dictybase.api.jsonapi.DataCollection) | [.google.protobuf.Empty](#google.protobuf.Empty) | Delete relationship links with roles |
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

