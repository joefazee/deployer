# How to deploy your Go application using Github Actions

This is example program that we use to show how we can do automated deployment with Github actions

## App requirements

- Postgres
- Redis
- migrate (github.com/golang-migrate/migrate)

## How to run

Running the program requires go 1.20. Clone the program and run the following commands

#### 1. Run docker-compose

```
docker compose up -d
```

#### 2. Run migrations

```
make migrations/up
```

#### 3. Run the program

```
make run
```

You can open the `Makefile` to see all the commands
