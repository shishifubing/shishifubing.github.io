locals {
  keys_path             = "${path.module}/packer/authorized_keys"
  ssh_authorized_keys   = <<-EOT
    root:${file("${local.keys_path}/id_ci.pub")}
    root:${file("${local.keys_path}/id_main.pub")}
    root:${file("${local.keys_path}/id_rsa.pub")}
  EOT
  bucket_terraform_name = "${replace(var.domain, ".", "-")}-terraform"
  bucket_vault_name     = "${replace(var.domain, ".", "-")}-vault"
  cluster_id            = module.main.cluster.id
}

# manage the s3 bucket
module "bucket" {
  source = "./modules/bucket"
  providers = {
    yandex = yandex.bucket
  }

  folder_id          = var.folder_id_bucket
  terraform_name     = local.bucket_terraform_name
  vault_name         = local.bucket_vault_name
  vault_max_size     = 1024 * 1024 * 300 # 300 megabytes
  terraform_max_size = 1024 * 1024 * 300 # 300 megabytes
}

# setup infrastructure, create a kubernetes cluster
module "main" {
  source = "./modules/main"

  folder_id           = var.folder_id
  cloud_id            = var.cloud_id
  authorized_key_path = var.authorized_key
  zone                = var.zone
  domain              = var.domain
  domain_internal     = var.domain_internal
  domain_top_redirect = var.domain_top_redirect
  kubernetes_version  = var.kubernetes_version
  ssh_authorized_keys = local.ssh_authorized_keys
  user_server         = var.user_server
}

# setup the kubernetes cluster
module "cluster" {
  source = "./modules/cluster"

  folder_id                = var.folder_id
  cluster_id               = local.cluster_id
  zone                     = var.zone
  vault_bucket_name        = local.bucket_vault_name
  vault_backend_static_key = module.bucket.vault_editor_static_key
}
