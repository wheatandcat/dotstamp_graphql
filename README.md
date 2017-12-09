# dotstamp_graphql

## About
https://dotstamp.com/

## setup

### install
dep ensure

### start
docker-compose up -d

### DB
mysql -h 0.0.0.0 -P 3306 -u root -e 'CREATE DATABASE stamp_test;'
goose -env test up

### stop
docker-compose stop
