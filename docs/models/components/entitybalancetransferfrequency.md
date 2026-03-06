# EntityBalanceTransferFrequency

The frequency with which the available amount on the balance will be settled to the configured transfer
destination.

Settlements created during weekends or on bank holidays will take place on the next business day.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.EntityBalanceTransferFrequencyEveryDay

// Open enum: custom values can be created with a direct type cast
custom := components.EntityBalanceTransferFrequency("custom_value")
```


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `EntityBalanceTransferFrequencyEveryDay`       | every-day                                      |
| `EntityBalanceTransferFrequencyDaily`          | daily                                          |
| `EntityBalanceTransferFrequencyEveryMonday`    | every-monday                                   |
| `EntityBalanceTransferFrequencyEveryTuesday`   | every-tuesday                                  |
| `EntityBalanceTransferFrequencyEveryWednesday` | every-wednesday                                |
| `EntityBalanceTransferFrequencyEveryThursday`  | every-thursday                                 |
| `EntityBalanceTransferFrequencyEveryFriday`    | every-friday                                   |
| `EntityBalanceTransferFrequencyMonthly`        | monthly                                        |
| `EntityBalanceTransferFrequencyRevenueDay`     | revenue-day                                    |
| `EntityBalanceTransferFrequencyNever`          | never                                          |