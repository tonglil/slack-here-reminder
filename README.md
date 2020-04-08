# Slack `@here` Reminder

This is a Slack Outgoing Webhook integration to respond to `@here`, `@channel`, and `@everyone` usage to reminder the user to mention specific users or teams instead.

## Configuration

The following environment variables can be set:

- `PORT`: specify a port *when running as a web service* (optional)
- `WEBHOOK_TOKEN`: specify the token from the Outgoing Webhook to verify requests from Slack (optional)
- `CHANNEL_NAMES`: specify the channel names to monitor and remind, separated by commas (optional)

## Deploying

The app can be run as a listening server (`main.go`) or as a function (`function/function.go`).
