# trading-bot Discord Bot

## This is a discord bot written in go with the purpose of filtering out the unwanted UQ events in pso2.

## Dependencies
* Postgres - I recommend using docker to host a local instance, https://hub.docker.com/_/postgres
* Discord bot with token - https://discord.com/developers/applications

## Running
The following env vars are needed to run
```
export BotToken=*

export db_user=postgres
export db_pw=*
export db_host=127.0.0.1
export db_schema=*
export db_name=*devdb*
```
Be sure to fill in the * with relevant fields for your environment

Once the appropriate environment variables are exported, simply run `go run cmd/trading-bot/main.go`

### Postgres
You will also need postgres running, with the `trading-bot.sql` file configuration ran.

## Usage