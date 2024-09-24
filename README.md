# Technical Test 2

This is a technical test project that demonstrates the use of Go, Gin Framework, MySQL, GORM, JWT for authentication, Domain Driven Design, and hashed pin using bcrypt.

## Technologies Used

- **Go**: The programming language used for building the backend.
- **Gin Framework**: A web framework used to create the API endpoints.
- **MySQL**: The relational database used for storing user and transaction data.
- **GORM**: An ORM library used for interacting with the MySQL database in Go.
- **JWT for Authentication**: JSON Web Tokens are used to authenticate and authorize users securely.
- **Domain Driven Design (DDD)**: The project structure follows the principles of DDD to keep the business logic separated from the application logic.
- **Bcrypt for Hashed PINs**: User PINs are hashed using bcrypt for secure password storage.
- **Goroutines**: Used for handling background tasks such as asynchronous balance transfers.

## Features

- **User Registration**: New users can register with a hashed PIN.
- **User Login**: Users can log in using their phone number and PIN to receive JWT tokens.
- **User Profile Update**: Users can update their profile details such as first name, last name, and address.
- **Top-Up Balance**: Users can top up their balance and view their transaction history.
- **Payment**: Users can make payments, which will deduct from their balance.
- **Transfer Balance**: Users can transfer balance to another user, and both the sender and receiver's balances will be updated accordingly.
- **Transaction History**: Users can view their transaction history including top-ups, payments, and transfers.

## How It Works

1. **Authentication**:

   - When a user registers, their PIN is hashed using bcrypt and stored securely in the database.
   - After successful registration or login, a JWT token is generated which is used to authenticate the user for subsequent requests.

2. **Transactions**:

   - Users can top up their balance, make payments, and transfer money to other users.
   - Each transaction is recorded with details such as the amount, balance before, balance after, and the type of transaction (CREDIT or DEBIT).

3. **Database Migrations**:

   - The database schema is managed using GORM's auto-migration feature. Migrations are automatically applied when the application starts.

4. **Background Task for Transfer**:
   - The transfer money process is handled in the background to simulate a real-world transaction system where transfers are processed asynchronously by using Gorountines.

## Getting Started

### Prerequisites

- Go (version 1.16 or higher)
- MySQL
- Git

### Installation

1. **Clone the repository**:

   ```bash
   git clone https://github.com/alifmaulidanar/technical-test-2
   cd technical-test-2
   ```

2. **Set up the database**:

   - Create a MySQL database.
   - Update the `database/database.go` file with your database credentials.

3. **Run the application**:

   - Install Go dependencies:
     ```bash
     go mod tidy
     ```
   - Run the application:
     ```bash
     go run main.go
     ```

4. **API Endpoints**:
   - The application exposes the following endpoints:
     - `POST /register`: Register a new user.
     - `POST /login`: Login a user and receive JWT token.
     - `PUT /profile`: Update user profile (protected by JWT).
     - `POST /topup`: Top-up balance (protected by JWT).
     - `POST /pay`: Make a payment (protected by JWT).
     - `POST /transfer`: Transfer balance to another user (protected by JWT).
     - `GET /transactions`: Get transaction history (protected by JWT).

## Example Data

To populate the database with some example data for testing, you can use the provided `example_data.sql` file. Simply import this file into your MySQL database:

```bash
mysql -u [username] -p [database_name] < example_data.sql
```
