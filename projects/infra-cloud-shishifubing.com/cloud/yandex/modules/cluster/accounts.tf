# accounts

resource "yandex_iam_service_account" "ingress" {
  name        = "cluster-ingress"
  description = "service account for the ingress (terraform)"
}

resource "yandex_iam_service_account" "vault" {
  name        = "cluster-vault"
  description = "service account for the vault (terraform)"
}

# roles

resource "yandex_kms_symmetric_key_iam_binding" "vault" {
  symmetric_key_id = yandex_kms_symmetric_key.vault.id
  role             = "kms.keys.encrypterDecrypter"
  members = [
    "serviceAccount:${yandex_iam_service_account.vault.id}",
  ]
}

resource "yandex_resourcemanager_folder_iam_binding" "storage_editor" {
  folder_id = var.folder_id
  role      = "storage.editor"
  members = [
    "serviceAccount:${yandex_iam_service_account.vault.id}"
  ]
}

resource "yandex_resourcemanager_folder_iam_binding" "alb_editor" {
  folder_id = var.folder_id
  role      = "alb.editor"
  members = [
    "serviceAccount:${yandex_iam_service_account.ingress.id}"
  ]
}

resource "yandex_resourcemanager_folder_iam_binding" "vpc_publicAdmin" {
  folder_id = var.folder_id
  role      = "vpc.publicAdmin"
  members = [
    "serviceAccount:${yandex_iam_service_account.ingress.id}"
  ]
}

resource "yandex_resourcemanager_folder_iam_binding" "certificate_admin" {
  folder_id = var.folder_id
  role      = "certificate-manager.admin"
  members = [
    "serviceAccount:${yandex_iam_service_account.ingress.id}"
  ]
}


resource "yandex_resourcemanager_folder_iam_binding" "compute_viewer" {
  folder_id = var.folder_id
  role      = "compute.viewer"
  members = [
    "serviceAccount:${yandex_iam_service_account.ingress.id}"
  ]
}

###
