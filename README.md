# api
Criar container no docker<br>
docker run --name mysql -e MYSQL_ROOT_PASSWORD=docker -p 3306:3306 -d mysql

Incializar as migrations (Não soube fazer com golang mas deixei uns arquivo com os comandos sql utilizados: api/apihg/mirations/...)

Iniciar o serviço
no diretório ./api/apihg/
go run main.go
