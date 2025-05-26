# DarkThroneApi

> **Disclaimer:** This project is provided as-is. The author is not responsible for how this library is used. Users are solely responsible for ensuring their usage complies with the Dark Throne Reborn site's Terms and Conditions.

DarkThroneApi is a Go client library for interacting with the Dark Throne Reborn MMO API. It provides convenient methods for authentication, player management, banking, and more.

## Features
- User authentication (login, register, logout)
- Player management (fetch, create, assume, unassume)
- Banking operations (deposit, withdraw gold)
- Structure upgrades and proficiency points (planned)
- Configurable logging
- Designed for automation and integration

## Installation

```
go get github.com/Rihoj/DarkThroneApi
```

## Usage

```go
import "github.com/Rihoj/DarkThroneApi"

func main() {
    api := DarkThroneApi.New(&DarkThroneApi.Config{Logger: nil})
    // Use api methods, e.g. api.DepositGold(...)
}
```

See GoDocs for full API reference.

## Development
- Requires Go 1.24+
- Run tests: `go test ./...`
- Lint: `staticcheck ./...`

## License
MIT
