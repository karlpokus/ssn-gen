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

Benchmarks for the `add-pipeline` branch

```bash
$ /usr/bin/time -lp ./ssn-gen-pipe --ssns=1000000 > file && wc -l file

# no data in mem
real         4.03
user         2.89
sys          3.69
7958528  maximum resident set size # ~7.9 MB
...
1000000 file

# all data in mem
real         1.84
user         2.31
sys          0.85
92061696  maximum resident set size # ~92 MB
...
999999 file
```

# license
MIT
