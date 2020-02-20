# PoolHashrate Service

This is the PoolHashrate service

Generated with

```
micro new github.com/jiangjinyuan/go-micro-demo/poolHashrate-web --namespace=btccom.explorer --alias=poolHashrate --type=web
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: btccom.explorer.web.poolHashrate
- Type: web
- Alias: poolHashrate

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
./poolHashrate-web
```

Build a docker image
```
make docker
```