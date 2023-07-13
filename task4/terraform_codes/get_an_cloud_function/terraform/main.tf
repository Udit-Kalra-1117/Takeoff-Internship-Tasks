//pushing the object into the already created bucket
resource "google_storage_bucket_object" "get_an-code" {
    name = "get-an-cloud-function-code"
    bucket = "rest-crud_api-function-code-bucket"
    source = "../artifacts/get_an_function.zip"
}

//creating a gen 2 cloud function
resource "google_cloudfunctions2_function" "get-an-function" {
    name = "get_an_element_function"
    location = var.region
    description = "function to get an entry by passing the ID parameter from the database"

    build_config {
        runtime = "go120"
        entry_point = "getAnEMP"
        source {
            storage_source {
                bucket = "rest-crud_api-function-code-bucket"
                object = google_storage_bucket_object.get_an-code.name
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
    location = google_cloudfunctions2_function.get-an-function.location
    cloud_function = google_cloudfunctions2_function.get-an-function.name
    role = "roles/cloudfunctions.invoker"
    member = "allUsers"
}