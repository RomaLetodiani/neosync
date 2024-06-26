-- name: GetJobById :one
SELECT * from neosync_api.jobs WHERE id = $1;

-- name: GetJobByNameAndAccount :one
SELECT j.* from neosync_api.jobs j
INNER JOIN neosync_api.accounts a ON a.id = j.account_id
WHERE a.id = sqlc.arg('accountId') AND j.name = sqlc.arg('jobName');

-- name: GetJobsByAccount :many
SELECT j.* from neosync_api.jobs j
INNER JOIN neosync_api.accounts a ON a.id = j.account_id
WHERE a.id = sqlc.arg('accountId')
ORDER BY j.created_at DESC;

-- name: RemoveJobById :exec
DELETE FROM neosync_api.jobs WHERE id = $1;

-- name: CreateJob :one
INSERT INTO neosync_api.jobs (
  name, account_id, status, connection_options, mappings,
  cron_schedule, created_by_id, updated_by_id, workflow_options, sync_options
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
)
RETURNING *;

-- name: DeleteJob :exec
DELETE FROM neosync_api.jobs WHERE id = $1;

-- name: UpdateJobSchedule :one
UPDATE neosync_api.jobs
SET cron_schedule = $1,
updated_by_id = $2
WHERE id = $3
RETURNING *;

-- name: SetJobWorkflowOptions :one
UPDATE neosync_api.jobs
SET workflow_options = $1,
updated_by_id = $2
WHERE id = $3
RETURNING *;

-- name: SetJobSyncOptions :one
UPDATE neosync_api.jobs
SET sync_options = $1,
updated_by_id = $2
WHERE id = $3
RETURNING *;

-- name: UpdateJobMappings :one
UPDATE neosync_api.jobs
SET mappings = $1,
updated_by_id = $2
WHERE id = $3
RETURNING *;

-- name: UpdateJobSource :one
UPDATE neosync_api.jobs
SET connection_options = $1,
updated_by_id = $2
WHERE id = $3
RETURNING *;

-- name: IsJobNameAvailable :one
SELECT count(j.id) from neosync_api.jobs j
INNER JOIN neosync_api.accounts a ON a.id = j.account_id
WHERE a.id = sqlc.arg('accountId') AND j.name = sqlc.arg('jobName');

-- name: CreateJobConnectionDestination :one
INSERT INTO neosync_api.job_destination_connection_associations (
  job_id, connection_id, options
) VALUES (
  $1, $2, $3
)
ON CONFLICT(job_id, connection_id)
DO NOTHING
RETURNING *;

-- name: CreateJobConnectionDestinations :copyfrom
INSERT INTO neosync_api.job_destination_connection_associations (
  job_id, connection_id, options
) VALUES (
  $1, $2, $3
);

-- name: GetJobConnectionDestination :one
SELECT jdca.* from neosync_api.job_destination_connection_associations jdca
WHERE jdca.id = $1;

-- name: GetJobConnectionDestinations :many
SELECT jdca.* from neosync_api.job_destination_connection_associations jdca
INNER JOIN neosync_api.jobs j ON j.id = jdca.job_id
WHERE j.id = $1
ORDER BY jdca.created_at;

-- name: GetJobConnectionDestinationsByJobIds :many
SELECT jdca.* from neosync_api.job_destination_connection_associations jdca
INNER JOIN neosync_api.jobs j ON j.id = jdca.job_id
WHERE j.id = ANY(sqlc.arg('jobIds')::uuid[])
ORDER BY j.created_at, jdca.created_at;

-- name: RemoveJobConnectionDestinations :exec
DELETE FROM neosync_api.job_destination_connection_associations
WHERE id = ANY(sqlc.arg('jobIds')::uuid[]);

-- name: UpdateJobConnectionDestination :one
UPDATE neosync_api.job_destination_connection_associations
SET options = $1,
connection_id = $2
WHERE id = $3
RETURNING *;

-- name: RemoveJobConnectionDestination :exec
DELETE FROM neosync_api.job_destination_connection_associations WHERE id = $1;
