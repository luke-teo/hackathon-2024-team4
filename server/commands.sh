#!/bin/bash
#
show_divider() {
    echo '------------------------------------------------------------'
}

use_env() {
    echo "Using .env.${1} as ENV"
    set -o allexport
    source ".env.${1}"
    set +o allexport
}

cli_help() {
  echo "
  Available commands

  codegen:handler                       Autogenerates stubs for handlers
  dev_packages:install                  Install all dev only packages
  format                                Formats output
  hooks:install                         Install git hooks to format code on commit
  init:db                               Initializes databases in local docker volumes
  lint                                  Lints existing code
  migration:codegen {migration_name}    Autogenerates empty migration files
  migration:up                          Runs migrations and autogenerates DB models
  migration:down                        Rolls back 1 migration
  migration:status                      Shows status of past migrations
  openapi:codegen                       Autogenerates OpenAPI Client/Server code
  openapi:compare                       Compare two OpenAPI spec files for breaking changes
  openapi:format                        Format OpenAPI file and strip unwanted symbols
  webserver                             Runs webserver with .env.local
  console                               Runs console for running one-time commands
  worker                                Runs worker process to process background tasks 
  "
}

show_divider

case "${1}" in
  "help")
    cli_help
    exit 0
    ;;

  "codegen:handler")
    echo 'Analyzing oapi.StrictServerInterface for methods that need to be implemented'
    ./bin/impl 'h *Handler' oapi.StrictServerInterface
    echo 'Please check which methods need to be implement (hint: grep the output)'
    echo '(e.g. ./commands.sh codegen:handler | grep api/v1/user -A 3)'
    exit 0
    ;;

  "dev_packages:install")
    echo 'Installing tools for local development'
    echo 'Tools will be installed in the ./bin folder'

    mkdir -p bin

    export GOBIN="${PWD}/bin"

    # CLI tool to run migrations against a DB
    go install github.com/pressly/goose/v3/cmd/goose@v3.20.0

    # parse OpenAPI specs and generate client and server code
    go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v2.3.0

    # read DB and autogenerate model structs
    go install github.com/go-jet/jet/v2/cmd/jet@v2.11.1

    # code coverage
    go install golang.org/x/tools/cmd/cover@latest

    # formatting
    go install golang.org/x/tools/cmd/goimports@v0.22.0
    go install github.com/segmentio/golines@v0.12.2
    echo 'All tools installed successfully'

    exit 0
    ;;

  "format")
    echo 'Formatting all code';
    ./bin/goimports -w .
    ./bin/golines -w .
    gofmt -w .
    echo 'Finished formatting all code';
    exit 0
    ;;

  "lint")
    docker run --rm -v $(pwd):/app -v ~/.cache/golangci-lint/v1.54.2:/root/.cache -w /app golangci/golangci-lint:v1.54.2 golangci-lint run
    ;;

  "openapi:codegen")
    echo 'Generating OpenAPI client/server code from ./openapi.yaml';
    ./bin/oapi-codegen -config ./generated/oapi/codegen.yaml ./openapi.yaml
    echo 'Finished generating openapi code';
    exit 0
    ;;

  "openapi:compare")
    # to run comparison, make create a new input.yaml file and put your new openapi.yaml into it
    # this command can also be ran with remote URLs
    docker run --rm -t \
        -v $(pwd):/specs:ro \
        openapitools/openapi-diff:latest /specs/openapi.yaml /specs/input.yaml
    ;;

  "openapi:format")
    docker run --rm -t \
        -v $(pwd):/usr/src/myapp -w /usr/src/myapp python:3.11 \
        bash -c "python -m pip install pyyaml && python ./scripts/format_yaml.py -i openapi.yaml -o openapi_v2.yaml -r x-stoplight"
    ;;

  "init:db")
    echo 'Initalize databases';
    docker-compose up -d

    use_env "local"
    docker-compose exec db psql -U postgres -c "DROP DATABASE IF EXISTS ${DATABASE_NAME};"
    docker-compose exec db psql -U postgres -c "CREATE DATABASE ${DATABASE_NAME};"

    use_env "test"
    docker-compose exec db psql -U postgres -c "DROP DATABASE IF EXISTS ${DATABASE_NAME};"
    docker-compose exec db psql -U postgres -c "CREATE DATABASE ${DATABASE_NAME};"
    exit 0
    ;;

  "console")
    echo 'Starting console with .env.local';
    use_env "local"
    go run ./cmd/console ${*:2}
    exit 0
    ;;

  "webserver")
    echo 'Starting server with .env.local';
    use_env "local"
    go run ./cmd/webserver
    exit 0
    ;;

  "worker")
    echo 'Starting worker with .env.local';
    use_env "local"
    go run ./cmd/worker
    exit 0
    ;;

  "migration:codegen")
    mkdir -p db/migrations
    # create a stub migration files
    # -s means to create a sequential migraiton instead of timestamped
    ./bin/goose -s -dir db/migrations create "${2}" sql
    exit 0
    ;;

  "migration:status")
    use_env "local"
    ./bin/goose -dir db/migrations postgres "$DATABASE_URL" status
    exit 0
    ;;

  "migration:up")
    # run migrations for local db
    echo 'Running migrations from db/migrations';

    show_divider
    echo 'Migrating local DB';
    use_env "local"
    ./bin/goose -dir db/migrations postgres "$DATABASE_URL" up

    show_divider

    echo 'Autogenerating DB models';
    use_env "local"
    ./bin/jet -dsn="$DATABASE_URL" -schema=public -path=./generated/db
    # run formatting on autogenerated code
    bash ./commands.sh format

    show_divider

    echo 'Migrating local DB for tests';
    use_env "test"
    ./bin/goose -dir db/migrations postgres "$DATABASE_URL" up
    echo 'Finished running migrations';
    exit 0
    ;;

  "migration:down")
    echo 'Migrating local DB';
    use_env "local"
    ./bin/goose -dir db/migrations postgres "$DATABASE_URL" down

    echo 'Migrating local DB for tests';
    use_env "test"
    ./bin/goose -dir db/migrations postgres "$DATABASE_URL" down
    echo 'Finished running migrations';
    exit 0
    ;;

  "hooks:install")
    # copy git hook to hooks directory
    cp ./pre-commit ./.git/hooks/pre-commit
    chmod +x .git/hooks/pre-commit
    echo "Copied git hook to ./.git/hooks/pre-commit"
    exit 0
    ;;

  "test")
    use_env "test"
    go test --cover -covermode=count -coverpkg=./...  ./test/... -coverprofile=cover.out
    go tool cover -html=cover.out
    ;;

  *)
    echo 'ERROR: No command provided'
    cli_help
    exit 1
    ;;
esac

