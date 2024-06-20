# Golang Playground App

## Commands

```bash
# run other commands
./commands.sh --help
```

## Using this template to start a new project

1. Replace all instances of `go_chi_template` with your intended project name/URL
2. Remove the "db/go_chi_template" folder
3. Remove the existing handlers except "base_handler.go"
4. Remove the existing internal\app code
5. Remove the existing tests

## VSCode integration
Please add the following in your project specific .vscode/settings.json file.

```
{
    "go.testEnvFile": "${workspaceFolder}/.env.test",
    "go.lintTool": "golangci-lint",
    "go.lintFlags": [
      "--fast"
    ]
}

```


## Setup

Ensure you have the following software installed beforehand
- Docker Compose
- Go (Recommended to install via [asdf](https://asdf-vm.com/))

Run the follow commands to set up the project
```bash
# starts database in docker
docker-compose up -d

# initialize the DB
./commands.sh dev_packages:install
./commands.sh init:db
./commands.sh migration:run

# Install Go Dependencies
go mod download

# after setup is completed. you can start the server via
./commands.sh serve:local
```

## Workflow

### DB changes

1. Create a new migration (refer to `./commands.sh`)
2. Edit the generated `up` and `down` migrations
3. Run the migrations and generate model files (refer to `./commands.sh`)

### API changes

1. Edit the `./openapi.yaml` file with any editor you prefer
    - try to only add 1 endpoint at a time
    - if using certain editors like Stoplight Studio, strip the stoplight tags before commiting the file
2. Autogenerate the OpenAPI go-chi router code (refer to `./commands.sh`)
3. Add handler for new endpoint (trying to run the server will raise compilation errors if you don't)
    - You can generate handler stub via `./commands.sh`

## IDE integration
1. A VSCode launch.json is included, you can use it to launch the server and/or run tests
2. The Golang app does not run in a docker container so you can use breakpoints in your IDE as needed

## Background Jobs
1. Refer to `/internal/app/task/tenant_cleanup.go` for a sample of creating a new task
2. Similar to webserver handlers, new tasks need to be registered at `/internal/worker/worker.go` (line 50)
3. You can test the background jobs by
    - start the worker process `./commands.sh worker`
    - dispatch a test job via `./commands.sh console tenant:cleanup "Housekeeping"`
    - view the logs to see what happened

