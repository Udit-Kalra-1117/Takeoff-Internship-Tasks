output "function_uri" {
    value = google_cloudfunctions2_function.update-function.service_config[0].uri
}