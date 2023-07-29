# Golang Task Management System

A robust containerized task management system built with Golang, Gin, and GORM. This system handles user authentication, authorization, and access management, allowing users to create, update, and manage tasks securely.

## Key Features

- Secure user registration and authentication using JWT (JSON Web Tokens).
- Protection against vulnerabilities like SQL injection attacks.
- Support for bulk upload of tasks using CSV.

## Prerequisites

Make sure you have the following installed:

- Docker and Docker Compose
- Golang

## Getting Started

1. Clone the repository:

```bash
git clone https://github.com/whoiaditya/golang-task-management-system.git
cd golang-task-management-system
```

2. Set up the environment variables:

Copy the `.env.example` file to `.env` and update the values with your own configuration:

```bash
cp .env.example .env
```

3. Build and run the Docker containers:

```bash
docker-compose up -d
```

## API Endpoints

The following API endpoints are available:

- **POST /user/register**: Register a new user.
- **POST /user/login**: Log in with registered user credentials and receive a JWT token.
- **PUT /user/logout**: Log out and invalidate the JWT token.
- **DELETE /user/delete**: Invalidate the JWT token and Delete User.
  
- **GET /tasks/:id**: Get a single task by ID.
- **GET /tasks/**: Get all tasks.
- **POST /tasks/create**: Create a new task.
- **POST /tasks/bulkupload**: Create a new tasks using bulk upload (explained below).
- **PUT /tasks/update/:id**: Update an existing task by ID.
- **DELETE /tasks/delete/:id**: Delete a task by ID.

## Bulk Upload

To bulk upload tasks from a CSV file, create a CSV file named `data.csv` with key: `taskBulkUpload` with the following format:

```
Title,Description,Planned Start Date,Planned Start Time,Planned End Date,Planned End Time,Seconds
Task 1,Description 1,2023-07-20,09:00:00,2023-07-20,18:00:00,120000
Task 2,Description 2,2023-07-21,09:30:00,2023-07-21,17:30:00,150000
Task 3,Description 3,2023-07-22,10:00:00,2023-07-22,16:00:00,200000
Task 4,Description 4,2023-07-23,10:30:00,2023-07-23,15:30:00,250000
Task 5,Description 5,2023-07-24,11:00:00,2023-07-24,15:00:00,300000
Task 6,Description 6,2023-07-25,11:30:00,2023-07-25,14:30:00,350000
Task 7,Description 7,2023-07-26,12:00:00,2023-07-26,14:00:00,400000
Task 8,Description 8,2023-07-27,12:30:00,2023-07-27,13:30:00,450000
Task 9,Description 9,2023-07-28,13:00:00,2023-07-28,13:00:00,500000
Task 10,Description 10,2023-07-29,13:30:00,2023-07-29,12:30:00,550000
```
## Documentation
Refer to the [postman documentation](https://documenter.getpostman.com/view/16151723/2s946mZ9Zr)

<p align="center">
	With :heart: by <a href="https://github.com/whoisaditya" target="_blank">Aditya Mitra</a>
</p>
