resource "google_compute_instance" "elastic" {
  name         = "elastic"
  machine_type = "n2-standard-2"
  zone         = "asia-northeast1-a"

  boot_disk {
    initialize_params {
      image = "ubuntu-os-cloud/ubuntu-minimal-2204-lts"
    }
  }

  attached_disk {
    source = google_compute_disk.elastic_data.self_link
  }

  network_interface {
    subnetwork = google_compute_subnetwork.primary.name
  }

  metadata = {
    user-data       = file("${path.module}/user-data")
    elastic-version = var.es_version
  }

  tags = ["elastic-ingress-fw", "elastic-egress-fw"]
}

resource "google_compute_disk" "elastic_data" {
  name = "elastic-data"
  type = "pd-ssd"
  zone = "asia-northeast1-a"
  size = "20"
}