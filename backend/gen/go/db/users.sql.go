// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: users.sql

package db_queries

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	pg_models "github.com/nucleuscloud/neosync/backend/sql/postgresql/models"
)

const createAccountInvite = `-- name: CreateAccountInvite :one
INSERT INTO neosync_api.account_invites (
  account_id, sender_user_id, email, expires_at
) VALUES (
  $1, $2, $3, $4
)
RETURNING id, account_id, sender_user_id, email, token, accepted, created_at, updated_at, expires_at
`

type CreateAccountInviteParams struct {
	AccountID    pgtype.UUID
	SenderUserID pgtype.UUID
	Email        string
	ExpiresAt    pgtype.Timestamp
}

func (q *Queries) CreateAccountInvite(ctx context.Context, db DBTX, arg CreateAccountInviteParams) (NeosyncApiAccountInvite, error) {
	row := db.QueryRow(ctx, createAccountInvite,
		arg.AccountID,
		arg.SenderUserID,
		arg.Email,
		arg.ExpiresAt,
	)
	var i NeosyncApiAccountInvite
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.SenderUserID,
		&i.Email,
		&i.Token,
		&i.Accepted,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ExpiresAt,
	)
	return i, err
}

const createAccountUserAssociation = `-- name: CreateAccountUserAssociation :one
INSERT INTO neosync_api.account_user_associations (
  account_id, user_id
) VALUES (
  $1, $2
)
RETURNING id, account_id, user_id, created_at, updated_at
`

type CreateAccountUserAssociationParams struct {
	AccountID pgtype.UUID
	UserID    pgtype.UUID
}

func (q *Queries) CreateAccountUserAssociation(ctx context.Context, db DBTX, arg CreateAccountUserAssociationParams) (NeosyncApiAccountUserAssociation, error) {
	row := db.QueryRow(ctx, createAccountUserAssociation, arg.AccountID, arg.UserID)
	var i NeosyncApiAccountUserAssociation
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createAuth0IdentityProviderAssociation = `-- name: CreateAuth0IdentityProviderAssociation :one
INSERT INTO neosync_api.user_identity_provider_associations (
  user_id, auth0_provider_id
) VALUES (
  $1, $2
)
RETURNING id, user_id, auth0_provider_id, created_at, updated_at
`

type CreateAuth0IdentityProviderAssociationParams struct {
	UserID          pgtype.UUID
	Auth0ProviderID string
}

func (q *Queries) CreateAuth0IdentityProviderAssociation(ctx context.Context, db DBTX, arg CreateAuth0IdentityProviderAssociationParams) (NeosyncApiUserIdentityProviderAssociation, error) {
	row := db.QueryRow(ctx, createAuth0IdentityProviderAssociation, arg.UserID, arg.Auth0ProviderID)
	var i NeosyncApiUserIdentityProviderAssociation
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Auth0ProviderID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createMachineUser = `-- name: CreateMachineUser :one
INSERT INTO neosync_api.users (
  id, created_at, updated_at, user_type
) VALUES (
  DEFAULT, DEFAULT, DEFAULT, 1
)
RETURNING id, created_at, updated_at, user_type
`

func (q *Queries) CreateMachineUser(ctx context.Context, db DBTX) (NeosyncApiUser, error) {
	row := db.QueryRow(ctx, createMachineUser)
	var i NeosyncApiUser
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserType,
	)
	return i, err
}

const createNonMachineUser = `-- name: CreateNonMachineUser :one
INSERT INTO neosync_api.users (
  id, created_at, updated_at, user_type
) VALUES (
  DEFAULT, DEFAULT, DEFAULT, 0
)
RETURNING id, created_at, updated_at, user_type
`

func (q *Queries) CreateNonMachineUser(ctx context.Context, db DBTX) (NeosyncApiUser, error) {
	row := db.QueryRow(ctx, createNonMachineUser)
	var i NeosyncApiUser
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserType,
	)
	return i, err
}

const createPersonalAccount = `-- name: CreatePersonalAccount :one
INSERT INTO neosync_api.accounts (
  account_type, account_slug
) VALUES (
  0, $1
)
RETURNING id, created_at, updated_at, account_type, account_slug, temporal_config
`

func (q *Queries) CreatePersonalAccount(ctx context.Context, db DBTX, accountSlug string) (NeosyncApiAccount, error) {
	row := db.QueryRow(ctx, createPersonalAccount, accountSlug)
	var i NeosyncApiAccount
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.AccountType,
		&i.AccountSlug,
		&i.TemporalConfig,
	)
	return i, err
}

const createTeamAccount = `-- name: CreateTeamAccount :one
INSERT INTO neosync_api.accounts (
  account_type, account_slug
) VALUES (
  1, $1
)
RETURNING id, created_at, updated_at, account_type, account_slug, temporal_config
`

func (q *Queries) CreateTeamAccount(ctx context.Context, db DBTX, accountSlug string) (NeosyncApiAccount, error) {
	row := db.QueryRow(ctx, createTeamAccount, accountSlug)
	var i NeosyncApiAccount
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.AccountType,
		&i.AccountSlug,
		&i.TemporalConfig,
	)
	return i, err
}

const getAccount = `-- name: GetAccount :one
SELECT id, created_at, updated_at, account_type, account_slug, temporal_config from neosync_api.accounts
WHERE id = $1
`

func (q *Queries) GetAccount(ctx context.Context, db DBTX, id pgtype.UUID) (NeosyncApiAccount, error) {
	row := db.QueryRow(ctx, getAccount, id)
	var i NeosyncApiAccount
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.AccountType,
		&i.AccountSlug,
		&i.TemporalConfig,
	)
	return i, err
}

const getAccountInvite = `-- name: GetAccountInvite :one
SELECT id, account_id, sender_user_id, email, token, accepted, created_at, updated_at, expires_at FROM neosync_api.account_invites
WHERE id = $1
`

func (q *Queries) GetAccountInvite(ctx context.Context, db DBTX, id pgtype.UUID) (NeosyncApiAccountInvite, error) {
	row := db.QueryRow(ctx, getAccountInvite, id)
	var i NeosyncApiAccountInvite
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.SenderUserID,
		&i.Email,
		&i.Token,
		&i.Accepted,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ExpiresAt,
	)
	return i, err
}

const getAccountInviteByToken = `-- name: GetAccountInviteByToken :one
SELECT id, account_id, sender_user_id, email, token, accepted, created_at, updated_at, expires_at FROM neosync_api.account_invites
WHERE token = $1
`

func (q *Queries) GetAccountInviteByToken(ctx context.Context, db DBTX, token string) (NeosyncApiAccountInvite, error) {
	row := db.QueryRow(ctx, getAccountInviteByToken, token)
	var i NeosyncApiAccountInvite
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.SenderUserID,
		&i.Email,
		&i.Token,
		&i.Accepted,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ExpiresAt,
	)
	return i, err
}

const getAccountUserAssociation = `-- name: GetAccountUserAssociation :one
SELECT aua.id, aua.account_id, aua.user_id, aua.created_at, aua.updated_at from neosync_api.account_user_associations aua
INNER JOIN neosync_api.accounts a ON a.id = aua.account_id
INNER JOIN neosync_api.users u ON u.id = aua.user_id
WHERE a.id = $1 AND u.id = $2
`

type GetAccountUserAssociationParams struct {
	AccountId pgtype.UUID
	UserId    pgtype.UUID
}

func (q *Queries) GetAccountUserAssociation(ctx context.Context, db DBTX, arg GetAccountUserAssociationParams) (NeosyncApiAccountUserAssociation, error) {
	row := db.QueryRow(ctx, getAccountUserAssociation, arg.AccountId, arg.UserId)
	var i NeosyncApiAccountUserAssociation
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getAccountsByUser = `-- name: GetAccountsByUser :many
SELECT a.id, a.created_at, a.updated_at, a.account_type, a.account_slug, a.temporal_config from neosync_api.accounts a
INNER JOIN neosync_api.account_user_associations aua ON aua.account_id = a.id
INNER JOIN neosync_api.users u ON u.id = aua.user_id
WHERE u.id = $1
`

func (q *Queries) GetAccountsByUser(ctx context.Context, db DBTX, id pgtype.UUID) ([]NeosyncApiAccount, error) {
	rows, err := db.Query(ctx, getAccountsByUser, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []NeosyncApiAccount
	for rows.Next() {
		var i NeosyncApiAccount
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.AccountType,
			&i.AccountSlug,
			&i.TemporalConfig,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getActiveAccountInvites = `-- name: GetActiveAccountInvites :many
SELECT id, account_id, sender_user_id, email, token, accepted, created_at, updated_at, expires_at FROM neosync_api.account_invites
WHERE account_id = $1 AND expires_at > CURRENT_TIMESTAMP AND accepted = false
`

func (q *Queries) GetActiveAccountInvites(ctx context.Context, db DBTX, accountid pgtype.UUID) ([]NeosyncApiAccountInvite, error) {
	rows, err := db.Query(ctx, getActiveAccountInvites, accountid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []NeosyncApiAccountInvite
	for rows.Next() {
		var i NeosyncApiAccountInvite
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.SenderUserID,
			&i.Email,
			&i.Token,
			&i.Accepted,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ExpiresAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAnonymousUser = `-- name: GetAnonymousUser :one
SELECT id, created_at, updated_at, user_type from neosync_api.users
WHERE id = '00000000-0000-0000-0000-000000000000'
`

func (q *Queries) GetAnonymousUser(ctx context.Context, db DBTX) (NeosyncApiUser, error) {
	row := db.QueryRow(ctx, getAnonymousUser)
	var i NeosyncApiUser
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserType,
	)
	return i, err
}

const getPersonalAccountByUserId = `-- name: GetPersonalAccountByUserId :one
SELECT a.id, a.created_at, a.updated_at, a.account_type, a.account_slug, a.temporal_config from neosync_api.accounts a
INNER JOIN neosync_api.account_user_associations aua ON aua.account_id = a.id
INNER JOIN neosync_api.users u ON u.id = aua.user_id
WHERE u.id = $1 AND a.account_type = 0
`

func (q *Queries) GetPersonalAccountByUserId(ctx context.Context, db DBTX, userid pgtype.UUID) (NeosyncApiAccount, error) {
	row := db.QueryRow(ctx, getPersonalAccountByUserId, userid)
	var i NeosyncApiAccount
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.AccountType,
		&i.AccountSlug,
		&i.TemporalConfig,
	)
	return i, err
}

const getTeamAccountsByUserId = `-- name: GetTeamAccountsByUserId :many
SELECT a.id, a.created_at, a.updated_at, a.account_type, a.account_slug, a.temporal_config from neosync_api.accounts a
INNER JOIN neosync_api.account_user_associations aua ON aua.account_id = a.id
INNER JOIN neosync_api.users u ON u.id = aua.user_id
WHERE u.id = $1 AND a.account_type = 1
`

func (q *Queries) GetTeamAccountsByUserId(ctx context.Context, db DBTX, userid pgtype.UUID) ([]NeosyncApiAccount, error) {
	rows, err := db.Query(ctx, getTeamAccountsByUserId, userid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []NeosyncApiAccount
	for rows.Next() {
		var i NeosyncApiAccount
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.AccountType,
			&i.AccountSlug,
			&i.TemporalConfig,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTemporalConfigByAccount = `-- name: GetTemporalConfigByAccount :one
SELECT temporal_config
FROM neosync_api.accounts
WHERE id = $1
`

func (q *Queries) GetTemporalConfigByAccount(ctx context.Context, db DBTX, id pgtype.UUID) (*pg_models.TemporalConfig, error) {
	row := db.QueryRow(ctx, getTemporalConfigByAccount, id)
	var temporal_config *pg_models.TemporalConfig
	err := row.Scan(&temporal_config)
	return temporal_config, err
}

const getTemporalConfigByUserAccount = `-- name: GetTemporalConfigByUserAccount :one
SELECT a.temporal_config
FROM neosync_api.accounts a
INNER JOIN neosync_api.account_user_associations aua ON aua.account_id = a.id
INNER JOIN neosync_api.users u ON u.id = aua.user_id
WHERE a.id = $1 AND u.id = $2
`

type GetTemporalConfigByUserAccountParams struct {
	AccountId pgtype.UUID
	UserId    pgtype.UUID
}

func (q *Queries) GetTemporalConfigByUserAccount(ctx context.Context, db DBTX, arg GetTemporalConfigByUserAccountParams) (*pg_models.TemporalConfig, error) {
	row := db.QueryRow(ctx, getTemporalConfigByUserAccount, arg.AccountId, arg.UserId)
	var temporal_config *pg_models.TemporalConfig
	err := row.Scan(&temporal_config)
	return temporal_config, err
}

const getUser = `-- name: GetUser :one
SELECT id, created_at, updated_at, user_type FROM neosync_api.users
WHERE id = $1
`

func (q *Queries) GetUser(ctx context.Context, db DBTX, id pgtype.UUID) (NeosyncApiUser, error) {
	row := db.QueryRow(ctx, getUser, id)
	var i NeosyncApiUser
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserType,
	)
	return i, err
}

const getUserAssociationByAuth0Id = `-- name: GetUserAssociationByAuth0Id :one
SELECT id, user_id, auth0_provider_id, created_at, updated_at from neosync_api.user_identity_provider_associations
WHERE auth0_provider_id = $1
`

func (q *Queries) GetUserAssociationByAuth0Id(ctx context.Context, db DBTX, auth0ProviderID string) (NeosyncApiUserIdentityProviderAssociation, error) {
	row := db.QueryRow(ctx, getUserAssociationByAuth0Id, auth0ProviderID)
	var i NeosyncApiUserIdentityProviderAssociation
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Auth0ProviderID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByAuth0Id = `-- name: GetUserByAuth0Id :one
SELECT u.id, u.created_at, u.updated_at, u.user_type from neosync_api.users u
INNER JOIN neosync_api.user_identity_provider_associations uipa ON uipa.user_id = u.id
WHERE uipa.auth0_provider_id = $1 and u.user_type = 0
`

func (q *Queries) GetUserByAuth0Id(ctx context.Context, db DBTX, auth0ProviderID string) (NeosyncApiUser, error) {
	row := db.QueryRow(ctx, getUserByAuth0Id, auth0ProviderID)
	var i NeosyncApiUser
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserType,
	)
	return i, err
}

const getUserIdentitiesByTeamAccount = `-- name: GetUserIdentitiesByTeamAccount :many
SELECT aipa.id, aipa.user_id, aipa.auth0_provider_id, aipa.created_at, aipa.updated_at FROM neosync_api.user_identity_provider_associations aipa
JOIN neosync_api.account_user_associations aua ON aua.user_id = aipa.user_id
JOIN neosync_api.accounts a ON a.id = aua.account_id
WHERE aua.account_id = $1 AND a.account_type = 1
`

func (q *Queries) GetUserIdentitiesByTeamAccount(ctx context.Context, db DBTX, accountid pgtype.UUID) ([]NeosyncApiUserIdentityProviderAssociation, error) {
	rows, err := db.Query(ctx, getUserIdentitiesByTeamAccount, accountid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []NeosyncApiUserIdentityProviderAssociation
	for rows.Next() {
		var i NeosyncApiUserIdentityProviderAssociation
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Auth0ProviderID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserIdentityAssociationsByUserIds = `-- name: GetUserIdentityAssociationsByUserIds :many
SELECT id, user_id, auth0_provider_id, created_at, updated_at from neosync_api.user_identity_provider_associations
WHERE user_id = ANY($1::uuid[])
`

func (q *Queries) GetUserIdentityAssociationsByUserIds(ctx context.Context, db DBTX, dollar_1 []pgtype.UUID) ([]NeosyncApiUserIdentityProviderAssociation, error) {
	rows, err := db.Query(ctx, getUserIdentityAssociationsByUserIds, dollar_1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []NeosyncApiUserIdentityProviderAssociation
	for rows.Next() {
		var i NeosyncApiUserIdentityProviderAssociation
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Auth0ProviderID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserIdentityByUserId = `-- name: GetUserIdentityByUserId :one
SELECT aipa.id, aipa.user_id, aipa.auth0_provider_id, aipa.created_at, aipa.updated_at FROM neosync_api.user_identity_provider_associations aipa
WHERE aipa.user_id = $1
`

func (q *Queries) GetUserIdentityByUserId(ctx context.Context, db DBTX, userID pgtype.UUID) (NeosyncApiUserIdentityProviderAssociation, error) {
	row := db.QueryRow(ctx, getUserIdentityByUserId, userID)
	var i NeosyncApiUserIdentityProviderAssociation
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Auth0ProviderID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const isUserInAccount = `-- name: IsUserInAccount :one
SELECT count(aua.id) from neosync_api.account_user_associations aua
INNER JOIN neosync_api.accounts a ON a.id = aua.account_id
INNER JOIN neosync_api.users u ON u.id = aua.user_id
WHERE a.id = $1 AND u.id = $2
`

type IsUserInAccountParams struct {
	AccountId pgtype.UUID
	UserId    pgtype.UUID
}

func (q *Queries) IsUserInAccount(ctx context.Context, db DBTX, arg IsUserInAccountParams) (int64, error) {
	row := db.QueryRow(ctx, isUserInAccount, arg.AccountId, arg.UserId)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const removeAccountInvite = `-- name: RemoveAccountInvite :exec
DELETE FROM neosync_api.account_invites
WHERE id = $1
`

func (q *Queries) RemoveAccountInvite(ctx context.Context, db DBTX, id pgtype.UUID) error {
	_, err := db.Exec(ctx, removeAccountInvite, id)
	return err
}

const removeAccountUser = `-- name: RemoveAccountUser :exec
DELETE FROM neosync_api.account_user_associations
WHERE account_id = $1 AND user_id = $2
`

type RemoveAccountUserParams struct {
	AccountId pgtype.UUID
	UserId    pgtype.UUID
}

func (q *Queries) RemoveAccountUser(ctx context.Context, db DBTX, arg RemoveAccountUserParams) error {
	_, err := db.Exec(ctx, removeAccountUser, arg.AccountId, arg.UserId)
	return err
}

const setAnonymousUser = `-- name: SetAnonymousUser :one
INSERT INTO neosync_api.users (
  id, created_at, updated_at
) VALUES (
  '00000000-0000-0000-0000-000000000000', DEFAULT, DEFAULT
)
ON CONFLICT (id)
DO
  UPDATE SET updated_at = current_timestamp
RETURNING id, created_at, updated_at, user_type
`

func (q *Queries) SetAnonymousUser(ctx context.Context, db DBTX) (NeosyncApiUser, error) {
	row := db.QueryRow(ctx, setAnonymousUser)
	var i NeosyncApiUser
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserType,
	)
	return i, err
}

const updateAccountInviteToAccepted = `-- name: UpdateAccountInviteToAccepted :one
UPDATE neosync_api.account_invites
SET accepted = true
WHERE id = $1
RETURNING id, account_id, sender_user_id, email, token, accepted, created_at, updated_at, expires_at
`

func (q *Queries) UpdateAccountInviteToAccepted(ctx context.Context, db DBTX, id pgtype.UUID) (NeosyncApiAccountInvite, error) {
	row := db.QueryRow(ctx, updateAccountInviteToAccepted, id)
	var i NeosyncApiAccountInvite
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.SenderUserID,
		&i.Email,
		&i.Token,
		&i.Accepted,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ExpiresAt,
	)
	return i, err
}

const updateActiveAccountInvitesToExpired = `-- name: UpdateActiveAccountInvitesToExpired :one
UPDATE neosync_api.account_invites
SET expires_at = CURRENT_TIMESTAMP
WHERE account_id = $1 AND email = $2 AND expires_at > CURRENT_TIMESTAMP
RETURNING id, account_id, sender_user_id, email, token, accepted, created_at, updated_at, expires_at
`

type UpdateActiveAccountInvitesToExpiredParams struct {
	AccountId pgtype.UUID
	Email     string
}

func (q *Queries) UpdateActiveAccountInvitesToExpired(ctx context.Context, db DBTX, arg UpdateActiveAccountInvitesToExpiredParams) (NeosyncApiAccountInvite, error) {
	row := db.QueryRow(ctx, updateActiveAccountInvitesToExpired, arg.AccountId, arg.Email)
	var i NeosyncApiAccountInvite
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.SenderUserID,
		&i.Email,
		&i.Token,
		&i.Accepted,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ExpiresAt,
	)
	return i, err
}

const updateTemporalConfigByAccount = `-- name: UpdateTemporalConfigByAccount :one
UPDATE neosync_api.accounts
SET temporal_config = $1
WHERE id = $2
RETURNING id, created_at, updated_at, account_type, account_slug, temporal_config
`

type UpdateTemporalConfigByAccountParams struct {
	TemporalConfig *pg_models.TemporalConfig
	AccountId      pgtype.UUID
}

func (q *Queries) UpdateTemporalConfigByAccount(ctx context.Context, db DBTX, arg UpdateTemporalConfigByAccountParams) (NeosyncApiAccount, error) {
	row := db.QueryRow(ctx, updateTemporalConfigByAccount, arg.TemporalConfig, arg.AccountId)
	var i NeosyncApiAccount
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.AccountType,
		&i.AccountSlug,
		&i.TemporalConfig,
	)
	return i, err
}
