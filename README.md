# File Processor

File Processor is a Go application that processes a CSV file containing debit and credit transactions. It generates a summary with information about the transactions and sends an email with the summary to a recipient. Additionally, it saves the transactions to a PostgreSQL database.

## Code Interface

These are the main components of the application:

- `fileParser`: Contains the logic of the file parsing.
- `summaryGenerator`: Contains the logic of summary generation.
- `emailSender`: Includes functions for sending the summary email.
- `transactionsRepository`: Contains functions for saving and retrieving transactions from the database.
- `fileProcessor`: Contains the main logic of the transaction processing and database operations.

The code performs the following steps:

1. Reads a CSV file containing transaction data (format: "Id,Date,Transaction").
2. Processes the transactions and generates the summary information.
3. Sends an email with the summary to a specified recipient.
4. Saves the transactions to a PostgreSQL database.
5. Retrieves and prints the transactions from the database.

## Prerequisites

- Go 1.20 or higher
- Docker
- Docker Compose

## Execution Instructions

1. Clone the repository:

```bash
   git clone https://github.com/fedtorres/file-processor.git
```

2. Navigate to the project directory

```bash
   cd file-processor
```

3. Open the `Dockerfile` and set the file path and the recipient email as parameters the `CMD` instruction

```bash
   CMD ["./file-processor", "/app/data/your-file.csv", "recipient@example.com"]
```

4. Open the `docker-compose.yml` file and set the environment variables:

- `DB_HOST`: Hostname or IP address of the PostgreSQL database server.
- `DB_PORT`: Port number on which the PostgreSQL database server is listening.
- `DB_USER`, `POSTGRES_USER`: Username for connecting to the PostgreSQL database.
- `DB_PASSWORD`, `POSTGRES_PASSWORD`: Password for the PostgreSQL database user.
- `DB_NAME`, `POSTGRES_DB`: Name of the PostgreSQL database.
- `EMAIL_FROM`: Email address of the sender.
- `EMAIL_PASSWORD`: Password for the email account.

5. Place the CSV file containing the transactions in the `data/` directory.

6. Build and start the Docker containers using Docker Compose:

```bash
   docker-compose up --build
```
This command will build the Docker image and start the File Processor and PostgreSQL database containers.

7. File Processor will process the CSV file, send the summary email, save the transactions to the database, and retrieve the summary from the database. The summary information will be printed to the console.