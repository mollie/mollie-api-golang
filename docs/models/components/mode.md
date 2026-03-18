# Mode

Whether this entity was created in live mode or in test mode.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.ModeLive

// Open enum: custom values can be created with a direct type cast
custom := components.Mode("custom_value")
```


## Values

| Name       | Value      |
| ---------- | ---------- |
| `ModeLive` | live       |
| `ModeTest` | test       |