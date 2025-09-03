# CancelSubscriptionMetadata

Provide any data you like, for example a string or a JSON object. We will save the data alongside the entity.
Whenever you fetch the entity with our API, we will also include the metadata. You can use up to approximately
1kB.

Any metadata added to the subscription will be automatically forwarded to the payments generated for it.


## Supported Types

### 

```go
cancelSubscriptionMetadata := operations.CreateCancelSubscriptionMetadataStr(string{/* values here */})
```

### 

```go
cancelSubscriptionMetadata := operations.CreateCancelSubscriptionMetadataMapOfAny(map[string]any{/* values here */})
```

### 

```go
cancelSubscriptionMetadata := operations.CreateCancelSubscriptionMetadataArrayOfStr([]string{/* values here */})
```

