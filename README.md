# Simple Logger (Go)
Implementation of a basic logger using the built in log tool.

# Usage
```Go
    // Create new instance
    // Logging levels:
    // - None
    // - Error
    // - Warning
    // - Info
    logger := simplelogger.NewSimpleLogger(simplelogger.Info)
    logger.Error("This is an error")
    logger.Warning("This is a warning")
    logger.Info("This is an information")
    
    // Set different logging levels
    logger.SetLoggingLevel(simplelogger.None)
    logger.SetLoggingLevel(simplelogger.Error)
    logger.SetLoggingLevel(simplelogger.Warning)
    logger.SetLoggingLevel(simplelogger.Info)
```

