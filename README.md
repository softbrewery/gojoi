# gojoi

[![Build Status](https://travis-ci.org/softbrewery/gojoi.svg?branch=master)](https://travis-ci.org/softbrewery/gojoi)

## Install
```shell
$ go get github.com/softbrewery/gojoi
```

## Api

### `Any`

Generates a schema object that matches any data type.

```go
schema := joi.Any()
```

#### `Any().Kind()`

Gets the type of the schema

```go
schema := joi.Any()
kind := schema.Kind() // kind == "interface"
```
```go
schema := joi.String()
kind := schema.Kind() // kind == "string"
```

#### `Any().Allow(values ...interface{})`

Whitelists a value where:
- `value` - the allowed value which can be of any type and will be matched against the validated value before applying any other rules.
  `value` can be a slice of values, or multiple values can be passed as individual arguments.

Pass in single or multiple arguments
```go
schema1 := joi.Any().Allow("id", "name")
schema2 := joi.Any().Allow(0, 10, 200)
schema3 := joi.Any().Allow(true)
```
Pass in a slice
```go
data := []string{"id", "name", "isbn"}
schema := joi.Any().Allow(data...)
```

#### `Any().Disallow(values ...interface{})`

Blacklists a value where:
- `value` - the forbidden value which can be of any type and will be matched against the validated value before applying any other rules.
  `value` can be an array of values, or multiple values can be passed as individual arguments.

Pass in single or multiple arguments
```go
schema1 := joi.Any().Disallow("id", "name")
schema2 := joi.Any().Disallow(0, 10, 200)
schema3 := joi.Any().Disallow(true)
```
Pass in a slice
```go
data := []string{"id", "name", "isbn"}
schema := joi.Any().Disallow(data...)
```

#### `Any().Required()`

Marks a key as required which will not allow `nil` as value. All keys are optional by default.

```go
schema := joi.Any().Required()
```

#### `Any().Forbidden()`

Marks a key as forbidden which will not allow any value except `nil`. Used to explicitly forbid keys.

```go
schema := joi.Any().Forbidden()
```

#### `Any().Description(desc string)`

Annotates the key where:
- `desc` - the description string.

```go
schema := joi.Any().Description("my description")
```

#### `Any().Transform(stage TransformStage, fn TransformFunc)`

Allows to inject a custom tranformation function, which can be used to
- Define custom validators
- Modify the value being validated


TransformFunc type definition
```go
type TransformFunc func(oldValue interface{}) (newValue interface{}, error)
```

```go
// Tranform function
fn := func(value interface{}) (interface{}, error) {

    cValue, ok := value.(string)
    if !ok {
        return nil, errors.New("Failed to cast type")
    }

    if cValue == "id" {
        cValue = "name"
    }
    
    return cValue, nil
}

// Build schema
schema := joi.Any().Allow("name").Transform(TransformStagePRE, fn)

// Validate
err := schema.Validate("id") // err == nil
```