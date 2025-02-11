# LockNote

LockNote is a simple API that allows users to securely store and manage their notes.

This project was created as an experiment about how to build a REST API using [Paseto](https://github.com/paragonie/paseto) for authentication and authorization, instead of JWT. Paseto is a specification for secure stateless tokens and serves as a more secure alternative to JWT. This project also uses [Zap](https://github.com/uber-go/zap) and [Echo](https://github.com/labstack/echo) as the web framework.
## Features

- Create user accounts
- Create, read, update, and delete notes
- Authentication of users using Paseto

## Prerequisites

- [Go](https://go.dev/doc/install)
- [PostgreSQL](https://www.postgresql.org/download/)
- [Docker](https://www.docker.com/products/docker-desktop)
- [Air](https://github.com/air-verse/air)
- [Echo](https://github.com/labstack/echo)
- [Make](https://www.gnu.org/software/make/)
- [Migrate](https://github.com/golang-migrate/migrate)
- [Mockery](https://github.com/vektra/mockery)
- [Sqlc](https://github.com/kyleconroy/sqlc)
- [Zap](https://github.com/uber-go/zap)

## License

LockNote is released under the MIT License. See the [LICENSE](LICENSE) file for more information.