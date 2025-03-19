### Installation

Instructions on how to install the project. This might include downloading from a package manager, cloning the repository, or other setup steps.

```bash
go install github.com/swaggo/swag/cmd/swag@latest
go install github.com/air-verse/air@latest
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

### Usage

Examples and instructions for using the project. This could include code snippets, commands, or configuration details.

```bash
make mod
make swag
```

### Development

Guidelines for contributing to the project. This might include information on how to run the project locally, testing and debugging, or how to build and release the project.

```bash
make dev
```

### Testing

Instructions on how to test the project. This might include information on how to run unit tests, integration tests, or end-to-end tests.

Run all tests and generate coverage report in text format in the console.

```bash
make test
```

## Configuration

Details on how to configure the project for different environments or use cases. This might include environment variables, configuration files, or settings.

Rename the `.env.sample` file to `.env`
To add more environment variables, add them to the `.env`. The format is as follows:

```bash
VAR=value
```

## License

This project is licensed under the [License Name] License - see the [LICENSE.md](LICENSE.md) file for details.
