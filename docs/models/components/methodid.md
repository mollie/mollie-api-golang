# MethodID

Normally, a payment method screen is shown. However, when using this parameter, you can choose a specific payment
method and your customer will skip the selection screen and is sent directly to the chosen payment method. The
parameter enables you to fully integrate the payment method selection into your website.

You can also specify the methods in an array. By doing so we will still show the payment method selection screen
but will only show the methods specified in the array. For example, you can use this functionality to only show
payment methods from a specific country to your customer `['bancontact', 'belfius']`.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.MethodIDAlma

// Open enum: custom values can be created with a direct type cast
custom := components.MethodID("custom_value")
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `MethodIDAlma`         | alma                   |
| `MethodIDApplepay`     | applepay               |
| `MethodIDBacs`         | bacs                   |
| `MethodIDBancomatpay`  | bancomatpay            |
| `MethodIDBancontact`   | bancontact             |
| `MethodIDBanktransfer` | banktransfer           |
| `MethodIDBelfius`      | belfius                |
| `MethodIDBillie`       | billie                 |
| `MethodIDBillink`      | billink                |
| `MethodIDBizum`        | bizum                  |
| `MethodIDBlik`         | blik                   |
| `MethodIDCreditcard`   | creditcard             |
| `MethodIDDirectdebit`  | directdebit            |
| `MethodIDEps`          | eps                    |
| `MethodIDGiftcard`     | giftcard               |
| `MethodIDGooglepay`    | googlepay              |
| `MethodIDIdeal`        | ideal                  |
| `MethodIDIn3`          | in3                    |
| `MethodIDKbc`          | kbc                    |
| `MethodIDKlarna`       | klarna                 |
| `MethodIDMbway`        | mbway                  |
| `MethodIDMobilepay`    | mobilepay              |
| `MethodIDMultibanco`   | multibanco             |
| `MethodIDMybank`       | mybank                 |
| `MethodIDPaybybank`    | paybybank              |
| `MethodIDPaypal`       | paypal                 |
| `MethodIDPaysafecard`  | paysafecard            |
| `MethodIDPointofsale`  | pointofsale            |
| `MethodIDPrzelewy24`   | przelewy24             |
| `MethodIDRiverty`      | riverty                |
| `MethodIDSatispay`     | satispay               |
| `MethodIDSwish`        | swish                  |
| `MethodIDTrustly`      | trustly                |
| `MethodIDTwint`        | twint                  |
| `MethodIDVipps`        | vipps                  |
| `MethodIDVoucher`      | voucher                |
| `MethodIDWero`         | wero                   |