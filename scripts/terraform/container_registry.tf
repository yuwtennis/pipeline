resource "google_container_registry" "registry" {
  location = "ASIA"
}

resource "google_storage_bucket_iam_member" "viewer" {
  bucket = google_container_registry.registry.id
  role = "roles/storage.objectViewer"
  member = "serviceAccount:462754264937-compute@developer.gserviceaccount.com"
}