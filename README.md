# 🎬 Tickitz API Go 🎟️

<div align="center">
  
  <!-- ![Tickitz Logo](https://via.placeholder.com/200x80?text=Tickitz) -->
  
  ### A powerful movie ticket booking API built with Go
  
  [![Go Version](https://img.shields.io/badge/Go-1.23.5-00ADD8?style=flat-square&logo=go)](https://golang.org/)
  [![License](https://img.shields.io/badge/License-MIT-blue?style=flat-square)](LICENSE)
  [![Swagger](https://img.shields.io/badge/Swagger-Documented-85EA2D?style=flat-square&logo=swagger)](./api/swagger.json)
  
</div>

## ✨ Features

- 🎞️ **Movie Management** - Browse, search, and filter movies
- 🗓️ **Schedule Management** - View available showtimes and cinemas
- 🔐 **Authentication** - Secure user registration and login
- 👤 **User Profiles** - Personalized user experience
- 🎫 **Ticket Booking** - Reserve seats and book tickets
- 💰 **Payment Processing** - Handle ticket payments
- 👑 **Admin Dashboard** - Manage movies, schedules, and more

## 🚀 Getting Started

### Prerequisites

- Go 1.23.5 or later
- PostgreSQL
- Git

### Installation

1. **Clone the repository**

```bash
git clone https://github.com/mhakimsaputra17/tickitz-api-go.git
cd tickitz-api-go
```

2. **Set up environment variables**

```bash
cp .env.example .env
# Edit .env file with your configuration
```

3. **Install dependencies**

```bash
go mod download
```

4. **Run the application**

```bash
go run main.go
```

## 📚 API Documentation

API documentation is available via Swagger UI. After running the application, visit:

```
http://localhost:8080/swagger/index.html
```

For a detailed API specification, see the [Swagger JSON file](./api/swagger.json).

## 🛠️ Tech Stack

- **[Go](https://golang.org/)** - Programming language
- **[Gin](https://gin-gonic.com/)** - Web framework
- **[PostgreSQL](https://www.postgresql.org/)** - Database
- **[pgx](https://github.com/jackc/pgx)** - PostgreSQL driver
- **[JWT](https://github.com/golang-jwt/jwt)** - Authentication
- **[Swagger](https://swagger.io/)** - API documentation

## 📁 Project Structure

```
tickitz-api-go/
├── api/            # API documentation
├── cmd/            # Application entrypoints
├── config/         # Configuration files
├── internal/       # Internal packages
│   ├── handler/    # HTTP handlers
│   ├── middleware/ # HTTP middlewares
│   ├── model/      # Data models
│   ├── repository/ # Database operations
│   └── service/    # Business logic
├── pkg/            # Shareable packages
└── scripts/        # Utility scripts
```

## 👨‍💻 Development

### Running in Development Mode

```bash
go run main.go
```

### Running Tests

```bash
go test ./...
```

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 📧 Contact

- **Project Maintainer** - [mhakimsaputra17](mailto:mhakimsaputra17@gmail.com)
- **Issues and Feature Requests** - [GitHub Issues](https://github.com/mhakimsaputra17/tickitz-api-go/issues)

---

<div align="center">
  
  Made with ❤️ by mhakimsaputra17
  
</div>
