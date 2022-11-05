package cvar_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vikpe/cparse/pkg/cvar"
)

func TestFromFilePath(t *testing.T) {
	t.Run("invalid file path", func(t *testing.T) {
		cvars, err := cvar.FromFile("./test_files/__INVALID_FILE_PATH__.c")
		expect := []cvar.Cvar{}
		assert.Equal(t, expect, cvars)
		assert.ErrorContains(t, err, "no such file or directory")
	})

	t.Run("valid file path", func(t *testing.T) {
		expect := []cvar.Cvar{
			{
				Name:         "sv_login",
				DefaultValue: "0",
				Description:  "if enabled, login required",
			},
			{
				Name:         "sv_login_web",
				DefaultValue: "1",
				Description:  "0=local files, 1=auth via website (bans can be in local files), 2=mandatory auth (must have account in local files)",
			},
			{
				Name:         "sv_local_addr",
				DefaultValue: "",
				Flags:        "CVAR_ROM",
			},
			{
				Name:         "pm_pground",
				DefaultValue: "",
				Flags:        "CVAR_SERVERINFO|CVAR_ROM",
			},
			{
				Name:         "qtv_pendingtimeout",
				DefaultValue: "5",
				Description:  "5  seconds must be enough",
			},
		}

		cvars, err := cvar.FromFile("./test_files/sv_login.c")
		assert.Equal(t, expect, cvars)
		assert.Nil(t, err)
	})
}

func TestFromLine(t *testing.T) {
	testCases := map[string]cvar.Cvar{
		"":                                  {},
		"foo":                               {},
		"foo sv_login":                      {},
		"cvar_t sv_login":                   {},
		"cvar_t sv_login {":                 {},
		"cvar_t sv_login }":                 {},
		"cvar_t sv_login {  }":              {},
		`foo sv_login = { "sv_login" };`:    {},
		`cvar_t sv_login = { "sv_login" };`: {Name: "sv_login"},
		`static cvar_t sv_login = { "sv_login" };`:                                     {Name: "sv_login"},
		`cvar_t sv_login = { "sv_login", "1" };`:                                       {Name: "sv_login", DefaultValue: "1"},
		`cvar_t sv_login = { "sv_login", "1", CVAR_ROM };`:                             {Name: "sv_login", DefaultValue: "1", Flags: "CVAR_ROM"},
		`cvar_t sv_login = { "sv_login", "1", CVAR_SERVERINFO | CVAR_ROM, callback };`: {Name: "sv_login", DefaultValue: "1", Flags: "CVAR_SERVERINFO | CVAR_ROM", OnChange: "callback"},
		`cvar_t sv_login = { "sv_login", "1" }; // allow logins`:                       {Name: "sv_login", DefaultValue: "1", Description: "allow logins"},
	}

	testCaseNumber := 1
	for line, expect := range testCases {
		t.Run(fmt.Sprintf("testcase %d", testCaseNumber), func(t *testing.T) {
			cv, _ := cvar.FromLine(line)
			assert.Equal(t, expect, cv)
			testCaseNumber++
		})
	}
}
