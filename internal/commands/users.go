package commands

import (
	"github.com/urfave/cli"

	"github.com/photoprism/photoprism/internal/acl"
)

// Usage hints for the user management subcommands.
const (
	UserDisplayNameUsage  = "full `NAME` for display in the interface"
	UserEmailUsage        = "unique `EMAIL` address of the user"
	UserPasswordUsage     = "`PASSWORD` for authentication"
	UserRoleUsage         = "user account `ROLE`"
	UserAttrUsage         = "custom user account `ATTRIBUTES`"
	UserAdminUsage        = "make user super admin with full access"
	UserDisableLoginUsage = "disable login and use of the web interface"
	UserCanSyncUsage      = "allow to sync files via WebDAV"
)

// UsersCommand registers the user management subcommands.
var UsersCommand = cli.Command{
	Name:  "users",
	Usage: "User management subcommands",
	Subcommands: []cli.Command{
		UsersListCommand,
		UsersAddCommand,
		UsersShowCommand,
		UsersModCommand,
		UsersRemoveCommand,
	},
}

// UserFlags specifies the add and modify user command flags.
var UserFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "displayname, n",
		Usage: UserDisplayNameUsage,
	},
	cli.StringFlag{
		Name:  "email, m",
		Usage: UserEmailUsage,
	},
	cli.StringFlag{
		Name:  "password, p",
		Usage: UserPasswordUsage,
	},
	cli.StringFlag{
		Name:  "role, r",
		Usage: UserRoleUsage,
		Value: acl.RoleAdmin.String(),
	},
	cli.StringFlag{
		Name:  "attr, a",
		Usage: UserAttrUsage,
	},
	cli.BoolFlag{
		Name:  "superadmin, s",
		Usage: UserAdminUsage,
	},
	cli.BoolFlag{
		Name:  "disable-login, d",
		Usage: UserDisableLoginUsage,
	},
	cli.BoolFlag{
		Name:  "can-sync, w",
		Usage: UserCanSyncUsage,
	},
}
