# BalancePrepaymentPartType

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.BalancePrepaymentPartTypeFee

// Open enum: custom values can be created with a direct type cast
custom := components.BalancePrepaymentPartType("custom_value")
```


## Values

| Name                                               | Value                                              |
| -------------------------------------------------- | -------------------------------------------------- |
| `BalancePrepaymentPartTypeFee`                     | fee                                                |
| `BalancePrepaymentPartTypeFeeReimbursement`        | fee-reimbursement                                  |
| `BalancePrepaymentPartTypeFeeDiscount`             | fee-discount                                       |
| `BalancePrepaymentPartTypeFeeVat`                  | fee-vat                                            |
| `BalancePrepaymentPartTypeFeeRoundingCompensation` | fee-rounding-compensation                          |