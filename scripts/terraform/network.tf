resource "google_compute_network" "dev" {
  name                    = "vpc-dev"
  description             = ""
  auto_create_subnetworks = false
  routing_mode            = "REGIONAL"
}

resource "google_compute_subnetwork" "primary" {
  name          = "dev-primary"
  region        = "asia-northeast1"
  ip_cidr_range = "192.168.1.0/24"
  network       = google_compute_network.dev.id
}

# nat
resource "google_compute_router" "dev" {
  name    = "my-router"
  region  = google_compute_subnetwork.primary.region
  network = google_compute_network.dev.id
}

resource "google_compute_router_nat" "nat" {
  name                               = "my-router-nat"
  router                             = google_compute_router.dev.name
  region                             = google_compute_router.dev.region
  nat_ip_allocate_option             = "AUTO_ONLY"
  source_subnetwork_ip_ranges_to_nat = "ALL_SUBNETWORKS_ALL_IP_RANGES"

  log_config {
    enable = true
    filter = "ERRORS_ONLY"
  }
}

# firewall
resource "google_compute_firewall" "elastic_ingress_fw" {
  name          = "elastic-ingress"
  network       = google_compute_network.dev.self_link
  target_tags   = ["elastic-ingress-fw"]
  source_ranges = ["0.0.0.0/0"]
  direction     = "INGRESS"

  dynamic "allow" {
    for_each = var.ingress_allow
    content {
      protocol = allow.value["protocol"]
      ports    = allow.value["ports"]
    }
  }
}

resource "google_compute_firewall" "elastic_egress_fw" {
  name          = "elastic-egress"
  network       = google_compute_network.dev.self_link
  target_tags   = ["elastic-egress-fw"]
  source_ranges = ["0.0.0.0/0"]
  direction     = "EGRESS"

  dynamic "allow" {
    for_each = var.egress_allow
    content {
      protocol = allow.value["protocol"]
      ports    = allow.value["ports"]
    }
  }
}