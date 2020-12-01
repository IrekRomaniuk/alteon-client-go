# alteon-client-go

Client for Radware Alteon VA (Standalone) API, tested with version 32.6.1.0. To be used with terraform-provider-alteon

```
go mod init alteon-client-go
go mod vendor
go build -o alteon-client-go 
rm alteon-client-go
```

### Env variables

The following env variables should be set and exported (IP address is an example only):

```
export ALTEON_USERNAME=admin
export ALTEON_PASSWORD=
export ALTEON_URI=https://13.72.75.201:8443/config
```

### Test requires real alteon server to be reachable

```
go test -run TestNewClient
go test -v -run TestNewClient
go test -cover
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