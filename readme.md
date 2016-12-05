# Brizo

[![wercker status](https://app.wercker.com/status/71f5e5d09990c22f394a998bef86d9af/s/master "wercker status")](https://app.wercker.com/project/byKey/71f5e5d09990c22f394a998bef86d9af)
[![codecov](https://codecov.io/gh/generationtux/brizo/branch/master/graph/badge.svg)](https://codecov.io/gh/generationtux/brizo)
[![Code Climate](https://codeclimate.com/github/generationtux/brizo/badges/gpa.svg)](https://codeclimate.com/github/generationtux/brizo)
[![Go Report Card](https://goreportcard.com/badge/github.com/generationtux/brizo)](https://goreportcard.com/report/github.com/generationtux/brizo)

PAAS that runs on top of [Kubernetes](http://kubernetes.io/).

## Development

You can configure the app for local development using a `.env` file in the root of the project. You should copy `.env.example` to `.env` to get started.
The example file will have all of the required config with sensible defaults.

#### Local database

A test database is provided using Docker. The `.env.example` file provided is configured to connect to this database by default. To start the database [install Docker] for 
your machine and run
```sh
$ docker-compose up -d
```

This will start a database listening on `localhost:33062` with the username `root` and password `secret`. It will also setup a database named `brizo`.

#### Autoreload

You can use [Fresh](https://github.com/pilu/fresh) to serve and autoreload the service during development.

Install Fresh
```sh
$ go get github.com/pilu/fresh
```

Then run the app
```sh
$ fresh
```

#### Dependency management

Dependencies are managed using [Govendor](https://github.com/kardianos/govendor). For simpliciy in version management
dependencies are committed with the repo.
