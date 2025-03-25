# Social Media API
<div style="display:flex">
  <img align="center" alt="GO" src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" />
  <img align="center" alt="mysql" src="https://img.shields.io/badge/MySQL-005C84?style=for-the-badge&logo=mysql&logoColor=white" />
  <img align="center" alt="jwt" src="https://img.shields.io/badge/JWT-000000?style=for-the-badge&logo=JSON%20web%20tokens&logoColor=white" />
</div>


## 💻 The project 
This is a RESTful API developed in Go that simulates a social media network. The API allows users to create accounts, publish posts, interact with other users through followers and likes, and manage their profiles.

## 💡 Technologies Used:

- **Go (Golang)** - Main programming language
- **Gorilla Mux** - HTTP router and dispatcher
- **MySQL** - Relational database
- **JWT (JSON Web Tokens)** - Authentication and security


## 📘 Dependencies

The application uses the following Go libraries:
- github.com/badoux/checkmail - Email validation
- github.com/go-sql-driver/mysql - MySQL driver for Go
- github.com/gorilla/mux - HTTP routing
- github.com/joho/godotenv - Environment variable management
- golang.org/x/crypto - Cryptographic algorithms
- github.com/dgrijalva/jwt-go - JWT generation and validation

## 🚀 Features

**Users**
- Create a user account
- Retrieve a list of users
- Get user details by ID
- Update profile information
- Delete user account
- Follow and unfollow other users
- List followers and following users
- Update password

**Posts**
- Create a post
- Retrieve all posts and user-specific posts
- Get a post by ID
- Update a post
- Delete a post
- Like and unlike posts

## 🔧 Setup and Execution
1. Clone the repository
```bash
$ git clone https://github.com/your-username/your-repository.git
$ cd your-repository
```
2. Configure the database
Create a MySQL database and update the connection settings.

3. Install dependencies
```bash
$ go mod tidy
```
4. Run the application

## ▶️ API Endpoints

Authentication
- POST /login - Authenticates a user and returns a JWT token

Users
- POST /users - Creates a new user
- GET /users - Retrieves all users
- GET /users/{userId} - Gets a user by ID
- PUT /users/{userId} - Updates user information
- DELETE /users/{userId} - Deletes a user
- POST /users/{userId}/follow - Follows a user
- POST /users/{userId}/unfollow - Unfollows a user
- GET /users/{userId}/followers - Lists followers
- GET /users/{userId}/following - Lists users being followed
- POST /users/{userId}/updatePassword - Updates password

Posts
- POST /posts - Creates a post
- GET /posts - Retrieves all posts
- GET /posts/{postId} - Gets a post by ID
- PUT /posts/{postId} - Updates a post
- DELETE /posts/{postId} - Deletes a post
- GET /users/{userId}/posts - Lists posts from a specific user
- POST /posts/{postId}/like - Likes a post
- POST /posts/{postId}/unlike - Unlikes a post

.env example
```bash
DB_USER=
DB_PASSWORD=
DB_NAME=
API_HOST=
SECRET_KEY=
```






