# CreateWebhookEventTypes


## Supported Types

### 

```go
createWebhookEventTypes := operations.CreateCreateWebhookEventTypesArrayOfWebhookEventTypes([]components.WebhookEventTypes{/* values here */})
```

### WebhookEventTypes

```go
createWebhookEventTypes := operations.CreateCreateWebhookEventTypesWebhookEventTypes(components.WebhookEventTypes{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createWebhookEventTypes.Type {
	case operations.CreateWebhookEventTypesTypeArrayOfWebhookEventTypes:
		// createWebhookEventTypes.ArrayOfWebhookEventTypes is populated
	case operations.CreateWebhookEventTypesTypeWebhookEventTypes:
		// createWebhookEventTypes.WebhookEventTypes is populated
}
```
