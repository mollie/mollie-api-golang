# ProfileReviewStatusResponse

The status of the requested changes.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.ProfileReviewStatusResponsePending

// Open enum: custom values can be created with a direct type cast
custom := components.ProfileReviewStatusResponse("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `ProfileReviewStatusResponsePending`  | pending                               |
| `ProfileReviewStatusResponseRejected` | rejected                              |