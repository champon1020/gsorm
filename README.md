# mgorm

## Test

### Run test
```
go test -v ./...
```

### Run test with coverage profile
```
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```
