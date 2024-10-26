# Receipt Processing Service

This is a Go-based web service that processes receipts to calculate and retrieve points based on specific criteria. The service exposes two main endpoints, one for processing a receipt and the other for retrieving the calculated points for a given receipt.

## Features

- **Process Receipts:** Accepts JSON data representing a receipt, calculates points based on defined rules, and returns a unique ID for the receipt.
- **Retrieve Points:** Returns the points associated with a specific receipt ID.

Dependencies
Gorilla Mux - for HTTP request routing
Google UUID - for generating unique IDs

Install these dependencies using go mod tidy.
