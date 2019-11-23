# catharsis

## what is this

```
the boilerplate for monorepo application
```

- Base project is https://github.com/golang-standards/project-layout

## development

### Preparation

- generate rsa pem file

```bash
> openssl genrsa -out ./secret/catharsis.rsa 1024
> openssl rsa -in ./secret/catharsis.rsa  -pubout > ./secret/catharsis.rsa.pub
```

- environment (using direnv)

```bash
> cp .envrc.tmpl .envrc
```

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

## structure

```bash
.
├── LICENSE.md
├── Makefile
├── README.md
├── authentication
│   ├── Makefile
│   ├── domain
│   │   ├── model
│   │   │   └── user.go
│   │   ├── repository
│   │   │   ├── mock
│   │   │   │   └── user.go
│   │   │   └── user.go
│   │   └── service
│   │       ├── authentication.go
│   │       └── mock
│   │           └── authentication.go
│   ├── handler
│   │   └── authentication.go
│   ├── infrastructure
│   │   ├── internal
│   │   │   └── entity
│   │   │       ├── base.go
│   │   │       └── user.go
│   │   ├── repository
│   │   │   └── user_mock.go
│   │   └── service
│   │       ├── authentication_impl.go
│   │       └── test
│   │           └── authentication_impl_test.go
│   └── intercepter
│       └── authentication.go
├── cmd
│   ├── README.md
│   └── authentication
│       ├── authentication-server
│       ├── dependency.go
│       ├── env.go
│       └── main.go
├── docker
│   ├── development
│   │   └── Dockerfile
│   └── production
│       └── authentication
│           └── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
├── pkg
│   ├── README.md
│   ├── jwtauth
│   │   ├── config.go
│   │   ├── context.go
│   │   ├── mock
│   │   │   ├── sign_service.go
│   │   │   └── verify_service.go
│   │   ├── model.go
│   │   ├── sign_client.go
│   │   ├── sign_service.go
│   │   ├── sign_service_impl.go
│   │   ├── verify_client.go
│   │   ├── verify_service.go
│   │   └── verify_service_impl.go
│   └── log
│       ├── config.go
│       └── logger.go
├── proto
│   └── authentication
│       ├── authentication.pb.go
│       └── authentication.proto
└── secret
    ├── catharsis.rsa
    └── catharsis.rsa.pub
```
