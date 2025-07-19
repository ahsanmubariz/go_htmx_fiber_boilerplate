# Go + Fiber + HTMX Boilerplate (Production Ready)

A production-ready boilerplate for building web applications using Go, the Fiber framework, and HTMX for dynamic frontends. This project follows a modular monolith architecture and includes key security features and session management.

## Features

- **Go + Fiber:** A high-performance and expressive backend.
- **HTMX:** Modern browser features directly from HTML.
- **PostgreSQL Database:** Integrated with GORM.
- **Modular Monolith:** Organized into modules for better separation of concerns.
- **MVC-like Pattern:** Controllers, Services, and Views (Templates).
- **Security Hardened:**
    - CSRF Protection Middleware
    - Secure Headers with Helmet
    - Input Validation
- **Session Management:** Using database-backed sessions.
- **Containerized:** Multi-stage `Dockerfile` for small, secure production images.
- **Live Reloading:** Using `air` for a fast development workflow.
- **Environment-based Config:** Using `.env` files for configuration.
- **CRUD Example:** A fully functional, HTMX-powered "Users" CRUD module.

## Prerequisites

- Go 1.22+
- Docker & Docker Compose
- [Air](https://github.com/cosmtrek/air) for live reloading (development only).

## Getting Started (Development)

1.  **Clone the repository:**
    ```bash
    git clone <your-repo-url>
    cd go-fiber-htmx-boilerplate
    ```

2.  **Start the database:**
    ```bash
    docker-compose up -d
    ```

3.  **Create your environment file:**
    ```bash
    cp .env.example .env
    ```
    *Important: For production, generate new, random `CSRF_SECRET` and `SESSION_KEY` values.*

4.  **Install dependencies:**
    ```bash
    go mod tidy
    ```

5.  **Run the application (with live reloading):**
    ```bash
    air
    ```

6.  **Open your browser:**
    Navigate to `http://localhost:8080`.

## Building for Production

1.  **Build the Docker image:**
    ```bash
    docker build -t go-fiber-htmx-app .
    ```

2.  **Run the container:**
    ```bash
    docker run -p 8080:8080 --env-file .env go-fiber-htmx-app
    ```
    Ensure your `.env` file is configured for your production environment.