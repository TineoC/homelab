# LDIF seed (high level)

## 01-base.ldif
dn: dc=example,dc=org
objectClass: top
objectClass: dcObject
objectClass: organization
o: Example Org
dc: example

dn: ou=People,dc=example,dc=org
objectClass: organizationalUnit
ou: People

dn: ou=Groups,dc=example,dc=org
objectClass: organizationalUnit
ou: Groups

dn: ou=Departments,dc=example,dc=org
objectClass: organizationalUnit
ou: Departments

dn: ou=Engineering,ou=Departments,dc=example,dc=org
objectClass: organizationalUnit
ou: Engineering

dn: ou=Data,ou=Departments,dc=example,dc=org
objectClass: organizationalUnit
ou: Data

dn: ou=Ops,ou=Departments,dc=example,dc=org
objectClass: organizationalUnit
ou: Ops

## 02-groups.ldif
dn: cn=eng-app-a,ou=Groups,dc=example,dc=org
objectClass: groupOfNames
cn: eng-app-a
member: cn=dummy,dc=example,dc=org

dn: cn=data-app-b,ou=Groups,dc=example,dc=org
objectClass: groupOfNames
cn: data-app-b
member: cn=dummy,dc=example,dc=org

dn: cn=ops-app-c,ou=Groups,dc=example,dc=org
objectClass: groupOfNames
cn: ops-app-c
member: cn=dummy,dc=example,dc=org

dn: cn=shared-app-d,ou=Groups,dc=example,dc=org
objectClass: groupOfNames
cn: shared-app-d
member: cn=dummy,dc=example,dc=org

## 03-users.ldif  (repeat for ~12 users)
dn: uid=alice.eng,ou=People,dc=example,dc=org
objectClass: inetOrgPerson
cn: Alice Engineer
sn: Engineer
uid: alice.eng
mail: alice.eng@example.org
userPassword: {SSHA}Vv1VbY1...   # (pre-hash; demo only)
ou: Engineering

dn: uid=bob.eng,ou=People,dc=example,dc=org
objectClass: inetOrgPerson
cn: Bob Engineer
sn: Engineer
uid: bob.eng
mail: bob.eng@example.org
userPassword: {SSHA}Hh8s... 
ou: Engineering

## 04-memberships.ldif
dn: cn=eng-app-a,ou=Groups,dc=example,dc=org
changetype: modify
add: member
member: uid=alice.eng,ou=People,dc=example,dc=org
member: uid=bob.eng,ou=People,dc=example,dc=org

dn: cn=data-app-b,ou=Groups,dc=example,dc=org
changetype: modify
add: member
member: uid=carol.data,ou=People,dc=example,dc=org

dn: cn=ops-app-c,ou=Groups,dc=example,dc=org
changetype: modify
add: member
member: uid=oscar.ops,ou=People,dc=example,dc=org

dn: cn=shared-app-d,ou=Groups,dc=example,dc=org
changetype: modify
add: member
member: uid=alice.eng,ou=People,dc=example,dc=org
member: uid=carol.data,ou=People,dc=example,dc=org
member: uid=oscar.ops,ou=People,dc=example,dc=org
