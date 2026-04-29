# Go SQL Employee Management API

A complete REST API project built using **Go (Golang)**, **Gin Framework**, and **SQLite Database** for managing employee records with full CRUD operations.

---

#  Features

✅ Add New Employee  
✅ View All Employees  
✅ View Single Employee  
✅ Update Employee Details  
✅ Delete Employee  
✅ SQLite Database Integration  
✅ REST API Architecture  
✅ Clean File Structure  
✅ Beginner Friendly Project  

---

# 🛠 Tech Stack

- Go (Golang)
- Gin Gonic Framework
- SQLite
- REST API

---

# 📁 Project Structure

```
employee-api/
│── go.mod
│── main.go
│── employees.db
│── database/
│   └── db.go
│── models/
│   └── employee.go
│── handlers/
│   └── employee_handler.go

```
# Installation
> 1 Clone Project
```
git clone https://github.com/yourusername/employee-api.git
cd employee-api
```
> 2 Install Dependencies
```
go mod tidy
```
> 3 Run Project
```
go run main.go
```
Server Starts On:
```
http://localhost:8080
```
# API Endpoints
> Create Employee
```
POST /employees
```
> Body
```
{
  "name":"Rishi",
  "email":"rishi@gmail.com",
  "position":"Developer",
  "salary":50000
}
```
> Get All Employees
```
GET /employees
```
> Get Single Employee
```
GET /employees/1
```
> Update Employee
```
PUT /employees/1
```
> Body
```
{
  "name":"Rishi Soni",
  "email":"rishi@gmail.com",
  "position":"Senior Developer",
  "salary":80000
}
```
> Delete Employee
```
DELETE /employees/1
```
# Database Schema
```
CREATE TABLE employees (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT,
    email TEXT,
    position TEXT,
    salary INTEGER
);
```
# Example Response
```
{
  "id":1,
  "name":"Rishi",
  "email":"rishi@gmail.com",
  "position":"Developer",
  "salary":50000
}
```
# Future Improvements
```
- JWT Authentication
- MySQL / PostgreSQL Support
- Docker Deployment
- Pagination
- Search Employee
- Admin Panel
```
# Contribution

> Pull requests are welcome.

> If you’d like to improve this project, fork the repo and submit a PR.

# License
> MIT License

# Author
```
rishi02soni
```



