# ListProfileResponseStatus

The profile status determines whether the profile is able to receive live payments.

* `unverified`: The profile has not been verified yet and can only be used to create test payments.
* `verified`: The profile has been verified and can be used to create live payments and test payments.
* `blocked`: The profile is blocked and can no longer be used or changed.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.ListProfileResponseStatusUnverified

// Open enum: custom values can be created with a direct type cast
custom := components.ListProfileResponseStatus("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `ListProfileResponseStatusUnverified` | unverified                            |
| `ListProfileResponseStatusVerified`   | verified                              |
| `ListProfileResponseStatusBlocked`    | blocked                               |