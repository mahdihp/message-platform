package shared

type TenantType string
type TenantStatus string

const (
	Platform TenantType = "platform"
	Reseller TenantType = "reseller"
	Customer TenantType = "customer"

	Draft     TenantStatus = "draft"
	Pending   TenantStatus = "pending"
	Verified  TenantStatus = "verified"
	Suspended TenantStatus = "suspended"
)
