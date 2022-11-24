-- plpgsql-language-server:use-keyword-query-parameter
-- name: GetUser :one
SELECT user_api_keys.user_api_key_id,
user_api_keys.api_key,
user_api_keys.expires_on,
user_api_keys.user_id,
(case when @testse::boolean = true then row_(users.*) end) as user FROM public.user_api_keys
left join users on users.user_id = user_api_keys.user_id
