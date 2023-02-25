terraform {

  backend "s3" {}

  required_providers {
    yandex = {
      source  = "yandex-cloud/yandex"
      version = "0.84.0"
    }
    kubernetes = {
      source  = "hashicorp/kubernetes"
      version = "2.16.1"
    }
    helm = {
      source  = "hashicorp/helm"
      version = "2.8.0"
    }
  }
  required_version = ">= 0.13"
}

provider "yandex" {
  zone                     = var.zone
  folder_id                = var.folder_id
  cloud_id                 = var.folder_id
  service_account_key_file = var.authorized_key

}

provider "yandex" {
  # For extended API usage, such as setting max_size, folder_id, anonymous_access_flags,
  # default_storage_class and https parameters for bucket, will be used default
  # authorization method, i.e. IAM / OAuth token from provider block will be used.
  # This might be a little bit confusing in cases when separate service account
  # is used for managing buckets because in this case buckets will be accessed
  # by two different accounts that might have different permissions for buckets.
  alias                    = "bucket"
  folder_id                = var.folder_id_bucket
  cloud_id                 = var.cloud_id
  service_account_key_file = var.authorized_key_bucket
  storage_access_key       = var.static_key_id_bucket
  storage_secret_key       = var.static_key_bucket
}

provider "kubernetes" {
  config_path    = pathexpand(var.kubernetes_config_path)
  config_context = local.cluster_id
}

provider "helm" {
  kubernetes {
    config_path    = pathexpand(var.kubernetes_config_path)
    config_context = local.cluster_id
  }

  registry {
    url      = "osi://cr.yandex"
    username = "json_key"
    password = var.authorized_key
  }
}
