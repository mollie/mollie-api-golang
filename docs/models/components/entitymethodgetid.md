# EntityMethodGetID

The unique identifier of the payment method. When used during [payment creation](create-payment), the payment
method selection screen will be skipped.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.EntityMethodGetIDAlma

// Open enum: custom values can be created with a direct type cast
custom := components.EntityMethodGetID("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `EntityMethodGetIDAlma`           | alma                              |
| `EntityMethodGetIDApplepay`       | applepay                          |
| `EntityMethodGetIDBacs`           | bacs                              |
| `EntityMethodGetIDBancomatpay`    | bancomatpay                       |
| `EntityMethodGetIDBancontact`     | bancontact                        |
| `EntityMethodGetIDBanktransfer`   | banktransfer                      |
| `EntityMethodGetIDBelfius`        | belfius                           |
| `EntityMethodGetIDBillie`         | billie                            |
| `EntityMethodGetIDBizum`          | bizum                             |
| `EntityMethodGetIDBlik`           | blik                              |
| `EntityMethodGetIDCreditcard`     | creditcard                        |
| `EntityMethodGetIDDirectdebit`    | directdebit                       |
| `EntityMethodGetIDEps`            | eps                               |
| `EntityMethodGetIDGiftcard`       | giftcard                          |
| `EntityMethodGetIDIdeal`          | ideal                             |
| `EntityMethodGetIDIn3`            | in3                               |
| `EntityMethodGetIDKbc`            | kbc                               |
| `EntityMethodGetIDKlarna`         | klarna                            |
| `EntityMethodGetIDMbway`          | mbway                             |
| `EntityMethodGetIDMobilepay`      | mobilepay                         |
| `EntityMethodGetIDMultibanco`     | multibanco                        |
| `EntityMethodGetIDMybank`         | mybank                            |
| `EntityMethodGetIDPaybybank`      | paybybank                         |
| `EntityMethodGetIDPaypal`         | paypal                            |
| `EntityMethodGetIDPaysafecard`    | paysafecard                       |
| `EntityMethodGetIDPointofsale`    | pointofsale                       |
| `EntityMethodGetIDPrzelewy24`     | przelewy24                        |
| `EntityMethodGetIDRiverty`        | riverty                           |
| `EntityMethodGetIDSatispay`       | satispay                          |
| `EntityMethodGetIDSwish`          | swish                             |
| `EntityMethodGetIDTrustly`        | trustly                           |
| `EntityMethodGetIDTwint`          | twint                             |
| `EntityMethodGetIDVipps`          | vipps                             |
| `EntityMethodGetIDVoucher`        | voucher                           |
| `EntityMethodGetIDKlarnapaylater` | klarnapaylater                    |
| `EntityMethodGetIDKlarnapaynow`   | klarnapaynow                      |
| `EntityMethodGetIDKlarnasliceit`  | klarnasliceit                     |
| `EntityMethodGetIDPayconiq`       | payconiq                          |