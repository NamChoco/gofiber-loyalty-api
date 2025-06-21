# Go Fiber PostgreSQL API - Loyalty Point System

A backend API built with [Go Fiber](https://gofiber.io), [GORM](https://gorm.io), and PostgreSQL for managing a loyalty points platform. \
This project is containerized using Docker and includes [pgAdmin 4](https://www.pgadmin.org/) for database administration.

---

## 🚀 Features

- ⚡ Fast and lightweight web framework using [Fiber](https://gofiber.io)
- 🔄 Full CRUD API with [GORM](https://gorm.io)
- 🐘 PostgreSQL + pgAdmin via Docker Compose
- 📦 Environment configuration using `.env`
- 🔧 Auto-migration for loyalty system models (Member, Reward, Transaction, Redemption)

---

## 📁 Project Structure

```
go-database/
│
├── controller/        # Fiber route handlers (e.g. GetMembers, CreateReward)
├── database/          # Database connection and migration logic
├── model/             # GORM models (Member, Reward, Transaction, Redemption)
├── router/            # API route definitions
├── .env               # Environment variables
├── docker-compose.yml
├── go.mod / go.sum    # Go modules
└── main.go            # Application entry point
```

---

## ⚙️ Getting Started

### Prerequisites

- Go 1.20+
- Docker & Docker Compose

### 1. Clone the Repository

```bash
git clone "Clone using the web URL."
cd project-name
```

### 2. Setup Environment Variables

Create `.env` file:

```dotenv
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=root
DB_NAME=mydb
JWT_SECRET=secret
```

### 3. Start PostgreSQL + pgAdmin with Docker

```bash
docker-compose up -d
```

- PostgreSQL: `localhost:5432`
- pgAdmin: [http://localhost:8081](http://localhost:8081)
  - Email: `admin@admin.com`
  - Password: `root`

### 4. Run the API

```bash
go run .
```

API server will run at: [http://localhost:8080](http://localhost:8080)

---

## 📚 Example API Endpoints

| Method | Endpoint             | Description                      |
|--------|----------------------|----------------------------------|
| GET    | `/api/members`       | Get all members                  |
| POST   | `/api/members`       | Create a new member              |
| GET    | `/api/rewards`       | Get all rewards                  |
| POST   | `/api/rewards`       | Add a new reward                 |
| GET    | `/api/transactions`  | Get all transactions             |
| POST   | `/api/transactions`  | Create a new transaction         |
| GET    | `/api/redemptions`   | Get all redemptions              |
| POST   | `/api/redemptions`   | Redeem points for a reward       |

---

## 🔪 Testing

You can use tools like:

- [Postman](https://www.postman.com/)
- `curl` CLI tool
---