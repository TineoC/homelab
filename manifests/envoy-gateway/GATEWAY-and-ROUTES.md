# Envoy Gateway + Gateway API resources

## Gateway (HTTPS)
- Listener: 443
- Hostname: `gw.example.com`
- TLS: secret `gw-tls`

## HTTPRoutes (one per app)
- Parent: the Gateway
- Rules: match `/app-a` (etc), forward to `Service/app-a`
- Filters/policies:
  - **JWT Auth**: issuer = `https://keycloak.iam.svc/auth/realms/demo`
    - JWKS URI auto-discovered or configured
    - Audience = client id (e.g., `app-a`)
  - **RBAC**: allow if `groups` contains the expected group(s)

## Control-plane exposure
- None by default. If needed, add a new HTTPRoute with the same authN/Z.
