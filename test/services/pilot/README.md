# To Test With Mocks

Start kubernetes cluster (say kind)
```bash
kind create cluster
```
Install Fybrik
```bash
helm repo add jetstack https://charts.jetstack.io
helm repo add hashicorp https://helm.releases.hashicorp.com
helm repo add fybrik-charts https://fybrik.github.io/charts
helm repo update

helm install cert-manager jetstack/cert-manager \
    --namespace cert-manager \
    --version v1.2.0 \
    --create-namespace \
    --set installCRDs=true \
    --wait --timeout 120s

kubectl create ns fybrik-system
```
Prepare a secret  named  wkc-credentials in fybrik-system namespace with wkc credentials: (the values should be replaced)

```bash
cat << EOF | kubectl apply -f -
apiVersion: v1
kind: Secret
metadata:
  name: wkc-credentials
  namespace: fybrik-system
type: kubernetes.io/Opaque
stringData:
  CP4D_USERNAME: "<CP4D_USERNAME>"
  CP4D_PASSWORD: "<CP4D_PASSWORD>"
  CP4D_SERVER_URL: "<CP4D_URL>"
EOF
```

Install Vault

```bash
helm repo add hashicorp https://helm.releases.hashicorp.com
helm install vault hashicorp/vault --version 0.9.1 --create-namespace -n fybrik-system \
    --set "server.dev.enabled=true" \
    --values values-for-vault.yaml \
    --wait --timeout 120s
```
where values-for-vault.yaml is from https://github.ibm.com/data-mesh-research/vault-plugin-secrets-wkc-reader/blob/master/helm-deployment/vault-single-cluster/values.yaml

```bash
wget https://raw.githubusercontent.com/IBM/the-mesh-for-data/a3f951087eada4aed4b1cee9390bed5d71c35970/third_party/vault/vault-single-cluster/vault-rbac.yaml
kubectl apply -f vault-rbac.yaml -n fybrik-system
```

Install Fybrik
```bash
git clone https://github.com/fybrik/fybrik.git
cd fybrik
helm install fybrik-crd charts/fybrik-crd -n fybrik-system --wait
helm install fybrik charts/fybrik --set global.tag=latest -n fybrik-system --wait

kubectl apply -f https://github.com/fybrik/arrow-flight-module/releases/latest/download/module.yaml -n fybrik-system
```

Install WKC connector (Ref: https://github.ibm.com/data-mesh-research/WKC-connector)


Define data access policies

Define an OpenPolicyAgent policy to redact the nameOrig column for datasets tagged as finance. Below is the policy (written in Rego language):

```bash
package dataapi.authz

import data.data_policies as dp

transform[action] {
  description := "Redact sensitive columns in finance datasets"
  dp.AccessType() == "READ"
  dp.check_intent("Fraud Detection")
  column_names := dp.column_with_any_name({"nameOrig"})
  action = dp.build_redact_column_action(column_names[_], dp.build_policy_from_description(description))
}
```

In this sample only the policy above is applied. Copy the policy to a file named sample-policy.rego and then run:

```bash
kubectl -n fybrik-system create configmap sample-policy --from-file=sample-policy.rego
kubectl -n fybrik-system label configmap sample-policy openpolicyagent.org/policy=rego
while [[ $(kubectl get cm sample-policy -n fybrik-system -o 'jsonpath={.metadata.annotations.openpolicyagent\.org/policy-status}') != '{"status":"ok"}' ]]; do echo "waiting for policy to be applied" && sleep 5; done
```

WKC Configuration:

Upload data.csv to COS. data.csv contains the first 100 rows from the following data set created by NTNU, and it is shared under the CC BY-SA 4.0 license. Please check https://fybrik.io/v0.4/samples/notebook/ under the section "Prepare a dataset to be accessed by the notebook".

In WKC: add a connection to the bucket in COS that stores that data and an asset from the connection.
Please fill access_key and secret_key in the connection details.

in WKC: add a tag to the asset:
residency = Netherlands

(with a spaces arround "=")

```bash
kubectl create namespace cp4d
kubectl apply -f cp4d-secret.yaml
kubectl apply -f wkc-credentials-fybrik-system.yaml
```
where  cp4d-secret.yaml is
```bash
apiVersion: v1
data:
  WKC_ownerId: <OWNERID_BASE64ENCODED>
  WKC_password: <WKCPASS_BASE64ENCODED>
  WKC_username: <WKCUSER_BASE64ENCODED>
kind: Secret
metadata:
  name: wkc-creds
  namespace: cp4d
```
<OWNERID_BASE64ENCODED> is got
```bash
echo <OWNERID> | base64
```
Similarly for <WKCPASS_BASE64ENCODED> and <WKCUSER_BASE64ENCODED>


where wkc-credentials-fybrik-system.yaml is
```bash
apiVersion: v1
kind: Secret
metadata:
  name: wkc-credentials
  namespace: fybrik-system
type: kubernetes.io/Opaque
stringData:
  CP4D_USERNAME: <CP4DUSERNAME>
  CP4D_PASSWORD: <CP4DPASS>
  CP4D_SERVER_URL: "<CP4DURL>"
```


Now go to test/services/pilot folder
Change the mock files appropriately

```bash
make dockerall
make deploy
```