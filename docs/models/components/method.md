# Method

Normally, a payment method screen is shown. However, when using this parameter, you can choose a specific
payment method and your customer will skip the selection screen and is sent directly to the chosen payment
method. The parameter enables you to fully integrate the payment method selection into your website.

You can also specify the methods in an array. By doing so we will still show the payment method selection
screen but will only show the methods specified in the array. For example, you can use this functionality
to only show payment methods from a specific country to your customer `['bancontact', 'belfius']`.


## Supported Types

### MethodEnum

```go
method := components.CreateMethodMethodEnum(components.MethodEnum{/* values here */})
```

### 

```go
method := components.CreateMethodArrayOfMethodEnum([]*components.MethodEnum{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch method.Type {
	case components.MethodTypeMethodEnum:
		// method.MethodEnum is populated
	case components.MethodTypeArrayOfMethodEnum:
		// method.ArrayOfMethodEnum is populated
}
```
