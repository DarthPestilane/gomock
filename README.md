# An easy fake data generator in golang [![Build Status](https://travis-ci.org/DarthPestilane/gofake.svg?branch=master)](https://travis-ci.org/DarthPestilane/gofake)

`go get -u github.com/DarthPestilane/gofake`

## example

```go
faker := gofake.NewFaker()

faker.Name() // => a random name
faker.Name("middle name") // => a random name with middle name
faker.NameFemale() // => a random female name
faker.NameFemale("middle name") // => a random female name with middle name

// ... for other functions checkout faker.go
```
