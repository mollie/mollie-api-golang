# ListEntityMethodAllID

The unique identifier of the payment method. When used during [payment creation](create-payment), the payment
method selection screen will be skipped.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.ListEntityMethodAllIDAlma

// Open enum: custom values can be created with a direct type cast
custom := components.ListEntityMethodAllID("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `ListEntityMethodAllIDAlma`         | alma                                |
| `ListEntityMethodAllIDApplepay`     | applepay                            |
| `ListEntityMethodAllIDBacs`         | bacs                                |
| `ListEntityMethodAllIDBancomatpay`  | bancomatpay                         |
| `ListEntityMethodAllIDBancontact`   | bancontact                          |
| `ListEntityMethodAllIDBanktransfer` | banktransfer                        |
| `ListEntityMethodAllIDBelfius`      | belfius                             |
| `ListEntityMethodAllIDBillie`       | billie                              |
| `ListEntityMethodAllIDBizum`        | bizum                               |
| `ListEntityMethodAllIDBlik`         | blik                                |
| `ListEntityMethodAllIDCreditcard`   | creditcard                          |
| `ListEntityMethodAllIDDirectdebit`  | directdebit                         |
| `ListEntityMethodAllIDEps`          | eps                                 |
| `ListEntityMethodAllIDGiftcard`     | giftcard                            |
| `ListEntityMethodAllIDIdeal`        | ideal                               |
| `ListEntityMethodAllIDIn3`          | in3                                 |
| `ListEntityMethodAllIDKbc`          | kbc                                 |
| `ListEntityMethodAllIDKlarna`       | klarna                              |
| `ListEntityMethodAllIDMbway`        | mbway                               |
| `ListEntityMethodAllIDMobilepay`    | mobilepay                           |
| `ListEntityMethodAllIDMultibanco`   | multibanco                          |
| `ListEntityMethodAllIDMybank`       | mybank                              |
| `ListEntityMethodAllIDPaybybank`    | paybybank                           |
| `ListEntityMethodAllIDPaypal`       | paypal                              |
| `ListEntityMethodAllIDPaysafecard`  | paysafecard                         |
| `ListEntityMethodAllIDPrzelewy24`   | przelewy24                          |
| `ListEntityMethodAllIDRiverty`      | riverty                             |
| `ListEntityMethodAllIDSatispay`     | satispay                            |
| `ListEntityMethodAllIDSwish`        | swish                               |
| `ListEntityMethodAllIDTrustly`      | trustly                             |
| `ListEntityMethodAllIDTwint`        | twint                               |
| `ListEntityMethodAllIDVipps`        | vipps                               |
| `ListEntityMethodAllIDVoucher`      | voucher                             |