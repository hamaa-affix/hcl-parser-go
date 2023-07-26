resource "datadog_monitor" "keepalive" {
  name    = "[EC2] Linux: 死活監視"
  type    = "service check"
  message = "hoge"
  query   = "hogew"
#   monitor_thresholds {
#     ok       = 1
#     warning  = 1
#     critical = 1
#   }
  notify_no_data    = true
  no_data_timeframe = 2
  new_host_delay    = 300
  renotify_interval = 0
  timeout_h         = 0
  include_tags      = true
  notify_audit      = false
  tags              = ["service:hofe"]
}
