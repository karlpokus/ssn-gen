# ssn-gen
Generate safe swedish personnummer in go. By safe we mean personnummer with the last 4 digits starting with 99|98 which are not in use.

# Usage
Writes n safe line-separated ssns to stdout

```bash
$ go run ./cmd/main [--ssns int]
```

# Benchmarks
1M ssns. A [nodeJS implementation](https://github.com/karlpokus/pernr/tree/fix/no-limit) for comparison.

```bash
$ /usr/bin/time -lp ./ssn-gen --ssns=1000000 > file
real         0.81
user         1.01
sys          0.07
 114921472  maximum resident set size # ~115 MB
```

note: see other branches in this repo for other implementations and benchmarks

# todos
- [x] safe ssn
- [ ] opt to set age
- [ ] tests

# license
MIT
