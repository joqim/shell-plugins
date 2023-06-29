package railway

import (
	"testing"
	
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/plugintest"
	"github.com/1Password/shell-plugins/sdk/schema/fieldname"
)
	
func TestTokenProvisioner(t *testing.T) {
	plugintest.TestProvisioner(t, Token().DefaultProvisioner, map[string]plugintest.ProvisionCase{
		"default": {
			ItemFields: map[sdk.FieldName]string{ // TODO: Check if this is correct
				fieldname.Token: "rw_RyA8L0HN6TTqsXKTXuBzefBQDq4P4AiQMv5LBiWKwdqA9HpSrJchCc2EkTIb3KJSk9HkyMS0kt8K1W9Rqa1yJo8mRBIJx41T4qcOEPkEvvaiHUsc7pnh3ULl5BgfUNUbk4F7uPwn2qc0XX1sJFtZWOIWK3HdgBiXRfXXfmiOVT2NI8dtaMBCy0loPZuZKtHtgtB3KOL4yOxVBjUxEKP8kCEPGMZQG888xsa3EiOQ9Ov6r9tjOgcuASfxkVxwYKI3ZCTsJHmUofgPLprIVRHlDxIb7RRdC5BNYpt2lflBMzEXAMPLE",
			},
			ExpectedOutput: sdk.ProvisionOutput{
				Environment: map[string]string{
					"RAILWAY_TOKEN": "rw_RyA8L0HN6TTqsXKTXuBzefBQDq4P4AiQMv5LBiWKwdqA9HpSrJchCc2EkTIb3KJSk9HkyMS0kt8K1W9Rqa1yJo8mRBIJx41T4qcOEPkEvvaiHUsc7pnh3ULl5BgfUNUbk4F7uPwn2qc0XX1sJFtZWOIWK3HdgBiXRfXXfmiOVT2NI8dtaMBCy0loPZuZKtHtgtB3KOL4yOxVBjUxEKP8kCEPGMZQG888xsa3EiOQ9Ov6r9tjOgcuASfxkVxwYKI3ZCTsJHmUofgPLprIVRHlDxIb7RRdC5BNYpt2lflBMzEXAMPLE",
				},
			},
		},
	})
}

func TestTokenImporter(t *testing.T) {
	plugintest.TestImporter(t, Token().Importer, map[string]plugintest.ImportCase{
		"environment": {
			Environment: map[string]string{ // TODO: Check if this is correct
				"RAILWAY_TOKEN": "rw_RyA8L0HN6TTqsXKTXuBzefBQDq4P4AiQMv5LBiWKwdqA9HpSrJchCc2EkTIb3KJSk9HkyMS0kt8K1W9Rqa1yJo8mRBIJx41T4qcOEPkEvvaiHUsc7pnh3ULl5BgfUNUbk4F7uPwn2qc0XX1sJFtZWOIWK3HdgBiXRfXXfmiOVT2NI8dtaMBCy0loPZuZKtHtgtB3KOL4yOxVBjUxEKP8kCEPGMZQG888xsa3EiOQ9Ov6r9tjOgcuASfxkVxwYKI3ZCTsJHmUofgPLprIVRHlDxIb7RRdC5BNYpt2lflBMzEXAMPLE",
			},
			ExpectedCandidates: []sdk.ImportCandidate{
				{
					Fields: map[sdk.FieldName]string{
						fieldname.Token: "rw_RyA8L0HN6TTqsXKTXuBzefBQDq4P4AiQMv5LBiWKwdqA9HpSrJchCc2EkTIb3KJSk9HkyMS0kt8K1W9Rqa1yJo8mRBIJx41T4qcOEPkEvvaiHUsc7pnh3ULl5BgfUNUbk4F7uPwn2qc0XX1sJFtZWOIWK3HdgBiXRfXXfmiOVT2NI8dtaMBCy0loPZuZKtHtgtB3KOL4yOxVBjUxEKP8kCEPGMZQG888xsa3EiOQ9Ov6r9tjOgcuASfxkVxwYKI3ZCTsJHmUofgPLprIVRHlDxIb7RRdC5BNYpt2lflBMzEXAMPLE",
					},
				},
			},
		},
		// TODO: If you implemented a config file importer, add a test file example in railway/test-fixtures
		// and fill the necessary details in the test template below.
		"config file": {
			Files: map[string]string{
				// "~/path/to/config.yml": plugintest.LoadFixture(t, "config.yml"),
			},
			ExpectedCandidates: []sdk.ImportCandidate{
			// 	{
			// 		Fields: map[sdk.FieldName]string{
			// 			fieldname.Token: "rw_RyA8L0HN6TTqsXKTXuBzefBQDq4P4AiQMv5LBiWKwdqA9HpSrJchCc2EkTIb3KJSk9HkyMS0kt8K1W9Rqa1yJo8mRBIJx41T4qcOEPkEvvaiHUsc7pnh3ULl5BgfUNUbk4F7uPwn2qc0XX1sJFtZWOIWK3HdgBiXRfXXfmiOVT2NI8dtaMBCy0loPZuZKtHtgtB3KOL4yOxVBjUxEKP8kCEPGMZQG888xsa3EiOQ9Ov6r9tjOgcuASfxkVxwYKI3ZCTsJHmUofgPLprIVRHlDxIb7RRdC5BNYpt2lflBMzEXAMPLE",
			// 		},
			// 	},
			},
		},
	})
}
