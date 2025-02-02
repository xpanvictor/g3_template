# Go backend template

## Usage
```bash
chmod +x ./scripts/setup.sh
./setup.sh github.com/<username>/<module-name>

# run dev server after setting up permission
./scripts/runDevServer # uses docker compose
```

## Flow
- [x] Project Setup  
  - Initialize a Go project.  
  - Add required modules and libraries for:   
    - GraphQL  
    - Gin  
    - MongoDB  
    - Redis  
  - Authentication tools (Auth0 or JWT).  
- [x] Architecture
  - Decide on Clean Architecture or another modular pattern.
  - Create directory structure (cmd, internal, pkg, etc.).
- [x] GraphQL API
  - Choose a GraphQL library (e.g., gqlgen, graphql-go).
  - Define schema.
  - Implement resolvers.
- [ ] Routing
  - Integrate Gin for HTTP routing.
  - Create middleware for:
  - Logging
  - Error handling
  - Authentication. [TODO]
- [x] Database Integration
  - Set up MongoDB connection.
  - Define models and establish best practices for database operations.
- [ ] Caching with Redis
  - Configure Redis for:
  - Caching
  - Session management.
  - Write helper utilities for Redis operations.
- [ ] Authentication
  - Add support for:
  - JWT
  - Auth0.
  - Implement middleware for authorization.
- [x] Dockerization
  - Write a Dockerfile for the Go service.
  - Create docker-compose.yml for:
  - MongoDB
  - Redis
  - Backend.
- [x] Configuration
  - Use a configuration library (e.g., viper) to manage environment variables.
- [ ] Testing
  - Write unit tests.
  - Write integration tests.
  - Use tools like mock or testify for testing.
- [ ] Continuous Integration/Deployment
  - Add a CI/CD pipeline using:
  - GitHub Actions
  - Another preferred platform.