# Entity


## Supported Types

### PaymentLinkResponse

```go
entity := components.CreateEntityPaymentLinkResponse(components.PaymentLinkResponse{/* values here */})
```

### ProfileResponse

```go
entity := components.CreateEntityProfileResponse(components.ProfileResponse{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch entity.Type {
	case components.EntityTypePaymentLinkResponse:
		// entity.PaymentLinkResponse is populated
	case components.EntityTypeProfileResponse:
		// entity.ProfileResponse is populated
}
```
