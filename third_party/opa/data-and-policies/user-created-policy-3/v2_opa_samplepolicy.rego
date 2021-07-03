package dataapi.authz

import data.data_policies as dp

#Example of data policies that use "data_policies" package to create easily data policies that deny access or transform the data accordingly

#for transactions dataset
deny[action] {
	description = "test for transactions dataset with deny"

    #user context and access type check
    input.action.action_type == "read"

    action = "{\"action\":\"DenyAccess\", \"policy\": \"Deny access to dataset\"}" 
}


#for transactions dataset
transform[action] {
	description = "test for transactions dataset with deny"

    #user context and access type check
    input.action.action_type == "read"

    action = "{\"action\": {\"name\": \"RemoveColumn\", \"columns\": [\"ABCD\",\"DEFG\"]}, \"policy\": \"Remove access to some columns in the dataset\"}" 
}