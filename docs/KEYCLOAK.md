
---

# `docs/KEYCLOAK.md`
```md
# Keycloak (operator) setup

## Objects
- `Keycloak` (CR): single-node dev (demo), `iam` namespace.
- `KeycloakRealm`: `demo` realm.
- `KeycloakRealmRole` / group mappers: map LDAP groups to roles.
- `KeycloakClient` (OIDC) for each app: `app-a`, `app-b`, `app-c`, `app-d`.
- `KeycloakUserFederation` (LDAP):
  - vendor: `other`
  - connection: `ldaps://openldap.directory.svc:636`
  - bind DN: `cn=admin,dc=example,dc=org`
  - user/object classes: `inetOrgPerson`, `organizationalPerson`, `person`
  - group sync: `groupOfNames` (or `groupOfUniqueNames`)

## JWT settings (per client)
- Access type: `confidential` (for server-to-server) or `public` (SPAs).
- Redirect URIs (for demo): `https://gw.example.com/*`
- Standard OIDC scopes + add group/role mappers so `groups`/`roles` appear in the token.

## Argo CD flow
- `applications/iam/keycloak-operator.yaml`
- `applications/iam/keycloak-instance.yaml`
- `applications/iam/realm-and-clients.yaml`
- `applications/iam/ldap-federation.yaml`

## Testing tokens (demo)
Get token with password grant (demo only):
```bash
KC=https://keycloak.iam.svc/auth/realms/demo/protocol/openid-connect/token
CLIENT_ID=app-a
CLIENT_SECRET='<from secret>'
curl -s -X POST "$KC" \
  -d grant_type=password \
  -d username=alice.eng \
  -d password=alicepass \
  -d client_id="$CLIENT_ID" \
  -d client_secret="$CLIENT_SECRET" | jq -r .access_token
