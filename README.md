# Payment Microservice

This repository contains a payment microservice developed in Go, using PostgreSQL as the database and PgAdmin as the management tool. The development environment is managed with Docker and Docker Compose, including support for hot reload with the `air` tool.

## **Requirements**
Before starting, make sure you have installed on your machine:
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

## **Services**
The project uses the following Docker services:

1. **PostgreSQL (`postgres`)**
   - Database used by the payment microservice.
   - Runs on port `5432`.
   - Data persists in the `postgres_data` volume.

2. **PgAdmin (`pgadmin`)**
   - Graphical interface to manage the PostgreSQL database.
   - Accessible via browser at `http://localhost:5050`.
   - Default credentials:
     - **Email:** `admin@admin.com`
     - **Password:** `123456`

3. **Go Application (`app`)**
   - Implements the business logic of the payment microservice.
   - Uses `air` for hot reload.
   - Runs on port `8080`.

## **How to Start the System**
Follow the steps below to set up and run the system locally.

### **1. Clone the Repository**
```sh
git clone https://github.com/your-user/payment-ms.git
cd payment-ms
```

### **2. Create and Configure the `.env` File**
Create a `.env` file in the project's root directory with the database credentials:
```
POSTGRES_USER=eralves
POSTGRES_PASSWORD=123456
POSTGRES_DB=payment_db
PGADMIN_DEFAULT_EMAIL=eralves01@gmail.com
PGADMIN_DEFAULT_PASSWORD=123456
```

### **3. Start the Containers**
```sh
docker-compose up --build
```
The `--build` flag ensures the application image is properly rebuilt.

### **4. Access the Services**
- **Database:**
  - Host: `localhost`
  - Port: `5432`
  - User: `eralves`
  - Password: `123456`
  - Database: `payment_db`

- **PgAdmin:**
  - URL: [http://localhost:5050](http://localhost:5050)
  - Email: `eralves01@gmail.com`
  - Password: `123456`

- **Go API:**
  - Running at `http://localhost:8080`

### **5. Stop the Containers**
To stop the services, press `CTRL+C` or run:
```sh
docker-compose down
```

### **6. Run in Background (Daemon Mode)**
If you want to run the containers in the background, use:
```sh
docker-compose up -d
```

## **Hot Reload with Air**
The Go application uses `air` to automatically reload when code changes. If you want to run the application outside of Docker, install `air` locally:
```sh
go install github.com/air-verse/air@latest
```
Then, start the server with:
```sh
air
```

## **Tips**
- If you encounter errors while running the containers, try removing the volumes:
  ```sh
  docker-compose down -v
  ```
- To view container logs:
  ```sh
  docker-compose logs -f
  ```

Now you can run the system locally for development! 🚀

