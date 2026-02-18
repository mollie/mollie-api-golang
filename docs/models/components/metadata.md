# Metadata

Provide any data you like, for example a string or a JSON object. We will save the data alongside the entity. Whenever
you fetch the entity with our API, we will also include the metadata. You can use up to approximately 1kB.


## Supported Types

### 

```go
metadata := components.CreateMetadataStr(string{/* values here */})
```

### 

```go
metadata := components.CreateMetadataMapOfAny(map[string]any{/* values here */})
```

### 

```go
metadata := components.CreateMetadataArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metadata.Type {
	case components.MetadataTypeStr:
		// metadata.Str is populated
	case components.MetadataTypeMapOfAny:
		// metadata.MapOfAny is populated
	case components.MetadataTypeArrayOfStr:
		// metadata.ArrayOfStr is populated
}
```
