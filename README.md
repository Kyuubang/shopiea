# Shopiea

Shopiea is a RESTful API written in Go that provides scoring system functionalities for managing and scoring student assignments, authentication and registering lab and course, and generating reports. It currently does not have a dashboard UI, but provides an OpenAPI 3.0 specification for anyone to develop a dashboard.

## Features

- Authentication and authorization for users and administrators
- Registration of labs and courses
- Scoring of student assignments
- Generating reports for lab and course scores


## Installation

### Prerequisites

- Go 1.15 or higher
- PostgreSQL 10 or higher

Clone the repository: 
```bash
git clone https://github.com/Kyuubang/shopiea.git
cd shopiea
```

Install dependencies
```bash
go mod download
```

Run the database migrations
```bash
go run main.go -migrate
```

Source `.env` file
```bash
set -a; source .env; set +a
```

Run the server
```bash
go run main.go
```

or with make command
```bash
make run
```

## License

This project is licensed under the MIT License. See the LICENSE file for details.