# Project Management & Collaboration API

This project is a **Project Management & Collaboration API** built in Go.
The API enables users to manage projects, assign tasks, add team members, and track progress with features for user authentication, authorization, and role-based permissions.

## Tech
1. Golang
2. Gorilla Mux for handling requests
3. PostgreSQL database
4. Gorm ORM for interacting with database

## Features
1. **User Authentication and Authorization**  
   - JWT-based authentication for secure login and registration.
   - Role-based access control (Admin and Member roles).
   - Only admins can manage projects, while members can only view and update assigned tasks.

2. **Project and Task Management**  
   - CRUD operations for projects and tasks.
   - Nested routes to handle tasks within projects.
   - Role-based permissions for different operations.

3. **Team Collaboration**  
   - Add and remove team members from projects.
   - Assign tasks to team members and track status (To Do, In Progress, Done).

4. **Advanced Routing and Middleware**  
   - Custom middleware for JWT authentication and logging.

5. **Data Persistence**  
   - PostgreSQL as the database to manage users, projects, tasks, and team members.

---

## Sample Env file
- Create a `.env` file in the root directory and specify your database credentials:
```plaintext
   PORT=....

   JWT_SECRET_KEY=....

   DB_NAME=....
   DB_USERNAME=....
   DB_PASSWORD=....
   DB_PORT=....
   DB_HOST=....
   DB_SSL_MODE=....
```

