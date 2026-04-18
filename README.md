```
urlshortener/
├── cmd/
│   └── server/
│       ├── main.go
│       └── handlers/        # HTTP/gRPC entry points live here
│
├── internal/
│   ├── user/                # feature package — owns its own domain
│   │   ├── user.go          # model
│   │   ├── service.go       # business logic
│   │   └── repository.go    # interface + impl
│   │
│   ├── url/
│   │   ├── url.go
│   │   ├── service.go
│   │   └── repository.go
│   │
│   └── platform/            # enablers — db, cache, pubsub, logger
│       ├── postgres/
│       ├── redis/
│       └── pubsub/
│
└── pkg/                     # only if something is truly reusable externally
    └── shortcode/
```

```
urlshortener/
├── cmd/
│   ├── server/
│   │   ├── main.go
│   │   └── handlers/       # HTTP handlers
│   └── cli/
│       ├── main.go
│       └── commands/       # CLI commands (cobra, etc.)
│
├── internal/
│   ├── user/
│   ├── url/
│   └── platform/
│
└── pkg/
```
