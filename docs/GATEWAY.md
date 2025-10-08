# Gateway API with Envoy Gateway

## Overview
- Single `Gateway` (TLS) terminates HTTPS.
- `HTTPRoute` per app with:
  - **JWT auth** filter (issuer = Keycloak realm; JWKS from Keycloak).
  - **RBAC** filter that checks `groups` or `roles` in JWT.

> Note: We avoid Ingress entirely. All external traffic goes through this Gateway.

## Resources (high level)
- Managed by Argo CD app `control-plane/applications/networking/gateway.yml`.
- Helm chart lives at `control-plane/manifests/gateway`.
- `Gateway` in `gateway-system`
- `TLS` secret for the host (demo: self-signed)
- `HTTPRoute` for each `app-a..d` in `apps` ns, parentRef to the Gateway
- `Authentication`/`Authorization` policies (Envoy Gateway CRDs) referenced by routes

## Claims model
- `groups` (e.g., `eng-app-a`, `data-app-b`, `ops-app-c`, `shared-app-d`)
- Route → policy mapping:
  - `/app-a/*` → allow `eng-app-a`
  - `/app-b/*` → allow `data-app-b`
  - `/app-c/*` → allow `ops-app-c`
  - `/app-d/*` → allow any of the four groups

## Control-plane
- No routes by default.
- If a control-plane tool needs exposure, create an `HTTPRoute` with the same JWT+RBAC policy. Otherwise remain unreachable externally.

## Curl test (after DNS/TLS)
```bash
TOKEN="$(./hack/get-token-app-a.sh alice.eng alicepass)"
curl -s -H "Authorization: Bearer $TOKEN" https://gw.example.com/app-a/health
