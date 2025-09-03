# CreateCaptureMetadataResponse

Provide any data you like, for example a string or a JSON object. We will save the data alongside the entity. Whenever
you fetch the entity with our API, we will also include the metadata. You can use up to approximately 1kB.


## Supported Types

### 

```go
createCaptureMetadataResponse := operations.CreateCreateCaptureMetadataResponseStr(string{/* values here */})
```

### 

```go
createCaptureMetadataResponse := operations.CreateCreateCaptureMetadataResponseMapOfAny(map[string]any{/* values here */})
```

### 

```go
createCaptureMetadataResponse := operations.CreateCreateCaptureMetadataResponseArrayOfStr([]string{/* values here */})
```

