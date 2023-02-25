resource "yandex_storage_bucket" "terraform" {
  bucket     = var.terraform_name
  access_key = yandex_iam_service_account_static_access_key.admin.access_key
  secret_key = yandex_iam_service_account_static_access_key.admin.secret_key

  default_storage_class = "COLD"

  max_size = var.terraform_max_size

  versioning {
    enabled = true
  }

  grant {
    id          = yandex_iam_service_account.terraform_editor.id
    type        = "CanonicalUser"
    permissions = ["READ", "WRITE"]
  }

  # delete old versions after a month
  lifecycle_rule {
    id      = "clean"
    enabled = true

    noncurrent_version_expiration {
      days = 30
    }
  }
}

resource "yandex_storage_bucket" "vault" {
  bucket     = var.vault_name
  access_key = yandex_iam_service_account_static_access_key.admin.access_key
  secret_key = yandex_iam_service_account_static_access_key.admin.secret_key

  default_storage_class = "ICE" # "COLD"

  max_size = var.vault_max_size

  versioning {
    enabled = true
  }

  grant {
    id          = yandex_iam_service_account.vault_editor.id
    type        = "CanonicalUser"
    permissions = ["READ", "WRITE"]
  }

  # delete old versions after a month
  lifecycle_rule {
    id      = "clean"
    enabled = true

    noncurrent_version_expiration {
      days = 30
    }
  }
}
