//pushing the object into the already created bucket
resource "google_storage_bucket_object" "get_all-code" {
    name = "get-all-cloud-function-code"
    bucket = "rest-crud_api-function-code-bucket"
    source = "../artifacts/get_all_function.zip"
}

//creating a gen 2 cloud function
resource "google_cloudfunctions2_function" "get-all-function" {
    name = "get_all_elements_function"
    location = var.region
    description = "function to get all entries from the database"

    build_config {
        runtime = "go120"
        entry_point = "getAllEMP"
        source {
            storage_source {
                bucket = "rest-crud_api-function-code-bucket"
                object = google_storage_bucket_object.get_all-code.name
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
    location = google_cloudfunctions2_function.get-all-function.location
    cloud_function = google_cloudfunctions2_function.get-all-function.name
    role = "roles/cloudfunctions.invoker"
    member = "allUsers"
}