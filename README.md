[![Build Status](https://travis-ci.org/axelspringer/go-chronos.svg?branch=master)](https://travis-ci.org/axelspringer/go-chronos)
[![Go Report Card](https://goreportcard.com/badge/github.com/axelspringer/go-chronos)](https://goreportcard.com/report/github.com/axelspringer/go-chronos)
[![Taylor Swift](https://img.shields.io/badge/secured%20by-taylor%20swift-brightgreen.svg)](https://twitter.com/SwiftOnSecurity)
[![Volkswagen](https://auchenberg.github.io/volkswagen/volkswargen_ci.svg?v=1)](https://github.com/auchenberg/volkswagen)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

# go-chronos

go-chronos is a Go library for the [Chronos REST API](https://mesos.github.io/chronos/docs/api.html). Check the [examples](https://github.com/dghubble/go-twitter/blob/master/examples) to see how to access the Chronos API.

## Features

* Creating Jobs
* Deleting Jobs
* Search Jobs
* Kill Tasks
* List Jobs
* Marking a Job as successful

## Documentation

You can find the documentation hosted on [godoc.org](https://godoc.org/github.com/axelspringer/go-chronos).

## Install

```
go get github.com/axelspringer/go-chronos
```

## Examples

You find some examples in the `examples` folder.

## Development

Install neat tools and dependencies.

```
make deps && make restore
```

## License
[Apache 2.0](/LICENSE)
