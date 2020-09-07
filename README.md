# prox

This is a work in progress and not ready for production yet.
If you really want to speed things up, please consider contributing.

The idea was to create an easy to deploy application capable to serve semi-static pages really, really fast.
Prox wants to be a tool capable of mitigating 
    - really high load on landing pages during expected traffic spikes (Black Friday or similar campaigns).
    - high load on mostly static content sites (wordpress, magento, etc.) 

* easy to deploy
    - prox is a single binary application
    - prox relies on only a JSON route configuration file and one .env file for application config
    
* semi-static
    - every N seconds (configurable in routes.json), prox loads JSON data from a fast source (memcache, redis, etc.) and populates an HTML template with it
    - prox serves the generated HTML content (template + cache data) as a normal server would
    - client can receive data that is at most N seconds old (N is configurable per each defined endpoint) 
    - prox does not know or care who writes the data, the only convention is that data is in JSON format and matches with the provided template 

* really, really fast
    - with logs disabled, ab reports
~~~~
 ab -n 1000000 -c 100 http://localhost:50001/

Document Path:          /
Document Length:        56 bytes

Concurrency Level:      100
Time taken for tests:   40.941 seconds
Complete requests:      1000000
Failed requests:        0
Total transferred:      195000000 bytes
HTML transferred:       56000000 bytes
Requests per second:    24425.38 [#/sec] (mean)
Time per request:       4.094 [ms] (mean)
Time per request:       0.041 [ms] (mean, across all concurrent requests)
Transfer rate:          4651.32 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    1   0.5      1       7
Processing:     0    3   1.6      3      29
Waiting:        0    3   1.5      3      27
Total:          0    4   1.5      4      29

Percentage of the requests served within a certain time (ms)
  50%      4
  66%      4
  75%      5
  80%      5
  90%      6
  95%      7
  98%      8
  99%      9
 100%     29 (longest request)

~~~~     

### Contributors Stuff

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


## Providers

Providers are defined in the routes.json file and registered upon application start

Here is an example for a provider JSON structure:

```
{
    "name": "memcache",
    "connection_string": "memcached:11211",
    "user": "",
    "password": "",
    "collection": ""
}
```

## TESTS

Work in progress

~~~~
go test ./...
~~~~

## TODO

Unit tests
* Implement connectors for mongoDB, url(CDN) and file
* Implement templating system (support multiple template formats)
* Implement response codes other than 200 
* Implement routes with parameters
* Implement authentication support for each endpoint
* Remove dependency on .env file (replace with cli params)