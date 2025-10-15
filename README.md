<div align="center">

# GATOR
</div>

## Description

- CLI Tool for getting RSS Feeds
- Learning project which is part of Bootdev curicculum
- Lightweight


## Requirements

- GO - tested on version 1.25.3
- PostgreSQL


## Instalation

```bash
go install github.com/hrncacz/go-gator
```

## Configuration

- Create `.gatorconfig.json` file in your HOME directory
- Include connection URL to your PostgreSQL server
```json
{"db_url":"postgres://exampleUsername:examplePassword@localhost:5432?sslmode=disable"}
```

## Usage
- Call program `go-gator` followed by command and potentionaly arguments
- **Example**
```bash
go-gator register exampleUser
```

## Commands
- `login`
    - login user
    - Arg 1 - username of registered user
- `register`
    - register new user
    - Arg 1 - username
- `reset`
    - delete all the data inside DB
    - No arguments
- `users`
    - returns all registered users
    - No arguments 
- `agg`
    - infinite function which downloading data from registered feeds
    - Arg 1 - time duration between fetches in format `1s` `1m` `1h` 
- `addfeed`
    - adds feed to database with currently logged user
    - Arg 1 - name of remote server
    - Arg 2 - url to fetch RSS
- `feeds`
    - lists all registered feeds for current user
    - No arguments
- `follow`
    - create link between current user and registered feed
    - Arg 1 - url to RSS
- `following`
    - lists all feeds followed by current user
    - No arguments
- `unfollow`
    - removes follow of selected feed
    - Arg 1 - url to RSS
- `browse`
    - lists number of newest feeds
    - *optional Arg1* - number or newest feeds. Default set to 2

## ToDo

- [ ] Add some ToDos



