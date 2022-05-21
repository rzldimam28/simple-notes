<h1>Simple Notes</h1>
<p>a simple RESTful API for crud database operation</p>

## 🛠 Tech Stack 

- Go
- MySQL
- Gorillamux

## 📄 Reference

- [Simple Bank API](https://github.com/matheusmosca/simple-bank)
- [Golang Clean Architecture](https://github.com/khannedy/golang-clean-architecture)
- [Golang Clean Architecture #2](https://github.com/bxcodec/go-clean-arch)

# 📌 Endpoints

## 🧍‍♂️ USERS

### `/users/login - POST` Login URL. Generate a Json Web Token. Example of request body:
```json
{
  "username": "imamrizaldi",
  "password": "donggala"
}
```
### `/users - POST` Create a new user. Example of request body:
```json
{
  "username": "johndoe",
  "password": "janedoe"
}
```
### `/users - GET` Fetch all users

## ✏ NOTES

### `/notes - POST` Create new note. Example of request body:
```json
{
  "title": "Golang",
  "content" : "Tutorial Golang Dasar"
}
```
### `/notes - GET` Fetch all notes
### `/notes/{id} - GET` Get note by ID
### `/notes/{id} - PUT` Update note. Example of request body:
```json
{
  "title": "Golang",
  "content" : "Tutorial Golang MySQL"
}
```
### `/users/notes/{id} - DELETE` Delete note
<br>

# 🔜 Feature to Add:
- JWT auth (done✅)
- Unit testing