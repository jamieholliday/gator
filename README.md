# gator

gator is a cli program for fetching feeds. To use you will need Postgres and go installed


## Installation

### Using Go Install

1. You can install the gator CLI directly using Go:

```
go install github.com/jamieholliday/gator@latest
```

Make sure your Go bin directory is in your PATH.

2. Build it `go build`

3. Create a config file at the root of your machine called
`~/.gatorconfig.json`. It should contain the url to your postgress db

{"db_url":"postgres://dbname:@localhost:5432/gator?sslmode=disable"}

4. Start your database
`Make db-start`

5. Use the following commands to interact with the cli

## Commands
- `gator register` - Create a new account
- `gator login` - Log in to your account
- `gator reset` - Reset database (admin function)
- `gator users` - List all users
- `gator agg` - Aggregate feeds
- `gator addfeed` - Add a new feed (requires login)
- `gator feeds` - List all available feeds
- `gator follow` - Follow a feed (requires login)
- `gator following` - List feeds you're following (requires login)
- `gator unfollow` - Unfollow a feed (requires login)
- `gator browse` - Browse your feed content (requires login)

