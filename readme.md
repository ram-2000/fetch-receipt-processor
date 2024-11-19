
# Receipt Service API

This project is a Receipt Processing API developed as part of a technical challenge. It implements an efficient backend in Go, integrates Memcached for in-memory key-value storage, and follows the principles of containerized development with Docker.

---

## Tech Stack Used

- **Programming Language:** Go (Golang)
- **In-Memory Storage:** Memcached
- **Containerization:** Docker, Docker Compose
- **Tools and Libraries:**
  - `github.com/bradfitz/gomemcache/memcache`: Memcached Go client
  - `github.com/google/uuid`: For generating unique receipt IDs
- **Logging:**
  - The application includes detailed logging for key operations like receipt processing, Memcached interactions, and error handling.

---

## API Overview

### 1. **Process Receipt**

- **Endpoint:** `/receipts/process`
- **Method:** `POST`
- **Description:** Accepts a JSON receipt, calculates points based on predefined rules, and stores the receipt in Memcached.

#### Request Body:
```json
{
  "retailer": "Target",
  "purchaseDate": "2022-01-01",
  "purchaseTime": "13:01",
  "items": [
    {
      "shortDescription": "Mountain Dew 12PK",
      "price": "6.49"
    }
  ],
  "total": "6.49"
}
```

#### Response:
```json
{
  "id": "c09aa3fe-69e7-47c7-bd68-97a376cd9043"
}
```

---

### 2. **Retrieve Points**

- **Endpoint:** `/receipts/{id}/points`
- **Method:** `GET`
- **Description:** Retrieves precomputed points for the receipt identified by the provided ID.

#### Response:
```json
{
  "points": 6
}
```

---

## How to Execute

### Prerequisites
- Docker
- Docker Compose

### Steps to Run

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd <repository-directory>
   ```

2. Build and start the application using Docker Compose:
   ```bash
   docker-compose up --build
   ```

3. The API server will start at `http://localhost:8080`.

---

## How to Test

### Using `curl`

1. **Process a Receipt**:
   ```bash
   curl -X POST -H "Content-Type: application/json"    -d '{
     "retailer": "Target",
     "purchaseDate": "2022-01-01",
     "purchaseTime": "13:01",
     "items": [
       {
         "shortDescription": "Mountain Dew 12PK",
         "price": "6.49"
       }
     ],
     "total": "6.49"
   }'    http://localhost:8080/receipts/process
   ```

2. **Retrieve Points**:
   ```bash
   curl http://localhost:8080/receipts/<id>/points
   ```

Replace `<id>` with the receipt ID returned from the `/receipts/process` response.

---

## Logs

Detailed logs are implemented for key application operations:

1. **Receipt Processing:**
   - Logs the successful processing and storage of receipts.
   - Logs any errors during data parsing or storage.

2. **Memcached Interactions:**
   - Logs successful connections and any errors during data retrieval or saving.

3. **Error Handling:**
   - Logs any unexpected errors with detailed information for debugging.

Logs are displayed in the container console during runtime for easy monitoring.

---

## Notes

- **Point Calculation Rules:**
  - One point for every alphanumeric character in the retailer name.
  - 50 points if the total is a round dollar amount with no cents.
  - 25 points if the total is a multiple of 0.25.
  - 5 points for every two items on the receipt.
  - If the item description length is a multiple of 3, multiply the price by 0.2, round up, and add the resulting points.
  - 6 points if the purchase day is odd.
  - 10 points if the purchase time is between 2:00 PM and 4:00 PM.

- **Data Persistence:**
  - Memcached is used for in-memory storage. Data will be lost if Memcached restarts.

- **Testing Tools:**
  - You can use tools like `Postman` or `curl` to test the API endpoints.

---

## License

This project is licensed under the MIT License.
