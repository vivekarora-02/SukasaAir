Sukasa Air Seat Reservation API

🚀 Overview

Sukasa Air is launching a market validation MVP for affordable airline tickets with high-quality facilities. This API provides authentication and seat reservation functionalities, ensuring reliable and high-demand handling.

📌 Features

User Authentication: Login with email to receive a session token.

Seat Reservation: Reserve seats (1-300) with passenger details.

Reset Reservations: Admin-only endpoint to reset all seat reservations.

High Availability: Ensures consistency in seat booking confirmations.

MongoDB Integration: Uses MongoDB as the primary database.

📜 API Endpoints

🔑 Authentication

POST /login

Request Body:

{
  "emailId": "user@example.com"
}

Response:

{
  "token": "session_token_here"
}

✈️ Seat Reservation

POST /seat/reserve

Request Body:

{
  "seatNumber": 25,
  "passengerPhone": "1234567890",
  "passengerName": "John Doe",
  "passengerAge": 30
}

Response:

{
  "status": "success",
  "message": "Seat 25 reserved successfully."
}

🛑 Reset Reservations (Admin Only)

POST /seat/reset

Request Body: None
Response:

{
  "status": "success",
  "message": "All seat reservations reset successfully."
}

⚙️ Tech Stack

Golang (Backend API)

MongoDB (Database)

Redis (Session management & caching)

Gin (Web framework for Golang)

Swagger (API Documentation)

🏗️ Setup & Installation

1️⃣ Clone the repository

git clone https://github.com/yourusername/SukasaAir.git
cd sukasa-air

2️⃣ Install dependencies

go mod tidy

3️⃣ Run the server

go run main.go

Server runs at http://localhost:8080

🛠️ Running Tests & Coverage

go test ./... -coverprofile=coverage.out

To generate a coverage report:

go tool cover -html=coverage.out -o coverage.html

📜 API Documentation

Swagger documentation is available at:

http://localhost:8080/swagger/index.html

✅ Project Constraints

Minimum 80% test coverage

Modular & maintainable code for scalability

Ensures booking consistency under high demand

Uses MongoDB for data storage

📌 Version Control & Contribution

Follow best practices for commits (git commit -m "feat: added seat reservation")

Use GitHub Issues & PRs for contributions

📜 License

MIT License

🛫 Sukasa Air - Book with confidence!

# Sukasa-Air
