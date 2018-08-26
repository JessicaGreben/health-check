# pratique

pratique /prah-teak/, _noun_

definiton: permission granted to a ship to have dealings with a port, given after quarantine or on showing a clean bill of health.

---

Pratique is a service can be launched that reports the overall health of a cluster.

## Getting Started

### CLI Prerequisites

* Golang installed

* Your $PATH configured:

```
$ export PATH=$GOROOT/bin:$GOPATH/bin:$PATH
```

### Download and run

In order to use the command line, compile it using the following command:

```
$ go get github.com/jessicagreben/pratique
```

Build and install the program:

```
go install github.com/jessicagreben/pratique
```

### `pratique` commands:

To view details and available commands:

```
$ pratique help
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

Please read [CONTRIBUTING.md](https://github.com/JessicaGreben/pratique/blob/master/CONTRIBUTING.md)for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [releases for this repository](https://github.com/JessicaGreben/pratique/releases). 

## Authors

See the list of [contributors](https://github.com/JessicaGreben/pratique/graphs/contributors) who have participated in this project.

## License

This project is licensed under the MIT License.
