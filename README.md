# CRUD App with Go Gin Framework

This is a simple CRUD (Create, Read, Update, Delete) application built using the Go programming language and the Gin web framework. The application manages users, roles, and the relationship between users and roles.

## Features

- **User Management**: Create, read, update, and delete users.
- **Role Management**: Manage roles using SQL scripts.
- **User-Role Association**: Assign roles to users and manage these associations using JSON files.

## Installation

1. **Clone the repository**:

   ```bash
   git clone https://github.com/your-username/crud_app.git
   cd crud_app
   ```

2. **Install dependencies**:

   Ensure you have [Go](https://golang.org/doc/install) installed, then run:

   ```bash
   go mod tidy
   ```

3. **Setup the Database**:

    - Create a PostgreSQL database and update the connection string in your configuration file (e.g., `config.yml`).

4. **Run the application**:

   ```bash
   go run main.go
   ```

## Project Structure

- `main.go`: The entry point of the application.
- `models/`: Contains the data models for Users, Roles, and UserRoles.
- `controllers/`: Contains the HTTP handlers for the CRUD operations.
- `repositories/`: Contains the database interaction logic.
- `config/`: Contains configuration files.
- `initDb/`: Contains SQL scripts for initializing the database.

## Database Initialization

The application uses three methods to initialize the database:

1. **Roles Table**:
    - The roles table is initialized using an SQL script file.
    - The script file is read, and the SQL commands are executed to insert data into the roles table.

2. **Users Table**:
    - Users are initialized using JSON data.
    - The JSON data is read and processed to insert user information into the users table.

3. **User-Roles Table**:
    - The user_roles table is initialized using a JSON file.
    - The JSON file is read, and the associations between users and roles are inserted into the user_roles table.

## Endpoints

- **User Endpoints**:
    - `POST /users`: Create a new user.
    - `GET /users/:id`: Retrieve a user by ID.
    - `PUT /users/:id`: Update a user by ID.
    - `DELETE /users/:id`: Delete a user by ID.

- **Role Endpoints**:
    - `POST /roles`: Create a new role using SQL script.
    - `GET /roles/:id`: Retrieve a role by ID.
    - `PUT /roles/:id`: Update a role by ID.
    - `DELETE /roles/:id`: Delete a role by ID.

- **User-Role Endpoints**:
    - `POST /user_roles`: Assign a role to a user using JSON data.
    - `GET /user_roles/:user_id`: Retrieve roles for a user by user ID.
    - `DELETE /user_roles/:user_id/:role_id`: Remove a role from a user.

## Sample Data

### Roles SQL Script

```sql
INSERT INTO roles (name) VALUES ('Admin');
INSERT INTO roles (name) VALUES ('User');
INSERT INTO roles (name) VALUES ('Guest');
```

### Users JSON

```json
[
  {
    "firstName": "John",
    "lastName": "Doe",
    "email": "john.doe@example.com",
    "tel": "1234567890",
    "roleId": 1
  },
  {
    "firstName": "Jane",
    "lastName": "Doe",
    "email": "jane.doe@example.com",
    "tel": "0987654321",
    "roleId": 2
  }
]
```

### User-Roles JSON

```json
[
  {
    "userId": 1,
    "roleId": 1
  },
  {
    "userId": 2,
    "roleId": 2
  }
]
```

## Usage

1. **Initialize Roles**:

    - Place the roles SQL script in the `InitDb/` directory.
    - The application reads this script at startup and initializes the roles table.

2. **Initialize Users**:

    - Place the users JSON data in a designated directory.
    - The application reads this JSON data and initializes the users table.

3. **Initialize User-Roles**:

    - Place the user-roles JSON file in a designated directory.
    - The application reads this JSON file and initializes the user_roles table.

## Contributing

Feel free to submit issues, fork the repository, and send pull requests. For major changes, please open an issue first to discuss what you would like to change.

## License

This project is licensed under the MIT License.
ile provides a comprehensive overview of your CRUD application, detailing the features, installation steps, project structure, database initialization methods, endpoints, sample data, and usage instructions.