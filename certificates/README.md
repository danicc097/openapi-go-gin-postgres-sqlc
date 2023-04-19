# Certificates

Create with `mkcert`, setup with ``project check-build-deps``.

```bash
mkcert -install
# replace root store in host with guest's,
# also located in "$(mkcert -CAROOT)/rootCA.pem"
code "$(mkcert -CAROOT)/rootCA.pem"
mkcert -install
```
