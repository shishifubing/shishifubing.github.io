
locals {
  master_balancer_external = yandex_vpc_address.cluster_master.external_ipv4_address.0.address
  master_fqdm              = "master.${var.domain}"
  zone_fixed               = substr(var.zone, 0, length(var.zone) - 2)
}

# networks
resource "yandex_vpc_network" "default" {
  name        = "default"
  description = "default network"
}

resource "yandex_vpc_subnet" "default" {
  name           = "default"
  description    = "default subnet"
  network_id     = yandex_vpc_network.default.id
  v4_cidr_blocks = ["10.130.0.0/24"]
}

resource "yandex_vpc_subnet" "cluster" {
  name           = "cluster"
  description    = "cluster subnet"
  network_id     = yandex_vpc_network.default.id
  v4_cidr_blocks = ["10.129.0.0/24"]
  route_table_id = yandex_vpc_route_table.cluster.id
}

resource "yandex_vpc_gateway" "cluster" {
  name        = "cluster"
  description = <<-EOT
    gateway for the cluster nodes
    without it they do not have internet access
  EOT
  shared_egress_gateway {
  }
}

resource "yandex_vpc_route_table" "cluster" {
  name       = "cluster"
  network_id = yandex_vpc_network.default.id

  static_route {
    destination_prefix = "0.0.0.0/0"
    gateway_id         = yandex_vpc_gateway.cluster.id
  }
}

# public DNS
resource "yandex_dns_zone" "public" {
  name        = "public"
  description = "public dns zone"
  zone        = "${var.domain}."
  public      = true
}

resource "yandex_dns_recordset" "public" {
  zone_id = yandex_dns_zone.public.id
  name    = "@"
  type    = "CNAME"
  ttl     = 200
  data    = [var.domain_top_redirect]
}

resource "yandex_dns_recordset" "public_www" {
  zone_id = yandex_dns_zone.public.id
  name    = "www"
  type    = "CNAME"
  ttl     = 200
  data    = [var.domain_top_redirect]
}

resource "yandex_dns_recordset" "bastion" {
  zone_id = yandex_dns_zone.public.id
  name    = "bastion"
  type    = "A"
  ttl     = 200
  data    = [local.bastion_nat]
}

resource "yandex_dns_recordset" "master_external" {
  zone_id = yandex_dns_zone.public.id
  name    = "master"
  type    = "A"
  ttl     = 200
  data    = [local.master_balancer_external]
}
###

# internal DNS
resource "yandex_dns_zone" "internal" {
  name        = "internal"
  description = "internal dns zone"
  zone        = "${var.domain_internal}."
  public      = false
  private_networks = [
    yandex_vpc_network.default.id
  ]
}

resource "yandex_dns_recordset" "master" {
  zone_id = yandex_dns_zone.internal.id
  name    = "master"
  type    = "A"
  ttl     = 200
  data    = [local.master_ip]
}

###

# load balancer for the master
resource "yandex_lb_network_load_balancer" "cluster_master" {
  name      = "cluster-master-load-balancer"
  type      = "external"
  region_id = local.zone_fixed

  listener {
    name = "main"
    port = 443
    external_address_spec {
      address = local.master_balancer_external
    }
  }

  attached_target_group {
    target_group_id = yandex_lb_target_group.cluster_master.id

    healthcheck {
      name = "main"
      tcp_options {
        port = 443
      }
    }
  }
}

resource "yandex_vpc_address" "cluster_master" {
  name        = "cluster-master"
  description = "external IP address for the master"

  external_ipv4_address {
    zone_id = var.zone
  }
}

resource "yandex_lb_target_group" "cluster_master" {
  name        = "cluster-master"
  description = "target group for the master"
  region_id   = local.zone_fixed
  target {
    subnet_id = yandex_vpc_subnet.cluster.id
    address   = local.master_ip
  }
}
