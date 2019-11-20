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
grpcurl -plaintext -d '{"user_id": "user_id", "password":"password"}' 127.0.0.1:50051 Authentication/SignIn
```
