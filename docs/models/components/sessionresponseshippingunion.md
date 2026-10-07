# SessionResponseShippingUnion

> 🚧 Private beta
>
> This property is currently in private beta, and the final specification may still change.

Shipping information for the Checkout Session. Provide either `options` or `callbackUrl`, not both.

The `lines` of the Checkout Session must not contain a line with type `shipping_fee`. When `shipping` is set,
`requiredCustomerDetails` must contain `shipping-address`.


## Supported Types

### SessionResponseShipping1

```go
sessionResponseShippingUnion := components.CreateSessionResponseShippingUnionSessionResponseShipping1(components.SessionResponseShipping1{/* values here */})
```

### SessionResponseShipping2

```go
sessionResponseShippingUnion := components.CreateSessionResponseShippingUnionSessionResponseShipping2(components.SessionResponseShipping2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch sessionResponseShippingUnion.Type {
	case components.SessionResponseShippingUnionTypeSessionResponseShipping1:
		// sessionResponseShippingUnion.SessionResponseShipping1 is populated
	case components.SessionResponseShippingUnionTypeSessionResponseShipping2:
		// sessionResponseShippingUnion.SessionResponseShipping2 is populated
}
```
