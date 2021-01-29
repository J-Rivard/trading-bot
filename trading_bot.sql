CREATE SCHEMA trading_bot;

CREATE TABLE trading_bot.users(
   id text PRIMARY KEY    NOT NULL,
   liquid_value         float8    NOT NULL,
   asset_value float8 NOT NULL,
   stock_data jsonb NOT NULL
);
