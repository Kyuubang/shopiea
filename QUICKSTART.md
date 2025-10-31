# Shopiea Quick Start Guide

Get started with Shopiea in minutes using this quick start guide.

## Option 1: Docker (Recommended for Quick Setup)

This is the fastest way to get Shopiea up and running.

### Prerequisites
- Docker
- Docker Compose

### Steps

1. Clone the repository:
```bash
git clone https://github.com/Kyuubang/shopiea.git
cd shopiea
```

2. Start the application:
```bash
docker-compose up -d
```

That's it! The application will be available at `http://localhost:9898`

The database will be automatically:
- Created with the correct schema
- Seeded with default data (admin user, roles, class, course, lab)

3. Test the API:
```bash
# Get application info
curl http://localhost:9898/info

# Login with default admin credentials
curl -X POST http://localhost:9898/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"admin123"}'
```

4. Stop the application:
```bash
docker-compose down
```

To remove all data including the database:
```bash
docker-compose down -v
```

## Option 2: Local Setup

For development or if you prefer running locally.

### Prerequisites
- Go 1.18 or higher
- PostgreSQL 15 or higher

### Steps

1. Clone the repository:
```bash
git clone https://github.com/Kyuubang/shopiea.git
cd shopiea
```

2. Start PostgreSQL (if not already running):
```bash
# Using Docker
docker run --name postgres -e POSTGRES_PASSWORD=mysecretpassword \
  -e POSTGRES_USER=shopiea -e POSTGRES_DB=shopiea \
  -p 5432:5432 -d postgres:15

# Or use your local PostgreSQL installation
```

3. Configure environment:
```bash
# Copy and modify .env if needed
cp .env .env.local

# Source the environment
set -a; source .env; set +a
```

4. Run database migration:
```bash
go run main.go -migrate
```

5. Seed the database:
```bash
go run main.go -seed
```

6. Start the application:
```bash
go run main.go
```

The API will be available at `http://localhost:9898`

## Option 3: Using Makefile

If you have Make installed:

```bash
# Run migrations
make migrate

# Seed the database
make seed

# Run the application
make run

# Or use Docker
make docker-up
make docker-down
```

## Default Credentials

After seeding, you can log in with:
- **Username**: `admin`
- **Password**: `admin123`

**Important**: Change this password in production!

## Quick API Test

1. Get a JWT token:
```bash
TOKEN=$(curl -s -X POST http://localhost:9898/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"admin123"}' | jq -r '.token')

echo "Token: $TOKEN"
```

2. List courses:
```bash
curl http://localhost:9898/v1/course \
  -H "Authorization: Bearer $TOKEN"
```

3. List labs:
```bash
curl http://localhost:9898/v1/labs \
  -H "Authorization: Bearer $TOKEN"
```

4. List users (admin only):
```bash
curl http://localhost:9898/v1/admin/user \
  -H "Authorization: Bearer $TOKEN"
```

## Seeded Data

After running the seed command, the database contains:

### Roles
- ID: 1, Name: admin
- ID: 2, Name: student

### Classes
- ID: 1, Name: Shopiea

### Users
- ID: 1, Username: admin, Password: admin123, Role: admin, Class: Shopiea

### Courses
- ID: 1, Name: golang

### Labs
- ID: 1, Name: golang-001, Course: golang

## Next Steps

1. **Change the default password**: Create a new admin user or change the password
2. **Create student users**: Use the admin API to create student accounts
3. **Add more courses and labs**: Use the admin API to expand your content
4. **Explore the API**: Check out the OpenAPI specification in `openapi-shopiea.json`

## Troubleshooting

### Cannot connect to database
- Verify PostgreSQL is running: `pg_isready -h localhost -p 5432`
- Check database credentials in `.env`
- Ensure the database exists: `psql -U shopiea -l`

### Port already in use
- Change the port in `.env` (SHOPIEA_PORT)
- Or stop the service using port 9898

### Docker issues
- Check Docker is running: `docker info`
- View logs: `docker-compose logs`
- Rebuild containers: `docker-compose up --build`

### Migration or seeding fails
- Ensure database is accessible
- Check for existing tables: `psql -U shopiea -d shopiea -c "\dt"`
- Drop and recreate database if needed (development only!)

## More Information

- Full documentation: [README.md](README.md)
- Atlas guide: [ATLAS_GUIDE.md](ATLAS_GUIDE.md)
- API specification: [openapi-shopiea.json](openapi-shopiea.json)
- GitHub: https://github.com/Kyuubang/shopiea
