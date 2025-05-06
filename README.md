# DevOps-Final

## Docker Documentation

### Project Structure
```
devOps/
├── main.go           # Main application file
├── Dockerfile        # Docker image configuration
├── docker-compose.yml # Docker compose configuration
├── Jenkinsfile       # Jenkins pipeline configuration
└── README.md         # Documentation
```

### Prerequisites
- Docker Desktop installed
- Docker Hub account
- Go 1.20 or later

### Docker Configuration

#### Dockerfile
The project uses a multi-stage build to create a lightweight container:
```dockerfile
# Build stage
FROM golang:1.20-alpine AS builder
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app

# Final stage
FROM alpine:latest
COPY --from=builder /build/app /usr/local/bin/app
RUN chmod +x /usr/local/bin/app
EXPOSE 4200
CMD ["/usr/local/bin/app"]
```

#### Docker Compose
```yaml
version: '3.8'
services:
  go-book-api:
    build: .
    ports:
      - "4200:4200"
    volumes:
      - .:/app
    environment:
      - GO_ENV=development
    networks:
      - go-book-api-network

networks:
  go-book-api-network:
    driver: bridge
```

### Building and Running

1. **Build Docker Image**
```bash
docker build -t whyyouask/devops-boks:latest .
```

2. **Run with Docker Compose**
```bash
docker-compose up -d
```

3. **Stop Containers**
```bash
docker-compose down
```

### Docker Hub Integration

1. **Login to Docker Hub**
```bash
docker login
```

2. **Push Image to Docker Hub**
```bash
docker push whyyouask/devops-boks:latest
```

3. **Pull Image from Docker Hub**
```bash
docker pull whyyouask/devops-boks:latest
```

### API Endpoints
- GET `http://localhost:4200/api/books` - Get all books
- GET `http://localhost:4200/api/books/:id` - Get book by ID
- POST `http://localhost:4200/api/books` - Create new book
- PUT `http://localhost:4200/api/books/:id` - Update book
- DELETE `http://localhost:4200/api/books/:id` - Delete book

### Environment Variables
- `GO_ENV`: Development environment setting
- Port: 4200 (configured in Dockerfile and docker-compose.yml)

### Docker Best Practices Used
1. Multi-stage builds for smaller image size
2. Non-root user in production
3. Proper layer caching
4. Environment variable configuration
5. Volume mounting for development
6. Network isolation using custom network

### Troubleshooting
1. If port 4200 is already in use:
```bash
docker-compose down
netstat -ano | findstr :4200
# Kill the process if needed
taskkill /PID <PID> /F
```

2. If Docker image build fails:
```bash
# Clean Docker cache
docker system prune -a
# Rebuild
docker-compose build --no-cache
```

### CI/CD Integration
The project includes Jenkins pipeline configuration for automated Docker builds and deployments. See Jenkinsfile for details.