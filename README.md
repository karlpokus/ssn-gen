# ssn-gen
Generate safe swedish personnummer in go. By safe we mean personnummer with the last 4 digits starting with 99|98 which are not in use.

# usage
docker
```bash
# builds an image tagged by `cat version`
$ ./build.sh
# run above version
$ ./run.sh
```

native
```bash
# start server
$ go run ./cmd/main
```

request n ssns
```bash
GET http://localhost:9012/ssn/<n>
```

# tests
```bash
$ go test
```

# todos
- [x] safe ssn
- [x] http server
- [ ] opt to set age
- [x] tests
- [x] rate limiting
- [ ] don't run as [root](https://medium.com/@chemidy/create-the-smallest-and-secured-golang-docker-image-based-on-scratch-4752223b7324)
- [ ] add [valid opt](https://ssn-machine.herokuapp.com/valid?ssn=<ssn>)

# license
MIT
