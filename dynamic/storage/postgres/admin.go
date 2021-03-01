package postgres

const (
	adminTable     = "admins"
	adminUserIDCol = "user_id"
)

type adminConstsType struct {
	Table     string
	UserIDCol string
}

var adminConsts = &adminConstsType{
	Table:     "admins",
	UserIDCol: "user_id",
}
