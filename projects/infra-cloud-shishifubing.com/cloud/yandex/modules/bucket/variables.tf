variable "folder_id" {
  description = "yandex cloud folder_id"
}

variable "terraform_name" {
  description = "name of the terraform bucket"
}

variable "terraform_max_size" {
  description = "max size of the bucket (bytes)"
  type        = number
}

variable "vault_name" {
  description = "name of the terraform bucket"
}

variable "vault_max_size" {
  description = "max size of the bucket (bytes)"
  type        = number
}
