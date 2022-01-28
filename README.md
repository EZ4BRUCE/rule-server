# rule-server

 docker run --name rule-server-mysql -p 3307:3306 -v mysql-data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=123456 -d mysql