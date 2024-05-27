# URL Shortener

This repository represents a simple url shortener.

## Init

```sh
  docker compose up --build
```

Initial port and host - __80__, __localhost__

## Functional requirements:
  - Create a link from an inputed link
  - Redirect requests from server to origin link

# Structure

  - Service represents clean architecture with layers as repository, controllers, service, entity

# Stack
  - MongoDB
  - Redis
  - Gin
  - Docker
  - Golang

## API Reference

#### To short the link

```http
  POST /shorten
```

| Parameter  | Type     | Description                        |
| :--------- | :------- | :--------------------------------- |
| `url`    | `string` | **Required**. Origin ling    |


Return `short_url`

#### Redirect to origin link

```http
  GET /:url
```
