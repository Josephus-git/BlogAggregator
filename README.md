# Gator CLI Tool 🐊

Gator is a command-line interface (CLI) tool designed to help users manage RSS feeds and posts. This project was built to enhance database interaction skills, demonstrating how applications can store and retrieve records, much like how banks handle their data.

## Features ✨

* **User Management**: Register and log in users. 👤
* **Feed Management**: Add, list, follow, and unfollow RSS feeds. 📰
* **Post Aggregation**: Automatically scrape and store posts from followed feeds. 📥
* **Post Browse**: View aggregated posts. 📖
* **Database Reset**: Clear all data from the database. 🗑️

## Installation 🚀

### Prerequisites 🛠️

Before running Gator, ensure you have the following installed and configured on your system:

* **Go (Golang)**: Version 1.18 or higher.
* **PostgreSQL**: A running PostgreSQL database instance.
* **Goose**: A database migration tool for Go. You can install it via:
    ```bash
    go install https://github.com/pressly/goose/v3/cmd/goose@latest
    ```

### Database Setup 🗄️
  **Clone the repository**:
    ```bash
    git clone <https://github.com/Josephus-git/gator.git>
    cd gator
    ```

Gator utilizes `goose` for managing database schema migrations. You need to set up your PostgreSQL database and apply the necessary migrations:

1.  **Configure your database connection string**:
    Navigate to the `sql/schema` directory within the cloned repository. You will find migration files there. By default, the application expects a PostgreSQL database at `postgres://postgres:postgres@localhost:5432/gator`. You might need to:
    * Create a database named `gator`.
    * Ensure a user `postgres` with password `postgres` exists, or adjust the connection string in the `goose` commands to match your database credentials.

2.  **Run database migrations**:
    From the sql/schema of your cloned repository, execute `goose` with your database connection string:
    ```bash
    goose postgres postgres://postgres:postgres@localhost:5432/gator up
    ```
    This command will apply all required database schemas and prepare your database for Gator.

3. **Configure connection credentials for the postgreSQL database**:
    Manually create a config file in your home directory, ~/.gatorconfig.json with the following content:
   ```bash
   {
       "db_url": "postgres postgres://postgres:postgres@localhost:5432/gator?sslmode=disable"
   }
   ```
     
### Finally ✅
  **Install the executable**:
    ```bash
    go install
    ```
    This command will compile the `gator` program and place the executable in your `GOPATH/bin` directory, making it available for use from your terminal.

## Usage 💻

Once Gator is installed and the database is set up, you can interact with it using various commands from your terminal.

**General Command Structure:**

```
gator <command> [arguments...]
```
## Commands 📜

### List all registered users
```
gator users
```

### Register a new user
```
gator register <username>
# e.g.
gator register skywalker
```

### Log in as an existing user
```
gator login <username>
# e.g.
gator login skywalker
```

### Add a new RSS feed
```
gator addfeed <feed_name> <feed_url>
# e.g.
gator addfeed GoBlog https://blog.golang.org/feed.atom
```

### Start aggregating feeds every N seconds
```
gator agg <seconds>
# e.g.
gator agg 30
```

### Browse aggregated posts
```
gator browse <limit>
# e.g.
gator browse 5
```

### List all feeds
```
gator feeds
```

### List feeds you are following
```
gator following
```

### Follow an existing feed by URL
```
gator follow <feed_url>
```

### Unfollow a feed by URL
```
gator unfollow <feed_url>
```

### Reset the entire database
```
gator reset
```

---
Feel free to explore the commands and contribute to the project!
**Happy Gating! 🎉**




