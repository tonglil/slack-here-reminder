package function

import "log"

// MonitoredChannels logs the monitored channels.
func MonitoredChannels() {
	log.Printf("monitoring channels: %s\n", channels)
}

// AllowedUsers logs the allowed user ids.
func AllowedUsers() {
	log.Printf("allowing users: %s\n", userIds)
}
