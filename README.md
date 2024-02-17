# zerolog-gelf-output

The `zerolog-gelf-output` module is designed to seamlessly integrate with Go's popular [zerolog](https://github.com/rs/zerolog) logging library, enabling applications to redirect their JSON log outputs to a GELF-compatible server such as Graylog. This allows for more structured and centralized logging, making it easier to monitor and analyze logs in real-time.

## Features

- **Easy Integration**: Works out of the box with the zerolog logger.
- **GELF Support**: Fully compatible with the GELF (Graylog Extended Log Format) specification.
- **Flexible Configuration**: Supports various configurations including log level filtering, custom field mapping, and more.
- **Network Resilience**: Implements retry and backoff strategies to handle network issues gracefully.
- **Secure Transport**: Supports TLS encryption for secure log transmission.

## Getting Started

### Prerequisites

- Go 1.13 or later
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
    "github.com/DreamDevLost/zerolog-gelf-output"
)

func main() {
    // Configure the GELF writer
    gelfWriter, err := zerolog_gelf_output.NewGelfWriter("udp", "your-gelf-server:12201")
    if err != nil {
        panic("Failed to create GELF writer: " + err.Error())
    }

    // Create a new zerolog logger with the GELF writer
    log := zerolog.New(gelfWriter).With().Timestamp().Logger()

    // Log a message as you normally would
    log.Info().Msg("This is a test message sent to GELF server")
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
