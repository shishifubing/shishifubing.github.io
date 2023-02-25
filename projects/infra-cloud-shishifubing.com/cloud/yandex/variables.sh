#!/usr/bin/env bash

dir="${HOME}/Credentials/yc"
version=$(jq -r ".FullSemVer" <(gitversion))
version="${version//./-}"

# second - secret key
# fist line in the file - key id
AWS_ACCESS_KEY_ID=$(sed -n 1p "${dir}/terraform-state-manager-token.txt")
AWS_SECRET_ACCESS_KEY=$(sed -n 2p "${dir}/terraform-state-manager-token.txt")

TF_VAR_authorized_key=$(<"${dir}/authorized_key.main.json")
TF_VAR_authorized_key_bucket=$(<"${dir}/authorized_key.bucket.json")
TF_VAR_static_key_id_bucket=$(sed -n 1p "${dir}/bucket-admin.txt")
TF_VAR_static_key_bucket=$(sed -n 2p "${dir}/bucket-admin.txt")
TF_VAR_kubernetes_config=$(<"${HOME}/.kube/config")

PKR_VAR_authorized_key_path="${dir}/authorized_key.main.json"
PKR_VAR_version="${version}"

export AWS_SECRET_ACCESS_KEY AWS_ACCESS_KEY_ID TF_VAR_authorized_key         \
    PKR_VAR_authorized_key_path PKR_VAR_oauth_key PKR_VAR_version            \
    TF_VAR_authorized_key_bucket YC_STORAGE_ACCESS_KEY YC_STORAGE_SECRET_KEY \
    TF_VAR_static_key_bucket TF_VAR_static_key_id_bucket                     \
    TF_VAR_kubernetes_config