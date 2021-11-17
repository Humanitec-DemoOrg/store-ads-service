# ads-service

A minimal back-end REST API service mock.

## Configuration

This service supports the following environment variables:

| Name | Default | Description |
| ---- | ------- | ----------- |
| HOST | `""` | The `host` name to bind to. |
| PORT | `"8080"` | The `port` to listen on. |
| SERVICE_NAME | | The service name (for demo purposes). |
| SERVICE_SECRET | | The service secret (for demo purposes). |
| DATABASE_NAME | | The database name (for demo purposes). |

## Endpoints

This service exposes the following API endpoints:

| Method | Path | Description |
| ------ | ---- | ----------- |
| GET    | `/` | Responds with `HTTP 403 - See Other` to `GET /ready`. |
| GET    | `/alive` | The `liveness` probe. Responds with `HTTP 200 - OK`. |
| GET    | `/ready` | The `readiness` probe. Responds with `HTTP 200 - OK` and `application/json` content. |

The `readiness` probe request/response example:

```
GET /ready HTTP/1.0
---

HTTP/1.1 200 OK
Content-Type: application/json

{
  "name": "ads-service",
  "secret": "U**********f",
  "database": "ads-service-db"
}
```
