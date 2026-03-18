# PartnerType

Indicates the type of partner. Will be `null` if the currently authenticated organization is not
enrolled as a partner.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/operations"
)

value := operations.PartnerTypeOauth
```


## Values

| Name                    | Value                   |
| ----------------------- | ----------------------- |
| `PartnerTypeOauth`      | oauth                   |
| `PartnerTypeSignuplink` | signuplink              |
| `PartnerTypeUseragent`  | useragent               |