# Servers Health Checker

This is my **first attempt at learning Golang**. I'm building this small project
to get hands-on experience with Go fundamentals and concurrent programming.

## What it does

The project is a simple HTTP endpoint health checker. It reads a list of
endpoints from a JSON file and checks them concurrently using goroutines. It
also explores using a `sync.WaitGroup`, channels, and an HTTP client with a
timeout.

This is a learning project, so the code and features may change as I learn more
about Go.

## Configuration

The JSON file should contain an `endpoints` array:

```json
{
    "endpoints": [
        "https://httpbin.org/status/200",
        "https://api.github.com",
        "https://www.cloudflare.com/cdn-cgi/trace"
    ]
}
```

## Run it

Make sure Go is installed, then run the project from its root directory:

```bash
go run . -file server-endpoints.json -type json
```

The `-file` flag sets the path to the JSON configuration file. JSON is currently
supported; YAML support is not implemented yet.

## What I'm learning

- Go syntax, functions, and structs
- Reading and parsing JSON files
- Making HTTP requests
- Goroutines and concurrent execution
- `sync.WaitGroup` and channels
- Handling errors and measuring request duration

## Why I made this

I want to learn Go by building practical projects instead of only reading
tutorials. This is my starting point, and I'll keep improving the project as I
become more comfortable with the language.
