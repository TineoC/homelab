# Keycloak CRs overview (operator-managed)

## Realm
- name: `demo`
- default roles & optional realm groups: `eng-app-a`, `data-app-b`, `ops-app-c`, `shared-app-d`

## LDAP Federation
- `UserFederation` provider of type `ldap`
- connectionUrl: `ldaps://openldap.directory.svc:636`
- usersDn: `ou=People,dc=example,dc=org`
- groupsDn: `ou=Groups,dc=example,dc=org`
- importEnabled: true
- syncRegistrations: false
- periodic full/changed sync schedulers enabled

## Clients (OIDC)
- `app-a`, `app-b`, `app-c`, `app-d`
- public or confidential (demo ok)
- redirectUris: `https://gw.example.com/*`
- protocolMappers:
  - Group mapper emitting short names in `groups`
  - (optional) Realm role mapper to `realm_access.roles`

## Users
- Do not create local users; rely on LDAP sync.
