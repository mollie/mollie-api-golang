# ListEntityBalanceTransferFrequency

The frequency with which the available amount on the balance will be settled to the configured transfer
destination.

Settlements created during weekends or on bank holidays will take place on the next business day.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.ListEntityBalanceTransferFrequencyEveryDay

// Open enum: custom values can be created with a direct type cast
custom := components.ListEntityBalanceTransferFrequency("custom_value")
```


## Values

| Name                                               | Value                                              |
| -------------------------------------------------- | -------------------------------------------------- |
| `ListEntityBalanceTransferFrequencyEveryDay`       | every-day                                          |
| `ListEntityBalanceTransferFrequencyDaily`          | daily                                              |
| `ListEntityBalanceTransferFrequencyEveryMonday`    | every-monday                                       |
| `ListEntityBalanceTransferFrequencyEveryTuesday`   | every-tuesday                                      |
| `ListEntityBalanceTransferFrequencyEveryWednesday` | every-wednesday                                    |
| `ListEntityBalanceTransferFrequencyEveryThursday`  | every-thursday                                     |
| `ListEntityBalanceTransferFrequencyEveryFriday`    | every-friday                                       |
| `ListEntityBalanceTransferFrequencyMonthly`        | monthly                                            |
| `ListEntityBalanceTransferFrequencyRevenueDay`     | revenue-day                                        |
| `ListEntityBalanceTransferFrequencyNever`          | never                                              |