# soporte_crud
 
El API provee la gestión de los diferentes procesos que requiere el módulo de soporte, permitiendo la administración de categorías de preguntas y preguntas frecuentes del sistema.
 
## Especificaciones Técnicas
 
### Tecnologías Implementadas y Versiones
 
* __[Golang](https://go.dev/doc/install)__
* __[Gorilla Mux](https://github.com/gorilla/mux)__
* __[PostgreSQL](https://www.postgresql.org/)__
* __[PgAdmin 4](https://www.pgadmin.org/)__
* __[Postman](https://www.postman.com/)__
### Variables de Entorno
 
```
SOPORTE_CRUD_PGDB=[nombre de la base de datos]
SOPORTE_CRUD_PGPASS=[password del usuario]
SOPORTE_CRUD_PGURLS=[direccion de la base de datos]
SOPORTE_CRUD_PGPORT=[Puerto de conexión con la base de datos]
SOPORTE_CRUD_PGUSER=[usuario con acceso a la base de datos]
SOPORTE_CRUD_PGSCHEMA=[esquema donde se ubican las tablas]
SOPORTE_CRUD_HTTP_PORT=[puerto de ejecucion]
```
 
NOTA: Las variables se pueden ver en el fichero `config/db.go` y están identificadas con `SOPORTE_CRUD_...`
 
### Ejecución del Proyecto
 
```
#1. Obtener el repositorio con Go
go get github.com/tu_usuario/soporte_crud
 
#2. Moverse a la carpeta del repositorio
cd $GOPATH/src/github.com/tu_usuario/soporte_crud
 
#3. Moverse a la rama develop
git pull origin develop && git checkout develop
 
#4. Alimentar todas las variables de entorno que utiliza el proyecto.
SOPORTE_CRUD_HTTP_PORT=8082 SOPORTE_CRUD_PGHOST=127.0.0.1 SOPORTE_CRUD_PGPORT=5432 SOPORTE_CRUD_PGUSER=postgres SOPORTE_CRUD_PGPASS=postgres SOPORTE_CRUD_PGDB=AgroCampo SOPORTE_CRUD_PGSCHEMA=soporte go run main.go
```
