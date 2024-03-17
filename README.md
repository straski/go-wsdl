# go-wsdl

A website downloader written in Go. By default it will download all resources on http://books.toscrape.com to a directory named `output`. 

## Usage

### Pre-built binaries

Choose and download a pre-built binary for your OS/arch:

| OS      | Arch  | Binary                          |
|---------|-------|---------------------------------|
| Linux   | 386   | [go-wsdl-linux-386](url)        |
| Linux   | amd64 | [go-wsdl-linux-amd64](url)  |
| MacOS   | arm64 | [go-wsdl-darwin-arm64](url) |
| MacOS   | amd64 | [go-wsdl-darwin-amd64](url) |
| Windows | 386   | [go-wsdl-win32.exe](url)    |
| Windows | amd64 | [go-wsdl-win64.exe](url)    |

In your terminal run the following command:

```./bin/go-wsdl-<OS>-<ARCH>```

If required add execution permission to the file:

```chmod +x bin/go-wsdl-<OS>-<ARCH>```

It takes two optional arguments:

- `url` (`u`) - the URL of the page to download (default is http://)
- `dir` (`u`) - the directory to save files to (default is `output`)

### Build from source

To build your binaries from source, use the provided Makefile:

```./make build```

This will build binaries for all supported architectures listed above. The binaries are stored in the `bin` dir.

### Run without build

To run the tool without building it:

```
go mod vendor
go mod verify
go run main.go 
```
