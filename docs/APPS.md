# Demo apps (apps ns)

## Apps
- `app-a`: httpbin-alike container responding on `/app-a/*`
- `app-b`: same; `/app-b/*`
- `app-c`
- `app-d`

Each has:
- Deployment (single replica ok)
- ClusterIP Service
- No Ingress
- `HTTPRoute` with host `gw.example.com` and prefix `/app-X`

## Access policy
- `app-a` → `eng-app-a` group
- `app-b` → `data-app-b`
- `app-c` → `ops-app-c`
- `app-d` → any of the above
