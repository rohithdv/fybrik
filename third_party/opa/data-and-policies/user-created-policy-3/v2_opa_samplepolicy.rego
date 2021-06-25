package dataapi.authz

import data.data_policies as dp

#Example of data policies that use "data_policies" package to create easily data policies that deny access or transform the data accordingly

transform[action] {
	description = "location data should be removed before copy"
    #user context and access type check
    input.request_context.intent == "Fraud Detection"
    dp.dataset_has_tag("Finance")
    column_names := dp.column_with_any_tag(["SPI", "SMI"])
    action = dp.build_redact_column_action(column_names[_], dp.build_policy_from_description(description))
}

#for transactions dataset
deny[action] {
	description = "test for transactions dataset with deny"

    #user context and access type check
    input.action.action_type == "read"

    action = dp.build_deny_access_action(dp.build_policy_from_description(description))
}