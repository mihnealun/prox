# prox

## Setup GoLand

### Enable modules vendoring mode
Enable `Preferences` > `Go Modules (vgo)` > `Vendoring Mode`.

### File watcher linter
Add a new file watcher in GoLand in `Preferences` > `File Watchers` > `Add New` and select `golangci-lint`.

## General Setup

### Install `pre-commit` if you need to run stuff via git hook events (configuration available in .pre-commit-config.yaml)
Run 
~~~~
curl https://pre-commit.com/install-local.py | python -
~~~~
Go to project dir and run:
~~~~
pre-commit install
~~~~


## Run

~~~~
docker-compose up
~~~~

The docker container will be automatically restarted every time you change something in a .go file 


## Logs

~~~~
tail -f var/log/all.log
~~~~


## Endpoints

Endpoints are defined in the routes.json file and registered upon application start

Here is an example for an endpoint JSON structure:

~~~~
{
      "name": "user_list",
      "path": "/user",
      "template": {
            "provider": "redis",
            "key": "user_template"
      },
      "data": {
            "provider": "redis",
            "key": "user",
            "ttl": 7
      }
}
~~~~

## TESTS

Work in progress

~~~~
go test ./...
~~~~

## TODO

Unit tests
Implement connectors for memcache, redis, mongo and file
 