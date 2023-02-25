resource "yandex_kms_symmetric_key" "cluster" {
  name              = "cluster"
  description       = "key for the default kubernetes cluster"
  default_algorithm = "AES_128"
  rotation_period   = "7000h" # ~10 month
}

resource "yandex_vpc_security_group" "allow_outgoing" {
  name        = "allow_outgoing"
  description = "allow all outgoing connections"
  network_id  = yandex_vpc_network.default.id

  egress {
    description    = "allow all outgoing connections"
    protocol       = "ANY"
    from_port      = 0
    to_port        = 65535
    v4_cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "yandex_vpc_security_group" "allow_incoming" {
  name        = "allow_incoming"
  description = "allow all incoming connections"
  network_id  = yandex_vpc_network.default.id

  ingress {
    description    = "allow all incoming connections"
    protocol       = "ANY"
    from_port      = 0
    to_port        = 65535
    v4_cidr_blocks = ["0.0.0.0/0"]
  }
}
