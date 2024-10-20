-----------------------------------------------------------------
--- Account -----------------------------------------------------
-----------------------------------------------------------------

-- name: CreateAccount :one
INSERT INTO account (
  uuid, name
) VALUES (
  ?, ?
)
RETURNING *;

-- name: UpdateAccount :exec
UPDATE account
SET name = ?
WHERE uuid = ?;

INSERT INTO account (
  uuid, name
) VALUES (
  ?, ?, ?
)
RETURNING *;

-----------------------------------------------------------------
--- Settings ----------------------------------------------------
-----------------------------------------------------------------

-- name: GetSettings :many
SELECT * FROM settings;

-- name: GetSettingsID :one
SELECT id FROM settings
WHERE width = ?;

-- name: ListSettings :many
SELECT * FROM settings
ORDER BY width;

-- name: CreateSetting :one
INSERT INTO settings (
  width, height, fontsize
) VALUES (
  ?, ?, ?
)
RETURNING *;

-- name: UpdateSetting :exec
UPDATE settings
SET width = ?,
height = ?,
fontsize = ?
WHERE id = ?;

-- name: DeleteAuthor :exec
DELETE FROM settings
WHERE id = ?;

-------------------------------------------------------