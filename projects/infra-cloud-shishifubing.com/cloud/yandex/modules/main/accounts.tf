resource "yandex_iam_service_account" "editor" {
  # used for cluster
  name        = "main-editor"
  description = "editor service account (terraform)"
}

resource "yandex_iam_service_account" "viewer" {
  # used for bastion
  name        = "main-viewer"
  description = "viewer service account (terraform)"
}

resource "yandex_iam_service_account" "admin" {
  name        = "main-admin"
  description = "admin service account (terraform)"
}

resource "yandex_resourcemanager_folder_iam_binding" "admin" {
  folder_id = var.folder_id
  role      = "admin"
  members = [
    "serviceAccount:${yandex_iam_service_account.admin.id}"
  ]
}

resource "yandex_resourcemanager_folder_iam_binding" "editor" {
  folder_id = var.folder_id
  role      = "editor"
  members = [
    "serviceAccount:${yandex_iam_service_account.editor.id}"
  ]
}

resource "yandex_resourcemanager_folder_iam_binding" "viewer" {
  folder_id = var.folder_id
  role      = "viewer"
  members = [
    "serviceAccount:${yandex_iam_service_account.viewer.id}"
  ]
}
