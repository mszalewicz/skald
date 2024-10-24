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

-- name: CountAccounts :one
SELECT count(*) FROM account;

-----------------------------------------------------------------
--- Settings ----------------------------------------------------
-----------------------------------------------------------------

-- name: GetSettings :many
SELECT * FROM settings;

-- name: GetSettingsID :one
SELECT id FROM settings
WHERE width = ?;

-- name: GetFontSizeByWidth :one
SELECT fontsize FROM settings
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

-- name: UpdateSettingFont :exec
UPDATE settings
SET fontsize = ?
WHERE width = ?;

-- name: DeleteSetting :exec
DELETE FROM settings
WHERE id = ?;

-- name: CountSetting :one
SELECT count(*) FROM settings
WHERE width = ?;

-------------------------------------------------------