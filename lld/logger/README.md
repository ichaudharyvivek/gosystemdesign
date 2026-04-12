````sh
# Logger Package

A lightweight logging framework for Go designed with clean separation of concerns and extensibility in mind.

## Functional Requirements

1. Support standard log levels:
   - `DEBUG`
   - `INFO`
   - `WARN`
   - `ERROR`
   - `FATAL`
2. Filter log messages based on a configurable minimum log level.
3. Support multiple output destinations (appenders), including console and file.
4. Allow a single log message to be sent to multiple appenders simultaneously.
5. Support asynchronous logging to prevent blocking the main application thread.
6. Allow client applications to configure the logger by specifying log level, formatters, and appenders.

## Non-Functional Requirements

- **Thread Safety:** Logging is safe in concurrent environments to prevent interleaved or lost messages.
- **Performance:** The logger is designed to have minimal overhead on application performance.
- **Extensibility:** The architecture supports plugging in custom formatters, filters, and appenders with minimal code changes.
- **Maintainability:** The codebase follows a clean design with clear separation of concerns between logger, formatter, appender, and internal models.
- **Ease of Use:** The client-facing API is simple and intuitive for developers.

## Package Architecture

The logger package is organized into core components and implementation details:

- `internal/logger` - core logger entry points and log entry creation.
- `internal/appender` - appender interfaces and concrete output destinations.
- `internal/formatter` - formatter interfaces and message formatting implementations.
- `internal/model` - log level and record model definitions.
- `cmd` - sample executable entry point.

### File Structure

```text
logger/
├── README.md
├── cmd/
│   └── main.go
├── go.mod
├── go.sum
└── internal/
    ├── appender/
    │   ├── appender.go
    │   ├── console.go
    │   └── file.go
    ├── formatter/
    │   ├── formatter.go
    │   ├── json.go
    │   └── text.go
    ├── logger/
    │   ├── entry.go
    │   └── logger.go
    └── model/
        ├── level.go
        └── record.go
````

## Class Diagram

```mermaid
classDiagram
    class Logger {
        - model.Level level
        - []appender.Appender appenders
        + SetLevel(level model.Level)
        + SetAppenders([]appender.Appender)
        + Debug() *Entry
        + Info() *Entry
        + Warn() *Entry
        + Error() *Entry
        + Fatal() *Entry
        + Close() error
        - log(level model.Level) *Entry
    }

    class Entry {
        - logger *Logger
        - level model.Level
        + Msg(msg string)
        + Msgf(format string, args ...interface{})
    }

    class Appender {
        <<interface>>
        + Append(record model.Record) error
        + Close() error
    }

    class ConsoleAppender {
        - io.Writer output
        - formatter.Formatter formatter
        + Append(record model.Record) error
        + Close() error
    }

    class FileAppender {
        - string path
        - *os.File file
        - formatter.Formatter formatter
        + Append(record model.Record) error
        + Close() error
    }

    class Formatter {
        <<interface>>
        + Format(record model.Record) string
    }

    class TextFormatter {
        + Format(record model.Record) string
    }

    class JSONFormatter {
        + Format(record model.Record) string
    }

    Logger "1" --> "*" Entry : creates
    Logger "1" o-- "*" Appender : uses
    Appender <|.. ConsoleAppender
    Appender <|.. FileAppender
    Formatter <|.. TextFormatter
    Formatter <|.. JSONFormatter
    ConsoleAppender o-- Formatter : formats
    FileAppender o-- Formatter : formats
```