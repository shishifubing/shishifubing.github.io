variable "namespaces" {
  description = "dictionary of namespaces for services"
  default     = {}
  type = object({
    vault      = optional(string, "vault")
    ingress    = optional(string, "ingress")
    monitoring = optional(string, "monitoring")
  })
}

variable "vault_bucket_name" {
  description = "vault's s3 bucket name"
}

variable "vault_backend_static_key" {
  description = "s3 static key for vault"
  sensitive   = true
}

variable "zone" {
  description = "yandex cloud zone"
}

variable "folder_id" {
  description = "yandex cloud folder id"
}

variable "cluster_id" {
  description = "yandex cloud cluster id"
}
