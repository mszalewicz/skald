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
set width = ?,
height = ?,
fontsize = ?
WHERE id = ?;

-- name: DeleteAuthor :exec
DELETE FROM settings
WHERE id = ?;