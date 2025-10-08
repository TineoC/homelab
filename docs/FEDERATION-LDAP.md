
---

# `docs/FEDERATION-LDAP.md`
```md
# OpenLDAP federation (AD alternative)

## Seeded tree
Base DN: `dc=example,dc=org`

- `ou=Departments,dc=example,dc=org`
  - `ou=Engineering`
  - `ou=Data`
  - `ou=Ops`
- `ou=Groups,dc=example,dc=org`
  - `cn=eng-app-a`
  - `cn=data-app-b`
  - `cn=ops-app-c`
  - `cn=shared-app-d`
- `ou=People,dc=example,dc=org`
  - Users: `uid=alice.eng`, `uid=bob.eng`, `uid=carol.data`, … (12 total)

## LDIF snippets (applied via ConfigMap + init)
- `01-base.ldif`: base, OUs
- `02-groups.ldif`: four groups
- `03-users.ldif`: 12 users with hashed passwords (`{SSHA}`), `memberOf` if overlay used
- `04-memberships.ldif`: add group members

## Keycloak LDAP config (important)
- User DN: `ou=People,dc=example,dc=org`
- Group DN: `ou=Groups,dc=example,dc=org`
- Username LDAP attr: `uid`
- RDN: `uid`
- Group membership attr: `member`
- Import users: ON; Sync: periodic
- Mappers:
  - Group mapper → realm groups (or client roles)
  - Optional attribute mappers for `departmentNumber`, `title`, etc.
