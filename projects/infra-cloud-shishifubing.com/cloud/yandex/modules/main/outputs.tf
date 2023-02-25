output "user_server" {
  value = var.user_server
}

output "cluster" {
  description = "kubernetes cluster object"
  value       = yandex_kubernetes_cluster.default
}
