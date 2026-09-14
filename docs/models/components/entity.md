# Entity


## Supported Types

### PaymentResponse

```go
entity := components.CreateEntityPaymentResponse(components.PaymentResponse{/* values here */})
```

### RefundResponse

```go
entity := components.CreateEntityRefundResponse(components.RefundResponse{/* values here */})
```

### EntityChargeback

```go
entity := components.CreateEntityEntityChargeback(components.EntityChargeback{/* values here */})
```

### CaptureResponse

```go
entity := components.CreateEntityCaptureResponse(components.CaptureResponse{/* values here */})
```

### PaymentLinkResponse

```go
entity := components.CreateEntityPaymentLinkResponse(components.PaymentLinkResponse{/* values here */})
```

### EntityPayoutResponse

```go
entity := components.CreateEntityEntityPayoutResponse(components.EntityPayoutResponse{/* values here */})
```

### SalesInvoiceResponse

```go
entity := components.CreateEntitySalesInvoiceResponse(components.SalesInvoiceResponse{/* values here */})
```

### TransferResponse

```go
entity := components.CreateEntityTransferResponse(components.TransferResponse{/* values here */})
```

### DraftTransferResponse

```go
entity := components.CreateEntityDraftTransferResponse(components.DraftTransferResponse{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch entity.Type {
	case components.EntityTypePaymentResponse:
		// entity.PaymentResponse is populated
	case components.EntityTypeRefundResponse:
		// entity.RefundResponse is populated
	case components.EntityTypeEntityChargeback:
		// entity.EntityChargeback is populated
	case components.EntityTypeCaptureResponse:
		// entity.CaptureResponse is populated
	case components.EntityTypePaymentLinkResponse:
		// entity.PaymentLinkResponse is populated
	case components.EntityTypeEntityPayoutResponse:
		// entity.EntityPayoutResponse is populated
	case components.EntityTypeSalesInvoiceResponse:
		// entity.SalesInvoiceResponse is populated
	case components.EntityTypeTransferResponse:
		// entity.TransferResponse is populated
	case components.EntityTypeDraftTransferResponse:
		// entity.DraftTransferResponse is populated
}
```
