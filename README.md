# Shortly

### About
A URL shortener service written in Go. The application has a PostgreSQL as backend store.

### Install
```
$ go get github.com/wang502/shortly
```

### Config PostgreSQL Host
- configure PostgreSQL url as environment variable
```
  export PG_URL="example url"
```

### Run
```
$ cd shortly
```

```
$ go build
```

```
$ ./shortly
```

### Test
- Post a URL you want to shorten
```
$ curl -X POST url=pinterest.com
```

- Open the browser, access the shortened URL
