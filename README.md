# api
Criar container no docker:<br>
docker run --name mysql -e MYSQL_ROOT_PASSWORD=docker -p 3306:3306 -d mysql

Incializar as migrations

Iniciar o serviço
no diretório ./api/apihg/
go run main.go

o serviço estará escutando na porta :8080
