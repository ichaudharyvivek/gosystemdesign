```sh
mylogger/
│
├── logger.go          // Logger (core entry point)
├── config.go          // Config struct
├── level.go           // LogLevel enum
├── event.go           // LogMessage / LogEvent
│
├── appender/
│   ├── appender.go    // Appender interface
│   ├── console.go     // ConsoleAppender
│   ├── file.go        // FileAppender
│
├── formatter/
│   ├── formatter.go   // Formatter interface
│   ├── json.go        // JSONFormatter
│   ├── text.go        // TextFormatter
│
├── async/
│   ├── dispatcher.go  // AsyncDispatcher / worker / queue
│
└── internal/          // (optional, if you want to hide internals)
    └── buffer.go      // ring buffer / queue impl (if custom)
```
