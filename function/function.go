// Package function contains an HTTP Cloud Function.
package function

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

var (
	token    string
	channels []string
	userIds  []string
	reHere   = regexp.MustCompile(`<!(here)>`)
	re       = regexp.MustCompile(`<!(here|channel|everyone)>`)
)

const (
	message = "<@%s> please do not use `%s` - instead mention a specific `@team` or user."
)

func init() {
	token = os.Getenv("WEBHOOK_TOKEN")
	if token == "" {
		log.Println("no WEBHOOK_TOKEN specified, processing all requests")
	}

	chs := os.Getenv("CHANNEL_NAMES")
	if chs == "" {
		log.Println("no CHANNEL_NAMES specified, monitoring all channels")
	} else {
		channels = strings.Split(chs, ",")
	}

	ids := os.Getenv("ALLOWED_USER_IDS")
	if ids == "" {
		log.Println("no ALLOWED_USER_IDS specified, monitoring all users")
	} else {
		userIds = strings.Split(ids, ",")
	}
}

// HereHandler responds to requests from Slack that mention @here.
func HereHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if token != "" && r.Form.Get("token") != token {
		http.Error(w, "Not authenticated", http.StatusForbidden)
		return
	}

	user := r.Form.Get("user_id")
	trigger := r.Form.Get("trigger_word")
	channel := r.Form.Get("channel_name")

	// Ignore unmonitored channels
	if ignoreChannel(channel, channels) {
		log.Printf("user '%s' used '%s' in '#%s' but it is not monitored", user, trigger, channel)
		return
	}

	// Ignore @here for allowed users
	if reHere.MatchString(trigger) && ignoreUser(user, userIds) {
		log.Printf("user '%s' used '%s' in '#%s' but is allowed", user, trigger, channel)
		return
	}

	if re.MatchString(trigger) {
		log.Printf("user '%s' used '%s' in '#%s'", user, trigger, channel)

		data := make(map[string]string)
		data["text"] = fmt.Sprintf(message, user, re.ReplaceAllString(trigger, "@$1"))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data)
	}
}

func ignoreChannel(channel string, channels []string) bool {
	return len(channels) > 0 && !inSet(channel, channels)
}

func ignoreUser(user string, userIds []string) bool {
	return len(userIds) > 0 && inSet(user, userIds)
}

func inSet(item string, set []string) bool {
	for _, i := range set {
		if item == i {
			return true
		}
	}

	return false
}
