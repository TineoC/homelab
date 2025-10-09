# Architecture

## Components
1) **OpenLDAP** (`directory` ns)
   - Acts as the AD alternative. Holds org tree: Departments → Groups → Users.
   - Exposes `ldaps://openldap.directory.svc:636`.

2) **Keycloak** (`keycloak` ns)
   - Bitnami chart, configured via GitOps (realm import + LDAP federation).
   - Maps LDAP groups → Keycloak roles/realm groups.
   - OIDC clients created per app: `app-a`, `app-b`, `app-c`, `app-d`.

3) **Envoy Gateway + Gateway API** (`envoy-gateway-system` ns)
   - Two Gateways:
     - `public-gw` → `*.tineochristopher.com`
     - `private-gw` → `*.internal.tineochristopher.com`
   - `HTTPRoute`s per app attach auth filters:
     - **JWT** validation (issuer: Keycloak; audience: the client)
     - **RBAC** using JWT claims (`realm_access.roles` or `groups`)

4) **Demo Apps** (`apps` ns)
   - Four Deployments/Services; no Ingress. Only HTTPRoutes.

5) **Control-Plane** (`control-plane` ns)
   - Internal tools (if any); no Ingress. Access only through the Gateway when explicitly routed. Default: blocked.

## Access model (example)
- Departments: `Engineering`, `Data`, `Ops`
- Groups (per department): `eng-app-a`, `data-app-b`, `ops-app-c`, `shared-app-d`
- Users: 12 dummy users spread across departments/groups.
- Authorization mapping:
  - `app-a` -> allow `eng-app-a`
  - `app-b` -> allow `data-app-b`
  - `app-c` -> allow `ops-app-c`
  - `app-d` -> allow any of `eng-app-a`, `data-app-b`, `ops-app-c` (shared)

## Token claims (Keycloak)
- `groups` claim contains LDAP group DNs mapped to realm groups (or short names).
- Alternative: use `realm_access.roles` with role mappers from LDAP groups.
