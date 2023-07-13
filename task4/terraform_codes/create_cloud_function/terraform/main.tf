//creating a bucket
resource "google_storage_bucket" "rest-crud_api-bucket" {
    name = "rest-crud_api-function-code-bucket"
    location = var.location
    storage_class = "STANDARD"
    uniform_bucket_level_access = true
    labels = {
        "key1" = "value1"
        "key2" = "value2"
    }
}

//pushing the object into the created bucket
resource "google_storage_bucket_object" "create-code" {
    name = "create-cloud-function-code"
    bucket = google_storage_bucket.rest-crud_api-bucket.name
    source = "../artifacts/create_function.zip"
}

//creating a gen 2 cloud function
resource "google_cloudfunctions2_function" "create-function" {
    name = "create_element_function"
    location = var.region
    description = "function to create an entry into the database"

    build_config {
        runtime = "go120"
        entry_point = "createEMP"
        source {
            storage_source {
                bucket = google_storage_bucket.rest-crud_api-bucket.name
                object = google_storage_bucket_object.create-code.name
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
    location = google_cloudfunctions2_function.create-function.location
    cloud_function = google_cloudfunctions2_function.create-function.name
    role = "roles/cloudfunctions.invoker"
    member = "allUsers"
}