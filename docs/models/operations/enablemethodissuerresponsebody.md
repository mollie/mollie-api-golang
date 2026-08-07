# EnableMethodIssuerResponseBody

The payment method issuer object.


## Supported Types

### Giftcard

```go
enableMethodIssuerResponseBody := operations.CreateEnableMethodIssuerResponseBodyGiftcard(components.Giftcard{/* values here */})
```

### Voucher

```go
enableMethodIssuerResponseBody := operations.CreateEnableMethodIssuerResponseBodyVoucher(components.Voucher{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch enableMethodIssuerResponseBody.Type {
	case operations.EnableMethodIssuerResponseBodyTypeGiftcard:
		// enableMethodIssuerResponseBody.Giftcard is populated
	case operations.EnableMethodIssuerResponseBodyTypeVoucher:
		// enableMethodIssuerResponseBody.Voucher is populated
}
```
