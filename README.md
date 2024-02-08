# go-estate-market

This project is a web application designed for managing and browsing real estate listings. It consists of a backend web server written in GoLang and a frontend client developed using TypeScript and Angular.
## Technologies Used:

- **GoLang**: Backend server development.

- **TypeScript**: Frontend development language.

- **Angular**: Frontend framework for building robust single-page applications.

- **HTML/CSS**: Markup and styling for the frontend interface.

- **PostgreSQL**: Database management for storing real estate listings, user information, etc.

- **Prometheus**: Statistics about HTTP requests' duration and database calls.

## Setup Instructions

### Backend Setup

- Ensure you have GoLang installed on your system.

- Clone the repository and navigate to the project root directory.

- Navigate to the **/backend/src** directory where the backend code resides.

- Build the backend Docker image using the following command:
```
docker build . -t backend-image
```

- Once the image is built successfully, you can run the backend server as a Docker container. Make sure to configure the database connection in the .env file before running the container.

### Frontend Setup:

- Make sure you have Node.js and npm installed.

- Navigate to the **/frontend** directory.

- Build the frontend Docker image using the following command:
```
docker build . -t frontend-image
```

- After the image is built, you can run the frontend client as a Docker container. Ensure to modify the API base URL in the environment configuration file if necessary.

### Run Docker images

- Navigate to **/backend** and run the following command:
```
docker-compose up -d
```
