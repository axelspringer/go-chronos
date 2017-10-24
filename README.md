[![Taylor Swift](https://img.shields.io/badge/secured%20by-taylor%20swift-brightgreen.svg)](https://twitter.com/SwiftOnSecurity)
[![Volkswagen](https://auchenberg.github.io/volkswagen/volkswargen_ci.svg?v=1)](https://github.com/auchenberg/volkswagen)
[![Build Status](https://travis-ci.org/axelspringer/go-chronos.svg?branch=master)](https://travis-ci.org/axelspringer/go-chronos)
[![MIT license](http://img.shields.io/badge/license-MIT-brightgreen.svg)](http://opensource.org/licenses/MIT)

# go-chronos

go-chronos is a Go library for the [Chronos REST API](https://mesos.github.io/chronos/docs/api.html). Check the [examples](https://github.com/dghubble/go-twitter/blob/master/examples) to see how to access the Chronos API.

## Features

* Creating Jobs
* Deleting Jobs
* Search Jobs
* Kill Tasks
* List Jobs
* Marking a Job as successful

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
[MIT](/LICENSE)
