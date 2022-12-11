-- plpgsql-language-server:use-keyword-query-parameter
-- name: GetUserNotifications :many
select
  user_notifications.*
  , notifications.notification_type
  , notifications.sender
  , notifications.title
  , notifications.body
  , notifications.label
  , notifications.link
from
  user_notifications
  inner join notifications using (notification_id)
where
  user_notifications.user_id = @user_id
  and notifications.notification_type = @notification_type::notification_type
  and (created_at > sqlc.narg('min_created_at')
    or sqlc.narg('min_created_at') is null) -- first search null, infinite query using last created_at
order by
  created_at desc
limit sqlc.narg('lim');

-- name: CreateNotification :exec
-- plpgsql-language-server:disable
insert into public.notifications (
  receiver_rank
  , title
  , body
  , label
  , link
  , created_at
  , sender
  , receiver
  , notification_type)
values (
  sqlc.narg('receiver_rank')
  , @title
  , @body
  , @label
  , @link
  , current_timestamp
  , @sender
  , sqlc.narg('receiver')
  , @notification_type);

-- name: DeleteNotification :exec
delete from notifications
where notification_id = @notification_id;
