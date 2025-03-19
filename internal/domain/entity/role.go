package entity

const (
	admin  = "admin"
	client = "client"
)

type Role string

const (
	ADMIN  Role = admin
	CLIENT Role = client
)

func (r Role) String() string {
	return string(r)
}

const SecurityDataKey = "securityData"
