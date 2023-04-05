# api

API Server for Zeta Library build with GoFiber (GoLang)

## Development

> Please the development server from the root directory.

## Improvements

- Use group routing for easier management

  Example:

  ```go
  api := fiber.New()
  collectionsApi := api.Group("/collections")
  ```
