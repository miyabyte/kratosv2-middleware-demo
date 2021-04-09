**GwIjMiddleware:**

将网关填充在header中的`X-POLARIS-IDENTITY-USER、X-POLARIS-IDENTITY-ORG`
  注入req的字段`GatewayUID、GatewayOrgID` 中。

使用`rotoc-go-inject-tag` 配合generate做标准化，从tag中取可避免强行依赖字段名做注入的不合理性。