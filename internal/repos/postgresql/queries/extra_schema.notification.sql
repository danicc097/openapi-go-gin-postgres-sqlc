-- plpgsql-language-server:use-keyword-query-parameter
-- name: GetExtraSchemaNotifications :many
select
  notification_type
  , sender
from
  extra_schema.notifications
where
  sender = @user_id
  and notification_type = @notification_type::extra_schema.notification_type;
