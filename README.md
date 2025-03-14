# Sukasa Air Seat Reservation System

## 📌 Overview  
Sukasa Air aims to validate the market value of cheap air tickets with high-quality facilities by developing a scalable reservation system. This application provides the following key functionalities:  
- **User authentication** (`/login`)  
- **Seat reservation** (`/seat/reserve`)  
- **Seat reset (Admin only)** (`/seat/reset`)  

## 📑 API Endpoints

| Endpoint        | Method | Input Body  | Response |
|---------------|--------|------------|---------|
| `/login`      | POST   | `{ "emailId": "user@example.com" }` | `{ "token": "session_token" }` |
| `/seat/reserve` | POST  | `{ "seatNumber": 5, "passengerPhone": "1234567890", "passengerName": "John Doe", "passengerAge": 25 }` | `{ "success": "Seat reserved successfully" }` |
| `/seat/reset` | POST   | `NA` | `{ "success": "All seat reservations reset successfully" }` (Admin Only) |

## 🏗️ Tech Stack  
- **Golang** (Gin Framework)  
- **MongoDB** (Seat reservations storage)  
- **Redis** (Session management)  
- **JWT Authentication**  

## 🔧 Setup & Installation  

### Prerequisites  
- Go 1.18+  
- MongoDB & Redis running locally  

### Steps  
1. Clone the repository:  
   ```sh  
   git clone https://github.com/yourusername/sukasa-air.git  
   cd sukasa-air  
   ```  
2. Install dependencies:  
   ```sh  
   go mod tidy  
   ```  
3. Start MongoDB & Redis (if not running):  
   ```sh  
   brew services start mongodb-community@7.0
   brew services start redis
   ```  
4. Run the server:  
   ```sh  
   go run main.go  
   ```  


## 📜 API Documentation  
Swagger docs available at:  
```
http://localhost:8080/swagger/index.html  
```  
