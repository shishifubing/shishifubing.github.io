resource "yandex_kms_symmetric_key" "vault" {
  name              = "cluster-vault"
  description       = "a key for the vault"
  default_algorithm = "AES_128"
  rotation_period   = "8760h" // equal to 1 year
}

resource "yandex_iam_service_account_key" "ingress" {
  service_account_id = yandex_iam_service_account.ingress.id
  description        = "authorized key for the cluster ingress (terraform)"
  key_algorithm      = "RSA_4096"
}


resource "yandex_iam_service_account_key" "vault" {
  service_account_id = yandex_iam_service_account.vault.id
  description        = "authorized key for the vault (terraform)"
  key_algorithm      = "RSA_4096"
}
