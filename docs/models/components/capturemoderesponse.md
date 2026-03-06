# CaptureModeResponse

Indicate if the funds should be captured immediately or if you want to [place a hold](https://docs.mollie.com/docs/place-a-hold-for-a-payment#/) 
and capture at a later time.

This field needs to be set to `manual` for method `riverty`.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.CaptureModeResponseAutomatic

// Open enum: custom values can be created with a direct type cast
custom := components.CaptureModeResponse("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `CaptureModeResponseAutomatic` | automatic                      |
| `CaptureModeResponseManual`    | manual                         |