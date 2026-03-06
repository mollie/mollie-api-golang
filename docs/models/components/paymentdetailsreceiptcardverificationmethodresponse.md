# PaymentDetailsReceiptCardVerificationMethodResponse

The method used to verify the cardholder's identity.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.PaymentDetailsReceiptCardVerificationMethodResponseNoCvmRequired

// Open enum: custom values can be created with a direct type cast
custom := components.PaymentDetailsReceiptCardVerificationMethodResponse("custom_value")
```


## Values

| Name                                                                       | Value                                                                      |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `PaymentDetailsReceiptCardVerificationMethodResponseNoCvmRequired`         | no-cvm-required                                                            |
| `PaymentDetailsReceiptCardVerificationMethodResponseOnlinePin`             | online-pin                                                                 |
| `PaymentDetailsReceiptCardVerificationMethodResponseOfflinePin`            | offline-pin                                                                |
| `PaymentDetailsReceiptCardVerificationMethodResponseConsumerDevice`        | consumer-device                                                            |
| `PaymentDetailsReceiptCardVerificationMethodResponseSignature`             | signature                                                                  |
| `PaymentDetailsReceiptCardVerificationMethodResponseSignatureAndOnlinePin` | signature-and-online-pin                                                   |
| `PaymentDetailsReceiptCardVerificationMethodResponseOnlinePinAndSignature` | online-pin-and-signature                                                   |
| `PaymentDetailsReceiptCardVerificationMethodResponseNone`                  | none                                                                       |
| `PaymentDetailsReceiptCardVerificationMethodResponseFailed`                | failed                                                                     |