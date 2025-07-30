# Traffic Splitter Demo with Envoy Proxy

This project demonstrates canary deployments using Envoy Proxy to split traffic between two versions of a microservice. It showcases how to gradually roll out a new version of a service by directing a percentage of traffic to it.

For the original problem statement and requirements, see [ProblemStatement.md](ProblemStatement.md).

## Overview

The demo consists of:

- **User Service v1**: Basic user service returning user information with version "v1"
- **User Service v2**: Enhanced user service with additional fields and version "v2"
- **Envoy Proxy**: Routes 70% of traffic to v1 and 30% to v2

## Prerequisites

- Docker and Docker Compose
- curl (for testing)
- jq (for formatting JSON output)

## Project Structure

```
.
├── ProblemStatement.md     # Original project requirements
├── README.md               # This documentation file
├── demo.sh                 # Demo script to show traffic splitting
├── docker-compose.yml      # Docker Compose configuration
├── envoy/
│   └── config.yaml         # Envoy proxy configuration
├── service/
│   ├── v1/                 # User Service v1
│   │   ├── Dockerfile
│   │   ├── go.mod
│   │   └── main.go
│   └── v2/                 # User Service v2
│       ├── Dockerfile
│       ├── go.mod
│       └── main.go
└── start.sh                # Helper script to start the services
```

## How to Run

1. **Clone the repository**:
   ```bash
   git clone <repository-url>
   cd trafficSplitter
   ```

2. **Start the services**:
   ```bash
   docker compose up --build
   ```

3. **Run the demo**:
   In a new terminal window, execute:
   ```bash
   chmod +x demo.sh
   ./demo.sh
   ```
   
   This will make 10 requests to the Envoy proxy and show the distribution of traffic between v1 and v2.

4. **Access Envoy's admin interface**:
   Open `http://localhost:9901` in your web browser to view Envoy's admin interface.

5. **Test manually**:
   ```bash
   # Make a request through Envoy
   curl http://localhost:8080/users/1 | jq
   
   # Direct request to v1
   curl http://localhost:8081/users/1 | jq
   
   # Direct request to v2
   curl http://localhost:8082/users/1 | jq
   ```

6. **Stop the services**:
   ```bash
   docker compose down
   ```

## Key Components

### User Services (v1 and v2)

- **v1**: Simple service returning basic user data
  - Endpoint: `/users/{id}`
  - Response: `{"id": 1, "name": "John", "version": "v1"}`

- **v2**: Enhanced service with additional fields
  - Endpoint: `/users/{id}`
  - Response: `{"id": 1, "name": "John Doe", "version": "v2", "email": "john@example.com"}`

### Envoy Proxy

- Routes 70% of traffic to v1 and 30% to v2
- Listens on port 8080
- Admin interface on port 9901
- Uses weighted clusters for traffic splitting

## Understanding the Configuration

### Traffic Splitting

The traffic splitting is configured in `envoy/config.yaml`:

```yaml
route:
  weighted_clusters:
    clusters:
      - name: service_v1
        weight: 70
      - name: service_v2
        weight: 30
```

This directs 70% of requests to the v1 service and 30% to the v2 service.

## Observability

To see detailed stats about traffic distribution:

```bash
curl http://localhost:9901/stats | grep "cluster.service_v" | grep "upstream_rq_completed"
```

This will show how many requests were processed by each service.

## Canary Deployment Benefits

This setup demonstrates several benefits of canary deployments:

1. **Risk Mitigation**: Test new versions with a small percentage of users
2. **Gradual Rollout**: Increase traffic to new versions incrementally
3. **Quick Rollback**: Easy to revert by adjusting weights if issues occur
4. **Real User Testing**: Collect feedback from real traffic patterns