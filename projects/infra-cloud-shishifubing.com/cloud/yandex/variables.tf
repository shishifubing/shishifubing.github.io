variable "folder_id" {
  description = "folder id in Yandex Cloud"
  default     = "b1gt03ltb5rv3iieiuv9"
}

variable "folder_id_bucket" {
  description = "bucket folder id in Yandex Cloud"
  default     = "b1gk0hd1lkhbe6qj3u2s"
}

variable "cloud_id" {
  description = "cloud id in yandex cloud"
  default     = "b1g1vmf1q6kf4ap0d710"
}

variable "authorized_key" {
  description = "authorized key for yandex cloud"
  sensitive   = true
  validation {
    condition     = can(jsondecode(var.authorized_key))
    error_message = "must be JSON"
  }
}

variable "authorized_key_bucket" {
  description = "authorized key for the bucket"
  sensitive   = true
  validation {
    condition     = can(jsondecode(var.authorized_key_bucket))
    error_message = "must be JSON"
  }
}

variable "static_key_bucket" {
  description = "secret key for the bucket admin account"
  sensitive   = true
}

variable "static_key_id_bucket" {
  description = "key id for the bucket admin account"
  sensitive   = true
}

variable "zone" {
  description = "yandex cloud zone"
  default     = "ru-central1-a"
}

variable "domain" {
  description = "domain for public Cloud DNS"
  default     = "shishifubing.com"
}

variable "user_server" {
  default = "shishifubing"
}

variable "domain_internal" {
  description = "domain for internal Cloud DNS"
  default     = "internal"
}

variable "domain_top_redirect" {
  description = "redirect for the top public DNS domain (CNAME)"
  default     = "shishifubing-com.github.io"
}

variable "kubernetes_version" {
  description = "kubernetes version"
  default     = "1.23"
}

variable "kubernetes_config_path" {
  description = "path to the kubernetes config file"
  default     = "~/.kube/config"
}
