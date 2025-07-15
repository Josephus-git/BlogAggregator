# Gator CLI Tool 🐊

Gator is a command-line interface (CLI) tool designed to help users manage RSS feeds and posts. This project was built to enhance database interaction skills, demonstrating how applications can store and retrieve records, much like how banks handle their data.

## Features ✨

* **User Management**: Register and log in users. 👤
* **Feed Management**: Add, list, follow, and unfollow RSS feeds. 📰
* **Post Aggregation**: Automatically scrape and store posts from followed feeds. 📥
* **Post Browse**: View aggregated posts. 📖
* **Database Reset**: Clear all data from the database. 🗑️

## Installation 🚀

To get started with Gator, follow these simple steps:

1.  **Clone the repository**:
    ```bash
    git clone <your-repository-url>
    cd <your-repository-directory>
    ```
    (Replace `<your-repository-url>` and `<your-repository-directory>` with the actual path to your repository.)

2.  **Install the executable**:
    ```bash
    go install
    ```
    This command will compile the `gator` program and place the executable in your `GOPATH/bin` directory, making it available for use from your terminal.

## Prerequisites 🛠️

Before running Gator, ensure you have the following installed and configured on your system:

* **Go (Golang)**: Version 1.18 or higher.
* **PostgreSQL**: A running PostgreSQL database instance.
* **Goose**: A database migration tool for Go. You can install it via:
    ```bash
    go install [github.com/pressly/goose/v3/cmd/goose@latest](https://github.com/pressly/goose/v3/cmd/goose@latest)
    ```

## Database Setup 🗄️

Gator utilizes `goose` for managing database schema migrations. You need to set up your PostgreSQL database and apply the necessary migrations:

1.  **Configure your database connection string**:
    Navigate to the `sql/schema` directory within the cloned repository. You will find migration files there. By default, the application expects a PostgreSQL database at `postgres://postgres:postgres@localhost:5432/gator`. You might need to:
    * Create a database named `gator`.
    * Ensure a user `postgres` with password `postgres` exists, or adjust the connection string in the `run()` function in `main.go` and in your `goose` commands to match your database credentials.

2.  **Run database migrations**:
    From the root of your cloned repository, execute `goose` with your database connection string:
    ```bash
    goose -dir sql/schema postgres "postgres://postgres:postgres@localhost:5432/gator" up
    ```
    This command will apply all required database schemas and prepare your database for Gator.

## Usage 💻

Once Gator is installed and the database is set up, you can interact with it using various commands from your terminal.

**General Command Structure:**

```bash
gator <command> [arguments...]