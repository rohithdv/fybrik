package dataapi.authz
import data.data_policies as dp

# METADATA
# scope: rule
# schemas:
#   - input: schema["input-schema"]
transform[action] {
	description = "Columns with Confidential tag to be redacted before read action"
    #user context and access type check
    dp.check_access_type(["READ"])
	dp.check_intent("Fraud Detection")
	dp.check_role("Data Scientist")
	dp.dataset_has_tag("residency = Turkey")
	dp.check_processingGeo_not("Turkey")
    column_names := dp.column_with_tag("Confidential")
    action = dp.build_redact_column_action(column_names[_], dp.build_policy_from_description(description))
}

deny[action] {
	description = "Deny if role is not Data Scientist when intent is Fraud Detection"
    #user context and access type check
    dp.check_access_type(["READ"])
	dp.check_intent("Fraud Detection")
	dp.check_role_not("Data Scientist")
	dp.dataset_has_tag("residency = Turkey")
    action = dp.build_deny_access_action(dp.build_policy_from_description(description))
}

deny[action] {
	description = "If columns have Confidential tag deny read action"
    #user context and access type check
    dp.check_access_type(["READ"])
	dp.check_intent("Customer Behaviour Analysis")
	dp.check_role("Business Analyst")
	dp.dataset_has_tag("residency = Turkey")
    dp.column_has_tag("Confidential")
    action = dp.build_deny_access_action(dp.build_policy_from_description(description))
}

deny[action] {
	description = "Deny if role is not Business Analyst when intent is Customer Behaviour Analysis"
    #user context and access type check
    dp.check_access_type(["READ"])
	dp.check_intent("Customer Behaviour Analysis")
	dp.check_role_not("Business Analyst")
	dp.dataset_has_tag("residency = Turkey")
	dp.check_processingGeo_not("Turkey")
    action = dp.build_deny_access_action(dp.build_policy_from_description(description))
}


deny[action] {
	description = "Deny if role is Data Scientist and intent is Fraud Detection but the processing geography is not Trukey"
    #user context and access type check
    dp.check_access_type(["READ"])
	dp.check_intent("Fraud Detection")
	dp.check_role_not("Data Scientist")
	dp.dataset_has_tag("residency = Turkey")
	dp.check_processingGeo_not("Turkey")
    action = dp.build_deny_access_action(dp.build_policy_from_description(description))
}

deny[action] {
	description = "If data residency is Turkey but processing geography is not Turkey then deny writing"
    #user context and access type check
    dp.check_access_type(["WRITE"])
	dp.dataset_has_tag("residency = Turkey")
	dp.check_processingGeo_not("Turkey")
    action = dp.build_deny_write_action(dp.build_policy_from_description(description))
}

deny[action] {
	description = "If data residency is not Turkey and processing geography is neither Turkey nor EEA then deny writing"
    #user context and access type check
    dp.check_access_type(["WRITE"])
	dp.dataset_has_tag_not("residency = Turkey")
	dp.check_processingGeo_not("Turkey")
	dp.check_processingGeo_not("EEA")
    action = dp.build_deny_write_action(dp.build_policy_from_description(description))
}