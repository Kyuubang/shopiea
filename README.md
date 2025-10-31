# Shopiea

Shopiea is a RESTful API written in Go that provides scoring system functionalities for managing and scoring student 
assignments, authentication and registering lab and course, and generating reports. It currently does not have a dashboard 
UI, but provides an OpenAPI 3.0 specification for anyone to develop a dashboard.

## Features

- Authentication and authorization for users and administrators
- Registration of labs and courses
- Scoring of student assignments
- Generating reports for lab and course scores
- Database seeding for quick setup
- Docker containerization support
- Atlas integration for database schema management

## Table of Contents

- [Prerequisites](#prerequisites)
- [Installation](#installation)
  - [Local Setup](#local-setup)
  - [Docker Setup](#docker-setup)
- [Database Management](#database-management)
  - [Database Migration](#database-migration)
  - [Database Seeding](#database-seeding)
  - [Using Atlas](#using-atlas)
- [Running the Application](#running-the-application)
- [Default Credentials](#default-credentials)
- [API Documentation](#api-documentation)
- [License](#license)

## Prerequisites

- Go 1.18 or higher
- PostgreSQL 15 or higher
- Docker and Docker Compose (optional, for containerized setup)
- Atlas CLI (optional, for advanced schema management)

## Installation

### Local Setup

1. Clone the repository:
```bash
git clone https://github.com/Kyuubang/shopiea.git
cd shopiea
```

2. Install dependencies:
```bash
go mod download
```

3. Configure environment variables:
   
   Create or modify the `.env` file with your database configuration:
```bash
SHOPIEA_MODE=development
SHOPIEA_ENV=development
SHOPIEA_HOST=0.0.0.0
SHOPIEA_PORT=9898
SHOPIEA_DB_HOST=localhost
SHOPIEA_DB_PORT=5432
SHOPIEA_DB_USER=shopiea
SHOPIEA_DB_PASSWORD=mysecretpassword
SHOPIEA_DB_NAME=shopiea
SHOPIEA_CASE_REPO=github.com/Kyuubang/philo-sample-case
SHOPIEA_CASE_BRANCH=master
SHOPIEA_INFRA_REPO=github.com/Kyuubang/philo-sample-infra
SHOPIEA_INFRA_BRANCH=master
```

4. Source the environment file:
```bash
set -a; source .env; set +a
```

### Docker Setup

For a quick setup using Docker, use Docker Compose:

```bash
# Start the application with database
docker-compose up -d

# The database will be automatically initialized with schema and seed data
```

This will:
- Start a PostgreSQL 15 database container
- Initialize the database schema
- Seed the database with default data
- Start the Shopiea application

To stop the containers:
```bash
docker-compose down
```

To remove all data (including the database volume):
```bash
docker-compose down -v
```

## Database Management

### Database Migration

Run database migrations to create the necessary tables:

```bash
go run main.go -migrate
```

This command will:
- Create all required tables (roles, classes, users, courses, labs, scores)
- Set up foreign key relationships
- Exit after migration is complete

### Database Seeding

After migrating the database, seed it with default data:

```bash
go run main.go -seed
```

This command will populate the database with:
- **Roles**: admin, student
- **Default Admin User**:
  - Username: `admin`
  - Password: `admin123`
  - Role: admin
  - Class: Shopiea
- **Default Class**: Shopiea
- **Default Course**: golang
- **Default Lab**: golang-001

The seeding operation is idempotent - it won't create duplicates if the data already exists.

### Using Atlas

Atlas is integrated for advanced database schema management. Install Atlas CLI:

```bash
# Install Atlas
go install ariga.io/atlas/cmd/atlas@latest

# Add to PATH if needed
export PATH=$PATH:~/go/bin
```

#### Inspect Current Schema

Inspect your current database schema:

```bash
atlas schema inspect \
  --url "postgres://shopiea:mysecretpassword@localhost:5432/shopiea?sslmode=disable" \
  --env local
```

#### Apply Schema Changes

Apply schema from the schema file:

```bash
atlas schema apply \
  --url "postgres://shopiea:mysecretpassword@localhost:5432/shopiea?sslmode=disable" \
  --to "file://db/schema.sql" \
  --env local
```

#### Validate Schema

Compare your database with the schema definition:

```bash
atlas schema diff \
  --from "postgres://shopiea:mysecretpassword@localhost:5432/shopiea?sslmode=disable" \
  --to "file://db/schema.sql"
```

## Running the Application

### Local Development

1. Make sure PostgreSQL is running and the database is migrated and seeded

2. Run the server:
```bash
go run main.go
```

Or use the Makefile:
```bash
make run
```

3. The API will be available at `http://localhost:9898`

### Production

1. Build the application:
```bash
go build -o bin/shopiea
```

2. Run the binary:
```bash
./bin/shopiea
```

For production deployment, make sure to:
- Set `SHOPIEA_MODE=release`
- Set `SHOPIEA_ENV=production`
- Use strong passwords
- Enable SSL/TLS for database connections
- Use secure authentication tokens

## Default Credentials

After seeding the database, you can log in with:

- **Username**: `admin`
- **Password**: `admin123`

**Important**: Change the default admin password in production environments!

## API Documentation

The API endpoints are documented in the OpenAPI 3.0 specification file: `openapi-shopiea.json`

### Key Endpoints

- `POST /auth/login` - User authentication
- `GET /info` - Application information
- `GET /v1/course` - List all courses (authenticated)
- `GET /v1/labs` - List all labs (authenticated)
- `POST /v1/score` - Submit a score (authenticated)
- `GET /v1/score` - Get scores (authenticated)
- `POST /v1/auth/check` - Verify authentication token

#### Admin Endpoints (requires admin role)

- `GET /v1/admin/user` - List all users
- `POST /v1/admin/user` - Create a new user
- `PUT /v1/admin/user` - Update a user
- `DELETE /v1/admin/user` - Delete a user
- `GET /v1/admin/class` - List all classes
- `POST /v1/admin/class` - Create a new class
- `PUT /v1/admin/class` - Update a class
- `DELETE /v1/admin/class` - Delete a class
- `POST /v1/admin/course` - Create a new course
- `PUT /v1/admin/course` - Update a course
- `DELETE /v1/admin/course` - Delete a course
- `POST /v1/admin/labs` - Create a new lab
- `PUT /v1/admin/labs` - Update a lab
- `DELETE /v1/admin/labs` - Delete a lab
- `GET /v1/admin/export` - Export scores

## Database Schema

The application uses the following database schema:

### Tables

- **roles**: User roles (admin, student)
- **classes**: Student classes
- **users**: User accounts with authentication
- **courses**: Available courses
- **labs**: Lab assignments for courses
- **scores**: Student scores for labs

### Relationships

- Users belong to a Role and a Class
- Labs belong to a Course
- Scores belong to a User and a Lab

## License

This project is licensed under the MIT License. See the [LICENSE](https://github.com/Kyuubang/shopiea/blob/master/LICENSE) file for details.
