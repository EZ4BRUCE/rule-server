# rule-server

 docker run --name rule-server-mysql -p 3307:3306 -v rule-data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=123456 -d mysql

 swagger的注释写入那里，为了方便说明，用了多个body，只有最后一个测试body会起作用

 $Env:GOOS = "linux"; 
$Env:GOARCH = "amd64"
go build
docker-compose up --build