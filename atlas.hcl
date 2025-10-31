// Define the environment for local development
env "local" {
  // Declare where the schema definition resides.
  src = "file://db/schema.sql"
  
  // Database URL for local development
  // Override with: atlas --env local --var url="postgres://user:pass@host:5432/db"
  url = getenv("SHOPIEA_DB_URL")
  
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
