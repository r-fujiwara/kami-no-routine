```
$ go run server.go
```

- `request` request

```
$ for i in {1..4096}; do curl localhost:8000/work -d name=$USER -d delay=$(expr $i % 11)s; done
```

```
$ curl localhost:8000/work -d name=$USER
```
