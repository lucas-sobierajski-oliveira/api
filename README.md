# api - GoLang
Projeto desenvolvido usando GoLang<br>

Criar container no docker:<br>
docker run --name mysql -e MYSQL_ROOT_PASSWORD=docker -p 3306:3306 -d mysql

Incializar as migrations para criar as tabelas e popular o banco de dados.

Iniciar o serviço<br>
no diretório ./api/apihg/<br>
go run main.go

o serviço estará escutando na porta :8080
