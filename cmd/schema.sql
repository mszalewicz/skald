CREATE TABLE IF NOT EXISTS settings (
    id        integer PRIMARY KEY,
    width     integer NOT NULL,
    height    integer NOT NULL,
    fontsize  integer NOT NULL
) strict;