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

resource "google_compute_firewall" "elastic_fw" {
  name          = "elastic"
  network       = google_compute_network.dev.self_link
  target_tags   = ["elastic_fw"]
  source_ranges = ["0.0.0.0/0"]

  dynamic "allow" {
    for_each = var.ingress_allow
    content {
      protocol = allow.value["protocol"]
      ports    = allow.value["ports"]
    }
  }
}