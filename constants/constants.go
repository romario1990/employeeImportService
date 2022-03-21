package constants

// HEADER if changed update the constants POSITIONHEADEREMAIL e POSITIONHEADERIDENTIFIER
var HEADER = [][]string{[]string{"Name", "Email", "Salary", "Identifier", "Phone", "Mobile"}}
var HEADERCONFIGURATION = []string{FULLNAME, FIRSTNAME, MIDDLENAME, LASTNAME, EMAIL, SALARY, IDENTIFIER, PHONE, MOBILE, INVALIDNAME}

const (
	POSITIONHEADEREMAIL      = 1
	POSITIONHEADERIDENTIFIER = 3
)

// file names
const (
	SUCCESSPATH          = "transfer/success/"
	SUCCESSNAMEFILE      = "employeesucess"
	SUCCESSPATHNAME      = SUCCESSPATH + SUCCESSNAMEFILE + EXTENSIONCSV
	ERRORPATH            = "transfer/error/"
	ERRORNAMEFILE        = "employeeinvalid"
	ERRORPATHNAME        = ERRORPATH + ERRORNAMEFILE + EXTENSIONCSV
	PATHPROCESSED        = "./transfer/processed/"
	PATHPROCESSEDERROR   = "./transfer/processedError/"
	PATHPPENDINGROCESSED = "./transfer/pending/"
)

// constants
const (
	FULLNAME     = "FullName"
	FIRSTNAME    = "FirstName"
	MIDDLENAME   = "MiddleName"
	LASTNAME     = "LastName"
	EMAIL        = "Email"
	SALARY       = "Salary"
	IDENTIFIER   = "Identifier"
	PHONE        = "Phone"
	MOBILE       = "Mobile"
	INVALIDNAME  = ""
	CSV          = "csv"
	EXTENSIONCSV = ".csv"
)

//[command] [flags]
const (
	// FILE is name of the file flag
	FILE      = "file"
	SHORTFILE = "f"
	// HASHEADER is the name of the flag for the header
	HASHEADER      = "header"
	SHORTHASHEADER = "e"
	// FILETYPE set file type in import
	FILETYPE      = "file type"
	SHORTFILETYPE = "t"
)
