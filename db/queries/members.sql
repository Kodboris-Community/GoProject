-- Create table entries and return one object after creation

-- name: CreateMember :one
INSERT  INTO members (
                       first_name,
                       last_name,
                       status,
                       comment

) VALUES (
          $1, $2, $3, $4
          -- return all values of alL columns
         ) RETURNING *;

-- name: GetMember :one
SELECT * FROM members
WHERE member_id = $1 LIMIT 1;  -- return a single record