# BalanceTransferPartyTypeResponse

Defines the type of the party. At the moment, only `organization` is supported.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.BalanceTransferPartyTypeResponseOrganization

// Open enum: custom values can be created with a direct type cast
custom := components.BalanceTransferPartyTypeResponse("custom_value")
```


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `BalanceTransferPartyTypeResponseOrganization` | organization                                   |