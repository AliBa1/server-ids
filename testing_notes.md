### Run all tests from root dir

```zsh
go test ./...
```

### Create and show coverage profile

```zsh
go test ./... -coverprofile cover.out
```

### Show coverage for each function (in terminal)

```zsh
go tool cover -func cover.out
```

### Show coverage for each function (in html)

```zsh
go tool cover -html cover.out
```

### Create coverage profile on build

```zsh
go build -cover
```

### Create coverage profile (specfic pkg)

```zsh
go test -coverprofile=... <pkg_target>
```
