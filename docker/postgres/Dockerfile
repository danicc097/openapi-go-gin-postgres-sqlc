FROM postgres:16.2

RUN apt-get update && apt-get install -y curl postgresql-$PG_MAJOR-plpgsql-check postgresql-$PG_MAJOR-rum postgresql-$PG_MAJOR-hypopg
# 0.3.1
ENV SUPA_AUDIT_COMMIT="0b87d13840e26fc68d0e01e9917c67c894121f6f"

RUN curl -SL "https://github.com/danicc097/supa_audit/archive/${SUPA_AUDIT_COMMIT}.tar.gz" -o /supa_audit.tar.gz \
  && tar -xzvf /supa_audit.tar.gz -C /usr/share/postgresql/$PG_MAJOR/extension/ --strip-components=1 \
  && rm /supa_audit.tar.gz \
  && chown -R postgres:postgres /usr/share/postgresql/$PG_MAJOR/extension/supa_audit*

RUN curl -SL "https://github.com/supabase/index_advisor/archive/refs/tags/v0.2.0.tar.gz" -o /index_advisor.tar.gz \
  && tar -xzvf /index_advisor.tar.gz -C /usr/share/postgresql/$PG_MAJOR/extension/ --strip-components=1 \
  && rm /index_advisor.tar.gz \
  && chown -R postgres:postgres /usr/share/postgresql/$PG_MAJOR/extension/index_advisor*
