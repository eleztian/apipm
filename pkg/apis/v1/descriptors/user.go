package descriptors

import (
	def "github.com/caicloud/nirvana/definition"
	"github.com/eleztian/apipm/pkg/app/user"
)

func init() {
	register([]def.Descriptor{{
		Path:        "/login",
		Definitions: []def.Definition{createLogin, deleteLogin},
	},
	}...)
}

var createLogin = def.Definition{
	Method:      def.Create,
	Summary:     "Login",
	Description: "User Login return tk or error",
	Function:    user.Login,
	Parameters: []def.Parameter{
		{
			Source:      def.Form,
			Name:        "account",
			Description: "Login account",
		},
		{
			Source:      def.Form,
			Name:        "password",
			Description: "Login password",
		},
	},
	Results: def.DataErrorResults("login result"),
}
var deleteLogin = def.Definition{
	Method:      def.Delete,
	Summary:     "logout",
	Description: "User Logout return ok and remove tk",
	Function:    user.Logout,
	Parameters:  nil,
	Results:     []def.Result{def.ErrorResult()},
}
