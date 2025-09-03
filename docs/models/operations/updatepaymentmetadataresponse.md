# UpdatePaymentMetadataResponse

Provide any data you like, for example a string or a JSON object. We will save the data alongside the entity. Whenever
you fetch the entity with our API, we will also include the metadata. You can use up to approximately 1kB.


## Supported Types

### 

```go
updatePaymentMetadataResponse := operations.CreateUpdatePaymentMetadataResponseStr(string{/* values here */})
```

### 

```go
updatePaymentMetadataResponse := operations.CreateUpdatePaymentMetadataResponseMapOfAny(map[string]any{/* values here */})
```

### 

```go
updatePaymentMetadataResponse := operations.CreateUpdatePaymentMetadataResponseArrayOfStr([]string{/* values here */})
```

