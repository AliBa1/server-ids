### Run all tests from root dir

```bash
go test ./...
```

### Create and show coverage profile

```bash
go test ./... -coverprofile cover.out
```

### Show coverage for each function (in terminal)

```bash
go tool cover -func cover.out
```

### Show coverage for each function (in html)

```bash
go tool cover -html cover.out
```

### Create coverage profile on build

```bash
go build -cover
```

### Create coverage profile (specfic pkg)

```bash
go test -coverprofile=... <pkg_target>
```
