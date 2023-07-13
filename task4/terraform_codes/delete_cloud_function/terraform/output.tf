output "function_uri" {
    value = google_cloudfunctions2_function.delete-function.service_config[0].uri
}