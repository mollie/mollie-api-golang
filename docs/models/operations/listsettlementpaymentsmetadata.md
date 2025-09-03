# ListSettlementPaymentsMetadata

Provide any data you like, for example a string or a JSON object. We will save the data alongside the entity. Whenever
you fetch the entity with our API, we will also include the metadata. You can use up to approximately 1kB.


## Supported Types

### 

```go
listSettlementPaymentsMetadata := operations.CreateListSettlementPaymentsMetadataStr(string{/* values here */})
```

### 

```go
listSettlementPaymentsMetadata := operations.CreateListSettlementPaymentsMetadataMapOfAny(map[string]any{/* values here */})
```

### 

```go
listSettlementPaymentsMetadata := operations.CreateListSettlementPaymentsMetadataArrayOfStr([]string{/* values here */})
```

