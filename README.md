# netscanner

Fast, light TCP port scanner written in Go using goroutines for concurrent scanning.

## Features

* Concurrent TCP port scanning
* Range parsing support (e.g. `1-100`)
* Service mapping for common ports

## Requirements

* Go 1.27.1 

## Installation

```bash
git clone [https://github.com/Resul4774/netscanner.git](https://github.com/Resul4774/netscanner.git)
cd netscanner
go build -o netscanner main.go