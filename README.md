# Go Ledger with fiber and docker

This application offers the following endpoints:

### [Account](docs/account.md)

| Path | Method | Description |
|------|--------|-------------|
| /accounts | GET | Return all accounts |
| /accounts/:document | GET | Get account by a document number |
| /accounts | POST | Create an account |
| /accounts | DELETE | Delete account softdelete |

### [Operations](docs/operations.md)

| Path | Method | Description |
|------|--------|-------------|
| /operations | GET | Return all operations |
| /operations | POST | Create an operations |
| /operations | DELETE | Delete an operations softdelete |

### [Transactions](docs/transactions.md)

| Path | Method | Description |
|------|--------|-------------|
| /transactions | POST | Create a transaction |

<br />

## Development environment setup

Requirements: go, docker, docker-container, postgres

Create and start a application and postgres instance with docker:

    docker-compose up -d

If any dependency is missing.

    go mod tidy

## Docs

- [Insomnia](docs/Insomnia_ledger_go.yaml)

- [Swagger](docs/swagger.yaml)

<br />

## Testing

#### unit

    cd pkg/shared/test
    go run test -v

#### Integration

    docker-compose up -d
    cd pkg/shared/test/integration
    go run test -v
