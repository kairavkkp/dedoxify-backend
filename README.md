To create backend APIs with ORM in Go without relying on full-fledged frameworks, you can take a modular approach by using libraries that provide routing, database connections, and ORM functionalities while retaining flexibility. Here's a stack that could work well:

### 1. **Routing**
   - **Chi**: A lightweight and idiomatic router that is compatible with Go's `net/http`. It provides minimalistic and expressive routing without adding too much overhead.
     - GitHub: [go-chi/chi](https://github.com/go-chi/chi)
   - **Mux**: Part of the Gorilla toolkit, it's a simple HTTP router and URL matcher, perfect for handling HTTP requests.
     - GitHub: [gorilla/mux](https://github.com/gorilla/mux)

### 2. **ORM**
   - **GORM**: A popular ORM for Go that provides a full-featured ORM with support for querying, relationships, hooks, transactions, etc. You can use only the parts you need and avoid others.
     - GitHub: [go-gorm/gorm](https://github.com/go-gorm/gorm)
   - **SQLx**: A minimal extension to Go's `database/sql` that provides easier handling of named parameters, struct scanning, and simple queries. It gives you more control over your SQL while abstracting some boilerplate.
     - GitHub: [jmoiron/sqlx](https://github.com/jmoiron/sqlx)

### 3. **Database Connection**
   - **Database/sql**: The standard Go library for SQL database interactions. Using this with a library like SQLx will provide low-level control over your queries and connection handling.
   - **PGX**: A pure Go PostgreSQL driver and toolkit, which offers better performance for PostgreSQL compared to `database/sql`. It can be used as a replacement for Postgres-based backends.
     - GitHub: [jackc/pgx](https://github.com/jackc/pgx)

### 4. **Middleware (Optional)**
   - **Negroni**: A tiny middleware handler that adds flexibility to your API by making it easier to manage middleware layers like logging, recovery, and authorization.
     - GitHub: [urfave/negroni](https://github.com/urfave/negroni)

### 5. **Validation**
   - **Go-playground/validator**: A library for validating Go structs and fields. It’s useful for validating API request payloads.
     - GitHub: [go-playground/validator](https://github.com/go-playground/validator)

### 6. **Authentication/Authorization**
   - **JWT-Go**: A library for working with JSON Web Tokens, useful for token-based authentication.
     - GitHub: [dgrijalva/jwt-go](https://github.com/dgrijalva/jwt-go)

By combining these libraries, you can create a modular API backend in Go, where you control the architecture while making use of specialized tools for routing, ORM, and validation.