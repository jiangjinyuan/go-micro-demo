# Greeter Service

This is the Greeter service

Generated with

```
micro new github.com/jiangjinyuan/go-micro-demo/poolHashrate/srv --namespace=go.micro --alias=greeter --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.greeter
- Type: srv
- Alias: greeter

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./greeter-srv
```

Build a docker image
```
make docker
```