# catharsis

## what is this

```
the boilerplate for monorepo application
```

- Base project is https://github.com/golang-standards/project-layout

## Apps

| Package                                           | Localhost              | Prodction         |
| :------------------------------------------------ | :--------------------- | :---------------- |
| **[[REST] bff](./cmd/bff)**                       | http://localhost:8081  | bff.\*            |
| **[[gRPC] authentication](./cmd/authentication)** | http://localhost:50051 | authentication.\* |

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

### server starting

- local

```bash
> realize start
```

- docker

```bash
# build image
> docker-compose build

# container start
> docker-compose up -d
```

- example of sign in grpcurl

```bash
> grpcurl -plaintext -d '{"user_id": "user_id", "password":"password"}' 127.0.0.1:50051 Authentication/SignIn
```

- example of create user in grpcurl

```bash
> grpcurl -plaintext -d '{"password":"password","display_name":"ゲストさn","icon_image_path":"icon_url","background_image_path":"background_url","profile":"text"}' localhost:50051 Authentication/CreateUser
```

- example of bff server

```bash
> curl --request POST 'http://localhost:8081/v1/me/sign_in' -d '{"user_id": "user_id", "password":"password"}'
```

### database

- generate server code by sql boiler

```bash
> sqlboiler -c ./db/authentication/sqlboiler.toml -o ./pkg/dbmodels/authentication psql
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
│   │   └── repository
│   │       ├── mock
│   │       │   └── user.go
│   │       └── user.go
│   ├── handler
│   │   └── authentication.go
│   ├── infrastructure
│   │   ├── internal
│   │   │   └── entity
│   │   │       ├── base.go
│   │   │       └── user.go
│   │   └── repository
│   │       └── user_mock.go
│   ├── intercepter
│   │   └── authentication.go
│   └── service
│       ├── authentication.go
│       ├── authentication_impl.go
│       ├── mock
│       │   └── authentication.go
│       └── test
│           └── authentication_impl_test.go
├── bff
│   ├── domain
│   │   ├── model
│   │   └── repository
│   │       └── authentication.go
│   ├── handler
│   │   ├── authentication.go
│   │   └── ping.go
│   ├── infrastructure
│   │   └── repository
│   │       └── authentication_impl.go
│   ├── middleware
│   │   └── access_control.go
│   └── service
│       ├── authentication.go
│       └── authentication_impl.go
├── cmd
│   ├── README.md
│   ├── authentication
│   │   ├── authentication-server
│   │   ├── dependency.go
│   │   ├── env.go
│   │   └── main.go
│   └── bff
│       ├── bff-server
│       ├── dependency.go
│       ├── env.go
│       ├── main.go
│       └── routing.go
├── docker
│   ├── development
│   │   └── Dockerfile
│   └── production
│       └── authentication
│           └── Dockerfile
├── docker-compose.yml
├── docs
├── go.mod
├── go.sum
├── pkg
│   ├── README.md
│   ├── errcode
│   │   ├── error.go
│   │   └── model.go
│   ├── httpaccesslog
│   │   └── middleware.go
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
│   ├── log
│   │   ├── config.go
│   │   └── logger.go
│   ├── parameter
│   │   ├── json.go
│   │   └── url.go
│   └── renderer
│       ├── handler.go
│       └── model.go
├── proto
│   ├── authentication.pb.go
│   └── authentication.proto
└── secret
    ├── catharsis.rsa
    └── catharsis.rsa.pub

```
