# UpdateSubscriptionMetadataRequest

Provide any data you like, for example a string or a JSON object. We will save the data alongside the
entity. Whenever you fetch the entity with our API, we will also include the metadata. You can use up to
approximately 1kB.

Any metadata added to the subscription will be automatically forwarded to the payments generated for it.


## Supported Types

### 

```go
updateSubscriptionMetadataRequest := operations.CreateUpdateSubscriptionMetadataRequestStr(string{/* values here */})
```

### 

```go
updateSubscriptionMetadataRequest := operations.CreateUpdateSubscriptionMetadataRequestMapOfAny(map[string]any{/* values here */})
```

### 

```go
updateSubscriptionMetadataRequest := operations.CreateUpdateSubscriptionMetadataRequestArrayOfStr([]string{/* values here */})
```

