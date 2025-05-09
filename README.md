# GoVaCatalog - Job Vacancy Aggregator

GoVaCatalog is a powerful job vacancy aggregator that collects and organizes job listings from multiple sources into a single, user-friendly platform. Whether you're a job seeker looking for opportunities or an employer posting positions, GoVaCatalog provides a seamless experience through both web interface and API endpoints.

## Features

- 📚 Aggregate job listings from multiple sources
- 🎨 Clean and intuitive user interface
- 🛠️ RESTful API for programmatic access
- 🔍 Advanced search and filtering
- 📊 Real-time job market analytics

## Tech Stack

- Backend: Go
- Frontend: now it's without frontend
- Database: PostgreSQL
- Containerization: Docker

## Getting Started

### Prerequisites

- Go 1.20 or higher
- Docker and Docker Compose (recommended for production)
- Modern web browser

### Installation

1.Clone the repository:

```bash
git clone https://github.com/4ktivated/GoVaCatalog.git
cd GoVaCatalog
```

2.Install dependencies:

```bash
go mod download
```

3.Set up production environment variables (copy `.env.example` to `.env` in `cmd` directory):

```bash
cp cmd/.env.example cmd/.env
```

### Production Mode

To start the application with production environment variables and Docker Compose:

## Running the Application

### Development Mode (Database Only)

To start only the database in development mode:

```bash
docker-compose -f docker-compose.dev.yaml up -d
go run cmd/main.go
```

### Production Mode (Database and Application)

To start both database and application in production mode:

```bash
docker-compose -f docker-compose.prod.yaml up -d
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/your-feature`)
3. Commit your changes (`git commit -m 'Add your feature description'`)
4. Push to the branch (`git push origin feature/your-feature`)
5. Open a Pull Request

Please ensure your code follows the project's coding standards and includes appropriate tests.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Support

For support, please:

1. Check the documentation
2. Search existing issues
3. Open a new issue if needed
4. dont contact us (we already know where to find you)
