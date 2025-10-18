# Kafikinha - Kafka Pub/Sub CLI

A simple command-line interface built with Go for studying and experimenting with Apache Kafka, Zap structured logging, and Cobra CLI framework.

## Overview

This project demonstrates the fundamental concepts of Kafka pub/sub messaging patterns through a clean Go application. It provides both producer and consumer functionality with proper error handling, structured logging, and command-line interface management.

## Features

- **Kafka Producer**: Send messages to Kafka topics with configurable brokers
- **Kafka Consumer**: Consume messages from Kafka topics with consumer group support
- **Structured Logging**: Uses Uber's Zap logger with development and production configurations
- **CLI Interface**: Built with Cobra for intuitive command-line operations
- **Docker Support**: Includes Docker Compose setup for local Kafka development

## Technologies Used

- **[Apache Kafka](https://kafka.apache.org/)**: Distributed streaming platform
- **[IBM Sarama](https://github.com/IBM/sarama)**: Go client library for Apache Kafka
- **[Zap Logger](https://github.com/uber-go/zap)**: Blazing fast, structured, leveled logging
- **[Cobra CLI](https://github.com/spf13/cobra)**: Modern CLI framework for Go applications
- **[Docker Compose](https://docs.docker.com/compose/)**: Multi-container Docker applications

## Project Structure

```
kafikinha/
├── cmd/                    # CLI commands
│   ├── root.go            # Root command and global flags
│   ├── producer.go        # Producer command implementation
│   └── consumer.go        # Consumer command implementation
├── pkg/                   # Application packages
│   ├── kafka/             # Kafka client implementations
│   │   ├── producer.go    # Async Kafka producer
│   │   ├── consumer.go    # Consumer group implementation
│   │   └── handler.go     # Message handler interface
│   ├── handlers/          # Message processing handlers
│   │   └── log_handler.go # Simple logging message handler
│   └── logger/            # Structured logging setup
│       └── logger.go      # Zap logger initialization
├── docker-compose.yml     # Local Kafka development environment
├── go.mod                 # Go module dependencies
└── main.go               # Application entry point
```

## Prerequisites

- Go 1.24+ installed
- Docker and Docker Compose (for local Kafka setup)

## Quick Start

### 1. Start Kafka Infrastructure

Start the local Kafka cluster using Docker Compose:

```bash
docker-compose up -d
```

This will start:
- **Zookeeper** on port `2181`
- **Kafka** on port `9092`

### 2. Build the Application

```bash
go build -o kafikinha
```

### 3. Run Producer

Send a message to a Kafka topic:

```bash
./kafikinha producer --message "Hello, Kafka!" --topic events --brokers localhost:9092
```

### 4. Run Consumer

Start consuming messages from the topic:

```bash
./kafikinha consumer --topic events --brokers localhost:9092 --group kafikinha-group
```

## CLI Usage

### Global Flags

- `--brokers`: Comma-separated list of Kafka brokers (default: `localhost:9092`)
- `--topic`: Kafka topic name (default: `events`)
- `--env`: Environment mode - `dev` or `prod` (default: `dev`)

### Producer Command

```bash
./kafikinha producer [flags]
```

**Flags:**
- `--message`: Message content to send to Kafka topic

**Example:**
```bash
./kafikinha producer --message "Processing user data" --topic user-events
```

### Consumer Command

```bash
./kafikinha consumer [flags]
```

**Flags:**
- `--group`: Kafka consumer group ID (default: `kafikinha-group`)

**Example:**
```bash
./kafikinha consumer --topic user-events --group processing-group
```

## Configuration

### Environment Modes

- **Development (`dev`)**: Uses Zap's development logger with human-readable output
- **Production (`prod`)**: Uses Zap's production logger with JSON structured output

### Kafka Configuration

The application uses sensible defaults for Kafka configuration:

**Producer:**
- Idempotent producer enabled
- Wait for all in-sync replicas acknowledgment
- Maximum 5 retry attempts
- Async message delivery with success/error handling

**Consumer:**
- Consumer group protocol for load balancing
- Automatic offset management
- Latest offset initial position
- Graceful shutdown handling

## Example Workflows

### Basic Message Flow
1. Start consumer in one terminal:
   ```bash
   ./kafikinha consumer --topic orders
   ```

2. Send messages from another terminal:
   ```bash
   ./kafikinha producer --message '{"orderId": 123, "status": "pending"}' --topic orders
   ```

### Multiple Consumers
Start multiple consumer instances with the same group ID to see load balancing in action:

```bash
# Terminal 1
./kafikinha consumer --topic orders --group order-processors

# Terminal 2  
./kafikinha consumer --topic orders --group order-processors
```

## Troubleshooting

### Common Issues

1. **Connection Refused**: Ensure Kafka is running via Docker Compose
2. **Topic Not Found**: Topics are auto-created by default in the Docker setup
3. **Consumer Lag**: Check if consumer group is processing messages efficiently

### Useful Commands

Check Kafka topics:
```bash
docker exec -it kafikinha-kafka-1 kafka-topics --bootstrap-server localhost:9092 --list
```

Monitor consumer groups:
```bash
docker exec -it kafikinha-kafka-1 kafka-consumer-groups --bootstrap-server localhost:9092 --describe --group kafikinha-group
```

## Contributing

This is a study project, but feel free to:
1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

This project is open source and available under the [MIT License](LICENSE).

## Acknowledgments

- [Apache Kafka Documentation](https://kafka.apache.org/documentation/)
- [Sarama Kafka Client](https://github.com/IBM/sarama)
- [Uber Zap Logger](https://github.com/uber-go/zap)
- [Cobra CLI Framework](https://github.com/spf13/cobra)