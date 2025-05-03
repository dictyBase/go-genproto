# DictyBase Feature Annotation and Annotation API

This documentation covers the Go protobuf generated code for the `feature_annotation` and `annotation` packages of the DictyBase API. These packages provide services and data structures for managing annotations of biological features and general annotation capabilities.

## Overview

The codebase represents Go gRPC protobuf generated code for feature annotation and annotation services used in dictyBase, a model organism database. It includes data models and service definitions for annotating biological features and managing annotations with tags derived from ontologies.

## Modules

### feature_annotation

Provides data structures and services for feature annotations in a biological context. Includes support for creating, retrieving, updating, and deleting feature annotations, managing tags, and filtering annotations by various criteria.

#### Key Components:

- **FeatureAnnotationService**: Service for managing feature annotations including CRUD operations and specialized queries.
- **FeatureAnnotation**: Core data structure representing an annotation with its metadata.
- **OrganismFeatureService**: Service for managing relations between features and organisms.

### annotation

Provides general annotation capabilities including tagging, grouping, and querying annotations. Includes support for managing annotation collections and working with ontology terms.

#### Key Components:

- **TaggedAnnotationService**: Service for managing tagged annotations with support for grouping and ontology integration.
- **TaggedAnnotation**: Core data structure representing an annotation with ontology-based tags.
- **TaggedAnnotationGroup**: Grouping mechanism for related annotations.

## Directory Structure

```
dictybaseapis/
  ├── annotation/
  │   ├── annotation.pb.go
  │   ├── annotation.validator.pb.go
  │   └── annotation_grpc.pb.go
  └── feature_annotation/
      ├── feature_annotation.pb.go
      ├── feature_annotation.validator.pb.go
      ├── feature_annotation_grpc.pb.go
      ├── feature_annotation_organism.pb.go
      ├── feature_annotation_organism.validator.pb.go
      └── feature_annotation_organism_grpc.pb.go
```

## Services

### FeatureAnnotationService

Service for managing feature annotations in a biological context. Provides CRUD operations and specialized query methods.

#### Methods:

| Method | Description | Input | Output |
|--------|-------------|-------|--------|
| CreateFeatureAnnotation | Create a feature annotation | NewFeatureAnnotation | FeatureAnnotation |
| GetFeatureAnnotation | Retrieves the specified feature annotation | FeatureAnnotationId | FeatureAnnotation |
| GetFeatureAnnotationByName | Retrieves the specified feature annotation by its name | FeatureName | FeatureAnnotation |
| UpdateFeatureAnnotation | Update an existing feature annotation. Any given tag will be appended to the existing tags | FeatureAnnotationUpdate | FeatureAnnotation |
| DeleteFeatureAnnotation | Delete an existing feature annotation | DeleteFeatureAnnotationRequest | Empty |
| AddTag | Add tag to an existing feature annotation | AddTagRequest | FeatureAnnotation |
| UpdateTag | Update an existing tag in a feature annotation | UpdateTagRequest | FeatureAnnotation |
| RemoveTag | Remove a tag from a feature annotation | RemoveTagRequest | FeatureAnnotation |
| ListFeatureAnnotationsByPubmedId | Retrieves a list of feature annotations by PubMed ID | PubmedId | FeatureAnnotationCollection |
| ListFeatureAnnotationsByDOI | Retrieves a list of feature annotations by DOI (Digital Object Identifier) | DOI | FeatureAnnotationCollection |

### OrganismFeatureService

Service for managing relationships between features and organisms.

#### Methods:

| Method | Description | Input | Output |
|--------|-------------|-------|--------|
| LinkFeatureToOrganism | Link a feature annotation to an organism | OrganismFeatureLink | Empty |
| GetFeatureOrganism | Get organism for a feature | FeatureAnnotationId | Organism |
| UpdateFeatureOrganism | Update feature's organism link | OrganismFeatureUpdate | Empty |
| RemoveFeatureOrganism | Remove feature's organism link | FeatureAnnotationId | Empty |

### TaggedAnnotationService

Service for managing tagged annotations with support for grouping and ontology integration.

#### Methods:

| Method | Description | Input | Output |
|--------|-------------|-------|--------|
| GetAnnotation | Retrieves the specified tagged annotation | AnnotationId | TaggedAnnotation |
| GetEntryAnnotation | Retrieves a single tagged annotation associated with a specific entry | EntryAnnotationRequest | TaggedAnnotation |
| ListAnnotations | List tagged annotations using pagination, ten entries are retrieved by default | ListParameters | TaggedAnnotationCollection |
| CreateAnnotation | Create a tagged annotation | NewTaggedAnnotation | TaggedAnnotation |
| UpdateAnnotation | Update an existing annotation, in this case a new annotation entry is created with a link to the previous annotation (copy on write) | TaggedAnnotationUpdate | TaggedAnnotation |
| DeleteAnnotation | Delete an existing annotation | DeleteAnnotationRequest | Empty |
| CreateAnnotationGroup | Creates an annotation group from bunch of existing tagged annotations | AnnotationIdList | TaggedAnnotationGroup |
| GetAnnotationGroup | Retrieves an annotation group | GroupEntryId | TaggedAnnotationGroup |
| AddToAnnotationGroup | Adds an existing annotation into an existing annotation group | AnnotationGroupId | TaggedAnnotationGroup |
| DeleteAnnotationGroup | Remove an annotation group | GroupEntryId | Empty |
| ListAnnotationGroups | List tagged annotation groups using pagination, ten entries are retrieved by default | ListGroupParameters | TaggedAnnotationGroupCollection |
| GetAnnotationTag | Retrieves tag information | TagRequest | AnnotationTag |
| OboJSONFileUpload | Upload obojson formatted file through client side streaming | FileUploadRequest (streaming) | FileUploadResponse |

## Key Data Models

### FeatureAnnotation

Represents a complete feature annotation record with all its metadata.

#### Fields:

| Field | Type | Description |
|-------|------|-------------|
| type | string | |
| id | string | unique identifier for the feature annotation |
| attributes | FeatureAnnotationAttributes | |
| created_by | string | email id of the user who created the content |
| updated_by | string | email id of the user who updated the content |
| created_at | Timestamp | |
| updated_at | Timestamp | |
| is_obsolete | bool | Toggle the obsolete status |
| version | int64 | Deprecated |

### TaggedAnnotation

Definition of a tag value based biological annotation where the tag always represents a term from ontology.

#### Fields:

| Field | Type | Description |
|-------|------|-------------|
| data | TaggedAnnotation.Data | |

### TaggedAnnotationAttributes

Definition of various tagged annotation attributes.

#### Fields:

| Field | Type | Description |
|-------|------|-------------|
| value | string | annotation in plain text format |
| editable_value | string | serialized text content in a format recognized by frontend rich text editor |
| created_by | string | Unique identifier(generally email) of the user who created the annotation |
| created_at | Timestamp | |
| tag | string | An identifiable tagname for the annotation, primarily a structured tag, generally an ontology term |
| version | int64 | version refers to the current version no |
| entry_id | string | unique identifier of a biological entity that is being annotated |
| ontology | string | Name of ontology in which the tag name is taken |
| rank | int64 | Ordering of annotation when an entry has multiple annotations with identical tag from the same ontology |
| is_obsolete | bool | Status for active or retired annotation |

## Filtering Annotations

The API provides robust filtering capabilities for querying annotations.

### ListParameters Filter Field

The `filter` field restricts the data returned by the collection. It accepts one or more filtering conditions using the syntax:

```
field_name operator expression
```

#### Available fields for filtering from `AnnotationAttributes`:

- entry_id - Identifier of the annotated entry (string)
- value - Annotation text content (string)
- created_by - Email of the creator (string)
- tag - Ontology term used as tag (string)
- ontology - Source ontology name (string)
- version - Version number (number)
- rank - Ordering of annotation (number)
- is_obsolete - Status of annotation (boolean)

#### String operators:
- =~ Contains substring
- !~ Does not contain substring
- === Equals exactly
- !== Not equals

#### Numeric operators:
- == Equals
- != Not equals
- > Greater than
- < Less than
- =< Less than or equal to
- >= Greater than or equal to

#### Boolean operators:
- == Equals
- != Not equals

#### Examples:

```
filter: "created_by===caboose@abc.com"
filter: "entry_id===DDB_G4839483"
filter: "value=~actin"
filter: "tag===GO:0005634"
filter: "ontology===cellular_component"
filter: "version>3"
filter: "rank==0"
filter: "is_obsolete==false"
filter: "value!~pseudogene"
filter: "created_by!==anonymous@dictybase.org"
```

#### Combining filters:

Filters can be combined using boolean operators:
- OR: represented by comma (,)
- AND: represented by semicolon (;)
- AND takes precedence over OR

Examples of combined filters:
```
filter: "value=~cytoskeleton;tag===cell membrane;ontology===cellular"
filter: "entry_id===DDB_G0285418;is_obsolete==false"
filter: "created_by===curator@dictybase.org;version>2;rank==0"
filter: "tag===GO:0005634,tag===GO:0005737" (entries with nuclear OR cytoplasmic localization)
filter: "ontology===molecular_function;value=~transport" (transport-related functions)
```