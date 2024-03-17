# go-wsdl

A website downloader written in Go. By default it will download all resources on http://books.toscrape.com to a directory named `output`, however this is configurable with flags.

# Usage

The easiest way to get started is by downloading the pre-built binaries for your system. Alternatively you can build binaries from source or run the code with on-the-fly compilation using `go run`.

## Pre-built binaries

### 1. Choose and download a pre-built binary for your OS and architecture:

| OS      | Arch  | Binary                                                                                                  |
|---------|-------|---------------------------------------------------------------------------------------------------------|
| Linux   | 386   | [go-wsdl-linux-386](https://github.com/straski/go-wsdl/releases/download/0.1.1/go-wsdl-linux-386)       |
| Linux   | amd64 | [go-wsdl-linux-amd64](https://github.com/straski/go-wsdl/releases/download/0.1.1/go-wsdl-linux-amd64)   |
| MacOS   | arm64 | [go-wsdl-darwin-arm64](https://github.com/straski/go-wsdl/releases/download/0.1.1/go-wsdl-darwin-arm64) |
| MacOS   | amd64 | [go-wsdl-darwin-amd64](https://github.com/straski/go-wsdl/releases/download/0.1.1/go-wsdl-darwin-amd64) |
| Windows | 386   | [go-wsdl-win32.exe](https://github.com/straski/go-wsdl/releases/download/0.1.1/go-wsdl-win32.exe)       |
| Windows | amd64 | [go-wsdl-win64.exe](https://github.com/straski/go-wsdl/releases/download/0.1.1/go-wsdl-win64.exe)       |

### 2. Run the command

> If required add execution permissions to the file: ```chmod +x go-wsdl-<OS>-<ARCH>```

In your terminal run the following command:

```./go-wsdl-<OS>-<ARCH>```

It takes two optional arguments:

| Name | Shorthand | Description                             | Default                              |
|------|-----------|-----------------------------------------|--------------------------------------|
| url  | u         | The URL to download the files from.     | http://books.toscrape.com/index.html |
| dir  | d         | The directory to download the files to. | output                               |

Usage with arguments:

```./go-wsdl-<OS>-<ARCH> --url <URL> --dir <DIR>```

### Windows

```go-wsdl-<OS>-<ARCH>.exe```

> If you use [MinGW](https://www.mingw-w64.org) (comes with [Git for Windows](https://git-scm.com/download/win)), run the commands as if you're on a Unix-based system. 

## Build from source

To build your binaries from source, use the provided Makefile:

```./make build```

This will build binaries for all supported architectures listed above. The binaries are stored in the `bin` dir.

## Run with ``go run``

To execute the source code on-the-fly:

```
go mod vendor
go mod verify
go run main.go 
```

# Tests

To run tests, execute:

```make test```
