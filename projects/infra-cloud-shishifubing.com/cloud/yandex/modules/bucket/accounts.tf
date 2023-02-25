resource "yandex_iam_service_account" "admin" {
  name        = "bucket-admin"
  description = "account that manages the bucket (terraform)"
}

resource "yandex_resourcemanager_folder_iam_binding" "admin" {
  folder_id = var.folder_id
  role      = "admin"
  members = [
    "serviceAccount:${yandex_iam_service_account.admin.id}"
  ]
}

resource "yandex_iam_service_account" "terraform_editor" {
  name        = "bucket-terraform-editor"
  description = <<EOT
    account that can download/upload files to the terraform bucket (terraform)
  EOT
}

resource "yandex_iam_service_account" "vault_editor" {
  name        = "bucket-vault-editor"
  description = <<EOT
    account that can download/upload files to the vault bucket (terraform)
  EOT
}

resource "yandex_resourcemanager_folder_iam_binding" "storage_editor" {
  folder_id = var.folder_id
  role      = "storage.editor"
  members = [
    "serviceAccount:${yandex_iam_service_account.terraform_editor.id}",
    "serviceAccount:${yandex_iam_service_account.vault_editor.id}"
  ]
}

resource "yandex_iam_service_account_static_access_key" "admin" {
  service_account_id = yandex_iam_service_account.admin.id
  description        = "static access key to manage buckets (terraform)"
}

resource "yandex_iam_service_account_static_access_key" "vault_editor" {
  service_account_id = yandex_iam_service_account.vault_editor.id
  description        = "static access key for vault (terraform)"
}
