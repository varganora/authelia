---
layout: default
title: MySQL
parent: Storage Backends
grand_parent: Configuration
nav_order: 2
---

# MySQL

The MySQL storage provider.

## Version support

We recommend using the latest version of MySQL that is officially supported by the MySQL developers. We also suggest
checking out [PostgreSQL](postgres.md) as an alternative.

The oldest version of MySQL that has been tested is 5.7. If using 5.7 you may be required to adjust the
`explicit_defaults_for_timestamp` setting. This will be evident when the container starts with an error similar to
`Error 1067: Invalid default value for 'exp'`. You can adjust this setting in the mysql.cnf file like so:

```cnf
[mysqld]
explicit_defaults_for_timestamp = 1
```

## Configuration

```yaml
storage:
  encryption_key: a_very_important_secret
  mysql:
    host: 127.0.0.1
    port: 3306
    database: authelia
    username: authelia
    password: mypassword
    timeout: 5s
```

## Options

### encryption_key
See the [encryption_key docs](./index.md#encryption_key).

### host
<div markdown="1">
type: string
{: .label .label-config .label-purple } 
default: localhost
{: .label .label-config .label-blue }
required: no
{: .label .label-config .label-green }
</div>

The database server host.

If utilising an IPv6 literal address it must be enclosed by square brackets and quoted:
```yaml
host: "[fd00:1111:2222:3333::1]"
```

### port
<div markdown="1">
type: integer
{: .label .label-config .label-purple } 
default: 3306
{: .label .label-config .label-blue }
required: no
{: .label .label-config .label-green }
</div>

The port the database server is listening on.

### database
<div markdown="1">
type: string
{: .label .label-config .label-purple }
required: yes
{: .label .label-config .label-red }
</div>

The database name on the database server that the assigned [user](#username) has access to for the purpose of
**Authelia**.

### username
<div markdown="1">
type: string
{: .label .label-config .label-purple }
required: yes
{: .label .label-config .label-red }
</div>

The username paired with the password used to connect to the database.

### password
<div markdown="1">
type: string
{: .label .label-config .label-purple }
required: yes
{: .label .label-config .label-red }
</div>

The password paired with the username used to connect to the database. Can also be defined using a
[secret](../secrets.md) which is also the recommended way when running as a container.

### timeout
<div markdown="1">
type: duration
{: .label .label-config .label-purple }
default: 5s
{: .label .label-config .label-blue }
required: no
{: .label .label-config .label-green }
</div>

The SQL connection timeout.
