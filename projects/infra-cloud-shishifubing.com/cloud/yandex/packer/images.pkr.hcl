locals {
  base = "debian-11-base"
  source = "debian-11"
  keys = "${path.root}/authorized_keys"
  init_base = templatefile("image_base.cloud-init.yml", {
    key_main    = file("${local.keys}/id_main.pub")
    key_personal = file("${local.keys}/id_rsa.pub")
    server_user = var.user_server
  })
}

source "yandex" "debian-11-base" {
  # it has to be a path, you cannot just use a key
  service_account_key_file = var.authorized_key_path
  folder_id            = var.folder_id
  zone                 = var.zone
  use_ipv4_nat         = var.use_nat
  ssh_username         = var.user_server
  ssh_private_key_file = var.ssh_key_connection

  disk_type           = "network-hdd"
  image_description   = "base debian image with configured ssh and users"
  image_family        = local.base
  # this resourse name cannot contain dots
  image_name          = "${local.base}-${var.version}"
  source_image_family = local.source
  metadata = {
    user-data = join("\n", [
      local.init_base
    ])
  }
  serial_log_file = "image_base.log"
}

build {
  sources = ["source.yandex.${local.base}"]

  provisioner "shell" {
    scripts = [
      "image_init_wait.sh",
      var.setup_script_path
    ]
  }

  post-processor "manifest" {
    output = "image_base.manifest.json"
    strip_path = true
  }
}
