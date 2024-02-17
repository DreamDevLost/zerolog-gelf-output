# zerolog-gelf-output

The `zerolog-gelf-output` module is designed to seamlessly integrate with Go's popular [zerolog](https://github.com/rs/zerolog) logging library, enabling applications to redirect their JSON log outputs to a GELF-compatible server such as Graylog. This allows for more structured and centralized logging, making it easier to monitor and analyze logs in real-time.

## Features

- **Easy Integration**: Works out of the box with the zerolog logger.
- **GELF Support**: Fully compatible with the GELF (Graylog Extended Log Format) specification.
<!-- - **Network Resilience**: Implements retry and backoff strategies to handle network issues gracefully. -->
<!-- - **Flexible Configuration**: Supports various configurations including log level filtering, custom field mapping, and more. -->
<!-- - **Secure Transport**: Supports TLS encryption for secure log transmission. -->

## Getting Started

### Prerequisites

- Go 1.15 or later
- A GELF-compatible server (e.g., Graylog) accessible over the network

### Installation

To install `zerolog-gelf-output`, use the `go get` command:

```sh
go get github.com/DreamDevLost/zerolog-gelf-output
```

### Usage

Here is a basic example of how to use `zerolog-gelf-output` in your Go application:

```go
package main

import (
    "github.com/rs/zerolog"
    zerologgelfoutput "github.com/DreamDevLost/zerolog-gelf-output"
)

func main() {
    // Configure the GELF writer
    gelfWriter, err := zerologgelfoutput.NewWithPassthrough("udp://localhost:12201", "app-name", "environment", "v1.0.0", zerolog.ConsoleWriter{Out: os.Stdout})
    if err != nil {
        panic("Failed to create GELF writer: " + err.Error())
    }

    // Create a new zerolog logger with the GELF writer
    log := zerolog.New(gelfWriter).With().Timestamp().Logger()

    // Log a message as you normally would
    log.Info().Str("info1", "info1").Int("info2", 2).Msg("This is a test message sent to GELF server")
}
```

### Configuration Options

The `NewGelfWriter` function supports various options to customize the behavior of the GELF writer:

- **Network Protocol**: Choose between `udp` and `tcp` for log transmission.
- **GELF Server Address**: Specify the address of your GELF server in the format `host:port`.
- **TLS Configuration**: If using `tcp`, you can provide TLS configuration for secure communication.
- **Retry Strategy**: Configure the retry attempts and backoff strategy for handling network failures.

## Advanced Usage

For more advanced use cases, refer to the [documentation](#) (link to detailed documentation). This includes topics such as custom field mappings, filtering logs based on level, and integrating with existing logging setups.

## Contributing

We welcome contributions! If you'd like to contribute, please follow the guidelines outlined in [CONTRIBUTING.md](CONTRIBUTING.md).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
