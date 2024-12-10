# receipts_api

## Receipt Processor

This is a Go-based web application designed to process receipts. The application uses the Gin web framework to handle HTTP requests and supports operations like generating receipts, processing transactions, and more.
#Table of Contents

- [Installation](#installation)
- [Usage] (#usage)
- [Testing] (#testing)

## Installation

Follow these steps to get your local development environment set up:
Clone the repository:

git clone https://github.com/yourusername/receipt-processor.git
cd receipt-processor

## Install dependencies:

Ensure you have Go installed on your machine (version 1.18 or above). Then, run:

go mod tidy

## Run the application:

To start the server, run:

go run main.go
The server will start on port 8080 by default. You can customize this by changing the configuration in the main.go file.

## Usage

Once the server is running, you can interact with it via HTTP requests. Here the basic endpoints:
Endpoints
POST /receipts/process takes in the receipt information to calculate the points
GET /receipt/:id Retrieves a receipt by its ID.

These endpoints can be tested using Postman or curl.

## Example Request

POST /receipt
{
"transaction_id": "12345",
"amount": 50.75,
"date": "2024-12-10",
"items": [
{"name": "Item 1", "price": 20.00},
{"name": "Item 2", "price": 30.75}
]
}

## Testing

To run tests, use the following command:
go test ./...

This will execute all unit tests defined in the project.

## API Documentation

If you're interested in detailed API documentation, it can be generated using GoDoc or other tools for API documentation.
