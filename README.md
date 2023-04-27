# Notifications System

This project is a simple web application that displays a list of notifications. It consists of a React frontend and a Golang backend with a PostgreSQL database for persistent storage.

## Project Structure

- `/frontend`: Contains the frontend React application
- `/backend`: Contains the backend Golang application

## Dependencies

- Docker
- Docker Compose

Optional (for manual setup without Docker):

- Node.js and npm (for the frontend)
- Go (for the backend)
- [Migrate CLI](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) (for the backend)

## Quick Start

1. Clone the repository:
```git clone https://github.com/broem/notification-system.git```

2. Change into the project directory:
```cd notification-system```

3. Build and run the project with Docker Compose:
```docker-compose up --build``` (add `-d` to run in detached mode)

4. Open the frontend in your browser at http://localhost:8081

## Frontend

The frontend is built using React and TypeScript.

### Development

To run the frontend in development mode, first ensure you have Node.js and npm installed. Then, follow these steps:

1. Change to the frontend directory:
```cd frontend```

2. Install dependencies:
```npm install```

3. Start the development server:
```npm start```

The frontend will be available at http://localhost:3000

### Building for Production

To build the frontend for production, first ensure you have Node.js and npm installed. Then, follow these steps:

1. Change to the frontend directory:
```cd frontend```

2. Install dependencies:
```npm install```

3. Build the frontend:
```npm run build```

The frontend will be built to the `build` directory.

## Backend

The backend is built using Golang, Echo, and PostgreSQL.

### Development

To run the backend in development mode, first ensure you have Go and Docker installed. Then, follow these steps:

1. Change to the backend directory:
```cd backend```

2. Start the database:
```docker-compose up -d db```

or 

```docker run --name notification-system-db -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres```

3. Create a `.env` file using the `.sample.env` file as a template:
```cp .sample.env .env``` (or `copy .sample.env .env` on Windows)

4. Install dependencies:
```go mod download```
```go build -o backend ./main.go```

5. Run the backend:
```./backend``` (or `backend.exe` on Windows)

6. The backend API will be accessible at [http://localhost:8080](http://localhost:8080) or at whatever port you specified in the `.env` file.

