# go-yml-etl
Write and run ETL rules in YML files


## Demo

### Create database

```
createdb web_app_development
```

### Create schema in the database

```
psql -d web_app_development -f psql_example.sql
```

### Build and run code

```
go build
./go-yml-etl
```

OR

```
go run main.go
```
