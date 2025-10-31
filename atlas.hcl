// Define the environment for local development
env "local" {
  // Declare where the schema definition resides.
  src = "file://db/schema.sql"
  
  // Database URL for local development
  url = "postgres://shopiea:mysecretpassword@172.17.0.2:5432/shopiea?sslmode=disable"
  
  // Migrations directory
  dev = "docker://postgres/15/dev"
  
  // Define the schema inspection behavior
  schemas = ["public"]
}

// Define the environment for production
env "prod" {
  src = "file://db/schema.sql"
  
  // Load database URL from environment variable
  url = getenv("DATABASE_URL")
  
  dev = "docker://postgres/15/dev"
  
  schemas = ["public"]
}
