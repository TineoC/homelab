# Gateway API with Envoy Gateway

## Overview
- Two `Gateway` resources terminate HTTPS:
  - `public-gw` (hostname `*.tineochristopher.com`)
  - `private-gw` (hostname `*.internal.tineochristopher.com`)
- `HTTPRoute` per app with:
  - **JWT auth** filter (issuer = Keycloak realm; JWKS from Keycloak).
  - **RBAC** filter that checks `groups` or `roles` in JWT.

> Note: We avoid Ingress entirely. All external traffic goes through this Gateway.

## Resources (high level)
- Managed by Argo CD application `gitops/applications/gateway.yaml`.
- Manifests in `gitops/manifests/gateway`.
- `HTTPRoute` per app (namespace `apps`, host-based routing).
- `SecurityPolicy` objects attach JWT validation + RBAC to each route.

## Claims model
- `groups` (e.g., `eng-app-a`, `data-app-b`, `ops-app-c`, `shared-app-d`)
- Host → policy mapping:
  - `app-a.tineochristopher.com` → allow `eng-app-a`
  - `app-b.tineochristopher.com` → allow `data-app-b`
  - `app-c.internal.tineochristopher.com` → allow `ops-app-c`
  - `app-d.internal.tineochristopher.com` → allow any of the four groups

## Control-plane
- No routes by default.
- If a control-plane tool needs exposure, create an `HTTPRoute` with the same JWT+RBAC policy. Otherwise remain unreachable externally.

## Curl test (after DNS/TLS)
```bash
kubectl -n envoy-gateway-system port-forward \
  svc/$(kubectl get svc -n envoy-gateway-system -l gateway.envoyproxy.io/managed-by=envoy-gateway -o jsonpath='{.items[0].metadata.name}') \
  18443:443

TOKEN="$(curl -sS -X POST http://keycloak.keycloak.svc.cluster.local:8080/realms/demo/protocol/openid-connect/token \
  -d grant_type=password \
  -d client_id=app-a \
  -d client_secret=app-a-client-secret \
  -d username=alice.eng \
  -d password=password | jq -r .access_token)"

curl -sk --resolve app-a.tineochristopher.com:18443:127.0.0.1 \
  -H "Host: app-a.tineochristopher.com" \
  -H "Authorization: Bearer ${TOKEN}" \
  https://app-a.tineochristopher.com:18443/
```
