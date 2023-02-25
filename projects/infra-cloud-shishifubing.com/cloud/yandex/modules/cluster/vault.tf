locals {
  vault_key = jsonencode({
    # I have to do it because of
    # key unmarshal fail: unknown field "pgp_key" in yandex.cloud.sdk.v1.Key
    id                 = yandex_iam_service_account_key.vault.id
    service_account_id = yandex_iam_service_account_key.vault.service_account_id
    created_at         = yandex_iam_service_account_key.vault.created_at
    key_algorithm      = yandex_iam_service_account_key.vault.key_algorithm
    public_key         = yandex_iam_service_account_key.vault.public_key
    private_key        = yandex_iam_service_account_key.vault.private_key
  })
}

resource "helm_release" "vault" {
  # https://cloud.yandex.ru/marketplace/products/yc/vault-yckms-k8s
  name        = "vault"
  description = "vault"

  namespace       = var.namespaces.vault
  atomic          = true
  cleanup_on_fail = true
  lint            = true

  version    = "0.19.0-1"
  chart      = "vault"
  repository = "oci://cr.yandex/yc-marketplace/yandex-cloud/vault/chart"

  # I cannot use `set` or `set_sensitive` - helm provider fucks up the formatting
  values = [
    "yandexKmsAuthJson: '${local.vault_key}'"
  ]

  set {
    name  = "ui.enabled"
    value = true
  }

  set_sensitive {
    name  = "yandexKmsKeyId"
    value = yandex_kms_symmetric_key.vault.id
  }

  set_sensitive {
    name = "extraSecretEnvironmentVars"
    value = jsonencode({
      AWS_S3_BUCKET         = var.vault_bucket_name
      AWS_ACCESS_KEY_ID     = var.vault_backend_static_key.access_key
      AWS_SECRET_ACCESS_KEY = var.vault_backend_static_key.secret_key
      AWS_REGION            = var.zone
    })
  }

}
