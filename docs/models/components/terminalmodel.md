# TerminalModel

The model of the terminal. For example for a PAX A920, this field's value will be `A920`.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.TerminalModelA35

// Open enum: custom values can be created with a direct type cast
custom := components.TerminalModel("custom_value")
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `TerminalModelA35`     | A35                    |
| `TerminalModelA77`     | A77                    |
| `TerminalModelA920`    | A920                   |
| `TerminalModelA920Pro` | A920Pro                |
| `TerminalModelIm30`    | IM30                   |
| `TerminalModelTap`     | Tap                    |