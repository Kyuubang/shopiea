# Atlas Integration Guide for Shopiea

This guide explains how to use Atlas for database schema management in the Shopiea application.

## What is Atlas?

Atlas is a modern database schema management tool that provides:
- Declarative schema definitions
- Schema versioning and migrations
- Schema inspection and validation
- CI/CD integration

## Installation

### Install Atlas CLI

```bash
# Via Go
go install ariga.io/atlas/cmd/atlas@latest

# Add to PATH
export PATH=$PATH:~/go/bin

# Verify installation
atlas version
```

## Configuration

The Atlas configuration is defined in `atlas.hcl` with two environments:

### Local Environment

```hcl
env "local" {
  src = "file://db/schema.sql"
  url = getenv("SHOPIEA_DB_URL")  // Reads from environment variable
  dev = "docker://postgres/15/dev"
  schemas = ["public"]
}
```

**Important**: Set the `SHOPIEA_DB_URL` environment variable before using Atlas:
```bash
export SHOPIEA_DB_URL="postgres://shopiea:mysecretpassword@localhost:5432/shopiea?sslmode=disable"
```

### Production Environment

```hcl
env "prod" {
  src = "file://db/schema.sql"
  url = getenv("DATABASE_URL")  // Reads from environment variable
  dev = "docker://postgres/15/dev"
  schemas = ["public"]
}
```

## Schema Files

### Schema Definition (`db/schema.sql`)

This file contains the declarative schema definition for all tables:
- roles
- classes
- users
- courses
- labs
- scores

### Seed Data (`db/seed.sql`)

This file contains initial seed data:
- Default roles (admin, student)
- Default admin user
- Default class (Shopiea)
- Default course (golang)
- Default lab (golang-001)

## Common Atlas Commands

### 1. Inspect Current Database Schema

View the current state of your database using the environment configuration:

```bash
# Ensure SHOPIEA_DB_URL is set in your environment
source .env
atlas schema inspect --env local
```

Or with explicit URL (not recommended for production):

```bash
atlas schema inspect \
  --url "postgres://shopiea:mysecretpassword@localhost:5432/shopiea?sslmode=disable"
```

### 2. Apply Schema Changes

Apply the schema from `db/schema.sql` to your database using the environment configuration:

```bash
# Ensure SHOPIEA_DB_URL is set in your environment
source .env
atlas schema apply --env local
```

Or with explicit URL (not recommended for production):

```bash
atlas schema apply \
  --url "postgres://shopiea:mysecretpassword@localhost:5432/shopiea?sslmode=disable" \
  --to "file://db/schema.sql" \
  --dev-url "docker://postgres/15/dev"
```

### 3. Compare Schemas (Diff)

Compare your database with the schema definition:

```bash
atlas schema diff \
  --from "postgres://shopiea:mysecretpassword@localhost:5432/shopiea?sslmode=disable" \
  --to "file://db/schema.sql"
```

### 4. Validate Schema

Check if your database matches the schema definition:

```bash
atlas schema apply \
  --url "postgres://shopiea:mysecretpassword@localhost:5432/shopiea?sslmode=disable" \
  --to "file://db/schema.sql" \
  --dry-run
```

### 5. Generate Migration Files

Create versioned migration files from schema changes:

```bash
atlas migrate diff migration_name \
  --env local \
  --to "file://db/schema.sql"
```

## Workflows

### Initial Setup

1. Start PostgreSQL database:
```bash
# Using Docker Compose
docker-compose up -d postgres

# Or standalone PostgreSQL
```

2. Apply schema using Atlas:
```bash
atlas schema apply --env local --auto-approve
```

3. Apply seed data:
```bash
psql -h localhost -U shopiea -d shopiea -f db/seed.sql
```

Or use the application's seed command:
```bash
go run main.go -seed
```

### Update Schema

1. Modify `db/schema.sql` with your changes

2. Review changes:
```bash
atlas schema diff \
  --from "postgres://shopiea:mysecretpassword@localhost:5432/shopiea?sslmode=disable" \
  --to "file://db/schema.sql"
```

3. Apply changes:
```bash
atlas schema apply --env local
```

### Production Deployment

1. Set the DATABASE_URL environment variable:
```bash
export DATABASE_URL="postgres://user:pass@host:5432/dbname?sslmode=require"
```

2. Apply schema to production:
```bash
atlas schema apply --env prod --auto-approve
```

3. Apply seed data (if needed):
```bash
psql $DATABASE_URL -f db/seed.sql
```

## Integration with CI/CD

### GitHub Actions Example

```yaml
name: Database Schema

on:
  push:
    branches: [main]
  pull_request:
    branches: [main]

jobs:
  schema-check:
    runs-on: ubuntu-latest
    services:
      postgres:
        image: postgres:15
        env:
          POSTGRES_PASSWORD: postgres
        options: >-
          --health-cmd pg_isready
          --health-interval 10s
          --health-timeout 5s
          --health-retries 5
        ports:
          - 5432:5432

    steps:
      - uses: actions/checkout@v3
      
      - name: Install Atlas
        run: |
          curl -sSf https://atlasgo.sh | sh
          
      - name: Validate Schema
        run: |
          atlas schema apply \
            --url "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" \
            --to "file://db/schema.sql" \
            --dry-run
```

## Comparison: Atlas vs GORM AutoMigrate

### GORM AutoMigrate (Current Method)

**Pros:**
- Simple to use
- Automatic from Go structs
- No external tools needed

**Cons:**
- Limited control over migrations
- Can't handle complex schema changes
- No version history
- No rollback capability

**Usage:**
```bash
go run main.go -migrate
```

### Atlas (New Method)

**Pros:**
- Full control over schema
- Versioned migrations
- Schema validation
- CI/CD integration
- Rollback support
- Better for production

**Cons:**
- Additional tool to learn
- Requires separate schema files
- More setup initially

**Usage:**
```bash
atlas schema apply --env local
```

## Best Practices

1. **Version Control**: Always commit `db/schema.sql` changes
2. **Review Changes**: Use `atlas schema diff` before applying
3. **Dry Run**: Test with `--dry-run` in production
4. **Backups**: Always backup before applying schema changes
5. **Documentation**: Document significant schema changes
6. **Testing**: Test migrations in staging before production
7. **Rollback Plan**: Have a rollback strategy for production

## Troubleshooting

### Connection Issues

If you can't connect to the database:

```bash
# Check if PostgreSQL is running
pg_isready -h localhost -p 5432

# Check connection with psql
psql -h localhost -U shopiea -d shopiea
```

### Schema Mismatch

If schema doesn't match:

```bash
# Inspect current state
atlas schema inspect --env local

# See differences
atlas schema diff \
  --from "postgres://shopiea:mysecretpassword@localhost:5432/shopiea?sslmode=disable" \
  --to "file://db/schema.sql"

# Force apply (use with caution)
atlas schema apply --env local --auto-approve
```

### Permission Errors

Ensure your database user has appropriate permissions:

```sql
GRANT ALL PRIVILEGES ON DATABASE shopiea TO shopiea;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO shopiea;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO shopiea;
```

## Additional Resources

- [Atlas Documentation](https://atlasgo.io/docs)
- [Atlas CLI Reference](https://atlasgo.io/cli-reference)
- [PostgreSQL with Atlas](https://atlasgo.io/getting-started/postgres)
- [Shopiea GitHub Repository](https://github.com/Kyuubang/shopiea)
