
---

# `docs/KEYCLOAK.md`
```md
# Keycloak setup

## Objects
- Bitnami `keycloak` release in the `keycloak` namespace.
- Realm export (`demo`) applied via GitOps (clients, roles, groups, LDAP federation + mappers).
- LDAP truststore secret (`ldap-truststore`) so Keycloak can talk LDAPS to OpenLDAP.

## JWT settings
- Clients are `confidential` with secrets (`app-*-client-secret`).
- Redirect URIs: `https://*.tineochristopher.com/*`.
- Protocol mappers:
  - `groups` → short group names (used by Envoy for RBAC).
  - `realm_access.roles` → realm role mirror.

## Argo CD flow
- `gitops/applications/iam.yaml` (syncs `gitops/manifests/iam`).
- `gitops/manifests/iam/realm-import.yaml` → realm, clients, LDAP federation (consumed by the Bitnami release).
- `gitops/manifests/iam/secrets.yaml` → LDAP truststore + admin credentials.

## Testing tokens (demo)
Get token with password grant (demo only):
```bash
KC=http://keycloak.keycloak.svc.cluster.local:8080/realms/demo/protocol/openid-connect/token
CLIENT_ID=app-a
CLIENT_SECRET='<from secret>'
curl -s -X POST "$KC" \
  -d grant_type=password \
  -d username=alice.eng \
  -d password=password \
  -d client_id="$CLIENT_ID" \
  -d client_secret="$CLIENT_SECRET" | jq -r .access_token
