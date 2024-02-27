# Campfire Bot for handling Rollbar webhooks

Handles webhook notifications described on https://docs.rollbar.com/docs/webhooks

## Usage

`kamal envify --skip-push`
Fill the .env with required env variables.
ROOM_URL is the bot room with secret token where the message should be posted

Deploy by `kamal deploy`

Point your Rollbar webhook url to "https://chatbot.example.com/rollbar" (replace domain with your own)
