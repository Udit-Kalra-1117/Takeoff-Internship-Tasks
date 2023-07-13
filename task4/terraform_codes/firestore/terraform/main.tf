//creating a firestore database
resource "google_firestore_database" "employee_database" {
    name = "(default)"
    location_id = var.multi_region
    type = "FIRESTORE_NATIVE"
}