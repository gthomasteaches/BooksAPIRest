# Simple REST API en Go

Esta es una aplicación sencilla escrita en Go que expone una API REST para gestionar libros, autores y categorías. La información se mantiene en memoria, sin uso de bases de datos, y está construida solamente con la biblioteca estándar de Go (`net/http`).

## Características

* CRUD de libros
* Registro y consulta de autores
* Registro y consulta de categorías
* Datos almacenados en memoria (sin persistencia)

---

## Requisitos

* Go 1.18 o superior

---

## Ejecución local

1. Clona este repositorio:

   ```bash
   git clone https://github.com/tu_usuario/tu_repo.git
   cd tu_repo
   ```

2. Ejecuta el servidor:

   ```bash
   go run main.go
   ```

3. Accede a los siguientes endpoints desde Postman, curl o navegador:

   * `GET http://localhost:8080/books`
   * `POST http://localhost:8080/books`
   * `PUT http://localhost:8080/books/{id}`
   * `DELETE http://localhost:8080/books/{id}`
   * `GET http://localhost:8080/authors`
   * `POST http://localhost:8080/authors`
   * `GET http://localhost:8080/categories`
   * `POST http://localhost:8080/categories`

---

## Compilación para otros sistemas operativos

### macOS (actual)

```bash
go build -o app
```

### Windows (64 bits)

```bash
GOOS=windows GOARCH=amd64 go build -o app.exe
```

### Linux (64 bits)

```bash
GOOS=linux GOARCH=amd64 go build -o app
```

> Puedes usar `GOARCH=386` para sistemas de 32 bits si lo necesitas.

---

## Pruebas con Postman

Incluye una colección de Postman para probar todos los endpoints de la API. Puedes importarla directamente desde el archivo:

* [`postman_collection.json`](./postman_collection.json)

---

## Licencia

MIT
