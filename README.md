# About
Conversion of unit library for golang


Fork form [martinlindhe](https://github.com/martinlindhe/unit)


Change Unit type from float64 to interface,  


Add my own data type.  


[![GoDoc](https://godoc.org/github.com/CokinLee/unit?status.svg)](https://godoc.org/github.com/martinlindhe/unit)
[![codecov.io](https://codecov.io/github/CokinLee/unit/coverage.svg?branch=master)](https://codecov.io/github/martinlindhe/unit?branch=master)


Conversion of unit library for golang


## Installation

```
go get -u github.com/CokinLee/unit
```


## General use

Basic usage:
```go
ft := 1 * unit.Foot
fmt.Println(ft.Feet(), "feet is", ft.Meters(), "meters")
```

To use your own data type, you need to convert to the base unit first (eg Length, Speed etc):
```go
type MyUnit int

n := MyUnit(2)
ft := Length(n) * Foot
fmt.Println(ft.Feet(), "feet is", ft.Meters(), "meters")
```


## Temperature

Cannot be used to scale directly like the other units.
Instead, use the From* functions to create a Temperature type:

```go
f := unit.FromFahrenheit(100)

fmt.Println("100 fahrenheit in celsius = ", f.Celsius())
```


## Future work

Please note the resulting precision is limited to the float64 type.
Big decimal version is being tracked in https://github.com/martinlindhe/unit/issues/3


## License

Under [MIT](LICENSE)
