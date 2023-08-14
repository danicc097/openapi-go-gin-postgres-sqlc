FROM postgres:13

RUN apt-get update && apt-get install -y postgresql-$PG_MAJOR-plpgsql-check postgresql-$PG_MAJOR-rum
# 0.3.1
ENV SUPA_AUDIT_COMMIT="0b87d13840e26fc68d0e01e9917c67c894121f6f"
ADD https://github.com/danicc097/supa_audit/archive/${SUPA_AUDIT_COMMIT}.tar.gz /supa_audit.tar.gz

RUN tar -xzvf /supa_audit.tar.gz && \
  mv /supa_audit-${SUPA_AUDIT_COMMIT}/* /usr/share/postgresql/$PG_MAJOR/extension/ && \
  rm -rf /supa_audit-${SUPA_AUDIT_COMMIT} && \
  rm /supa_audit.tar.gz && \
  chown -R postgres:postgres /usr/share/postgresql/$PG_MAJOR/extension/supa_audit*
