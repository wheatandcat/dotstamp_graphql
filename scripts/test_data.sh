mysql -h 127.0.0.1 -P 3306 -u root -N -e 'show tables' | xargs -IARG mysql -h 127.0.0.1 -P 3306 -u root -e 'truncate table ARG' stamp_test
mysql -h 127.0.0.1 -P 3306 -u root stamp_test < ./scripts/dump.sql
