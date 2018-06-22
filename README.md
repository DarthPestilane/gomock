# A mocker in golang.

`go get -u github.com/DarthPestilane/gomock`

## example

```go
mocker := gomock.NewMock()

mocker.Name() // => a random name
mocker.Name("middle name") // => a random name with middle name
mocker.NameFemale() // => a random female name
mocker.NameFemale("middle name") // => a random female name with middle name

// ... for other functions checkout mock_test.go
```
