# Microservices with Traffic Splitting

## Scenario
Build a simple user service with two versions to demonstrate canary deployment using Envoy's traffic splitting.

## What to Build

### User Service v1 (Go)
- GET `/users/{id}` - returns user info with field:
  ```json
  {"id": 1, "name": "John", "version": "v1"}
  ```

### User Service v2 (Go)
- Same endpoint but returns:
  ```json
  {"id": 1, "name": "John Doe", "version": "v2", "email": "john@example.com"}
  ```

### Envoy Proxy
- Routes 70% traffic to v1, 30% to v2
- Both services run on different ports (8081, 8082)
- Envoy exposes unified endpoint on port 8080

## Key Envoy Features Demonstrated
- **Traffic Splitting**: Weighted routing between service versions
- **Health Checking**: Envoy monitors both service instances
- **Load Balancing**: Round-robin within each version
- **Observability**: View traffic distribution via Envoy admin interface

## Demo
- Make 10 requests to localhost:8080/users/1
- Show ~7 responses from v1, ~3 from v2
- Show Envoy admin stats at localhost:9901

## Why This Works Well
- Single binary (Envoy), no Kubernetes needed
- Clear visual demonstration of traffic splitting
- Easy to explain canary deployment concept
- Shows practical production use case
