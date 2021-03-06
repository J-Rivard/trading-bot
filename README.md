# trading-bot Discord Bot

## This is a discord bot written in go with the purpose of simulating stock trading

## Dependencies
* Postgres - I recommend using docker to host a local instance, https://hub.docker.com/_/postgres
* Discord bot with token - https://discord.com/developers/applications
* Finnhub API token - https://finnhub.io/

## Running
The following env vars are needed to run
```
export BotToken=*
export StockAPIToken=*

export db_user=postgres
export db_pw=*
export db_host=127.0.0.1
export db_schema=*
export db_name=*devdb*
```
Be sure to fill in the * with relevant fields for your environment

Once the appropriate environment variables are exported, simply run `go run cmd/trading-bot/main.go`

### Postgres
You will also need postgres running, with the `trading_bot.sql` file configuration ran.

## Usage