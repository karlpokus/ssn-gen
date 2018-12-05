# ssn-gen
Swedish personnummer generator in go

# Usage
Writes n valid line-separated ssns to stdout

```bash
$ ./ssn-gen --ssns=n > file
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

# license
MIT
