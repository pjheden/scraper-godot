# Godot asset store

## TODOs

- Scraper
- Setup infra
- Web server
- DB layer

```mermaid
scraper --> db layer
web server <-- db layer
```

## How to run

### Scraper

```
go run cmd/scraper/main.go
```

## Generate templates

```
templ generate ./...
```
