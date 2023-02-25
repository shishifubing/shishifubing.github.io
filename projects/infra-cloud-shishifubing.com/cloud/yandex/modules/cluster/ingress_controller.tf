resource "helm_release" "ingress_controller" {
  # https://cloud.yandex.ru/marketplace/products/yc/alb-ingress-controller
  name        = "ingress-controller"
  description = "nginx ingress controller"

  namespace       = var.namespaces.ingress
  atomic          = true
  cleanup_on_fail = true
  lint            = true

  version    = "v0.1.9"
  chart      = "chart"
  repository = "oci://cr.yandex/yc-marketplace/yandex-cloud/yc-alb-ingress"

  # I cannot use `set` or `set_sensitive` - helm provider fucks up the formatting
  values = [
    "saKeySecretKey: '${jsonencode(yandex_iam_service_account_key.ingress)}'"
  ]

  set {
    name  = "folderId"
    value = var.folder_id
  }

  set {
    name  = "clusterId"
    value = var.cluster_id
  }

}
