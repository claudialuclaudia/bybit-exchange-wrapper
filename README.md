Exchange wrappers for Bybit Exchange

## Available Test Scripts

For each wrapper with test files written in the wrappers subdirectory you can run:

### `go test -v -run .`

This will run all tests in the directory and display all runtime logs in realtime with the `-v` verbose flag

### `go test -v -run . -testify.m {TestName}`

This command will search for the specified `TestName` within the directory and run only that particular test method

## Runtime and Tests setup

- Create a config.env file based off the config_template.env
- Stream example socket by running `go run .` from root
