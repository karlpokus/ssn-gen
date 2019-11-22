# ssn-gen
Generate safe swedish personnummer in go. By safe we mean personnummer with the last 4 digits starting with 99|98 which are not in use.

# Usage
```bash
# start server
$ go run ./cmd/main
# make request for n ssns
GET /ssn/<n>
```

# todos
- [x] safe ssn
- [x] http server
- [ ] opt to set age
- [ ] tests

# license
MIT
