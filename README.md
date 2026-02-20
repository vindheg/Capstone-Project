# Recipe Sharing API 🍳

A RESTful API built using Go (Golang), Gin, and GORM that allows users to share recipes with images, search recipes by ingredients, and rate recipes.

This project demonstrates backend API design, image upload handling, database integration, and structured project architecture.

---

## 🚀 Tech Stack

- Go (Golang)
- Gin Web Framework
- GORM ORM
- SQLite (or PostgreSQL if configured)
- UUID for unique identifiers

---

## 📁 Project Structure

```
recipe-api/
│
├── cmd/
│   └── main.go
│
├── internal/
│   ├── database/
│   ├── models/
│   ├── handlers/
│   └── routes/
│
├── uploads/
├── go.mod
├── README.md
└── PROMPTS.md
```

---

## 🗄️ Database Schema

### User
- ID (UUID, Primary Key)
- Username (Unique)

### Recipe
- ID (UUID, Primary Key)
- Title
- Description
- Ingredients
- ImageURL
- UserID (Foreign Key)

### Rating
- ID (UUID, Primary Key)
- UserID (Foreign Key)
- RecipeID (Foreign Key)
- Score (1–5)

---

## ⚙️ Setup Instructions

### 1️⃣ Clone Repository

```bash
git clone <your-repo-url>
cd recipe-api
```

---

### 2️⃣ Install Dependencies

```bash
go mod tidy
```

---

### 3️⃣ Run Application

If `main.go` is inside `cmd` folder:

```bash
go run cmd/main.go
```

Server runs on:

```
http://localhost:8080
```

---

## 📌 API Endpoints

### ➤ Create Recipe

**POST** `/recipes`

Form Data:
- title
- description
- ingredients
- image (file)

---

### ➤ Get All Recipes

**GET** `/recipes`

---

### ➤ Search Recipes by Ingredient

**GET** `/recipes/search?ingredient=chicken`

---

### ➤ Rate a Recipe

**POST** `/recipes/:id/rate`

Body:
```json
{
  "user_id": "uuid",
  "score": 5
}
```

---

## 🖼 Image Handling

- Images are uploaded using multipart/form-data.
- Stored locally in `/uploads` directory.
- Unique filenames generated using UUID.

---

## 🧠 Design Decisions

- Clean layered architecture (routes → handlers → models → database)
- UUID used to avoid ID collisions
- GORM used for ORM abstraction
- SQLite used for faster setup (can switch to PostgreSQL easily)
- Images stored locally to avoid cloud dependency

---

## 🔎 Future Improvements

- JWT Authentication
- Pagination
- Full-text search
- Average rating calculation endpoint
- Image resizing & compression
- Docker support

---

## 🧪 Testing

You can test using:

- Postman
- curl
- Browser (for GET endpoints)

---

## 📄 License

This project is developed for academic/capstone submission purposes.