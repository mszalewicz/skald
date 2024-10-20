CREATE TABLE IF NOT EXISTS settings (
    id        integer PRIMARY KEY,
    width     integer NOT NULL,
    height    integer NOT NULL,
    fontsize  integer NOT NULL
) strict;

CREATE TABLE IF NOT EXISTS account (
    uuid  text PRIMARY KEY,
    name  text NOT NULL
) strict;
