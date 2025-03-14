package dto

type ReserveSeatRequest struct {
	SeatNumber     int    `json:"seatNumber" binding:"required,min=1,max=300"`
	PassengerPhone string `json:"passengerPhone" binding:"required"`
	PassengerName  string `json:"passengerName" binding:"required"`
	PassengerAge   int    `json:"passengerAge" binding:"required,min=1"`
}

type ReserveSeatResponse struct {
	Message string `json:"message"`
}

type ResetSeatsResponse struct {
	Message string `json:"message"`
}

type Seat struct {
	SeatNumber int    `bson:"seatNumber"`
	Reserved   bool   `bson:"reserved"`
	Passenger  string `bson:"passenger,omitempty"`
	Phone      string `bson:"phone,omitempty"`
	Age        int    `bson:"age,omitempty"`
}
