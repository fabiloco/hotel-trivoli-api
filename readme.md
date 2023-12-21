# Hotel Trivoli API

Hotel Trivoli project API

### Configuration

Create a .env file in the root folder and filled the variables with your owns or ask for them.

#### Setup database

Make sure you have docker and docker-compose installed on your system. Run the following command on the root folder to setup the mariadb database.

```
docker compose up -d
```

#### Runing the app

You can run the app in dev mode using the following command:

```
go run main.go
```

You can also install [air](https://github.com/cosmtrek/air) to live-reload the app during dev time.

```
air
```

To build the app, run

```
go build main.go
./main
```
