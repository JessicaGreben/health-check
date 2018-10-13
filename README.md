# health-check

---

health-check is a service can be launched that reports the overall health of a cluster.

## Getting Started

### CLI Prerequisites

* Golang >= v1.11 [installed](https://golang.org/dl/)

* Your $PATH configured:

```
$ export PATH=$GOROOT/bin:$GOPATH/bin:$PATH
```

### Download and run

In order to use the command line, compile it using the following command:

```
$ go get github.com/jessicagreben/health-check
```

Enable Go [Modules](https://github.com/golang/go/wiki/Modules):

```
$ export GO111MODULE=on
```

Build the program with Go modules:

```
go build github.com/jessicagreben/health-check
```

or build the program with [vendor directory](https://github.com/golang/go/wiki/Modules#how-do-i-use-vendoring-with-modules-is-vendoring-going-away):


```
go build -mod=vendor github.com/jessicagreben/health-check
```


### `health-check` commands:

To view details and available commands:

```
$ health-check help
```

## Running the tests

To run tests in root package:

```
go test
```

To run tests in sub-packages:

```
go test -v ./...

```

To test with code coverage:

```
go test -cover
```

## Deployment

WIP

## Built With

* [Cobra](https://github.com/spf13/cobra) - A Commander for modern Go CLI interactions
* [Kubernetes client-go](https://github.com/kubernetes/client-go) - Go client for Kubernetes.

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md)for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [releases for this repository](https://github.com/JessicaGreben/health-check/releases). 

## Authors

See the list of [contributors](https://github.com/JessicaGreben/health-check/graphs/contributors) who have participated in this project.

## License

This project is licensed under the MIT License.
