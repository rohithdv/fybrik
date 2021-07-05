package dataapi.authz

transform[action] {
	description := "Columns with Confidential tag to be redacted before read action"
    #user context and access type check
	input.action.action_type == "read"
    input.request_context.intent == "Fraud Detection"
	input.request_context.role == "Data Scientist"
	input.resource.tags.tags.residency == "Turkey"
	not input.action.processingLocation == "Turkey"
	column_names := [input.resource.columns[i].name | input.resource.columns[i].tags.tags.Confidential == "true"]
    action = {"action": {"name":"RedactColumn", "columns": column_names}, "policy": description}
}






