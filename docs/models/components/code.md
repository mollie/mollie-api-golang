# Code

A machine-readable code that indicates the reason for the payment's status.


## Supported Types

### StatusReasonCardSchemeResponse

```go
code := components.CreateCodeStatusReasonCardSchemeResponse(components.StatusReasonCardSchemeResponse{/* values here */})
```

### StatusReasonMerchantResponse

```go
code := components.CreateCodeStatusReasonMerchantResponse(components.StatusReasonMerchantResponse{/* values here */})
```

### StatusReasonTerminalResponse

```go
code := components.CreateCodeStatusReasonTerminalResponse(components.StatusReasonTerminalResponse{/* values here */})
```

### StatusReasonVoucherResponse

```go
code := components.CreateCodeStatusReasonVoucherResponse(components.StatusReasonVoucherResponse{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch code.Type {
	case components.CodeTypeStatusReasonCardSchemeResponse:
		// code.StatusReasonCardSchemeResponse is populated
	case components.CodeTypeStatusReasonMerchantResponse:
		// code.StatusReasonMerchantResponse is populated
	case components.CodeTypeStatusReasonTerminalResponse:
		// code.StatusReasonTerminalResponse is populated
	case components.CodeTypeStatusReasonVoucherResponse:
		// code.StatusReasonVoucherResponse is populated
}
```
