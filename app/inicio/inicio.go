package inicio

import (
	"fmt"
	"log"

	"github.com/AlecAivazis/survey/v2"
	"golang.org/x/sys/windows"
)

// exporta essa funcao Inicio
func Insere_senha_inicio() string {
	fmt.Println("Estou na funcao Inicio")

	// Perguntar pela senha
	var senha string
	promptSenha := &survey.Password{
		Message: "Digite sua senha:",
	}
	survey.AskOne(promptSenha, &senha)

	if senha == "dsr@2017" {
		return "correta"
	} else {
		return "errada"
	}
	// return "estou fora do if"

}

func IsAdmin() (bool, error) {
	var sid *windows.SID

	// Although this looks scary, it is directly copied from the
	// official windows documentation. The Go API for this is a
	// direct wrap around the official C++ API.
	// See https://docs.microsoft.com/en-us/windows/desktop/api/securitybaseapi/nf-securitybaseapi-checktokenmembership
	err := windows.AllocateAndInitializeSid(
		&windows.SECURITY_NT_AUTHORITY,
		2,
		windows.SECURITY_BUILTIN_DOMAIN_RID,
		windows.DOMAIN_ALIAS_RID_ADMINS,
		0, 0, 0, 0, 0, 0,
		&sid)
	if err != nil {
		log.Fatalf("SID Error: %s", err)
		// return
	}
	defer windows.FreeSid(sid)

	// This appears to cast a null pointer so I'm not sure why this
	// works, but this guy says it does and it Works for Me™:
	// https://github.com/golang/go/issues/28804#issuecomment-438838144
	token := windows.Token(0)

	member, err := token.IsMember(sid)
	if err != nil {
		log.Fatalf("Token Membership Error: %s", err)
		// return
	}

	// Also note that an admin is _not_ necessarily considered
	// elevated.
	// For elevation see https://github.com/mozey/run-as-admin
	token.IsElevated()
	// fmt.Println("Elevated?", token.IsElevated())

	// fmt.Println("Admin?", member)
	return member, err
}
