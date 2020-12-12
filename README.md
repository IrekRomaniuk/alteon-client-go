# alteon-client-go

Client for Radware Alteon VA (Standalone) API, tested with version 32.6.1.0. To be used with terraform-provider-alteon

### Env variables

The following env variables should be set and exported (IP address is an example only):

```
export ALTEON_USERNAME=admin
export ALTEON_PASSWORD=
export ALTEON_URI=https://13.72.75.201:8443/config
```

### Test with real alteon server to be reachable

`onnline` flag to be added. if not mock server is to be used

```
go test -run TestNewClient -online
go test -v -run TestNewClient -online
go test -cover -online
go test -v -online
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

### Install if necessary

```
go mod init alteon-client-go
go mod vendor
go build -o alteon-client-go 
rm alteon-client-go
```