package function

import "log"

// MonitoredChannels logs the monitored channels.
func MonitoredChannels() {
	log.Printf("monitoring channels: %s\n", channels)
}
