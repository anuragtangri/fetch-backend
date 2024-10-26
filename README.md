# Receipt Processing Service

This is a Go-based web service that processes receipts to calculate and retrieve points based on specific criteria. The service exposes two main endpoints, one for processing a receipt and the other for retrieving the calculated points for a given receipt.

## Table of Contents

- [Features](#features)
- [Getting Started](#getting-started)
- [Endpoints](#endpoints)
- [Points Calculation](#points-calculation)
- [Example Usage](#example-usage)
- [Dependencies](#dependencies)
- [License](#license)

## Features

- **Process Receipts:** Accepts JSON data representing a receipt, calculates points based on defined rules, and returns a unique ID for the receipt.
- **Retrieve Points:** Returns the points associated with a specific receipt ID.

## Getting Started

### Prerequisites

Ensure you have Go installed. You can download it from [https://golang.org/dl/](https://golang.org/dl/).

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/receipt-processor.git
   cd receipt-processor
