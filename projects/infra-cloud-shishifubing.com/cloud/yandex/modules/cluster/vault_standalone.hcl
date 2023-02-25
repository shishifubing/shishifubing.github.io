  ui = true

  listener "tcp" {
    tls_disable = 1
    address = "[::]:8200"
    cluster_address = "[::]:8201"
  }

  storage "s3" {
    endpoint                    = "storage.yandexcloud.net"
  }

  seal "yandexcloudkms" {
    kms_key_id     = "{{ .Values.yandexKmsKeyId }}"
    service_account_key_file      = "/vault/userconfig/kms-creds/credentials.json"
  }