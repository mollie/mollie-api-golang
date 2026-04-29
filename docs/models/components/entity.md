# Entity


## Supported Types

### PaymentLinkResponse

```go
entity := components.CreateEntityPaymentLinkResponse(components.PaymentLinkResponse{/* values here */})
```

### SalesInvoiceResponse

```go
entity := components.CreateEntitySalesInvoiceResponse(components.SalesInvoiceResponse{/* values here */})
```

### TransferResponse

```go
entity := components.CreateEntityTransferResponse(components.TransferResponse{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch entity.Type {
	case components.EntityTypePaymentLinkResponse:
		// entity.PaymentLinkResponse is populated
	case components.EntityTypeSalesInvoiceResponse:
		// entity.SalesInvoiceResponse is populated
	case components.EntityTypeTransferResponse:
		// entity.TransferResponse is populated
}
```
