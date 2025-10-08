
---

# `docs/RBAC-MAPPING.md`
```md
# RBAC mapping

## LDAP → Keycloak
- LDAP groups: `eng-app-a`, `data-app-b`, `ops-app-c`, `shared-app-d`
- Keycloak realm groups (same names) via group mapper.

## Keycloak → JWT
- Add a Group Mapper (or Role Mapper) in each OIDC client to include:
  - `groups` array with the short group names (preferred)
  - OR `realm_access.roles` if using roles

## Gateway checks
- AuthN: validate JWT signature, issuer, audience.
- AuthZ: allow if `groups` contains an allowed group for the route/app.
