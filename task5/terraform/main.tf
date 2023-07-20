# //creating a firestore database
# resource "google_firestore_database" "grocery_items_database" {
#     name = "(default)"
#     location_id = var.multi_region
#     type = "FIRESTORE_NATIVE"
# }

// creating a bucket for storing images
resource "google_storage_bucket" "udit-grocery-items-images-task5-bucket" {
    name = "udit-grocery-items-images-task5-bucket"
    location = var.location
    storage_class = "STANDARD"
    uniform_bucket_level_access = true
    force_destroy = true
    labels = {
        "key1" = "value1"
        "key2" = "value2"
    }
}

// setting the bucket to be publicly accessible
resource "google_storage_bucket_iam_member" "member" {
    bucket = google_storage_bucket.udit-grocery-items-images-task5-bucket.name
    role = "roles/storage.objectViewer"
    member = "allUsers"
}

// creating a bucket for storing cloud functions zip files
resource "google_storage_bucket" "udit-grocery-items-cloud-functions-task5-bucket" {
    name = "udit-grocery-items-cloud-functions-task5-bucket"
    location = var.location
    storage_class = "STANDARD"
    uniform_bucket_level_access = true
    force_destroy = true
    labels = {
        "key1" = "value1"
        "key2" = "value2"
    }
}

// pushing the create an item object into the created bucket
resource "google_storage_bucket_object" "create-an-item-code" {
    name = "create-an-item-cloud-function-code"
    bucket = google_storage_bucket.udit-grocery-items-cloud-functions-task5-bucket.name
    source = "../artifacts/create_an_item_function-source.zip"
}

// creating the create an item gen 2 cloud function
resource "google_cloudfunctions2_function" "create-an-item-function" {
    name = "create_an_item"
    location = var.region
    description = "function to create an item entry into the database"

    build_config {
        runtime = "go120"
        entry_point = "CreateITEM"
        source {
            storage_source {
                bucket = google_storage_bucket.udit-grocery-items-cloud-functions-task5-bucket.name
                object = google_storage_bucket_object.create-an-item-code.name
            }
        }
    }
    service_config {
        max_instance_count = 1
        available_memory = "256M"
        timeout_seconds = 60
    }
}

// pushing the bulk create items object into the created bucket
resource "google_storage_bucket_object" "bulk-create-items-code" {
    name = "bulk-create-items-cloud-function-code"
    bucket = google_storage_bucket.udit-grocery-items-cloud-functions-task5-bucket.name
    source = "../artifacts/bulk_create_items_function-source.zip"
}

// creating the bulk create items gen 2 cloud function
resource "google_cloudfunctions2_function" "bulk-create-items-function" {
    name = "bulk_create_items"
    location = var.region
    description = "function to bulk create item entries into the database"

    build_config {
        runtime = "go120"
        entry_point = "BulkCreateItemsUsingCSV"
        source {
            storage_source {
                bucket = google_storage_bucket.udit-grocery-items-cloud-functions-task5-bucket.name
                object = google_storage_bucket_object.bulk-create-items-code.name
            }
        }
    }
    service_config {
        max_instance_count = 1
        available_memory = "256M"
        timeout_seconds = 60
    }
}

// pushing the update an item object into the created bucket
resource "google_storage_bucket_object" "update-an-item-code" {
    name = "update-an-item-cloud-function-code"
    bucket = google_storage_bucket.udit-grocery-items-cloud-functions-task5-bucket.name
    source = "../artifacts/update_an_item_function-source.zip"
}

// creating the update an item gen 2 cloud function
resource "google_cloudfunctions2_function" "update-an-item-function" {
    name = "update_an_item"
    location = var.region
    description = "function to update an item entry into the database"

    build_config {
        runtime = "go120"
        entry_point = "UpdateITEM"
        source {
            storage_source {
                bucket = google_storage_bucket.udit-grocery-items-cloud-functions-task5-bucket.name
                object = google_storage_bucket_object.update-an-item-code.name
            }
        }
    }
    service_config {
        max_instance_count = 1
        available_memory = "256M"
        timeout_seconds = 60
    }
}

// pushing the delete an item object into the created bucket
resource "google_storage_bucket_object" "delete-an-item-code" {
    name = "delete-an-item-by-ID-cloud-function-code"
    bucket = google_storage_bucket.udit-grocery-items-cloud-functions-task5-bucket.name
    source = "../artifacts/delete_an_item_by_ID_function-source.zip"
}

// creating the delete an item gen 2 cloud function
resource "google_cloudfunctions2_function" "delete-an-item-function" {
    name = "delete_an_item"
    location = var.region
    description = "function to delete an item entry by its ID into the database"

    build_config {
        runtime = "go120"
        entry_point = "DeleteItemByID"
        source {
            storage_source {
                bucket = google_storage_bucket.udit-grocery-items-cloud-functions-task5-bucket.name
                object = google_storage_bucket_object.delete-an-item-code.name
            }
        }
    }
    service_config {
        max_instance_count = 1
        available_memory = "256M"
        timeout_seconds = 60
    }
}

// pushing the get an item by ID object into the created bucket
resource "google_storage_bucket_object" "get-an-item-code" {
    name = "get-an-item-by-ID-cloud-function-code"
    bucket = google_storage_bucket.udit-grocery-items-cloud-functions-task5-bucket.name
    source = "../artifacts/get_an_item_function-source.zip"
}

// creating the get an item by ID gen 2 cloud function
resource "google_cloudfunctions2_function" "get-an-item-function" {
    name = "get_an_item"
    location = var.region
    description = "function to get the details of an item entry by its ID from the database"

    build_config {
        runtime = "go120"
        entry_point = "GetAnItemByID"
        source {
            storage_source {
                bucket = google_storage_bucket.udit-grocery-items-cloud-functions-task5-bucket.name
                object = google_storage_bucket_object.get-an-item-code.name
            }
        }
    }
    service_config {
        max_instance_count = 1
        available_memory = "256M"
        timeout_seconds = 60
    }
}

// pushing the get all items by query object into the created bucket
resource "google_storage_bucket_object" "get-all-items-code" {
    name = "get-all-items-by-query-cloud-function-code"
    bucket = google_storage_bucket.udit-grocery-items-cloud-functions-task5-bucket.name
    source = "../artifacts/get_all_items.zip"
}

// creating the get all item by query gen 2 cloud function
resource "google_cloudfunctions2_function" "get-all-items-function" {
    name = "get_all_items"
    location = var.region
    description = "function to get the details of all item entries by query from the database"

    build_config {
        runtime = "go120"
        entry_point = "GetAllItemsByQuery"
        source {
            storage_source {
                bucket = google_storage_bucket.udit-grocery-items-cloud-functions-task5-bucket.name
                object = google_storage_bucket_object.get-all-items-code.name
            }
        }
    }
    service_config {
        max_instance_count = 1
        available_memory = "256M"
        timeout_seconds = 60
    }
}