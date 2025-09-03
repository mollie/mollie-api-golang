# ListCustomerPaymentsMetadata

Provide any data you like, for example a string or a JSON object. We will save the data alongside the entity. Whenever
you fetch the entity with our API, we will also include the metadata. You can use up to approximately 1kB.


## Supported Types

### 

```go
listCustomerPaymentsMetadata := operations.CreateListCustomerPaymentsMetadataStr(string{/* values here */})
```

### 

```go
listCustomerPaymentsMetadata := operations.CreateListCustomerPaymentsMetadataMapOfAny(map[string]any{/* values here */})
```

### 

```go
listCustomerPaymentsMetadata := operations.CreateListCustomerPaymentsMetadataArrayOfStr([]string{/* values here */})
```

