# OrganizationVatRegulation

Mollie applies Dutch VAT for merchants based in The Netherlands, British VAT for merchants based in The United
Kingdom, and shifted VAT for merchants in the European Union.

The field is not present for merchants residing in other countries.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.OrganizationVatRegulationDutch

// Open enum: custom values can be created with a direct type cast
custom := components.OrganizationVatRegulation("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `OrganizationVatRegulationDutch`   | dutch                              |
| `OrganizationVatRegulationBritish` | british                            |
| `OrganizationVatRegulationShifted` | shifted                            |