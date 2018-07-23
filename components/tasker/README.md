# Tasker

A specialized task runner for Habitat build jobs

## Usage

`tasker build --config sample_config/build.sample.json`

`tasker promote --config sample_config/promote.sample.json`

** Note: the config files can also be passed in as a base64 encoded string (useful when running tasker as a container) with `--config-string`

## Running locally

`go run main.go <COMMAND> --config <FILE>`

## Running in Docker

* `docker build -t chuckleheads/tasker components/tasker`
* `docker run -v /hab/cache/keys:/hab/cache/keys -it --rm chuckleheads/tasker tasker build --config sample_config/build.sample.json`
* `docker run -v /var/run/docker.sock:/var/run/docker.sock -it --rm chuckleheads/tasker tasker export --config sample_config/export.sample.json`
