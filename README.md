# dotstamp_graphql

## About
https://dotstamp.com/

## setup

### install
dep ensure

### docker build
docker build --tag="dotstamp_graphql:latest" .

### start
docker-compose up -d

### DB
mysql -h 0.0.0.0 -P 3306 -u root -e 'CREATE DATABASE stamp_test;'
goose -env test up

mysql -h 0.0.0.0 -P 3306 -u root

### test data
sh scripts/test_data.sh

### stop
docker-compose stop
