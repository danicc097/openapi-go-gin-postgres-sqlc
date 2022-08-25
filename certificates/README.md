# Certificates

Create with `mkcert`:

```bash
mkcert --cert-file localhost.pem --key-file localhost-key.pem  localhost 127.0.0.1 ::1
mkcert -install
```

to work outside the guest machine, do the following in host, e.g. in Windows:

```powershell
mkcert.exe -install
# replace root store in host with guest's,
# also located in "$(mkcert -CAROOT)/rootCA.pem"
code "$(mkcert.exe -CAROOT)/rootCA.pem"
mkcert.exe -install
```
