# dotstamp_graphql

## About
https://dotstamp.com/

## setup

### install
```sh
dep ensure
```

### docker build
```sh
docker build --tag="dotstamp_graphql:latest" .
```
### start
```sh
docker-compose up -d
```

### DB
```sh
mysql -h 0.0.0.0 -P 3306 -u root -e 'CREATE DATABASE stamp_test;'
goose -env test up
mysql -h 0.0.0.0 -P 3306 -u root
```

### test data
```sh
sh scripts/test_data.sh
```

### stop
```sh
docker-compose stop
```

### development
```sh
go get github.com/pilu/fresh
fresh
```
