# GiftcardStatus

The status of the issuer.
If the status is `pending-issuer`, an additional action from your side may be required with the issuer.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.GiftcardStatusActivated

// Open enum: custom values can be created with a direct type cast
custom := components.GiftcardStatus("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `GiftcardStatusActivated`     | activated                     |
| `GiftcardStatusPendingIssuer` | pending-issuer                |