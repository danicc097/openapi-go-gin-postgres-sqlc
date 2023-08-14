FROM postgres:13

RUN apt-get update && apt-get install -y postgresql-$PG_MAJOR-plpgsql-check postgresql-$PG_MAJOR-rum
# 0.3.1
ADD https://github.com/supabase/supa_audit/archive/a3056834becfde737f491bd0e97ddf4c7959699d.tar.gz /supa_audit.tar.gz

RUN tar -xzvf /supa_audit.tar.gz && \
  mv /supa_audit-a3056834becfde737f491bd0e97ddf4c7959699d/* /usr/share/postgresql/$PG_MAJOR/extension/ && \
  rm -rf /supa_audit-a3056834becfde737f491bd0e97ddf4c7959699d && \
  rm /supa_audit.tar.gz && \
  chown -R postgres:postgres /usr/share/postgresql/$PG_MAJOR/extension/supa_audit*
