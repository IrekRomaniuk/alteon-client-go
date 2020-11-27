# alteon-client-go

```
go mod init alteon-client-go
go mod vendor
go build -o alteon-client-go 
rm alteon-client-go
```

### Test requires real alteon server to be reachable

```
go test -v
=== RUN   TestCreateItem
Server VTJXGEFNHO created: ok
--- PASS: TestCreateItem (1.37s)
=== RUN   TestGetItem
--- PASS: TestGetItem (0.14s)
=== RUN   TestUpdateItem
Server VTJXGEFNHO updated: ok
--- PASS: TestUpdateItem (1.13s)
=== RUN   TestDeleteItem
Server VTJXGEFNHO deleted
--- PASS: TestDeleteItem (0.11s)
PASS
ok      github.com/irekromaniuk/alteon-client-go        2.748s
```