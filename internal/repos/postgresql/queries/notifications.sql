-- plpgsql-language-server:use-keyword-query-parameter
-- name: GetUserNotifications :many
select
  user_notifications.*
  , notifications.notification_type
  , notifications.sender
  , notifications.title
  , notifications.body
  , notifications.labels
  , notifications.link
from
  user_notifications
  inner join notifications using (notification_id)
where
  user_notifications.user_id = @user_id
  and notifications.notification_type = @notification_type::notification_type
  and (user_notifications.created_at > sqlc.narg('min_created_at')
    or sqlc.narg('min_created_at') is null) -- first search null, infinite query using last created_at
order by
  created_at desc
limit sqlc.narg('lim');
