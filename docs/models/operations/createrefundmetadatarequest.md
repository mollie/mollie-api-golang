# CreateRefundMetadataRequest

Provide any data you like, for example a string or a JSON object. We will save the data alongside the entity. Whenever
you fetch the entity with our API, we will also include the metadata. You can use up to approximately 1kB.


## Supported Types

### 

```go
createRefundMetadataRequest := operations.CreateCreateRefundMetadataRequestStr(string{/* values here */})
```

### 

```go
createRefundMetadataRequest := operations.CreateCreateRefundMetadataRequestMapOfAny(map[string]any{/* values here */})
```

### 

```go
createRefundMetadataRequest := operations.CreateCreateRefundMetadataRequestArrayOfStr([]string{/* values here */})
```

