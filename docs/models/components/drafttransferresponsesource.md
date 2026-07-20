# DraftTransferResponseSource

Whether the draft transfer was created via this API, or created in Mollie Apps.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.DraftTransferResponseSourceAPI

// Open enum: custom values can be created with a direct type cast
custom := components.DraftTransferResponseSource("custom_value")
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `DraftTransferResponseSourceAPI`       | api                                    |
| `DraftTransferResponseSourceMollieApp` | mollie-app                             |