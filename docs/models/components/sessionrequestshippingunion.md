# SessionRequestShippingUnion

> 🚧 Private beta
>
> This property is currently in private beta, and the final specification may still change.

Shipping information for the Checkout Session. Provide either `options` or `callbackUrl`, not both.

The `lines` of the Checkout Session must not contain a line with type `shipping_fee`. When `shipping` is set,
`requiredCustomerDetails` must contain `shipping-address`.


## Supported Types

### SessionRequestShipping1

```go
sessionRequestShippingUnion := components.CreateSessionRequestShippingUnionSessionRequestShipping1(components.SessionRequestShipping1{/* values here */})
```

### SessionRequestShipping2

```go
sessionRequestShippingUnion := components.CreateSessionRequestShippingUnionSessionRequestShipping2(components.SessionRequestShipping2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch sessionRequestShippingUnion.Type {
	case components.SessionRequestShippingUnionTypeSessionRequestShipping1:
		// sessionRequestShippingUnion.SessionRequestShipping1 is populated
	case components.SessionRequestShippingUnionTypeSessionRequestShipping2:
		// sessionRequestShippingUnion.SessionRequestShipping2 is populated
}
```
