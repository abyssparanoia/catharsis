# catharsis

## what is this

```
the boilerplate for monorepo application
```

- Base project is https://github.com/golang-standards/project-layout

## development

### authentication server

#### local

```bash
> realize start
```

#### docker

```bash
# build image
> docker-compose build

# container start
> docker-compose up -d
```

#### sign in grpcurl sample

```bash
grpcurl -plaintext -d '{"user_id": "user_id", "password":"password"}' 127.0.0.1:50051 Authentication/SignIn
```

## production

### build

```bash
> docker build -f ./docker/production/authentication/Dockerfile .
```

### start

```bash
> docker container run -it 4069a0d5dfac -p 50051:50051
```
