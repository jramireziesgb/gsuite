package gsuite

import (
	"fmt"
	"math/rand"
	s "strings"
	"time"
)

// User Registro para almacenar los alumnos obtenidos de GSuite
type User struct {
	FirstName                  string
	LastName                   string
	EmailAddress               string
	Password                   string
	PasswordHashFunction       string
	OrgUnitPath                string
	NewPrimaryEmail            string
	RecoveryEmail              string
	HomeSecondaryEmail         string
	WorkSecondaryEmail         string
	RecoveryPhone              string
	WorkPhone                  string
	HomePhone                  string
	MobilePhone                string
	WorkAddress                string
	HomeAddress                string
	EmployeeID                 string
	EmployeeType               string
	EmployeeTitle              string
	ManagerEmail               string
	Department                 string
	CostCenter                 string
	BuildingID                 string
	FloorName                  string
	FloorSection               string
	ChangePasswordatNextSignIn string
	NewStatus                  string
}

// Gesuser Registro para almacenar los alumnos obtenidos de Séneca
type GesUser struct {
	Alumno                      string
	EstadoMatricula             string
	NIdEscolar                  string
	DNI                         string
	Direccion                   string
	Codigopostal                string
	Localidadderesidencia       string
	Fechadenacimiento           string
	Provinciaderesidencia       string
	Telefono                    string
	Telefonodeurgencia          string
	CorreoElectronico           string
	Curso                       string
	Ndelexpedientedelcentro     string
	Unidad                      string
	Primerapellido              string
	Segundoapellido             string
	Nombre                      string
	DNIPrimerturor              string
	PrimerapellidoPrimertutor   string
	SegundoapellidoPrimertutor  string
	NombrePrimertutor           string
	SexoPrimertutor             string
	DNISegundotutor             string
	PrimerapellidoSegundotutor  string
	SegundoapellidoSegundotutor string
	NombreSegundotutor          string
	SexoSegundotutor            string
	Localidaddenacimiento       string
	Añodelamatricula            string
	Ndematriculasenestecurso    string
	Observacionesdelamatricula  string
	Provincianacimiento         string
	Paisdenacimiento            string
	Edad                        string
	Nacionalidad                string
	Sexo                        string
	Fechadematricula            string
	NSegSocial                  string
}

func init() {
	rand.Seed(time.Now().Unix())
}

func createUserName(firstName string, lastName string) string {

	r := s.NewReplacer("á", "a", "é", "e", "í", "i", "ó", "o", "ú", "u", "ü", "u", "ñ", "n")
	first := r.Replace(s.ToLower(s.TrimSpace(firstName)))
	last := r.Replace(s.ToLower(s.TrimSpace(lastName)))

	firsts := s.Split(first, " ")
	lasts := s.Split(last, " ")

	username := ""

	for _, v := range firsts {
		username = username + string(v[0])
	}

	return username + lasts[0]
}

func createPassword(passwd string) string {

	return passwd
}

// NewUser Crea un usuario nuevo
func (a *User) NewUser(usuario *GesUser, sufix string, dominio string, unidad string) {

	if sufix != "" {
		sufix = "." + sufix
	}

	a.FirstName = usuario.Nombre
	a.LastName = usuario.Primerapellido + " " + usuario.Segundoapellido
	a.EmailAddress = createUserName(usuario.Nombre, usuario.Primerapellido) + sufix + "@" + dominio
	a.Password = createPassword(usuario.DNI)
	a.PasswordHashFunction = ""
	a.OrgUnitPath = unidad
	a.NewPrimaryEmail = ""
	a.RecoveryEmail = ""
	a.HomeSecondaryEmail = ""
	a.WorkSecondaryEmail = ""
	a.RecoveryPhone = ""
	a.WorkPhone = ""
	a.HomePhone = ""
	a.MobilePhone = ""
	a.WorkAddress = ""
	a.HomeAddress = ""
	a.EmployeeID = ""
	a.EmployeeType = ""
	a.EmployeeTitle = ""
	a.ManagerEmail = ""
	a.Department = ""
	a.CostCenter = ""
	a.BuildingID = ""
	a.FloorName = ""
	a.FloorSection = ""
	a.ChangePasswordatNextSignIn = "True"
	a.NewStatus = "Active"
}

func (a *User) String() string {

	return fmt.Sprintf("%s, %s, %s, %s, %s", a.FirstName, a.LastName, a.EmailAddress, a.Password, a.OrgUnitPath)
}
