CREATE SCHEMA pso2_filter;

CREATE TABLE pso2_filter.events(
   ID SERIAL PRIMARY KEY     NOT NULL,
   EVENT         TEXT    NOT NULL
);

CREATE TABLE pso2_filter.channels(
   ID SERIAL PRIMARY KEY     NOT NULL,
   CHANNELID           TEXT    NOT NULL
);

CREATE UNIQUE INDEX channel_idx ON pso2_filter.channels (channelid);