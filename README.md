# Go Ledger with fiber and docker

This application offers the following endpoints:

### [Account](docs/account)

| Path | Method | Description |
|------|--------|-------------|
| /accounts | GET | Return all accounts |
| /accounts/:document | GET | Get By a document number |
| /accounts | POST | Create an Account |
| /accounts | DELETE | Delete Account softdelete |

### [Operations](docs/operations)

| Path | Method | Description |
|------|--------|-------------|
| /operations | GET | Return all operations |
| /operations | POST | Create an operations |
| /operations | DELETE | Delete an operations softdelete |

### [Transactions](docs/transactions)

| Path | Method | Description |
|------|--------|-------------|
| /transactions | POST | Create an operations |

<br />

## Development environment setup

Requirements: go, docker, docker-container, postgres

Create and start a application and postgres instance with docker:

    docker-compose up -d

<br />

## Testing
