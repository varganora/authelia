---
layout: default
title: Miscellaneous
parent: Configuration
nav_order: 7
---

# Miscellaneous

Here are the main customizable options in Authelia that don't fit into their own sections.

## certificates_directory

This option defines the location of additional certificates to load into the trust chain specifically for Authelia.
This currently affects both the SMTP notifier and the LDAP authentication backend. The certificates should all be in the
PEM format and end with the extension `.pem`, `.crt`, or `.cer`. You can either add the individual certificates public
key or the CA public key which signed them (don't add the private key).

```yaml
certificates_directory: /config/certs/
```

## jwt_secret
<div markdown="1">
type: string
{: .label .label-config .label-purple } 
default: ""
{: .label .label-config .label-blue }
required: yes
{: .label .label-config .label-red }
</div>

Defines the secret used to craft JWT tokens leveraged by the identity
verification process. This can also be defined using a [secret](./secrets.md).

```yaml
jwt_secret: v3ry_important_s3cr3t
```

## default_redirection_url
<div markdown="1">
type: string
{: .label .label-config .label-purple } 
default: ""
{: .label .label-config .label-blue }
required: no
{: .label .label-config .label-green }
</div>

The default redirection URL is the URL where users are redirected when Authelia cannot detect the target URL where the
user was heading.

In a normal authentication workflow, a user tries to access a website and they get redirected to the sign-in portal in
order to authenticate. Since the user initially targeted a website, the portal knows where the user was heading and
can redirect them after the authentication process. However, when a user visits the sign in portal directly, the portal
considers the targeted website is the portal. In that case and if the default redirection URL is configured, the user is
redirected to that URL. If not defined, the user is not redirected after authentication.

```yaml
default_redirection_url: https://home.example.com:8080/
```

## default_2fa_method
<div markdown="1">
type: string
{: .label .label-config .label-purple } 
default: ""
{: .label .label-config .label-blue }
required: no
{: .label .label-config .label-green }
</div>

Sets the default second factor method for users. This must be blank or one of the enabled methods. New users will by
default have this method selected for them. In addition if this was configured to `webauthn` and a user had the `totp`
method, and the `totp` method was disabled in the configuration, the users' method would automatically update to the
`webauthn` method.

Options are:

- totp
- webauthn
- mobile_push

```yaml
default_2fa_method: totp
```