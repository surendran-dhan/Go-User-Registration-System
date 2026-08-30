# Go User Registration System

A simple full-stack user registration application built using **Golang, PostgreSQL, HTML, and CSS**.

This project demonstrates how a frontend form communicates with a Go backend and how Go connects to PostgreSQL to store user registration data.

## 🚀 Technologies Used

* **Golang** — Backend server and HTTP request handling
* **PostgreSQL** — Relational database
* **HTML** — Registration form
* **CSS** — Responsive dark/glass UI
* **SQL** — Database operations
* **Postman** — API/request testing

## 📌 Features

* User registration form
* Username and password input
* HTTP POST request handling
* Go backend server
* PostgreSQL database connection
* Store user data in PostgreSQL
* SQL INSERT operation
* Basic request validation
* Dark glassmorphism user interface

## 🔄 Application Flow

```text
User
  ↓
HTML + CSS Registration Form
  ↓
HTTP POST Request
  ↓
Golang Backend
  ↓
SQL INSERT
  ↓
PostgreSQL Database
```

## 🗄️ Database

The application uses a PostgreSQL database with a `users` table.

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(100),
    password VARCHAR(100)
);
```

## ⚙️ How to Run

### 1. Clone the repository

```bash
git clone <your-repository-url>
cd go-user-registration-system
```

### 2. Create the database

Create a PostgreSQL database:

```sql
CREATE DATABASE login_db;
```

Create the users table:

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(100),
    password VARCHAR(100)
);
```

### 3. Install Go dependency

```bash
go get github.com/lib/pq
```

### 4. Configure PostgreSQL

Update the PostgreSQL connection details in `main.go`:

```go
connStr := "host=localhost port=5432 user=postgres password=YOUR_PASSWORD dbname=login_db sslmode=disable"
```


### 5. Run the application

```bash
go run main.go
```

### 6. Open the application

```text
http://localhost:8080
```

## 📚 What I Learned

Through this project, I practiced:

* Creating an HTTP server using Go
* Handling HTTP POST requests
* Processing HTML form data
* Connecting Go with PostgreSQL
* Using `database/sql`
* Using `sql.Open()`
* Using `db.Ping()`
* Executing SQL queries with `db.Exec()`
* Using SQL parameter placeholders
* Structuring a simple full-stack application

## 🔮 Future Improvements

* Password hashing using bcrypt
* Login authentication
* JWT authentication
* Input validation
* GET user API
* Update and delete user APIs
* REST API architecture
* Improved error handling
* Environment variables for database credentials

## 👨‍💻 Project Goal
Output of Screenshot:
<img width="1361" height="636" alt="image" src="https://github.com/user-attachments/assets/719b05ff-fc73-4923-ac94-f2c7c4767dc4" />
<img width="1354" height="473" alt="image" src="https://github.com/user-attachments/assets/1320d017-f171-4fcc-813f-c0751d81ce13" />



The goal of this project is to demonstrate practical backend development skills using **Golang and PostgreSQL**, while understanding the complete flow from frontend form submission to database storage.
