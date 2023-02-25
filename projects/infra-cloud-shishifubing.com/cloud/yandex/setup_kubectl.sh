#!/usr/bin/env bash
set -Eeuxo pipefail

domain="${1:-$(terraform output -raw domain)}"
cloud_id="${2:-$(terraform output -raw cloud_id)}"
folder_id="${3:-$(terraform output -raw folder_id)}"
cluster_id="${4:-$(terraform output -raw cluster_id)}"

master_domain="master.${domain}"
credentials="${HOME}/Credentials/yc"
ca="${credentials}/${cluster_id}-ca.pem"
token="${credentials}/${cluster_id}-token-admin.txt"
admin="${cluster_id}-admin"

# get certificate authority
yc managed-kubernetes cluster get                          \
    "${cluster_id}"                                        \
    --cloud-id "${cloud_id}"                               \
    --folder-id "${folder_id}"                             \
    --format json                                          |
        jq -r ".master.master_auth.cluster_ca_certificate" |
        awk '{gsub(/\\n/,"\n")}1'                          \
    >"${ca}"

# setup kubectl
kubectl config set-cluster "${cluster_id}" \
    --certificate-authority="${ca}"        \
    --embed-certs                          \
    --server="https://${master_domain}"    \
    --tls-server-name="kubernetes"

kubectl config set-credentials "${admin}"                     \
    --exec-command="/usr/bin/yc"                              \
    --exec-arg="k8s"                                          \
    --exec-arg="create-token"                                 \
    --exec-arg="--profile=default"                            \
    --exec-api-version="client.authentication.k8s.io/v1beta1"

kubectl config set-context "${cluster_id}" \
    --cluster="${cluster_id}"              \
    --user="${admin}"

# check connectivity
kubectl config use-context "${cluster_id}"
kubectl cluster-info

# create static admin token
kubectl apply -f - <<EOF
apiVersion: v1
kind: ServiceAccount
metadata:
  name: admin-user
  namespace: kube-system
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: admin-user
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cluster-admin
subjects:
  - kind: ServiceAccount
    name: admin-user
    namespace: kube-system
EOF

secret=$(
    kubectl -n kube-system get secret |
        grep "admin-user" |
        awk '{print $1}'
)
set +x
token=$(
    kubectl -n kube-system get secret "${secret}" -o json |
        jq -r .data.token |
        base64 --d
)
echo "${token}" >"${HOME}/Credentials/yc/${cluster_id}-token-admin.txt"
set -x