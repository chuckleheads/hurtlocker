# Tasker

A specialized task runner for Habitat build jobs

## Usage

`tasker build --config sample_config/build.sample.yaml`

`tasker promote --config sample_config/promote.sample.yaml`

** Note: the config files can also be passed in as a base64 encoded string (useful when running tasker as a container) with `--config-string`

## Running locally

`go run main.go <COMMAND> --config <FILE>`
