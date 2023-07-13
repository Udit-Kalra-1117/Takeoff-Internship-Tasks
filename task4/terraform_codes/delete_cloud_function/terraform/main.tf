//pushing the object into the already created bucket
resource "google_storage_bucket_object" "delete-code" {
    name = "delete-by-ID-cloud-function-code"
    bucket = "rest-crud_api-function-code-bucket"
    source = "../artifacts/delete_function.zip"
}

//creating a gen 2 cloud function
resource "google_cloudfunctions2_function" "delete-function" {
    name = "delete_element_by_ID_function"
    location = var.region
    description = "function to delete an entry based on it ID from the database"

    build_config {
        runtime = "go120"
        entry_point = "deleteEMP"
        source {
            storage_source {
                bucket = "rest-crud_api-function-code-bucket"
                object = google_storage_bucket_object.delete-code.name
            }
        }
    }
    service_config {
        max_instance_count = 1
        available_memory = "256M"
        timeout_seconds = 60
    }
}

//creating an all users invoker function
resource "google_cloudfunctions2_function_iam_member" "invoker" {
    location = google_cloudfunctions2_function.delete-function.location
    cloud_function = google_cloudfunctions2_function.delete-function.name
    role = "roles/cloudfunctions.invoker"
    member = "allUsers"
}