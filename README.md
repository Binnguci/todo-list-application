# Todo List Application

A simple yet powerful **Todo List Application** built with Golang, GORM, and OAuth2 support. This application helps users manage their tasks efficiently by allowing them to create, organize, and track tasks with ease.

## Features

- **User Authentication**: Secure user authentication with OAuth2 (Google) and traditional username/password methods.
- **Task Management**: Create, view, update, and delete tasks.
- **Task Categories and Tags**: Organize tasks into categories and use tags for better management.
- **Task Prioritization and Deadlines**: Set priority levels and deadlines for tasks.
- **Responsive API**: RESTful API for easy integration with frontend or mobile apps.
- **Notifications**: Alerts for upcoming deadlines or important tasks.

## Tech Stack

- **Backend**: Golang with [Gin framework](https://gin-gonic.com/) and [GORM](https://gorm.io/)
- **Database**: MySQL or PostgreSQL (can be configured)
- **Authentication**: OAuth2 and JWT-based authentication
- **Containerization**: Docker for containerized deployment
- **Unit Testing**: Go testing framework

## Prerequisites

Make sure you have the following installed:

- [Go 1.23+](https://golang.org/doc/install)
- [Docker](https://www.docker.com/get-started)
- [MySQL](https://dev.mysql.com/downloads/mysql/) or [PostgreSQL](https://www.postgresql.org/download/)

## Getting Started

### Clone the repository

```bash
git clone https://github.com/your-username/todo-list-application.git
cd todo-list-application
