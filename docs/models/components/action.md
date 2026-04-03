# Action

The action performed on the unmatched credit transfer.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.ActionMatch

// Open enum: custom values can be created with a direct type cast
custom := components.Action("custom_value")
```


## Values

| Name           | Value          |
| -------------- | -------------- |
| `ActionMatch`  | match          |
| `ActionReturn` | return         |