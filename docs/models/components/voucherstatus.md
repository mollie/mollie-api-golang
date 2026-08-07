# VoucherStatus

The status of the issuer.
If the status is `pending-issuer`, an additional action from your side may be required with the issuer.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.VoucherStatusActivated

// Open enum: custom values can be created with a direct type cast
custom := components.VoucherStatus("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `VoucherStatusActivated`     | activated                    |
| `VoucherStatusPendingIssuer` | pending-issuer               |