.PHONY: mysqlup
mysqlup:
	docker run --name gofr-mysql -e MYSQL_ROOT_PASSWORD=root123 \
		-e MYSQL_DATABASE=test_db -p 3306:3306 -d mysql:8.0.30

.PHONY: mysqldown
mysqldown:
	docker stop gofr-mysql
	docker rm -f gofr-mysql

.PHONY: migrate
migrate:
	docker exec -it gofr-mysql mysql -uroot -proot123 test_db \
		-e "CREATE TABLE customers (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255) NOT NULL);"

.PHONY: run
run:
	go run .
