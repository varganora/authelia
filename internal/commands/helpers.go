package commands

import (
	"github.com/authelia/authelia/v4/internal/authentication"
	"github.com/authelia/authelia/v4/internal/authorization"
	"github.com/authelia/authelia/v4/internal/middlewares"
	"github.com/authelia/authelia/v4/internal/notification"
	"github.com/authelia/authelia/v4/internal/ntp"
	"github.com/authelia/authelia/v4/internal/oidc"
	"github.com/authelia/authelia/v4/internal/regulation"
	"github.com/authelia/authelia/v4/internal/session"
	"github.com/authelia/authelia/v4/internal/storage"
	"github.com/authelia/authelia/v4/internal/totp"
	"github.com/authelia/authelia/v4/internal/utils"
)

func getStorageProvider() (provider storage.Provider) {
	switch {
	case config.Storage.PostgreSQL != nil:
		return storage.NewPostgreSQLProvider(config)
	case config.Storage.MySQL != nil:
		return storage.NewMySQLProvider(config)
	case config.Storage.Local != nil:
		return storage.NewSQLiteProvider(config)
	default:
		return nil
	}
}

func getProviders() (providers middlewares.Providers, warnings []error, errors []error) {
	// TODO: Adjust this so the CertPool can be used like a provider.
	autheliaCertPool, warnings, errors := utils.NewX509CertPool(config.CertificatesDirectory)
	if len(warnings) != 0 || len(errors) != 0 {
		return providers, warnings, errors
	}

	storageProvider := getStorageProvider()

	var (
		userProvider authentication.UserProvider
		err          error
	)

	switch {
	case config.AuthenticationBackend.File != nil:
		userProvider = authentication.NewFileUserProvider(config.AuthenticationBackend.File)
	case config.AuthenticationBackend.LDAP != nil:
		userProvider = authentication.NewLDAPUserProvider(config.AuthenticationBackend, autheliaCertPool)
	}

	var notifier notification.Notifier

	switch {
	case config.Notifier.SMTP != nil:
		notifier = notification.NewSMTPNotifier(config.Notifier.SMTP, autheliaCertPool)
	case config.Notifier.FileSystem != nil:
		notifier = notification.NewFileNotifier(*config.Notifier.FileSystem)
	}

	ntpProvider := ntp.NewProvider(&config.NTP)

	clock := utils.RealClock{}
	authorizer := authorization.NewAuthorizer(config)
	sessionProvider := session.NewProvider(config.Session, autheliaCertPool)
	regulator := regulation.NewRegulator(config.Regulation, storageProvider, clock)

	oidcProvider, err := oidc.NewOpenIDConnectProvider(config.IdentityProviders.OIDC, storageProvider)
	if err != nil {
		errors = append(errors, err)
	}

	totpProvider := totp.NewTimeBasedProvider(config.TOTP)

	ppolicyProvider := middlewares.NewPasswordPolicyProvider(config.PasswordPolicy)

	return middlewares.Providers{
		Authorizer:      authorizer,
		UserProvider:    userProvider,
		Regulator:       regulator,
		OpenIDConnect:   oidcProvider,
		StorageProvider: storageProvider,
		NTP:             ntpProvider,
		Notifier:        notifier,
		SessionProvider: sessionProvider,
		TOTP:            totpProvider,
		PasswordPolicy:  ppolicyProvider,
	}, warnings, errors
}
