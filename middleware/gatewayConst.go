package middleware

const (

	/*
		- demo
		message SaveTemplateRequest {
		  string GatewayUID = 7;
		  string GatewayOrgID = 8;
	*/

	GatewayUID_Field = "GatewayUID"
	GatewayUID       = "X-POLARIS-IDENTITY-USER"

	GatewayOrgID_Field = "GatewayOrgID"
	GatewayOrgID       = "X-POLARIS-IDENTITY-ORG"
)
