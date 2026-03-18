# UpdateWebhookEventTypes


## Supported Types

### 

```go
updateWebhookEventTypes := operations.CreateUpdateWebhookEventTypesArrayOfWebhookEventTypes([]components.WebhookEventTypes{/* values here */})
```

### WebhookEventTypes

```go
updateWebhookEventTypes := operations.CreateUpdateWebhookEventTypesWebhookEventTypes(components.WebhookEventTypes{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch updateWebhookEventTypes.Type {
	case operations.UpdateWebhookEventTypesTypeArrayOfWebhookEventTypes:
		// updateWebhookEventTypes.ArrayOfWebhookEventTypes is populated
	case operations.UpdateWebhookEventTypesTypeWebhookEventTypes:
		// updateWebhookEventTypes.WebhookEventTypes is populated
}
```
