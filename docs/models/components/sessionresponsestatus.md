# SessionResponseStatus

The session's status.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.SessionResponseStatusOpen

// Open enum: custom values can be created with a direct type cast
custom := components.SessionResponseStatus("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `SessionResponseStatusOpen`      | open                             |
| `SessionResponseStatusCompleted` | completed                        |
| `SessionResponseStatusExpired`   | expired                          |