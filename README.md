# [Emailrepo.io](https://emailrep.io/) Go client

[![Go Report Card](https://goreportcard.com/badge/github.com/vertoforce/go-emailrep)](https://goreportcard.com/report/github.com/vertoforce/go-emailrep)
[![Documentation](https://pkg.go.dev/badge/github.com/vertoforce/go-emailrep)](https://pkg.go.dev/github.com/vertoforce/go-emailrep)

This is a simple library to interface with emailrepo.io.

* [x] Query an email
* [x] API Key support
* [ ] Submit an email

## Usage

```go
result, err := Query(context.Background(), "test@test.com", true)
```
