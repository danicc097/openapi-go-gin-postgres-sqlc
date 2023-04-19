# Certificates

Create with `mkcert` manually or setup with ``project check-build-deps``.

```bash
mkcert -install
# if using a VM, replace rootCA.pem in host with guest's,
# both located in "$(mkcert -CAROOT)/rootCA.pem"
# and then mkcert -install in host again
```
