// log_appender
package logger

// Start and Stop methods allow appenders to manage resources (e.g., open/close files, DB connections).
// Useful for setup/teardown when logger config changes or during graceful shutdown.
type Appender interface {
	Name() string
	Append(message *LogMessage) error
	// Start() error
	// Stop() error
}
