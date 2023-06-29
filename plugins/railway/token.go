package railway

import (
	"context"

	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/importer"
	"github.com/1Password/shell-plugins/sdk/provision"
	"github.com/1Password/shell-plugins/sdk/schema"
	"github.com/1Password/shell-plugins/sdk/schema/credname"
	"github.com/1Password/shell-plugins/sdk/schema/fieldname"
)

func Token() schema.CredentialType {
	return schema.CredentialType{
		Name:          credname.Token, // TODO: Register name in project://sdk/schema/credname/names.go
		DocsURL:       sdk.URL("https://railway.com/docs/token"), // TODO: Replace with actual URL
		ManagementURL: sdk.URL("https://console.railway.com/user/security/tokens"), // TODO: Replace with actual URL
		Fields: []schema.CredentialField{
			{
				Name:                fieldname.Token,
				MarkdownDescription: "Token used to authenticate to Railway.",
				Secret:              true,
				Composition: &schema.ValueComposition{
					Length: 308,
					Prefix: "rw_", // TODO: Check if this is correct
					Charset: schema.Charset{
						Uppercase: true,
						Lowercase: true,
						Digits:    true,
					},
				},
			},
		},
		DefaultProvisioner: provision.EnvVars(defaultEnvVarMapping),
		Importer: importer.TryAll(
			importer.TryEnvVarPair(defaultEnvVarMapping),
			TryRailwayConfigFile(),
		)}
}

var defaultEnvVarMapping = map[string]sdk.FieldName{
	"RAILWAY_TOKEN": fieldname.Token, // TODO: Check if this is correct
}

// TODO: Check if the platform stores the Token in a local config file, and if so,
// implement the function below to add support for importing it.
func TryRailwayConfigFile() sdk.Importer {
	return importer.TryFile("~/.railway/config.json", func(ctx context.Context, contents importer.FileContents, in sdk.ImportInput, out *sdk.ImportAttempt) {
		var config struct {
			User struct {
				Token string `json:"token"`
			} `json:"user"`
		}
		if err := contents.ToJSON(&config); err != nil {
			out.AddError(err)
			return
		}

		if config.User.Token == "" {
			return
		}

		out.AddCandidate(sdk.ImportCandidate{
			Fields: map[sdk.FieldName]string{
				fieldname.Token: config.User.Token,
			},
		})
	})
}
