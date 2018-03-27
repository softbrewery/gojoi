![gojoi Logo](https://raw.github.com/softbrewery/gojoi/master/images/gojoi.png)

# gojoi

Object schema description language and validator for golang. (Inspired by [Hapi.js Joi](https://github.com/hapijs/joi) library)

[![Build Status](https://travis-ci.org/softbrewery/gojoi.svg?branch=master)](https://travis-ci.org/softbrewery/gojoi)
[![codecov](https://codecov.io/gh/softbrewery/gojoi/branch/master/graph/badge.svg)](https://codecov.io/gh/softbrewery/gojoi)

## Install
```shell
$ go get github.com/softbrewery/gojoi
```

## Usage

Usage is a two steps process; 

First, a schema is constructed:

```go
schema := joi.String()
```

Then the value is validated against the schema:

```go
err := joi.Validate("hello", schema)
```
If the input is valid, then the error will be nil, otherwise it will be an Error object.

Example to validate slice of strings:
```go
schema := joi.Slice().Items(
    joi.String(),
)

data := []string{"hello", "world"}

err := joi.Validate(data, schema)

// err == nil
```

---

## Table of Contents - API

<!--ts-->
   * [gojoi](#gojoi)
      * [Install](#install)
      * [Usage](#usage)
      * [Table of Contents - API](#table-of-contents---api)
      * [Api](#api)
         * [Any](#any)
            * [Any().Kind()](#anykind)
            * [Any().Allow(values ...interface{})](#anyallowvalues-interface)
            * [Any().Disallow(values ...interface{})](#anydisallowvalues-interface)
            * [Any().Required()](#anyrequired)
            * [Any().Forbidden()](#anyforbidden)
            * [Any().Description(desc string)](#anydescriptiondesc-string)
            * [Any().Transform(stage TransformStage, fn TransformFunc)](#anytransformstage-transformstage-fn-transformfunc)
         * [String - inherits from <code>Any</code>](#string---inherits-from-any)
            * [String().Min(limit int)](#stringminlimit-int)
            * [String().Max(limit int)](#stringmaxlimit-int)
            * [String().Length(limit int)](#stringlengthlimit-int)
            * [String().UpperCase()](#stringuppercase)
            * [String().LowerCase()](#stringlowercase)
            * [String().Regex(expression string)](#stringregexexpression-string)
            * [String().CreditCard()](#stringcreditcard)
         * [Int - inherits from <code>Any</code>](#int---inherits-from-any)
            * [Int().Min(limit int)](#intminlimit-int)
            * [Int().Max(limit int)](#intmaxlimit-int)
            * [Int().Positive()](#intpositive)
            * [Int().Negative()](#intnegative)
         * [Bool - inherits from <code>Any</code>](#bool---inherits-from-any)
         * [Slice - inherits from <code>Any</code>](#slice---inherits-from-any)
            * [Slice().Items(schema Schema)](#sliceitemsschema-schema)
            * [Slice().Min(limit int)](#sliceminlimit-int)
            * [Slice().Max(limit int)](#slicemaxlimit-int)
            * [Slice().Length(limit int)](#slicelengthlimit-int)
         * [Struct - inherits from <code>Any</code>](#struct---inherits-from-any)
            * [Struct().Keys(keys StructKeys{...})](#structkeyskeys-structkeys)

<!-- Added by: steven, at: 2018-03-27T11:26+02:00 -->

<!--te-->

---

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

Allows to run custom tranformation functions where:
- `stage` - defines the stage that triggers this transform
- `fn` - function that will be executes

Allowed staged:
- `joi.TransformStagePRE` - Executes before the validation starts
- `joi.TransformStagePOST` - Executes after the validation has finished

Use this functionality to:
- Inject custom validators
- Transform or normalize values 

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
schema := joi.Any().Allow("name").Transform(joi.TransformStagePRE, fn)

// Validate
err := joi.Validate("id", schema) // err == nil
```

---

### `String` - inherits from `Any`

Generates a schema object that matches string data type.

Supports the same methods of the any() type.

```go
schema := joi.String()
```

#### `String().Min(limit int)`

Specifies the minimum number string characters where:

- `limit` - the minimum number of string characters required.

```go
schema := joi.String().Min(2)
```

#### `String().Max(limit int)`

Specifies the maximum number string characters where:

- `limit` - the maximum number of string characters allowed.

```go
schema := joi.String().Max(10)
```

#### `String().Length(limit int)`

Specifies the exact string length required where:

- `limit` - the required string length.

```go
schema := joi.String().Length(5)
```

#### `String().UpperCase()`

Requires the string value to be all uppercase.

```go
schema := joi.String().UpperCase()
```

#### `String().LowerCase()`

Requires the string value to be all lowercase.

```go
schema := joi.String().LowerCase()
```

#### `String().Regex(expression string)`

Defines a regular expression rule where:

- `expression` - a regular expression object the string value must match against.

```go
schema := joi.String().Regex("^[abc]+$")
```

#### `String().CreditCard()`

Requires the number to be a credit card number Using [Luhn Algorithm](http://en.wikipedia.org/wiki/Luhn_algorithm).

```go
schema := joi.String().CreditCard()
```

---

### `Int` - inherits from `Any`

Generates a schema object that matches int data type.

Supports the same methods of the any() type.

```go
schema := joi.Int()
```

#### `Int().Min(limit int)`

Specifies the minimum number where:

- `limit` - the minimum number required.

```go
schema := joi.Int().Min(2)
```

#### `Int().Max(limit int)`

Specifies the maximum number where:

- `limit` - the maximum number allowed.

```go
schema := joi.Int().Max(10)
```

#### `Int().Positive()`

Requires the number to be positive.

```go
schema := joi.Int().Positive()
```

#### `Int().Negative()`

Requires the number to be negative.

```go
schema := joi.Int().Negative()
```

---

### `Bool` - inherits from `Any`

Generates a schema object that matches bool data type.

Supports the same methods of the any() type.

```go
schema := joi.Bool()
```

---

### `Slice` - inherits from `Any`

Generates a schema object that matches slice [] data type.

Supports the same methods of the any() type.

```go
schema := joi.Slice()
```

#### `Slice().Items(schema Schema)`

Lists the types allowed for the array values where:

- `schema` - a joi schema object to validate each array item against. `schema` can be an array of values, or multiple values can be passed as individual arguments.

```go
schema := joi.Slice().Items(
    joi.String(),
).Max(10)
```

#### `Slice().Min(limit int)`

Specifies the minimum number of items in the slice where:

- `limit` - the lowest number of array items allowed.

```go
schema := joi.Slice().Min(2)
```

#### `Slice().Max(limit int)`

Specifies the maximum number of items in the slice where:

- `limit` - the highest number of array items allowed.

```go
schema := joi.Slice().Max(10)
```

#### `Slice().Length(limit int)`

Specifies the exact number of items in the slice where:

- `limit` - the number of array items allowed.

```go
schema := joi.Slice().Length(5)
```

---

### `Struct` - inherits from `Any`

Generates a schema object that matches struct data type.

Supports the same methods of the any() type.

```go
schema := joi.Struct()
```

#### `Struct().Keys(keys StructKeys{...})`

Sets the allowed object keys where:

- `keys` - object where each key is assigned a joi type object. keys is of `StructKeys`


StructKeys is of type: `map[string]Schema`

Basic example:
```go
schema := joi.Struct().Keys(StructKeys{
    "Name": joi.String(),
})
```

Advanced example:
```go
schema := joi.Struct().Keys(StructKeys{
    "ID": joi.Forbidden(),
    "Name": joi.String().Required(),
    "Tags": joi.Slice().Items(
        joi.String().UpperCase().Length(4),
    ).Max(10),
    "Meta": joi.Struct(StructKeys{
        "Active": joi.Bool(),
    }),
})
```