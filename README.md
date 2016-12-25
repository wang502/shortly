# Shortly

### About
A URL shortener service written in Go using net/http. The application uses PostgreSQL as backend store.

### Install

```
$ go get github.com/wang502/shortly
```

### Config PostgreSQL Host
- setting PostgreSQL url as environment variable by adding the following command in ***~/.bash_profile***
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
$ curl -X POST http://localhost:8080/shorten?url=www.pinterest.com
```

- Open the browser, access the shortened URL
