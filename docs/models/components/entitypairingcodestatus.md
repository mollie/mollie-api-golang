# EntityPairingCodeStatus

The status of the pairing code.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.EntityPairingCodeStatusActive

// Open enum: custom values can be created with a direct type cast
custom := components.EntityPairingCodeStatus("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `EntityPairingCodeStatusActive`  | active                           |
| `EntityPairingCodeStatusExpired` | expired                          |
| `EntityPairingCodeStatusRevoked` | revoked                          |