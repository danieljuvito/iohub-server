# Backend Service for Real-Time Stories (Golang + MongoDB)

This repository contains the backend service for a real-time stories application. Users can create and view stories, and
the stories automatically disappear after 24 hours. We've implemented user authentication, real-time updates, and
concurrency handling.

## Table of Contents

1. [Project Overview](#project-overview)
2. [Prerequisites](#prerequisites)
3. [Installation](#installation)
4. [Configuration](#configuration)
5. [API Endpoints](#api-endpoints)
6. [WebSocket](#websocket)
7. [Concurrency Handling](#concurrency-handling)
8. [Deployment](#deployment)
9. [Contributing](#contributing)
10. [License](#license)

## Project Overview

This backend service is built using the following technologies:

- **Go (Golang)**: We use the Echo framework for building RESTful APIs.
- **MongoDB**: Our database for storing user data and stories.
- **JWT (JSON Web Tokens)**: For user authentication.
- **WebSockets**: To achieve real-time updates.

## Prerequisites

Make sure you have the following installed:

- Go (latest version)
- MongoDB
- Git

## Installation

1. Clone this repository:

   ```bash
   git clone https://github.com/your-username/real-time-stories-backend.git
   cd real-time-stories-backend
   ```

2. Install Go dependencies:

   ```bash
   go mod download
   ```

3. Set up your MongoDB database and configure the connection in `config/config.go`.

## Configuration

Edit the `config/config.go` file to set your configuration options:

- MongoDB connection details
- JWT secret key
- Other environment-specific settings

## API Endpoints

- **User Registration**:
    - `POST /api/register`: Register a new user with email and password.
- **User Login**:
    - `POST /api/login`: Authenticate a user and return a JWT token.
- **Create Story**:
    - `POST /api/stories`: Create a new story.
- **Get Stories**:
    - `GET /api/stories`: Retrieve stories (visible for 24 hours).
- **Logout**:
    - `POST /api/logout`: Invalidate the user's token.

## WebSocket

We use WebSockets for real-time updates. Connect to the WebSocket endpoint (`/ws`) to receive new stories as they are
created.

## Concurrency Handling

We handle concurrent requests using MongoDB transactions. Ensure that your MongoDB server supports transactions.

## Deployment

You can deploy this application to platforms like Heroku, Vercel, or any cloud service provider. Don't forget to set
environment variables for production (e.g., JWT secret, database connection).

## Contributing

Contributions are welcome! Fork this repository, make your changes, and submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.