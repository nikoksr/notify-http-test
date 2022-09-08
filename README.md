### Description

This is a super basic HTTP server that offers a single POST endpoint at `/` and is intended to be used as a testing
webhook for Notify's coming generic http service.

### Usage

The server will listen on `:8080` by default. You can change this by setting the `NOTIFY_TEST_ADDR` environment
variable. Now, you can send a POST request to the server and see the request body printed to the console.

```bash
go run main.go
```

### Client examples

> Supported content types are `application/json`, `application/x-www-form-urlencoded` and `text/plain`. The following examples use curl for demonstration purposes, but of course the usage translates to Notify's generic http service.

#### JSON

```bash
curl http://localhost:8080 \
  -X POST \
  -H 'Content-Type: application/json' \
  -d '{"subject":"Test", "message":"Hello, world!"}'
```

#### Form

```bash
curl http://localhost:8080 \
  -X POST \
  -H 'Content-Type: application/x-www-form-urlencoded' \
  -d 'subject=Test&message=Hello, world!'
```

#### Plain text

```bash
curl http://localhost:8080 \
  -X POST \
  -H 'Content-Type: text/plain' \
  -d 'Test - Hello, world!'
```