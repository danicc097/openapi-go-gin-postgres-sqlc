# Certificates

Create with `mkcert`, see `bin/scripts/install-mkcert.sh`.

```powershell
mkcert.exe -install
# replace root store in host with guest's,
# also located in "$(mkcert -CAROOT)/rootCA.pem"
code "$(mkcert.exe -CAROOT)/rootCA.pem"
mkcert.exe -install
```
