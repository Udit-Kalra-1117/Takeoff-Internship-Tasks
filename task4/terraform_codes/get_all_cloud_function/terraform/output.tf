output "function_uri" {
    value = google_cloudfunctions2_function.get-all-function.service_config[0].uri
}