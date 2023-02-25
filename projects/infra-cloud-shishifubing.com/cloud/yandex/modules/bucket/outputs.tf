output "vault_editor_static_key" {
  description = "static access key for vault editor"
  sensitive   = true
  value       = yandex_iam_service_account_static_access_key.vault_editor
}
