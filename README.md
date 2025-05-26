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

See GoDocs for full API reference. If published, you can also browse the API at [pkg.go.dev](https://pkg.go.dev/github.com/Rihoj/DarkThroneApi).

## Linting & Commit Requirements

- **Node.js and npm are required for development tooling (commit hooks, semantic-release, etc.), but not for using the Go library itself.**
- All commits must follow [Conventional Commits](https://www.conventionalcommits.org/) style. This is enforced by [commitlint](https://github.com/conventional-changelog/commitlint) in CI and via a pre-commit hook.
- Linting is enforced on every commit using [`staticcheck`](https://staticcheck.io/) for Go code. Commits will be blocked if linting fails.
- You must have `staticcheck` installed and available in your `PATH`. Install it with:

```sh
go install honnef.co/go/tools/cmd/staticcheck@latest
```

- Node.js dev dependencies for commit linting and hooks are managed in `package.json`:
  - `@commitlint/cli`, `@commitlint/config-conventional`, `husky`

## Development
- Requires Go 1.24+
- Run tests: `go test ./...`
- Lint: `staticcheck ./...`

## License
MIT

---

### Suggestions for Contributors
- If you wish to contribute, please ensure you have Node.js, npm, and Go installed.
- Make sure your `staticcheck` binary is in your `PATH` (e.g., add `export PATH="$HOME/go/bin:$PATH"` to your shell profile if needed).
- For more information on GoDoc, see [pkg.go.dev](https://pkg.go.dev/github.com/Rihoj/DarkThroneApi).
- If you have questions or suggestions, feel free to open an issue or pull request.
