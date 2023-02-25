terraform {
  required_providers {
    kubernetes = {
      source = "hashicorp/kubernetes"
    }
    helm = {
      source = "hashicorp/helm"
    }
    yandex = {
      source = "yandex-cloud/yandex"
    }
  }
}
