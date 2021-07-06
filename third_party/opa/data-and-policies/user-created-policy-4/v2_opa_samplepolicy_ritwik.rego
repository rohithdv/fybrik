package dataapi.authz

verdict[output] {
	count(rule) == 0
	output = {"action": {"name":"DenyAccess"}, "policy": "Deny by default"}
}
verdict[output] {
	count(rule) > 0
	output = rule[_]
}
rule[{"action": {"name":"RedactColumn", "columns": column_names}, "policy": description}] {
	description := "Columns with Confidential tag to be redacted before read action"
    #user context and access type check
	input.action.action_type == "read"
    input.request_context.intent == "Fraud Detection"
	input.request_context.role == "Data Scientist"
	input.resource.tags.tags.residency == "Turkey"
	input.action.processingLocation != "Turkey"
	column_names := [input.resource.columns[i].name | input.resource.columns[i].tags.tags.Confidential == "true"]
}
rule[{"action": {"name":"DenyAccess"}, "policy": description}] {
	description := "Deny because the role is not Data Scientist when intent is Fraud Detection"
    #user context and access type check
    input.action.action_type == "read"
	input.request_context.intent == "Fraud Detection"
	input.request_context.role != "Data Scientist"
	input.resource.tags.tags.residency == "Turkey"
}




