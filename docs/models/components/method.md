# Method


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
