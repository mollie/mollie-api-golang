# OnboardingVatRegulation

Mollie applies Dutch VAT for merchants based in The Netherlands, British VAT for merchants based in
The United Kingdom, and shifted VAT for merchants in the European Union.

The field can be omitted for merchants residing in other countries.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.OnboardingVatRegulationDutch
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `OnboardingVatRegulationDutch`   | dutch                            |
| `OnboardingVatRegulationBritish` | british                          |
| `OnboardingVatRegulationShifted` | shifted                          |