## Qube FE Project - Backend
This project constitutes the backend implementation for the Qube FE (Frontend) Project.

## Functionality
- GET /api/v1/appliances: Retrieves a list of appliances based on optional query parameters deviceStatus and downloadStatus. (Mock response)
- GET /api/v1/appliance/{appliance-id}/info: Retrieves detailed information about a specific appliance identified by its ID (appliance-id). (Mock response)

> **_NOTE:_**  The server will start locally at http://localhost:8080.

## Setup and Deployment
1. Clone the repository.
2. Ensure Go is installed on your system.
3. Install dependencies using Go modules.
4. Run main.go to start the server.

```bash
go run main.go
```