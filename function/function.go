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
	re       = regexp.MustCompile(`<!(here|channel|everyone)>`)
)

const (
	message = "<@%s> please do not use `%s` - instead mention a specific `@team` or user."
)

func init() {
	token = os.Getenv("WEBHOOK_TOKEN")
	if token == "" {
		log.Println("no WEBHOOK_TOKEN specified, this will process all requests")
	}

	chs := os.Getenv("CHANNEL_NAMES")
	if chs == "" {
		log.Println("no CHANNEL_NAMES specified, this will monitor all channels")
	} else {
		channels = strings.Split(chs, ",")
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

	channel := r.Form.Get("channel_name")
	user := r.Form.Get("user_id")
	trigger := r.Form.Get("trigger_word")

	monitored := false
	for _, c := range channels {
		if channel == c {
			monitored = true
			break
		}
	}

	if len(channels) > 0 && !monitored {
		log.Printf("user '%s' used '%s' in '#%s' but it is not monitored", user, trigger, channel)
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

func inSet(item string, set []string) bool {
	for _, i := range set {
		if item == i {
			return true
		}
	}

	return false
}
