# PaymentDetailsSellerProtectionResponse

Indicates to what extent the payment is eligible for PayPal's Seller Protection. Only available for PayPal
payments, and if the information is made available by PayPal.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.PaymentDetailsSellerProtectionResponseEligible

// Open enum: custom values can be created with a direct type cast
custom := components.PaymentDetailsSellerProtectionResponse("custom_value")
```


## Values

| Name                                                                      | Value                                                                     |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `PaymentDetailsSellerProtectionResponseEligible`                          | Eligible                                                                  |
| `PaymentDetailsSellerProtectionResponseIneligible`                        | Ineligible                                                                |
| `PaymentDetailsSellerProtectionResponsePartiallyEligibleInrOnly`          | Partially Eligible - INR Only                                             |
| `PaymentDetailsSellerProtectionResponsePartiallyEligibleUnauthOnly`       | Partially Eligible - Unauth Only                                          |
| `PaymentDetailsSellerProtectionResponsePartiallyEligible`                 | Partially Eligible                                                        |
| `PaymentDetailsSellerProtectionResponseNone`                              | None                                                                      |
| `PaymentDetailsSellerProtectionResponseActive`                            | Active                                                                    |
| `PaymentDetailsSellerProtectionResponseFraudControlUnauthPremiumEligible` | Fraud Control - Unauth Premium Eligible                                   |