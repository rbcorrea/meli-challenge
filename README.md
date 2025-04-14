# URL Shortener API

A URL shortener service that uses Redis for caching, MongoDB for persistence, and RabbitMQ for message queuing.

## Features

- Shorten URLs with custom codes
- List shortened URLs with pagination
- Search URLs by code
- Delete URLs (soft delete)
- Automatic caching in Redis
- Persistent storage in MongoDB
- Message queuing with RabbitMQ

## Architecture

- **API**: Handles HTTP requests and responses
- **Worker**: Processes messages from RabbitMQ to save URLs in MongoDB
- **Redis**: Caches shortened URLs for fast access
- **MongoDB**: Persists shortened URLs
- **RabbitMQ**: Queues messages for background processing

## API Endpoints

### Shorten URL
```bash
POST /urls
Content-Type: application/json

{
  "url": "https://example.com"
}
```

Response:
```json
{
  "original_url": "https://example.com",
  "short_url": "https://me.li/123e4567-e89b-12d3-a456-426614174000"
}
```

### List URLs
```bash
GET /urls?page=1&pageSize=10
```

Response:
```json
{
  "urls": [
    {
      "original_url": "https://example.com",
      "short_url": "https://me.li/123e4567-e89b-12d3-a456-426614174000"
    }
  ],
  "total": 1,
  "page": 1,
  "pageSize": 10
}
```

### Search URL by Code
```bash
GET /urls/:code
```

Response:
```json
{
  "original_url": "https://example.com",
  "short_url": "https://me.li/123e4567-e89b-12d3-a456-426614174000",
  "code": "123e4567-e89b-12d3-a456-426614174000",
  "created_at": "2024-03-14T12:00:00Z",
  "is_active": true
}
```

### Delete URL
```bash
DELETE /urls/:code
```

Response:
```json
{
  "original_url": "https://example.com",
  "short_url": "https://me.li/123e4567-e89b-12d3-a456-426614174000"
}
```

## Running the Application

1. Start the services:
```bash
docker-compose up -d
```

2. The API will be available at `http://localhost:8080`

