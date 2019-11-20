# catharsis

## what is this

```
the boilerplate for monorepo application
```

- Base project is https://github.com/golang-standards/project-layout

## development

### authentication server

```bash
> go run ./cmd/authentication/main,go
```

### sign in grpcurl sample

```bash
grpcurl -plaintext -d '{"user_id": "tester", "password":"hogehoge"}' 127.0.0.1:50051 Authentication/SignIn
```
