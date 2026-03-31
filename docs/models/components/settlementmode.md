# SettlementMode

Whether this entity was created in live mode or in test mode. Settlements are always in live mode.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.SettlementModeLive

// Open enum: custom values can be created with a direct type cast
custom := components.SettlementMode("custom_value")
```


## Values

| Name                 | Value                |
| -------------------- | -------------------- |
| `SettlementModeLive` | live                 |