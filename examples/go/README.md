# Go bot example

This directory contains an example of a simple Campfire bot written in Go.

This bot implments a single endpoint, `/trace`. You can message this endpoint
with a URL. The bot will make a `GET` request to that URL, and respond with some
timings about how long parts of that request took: DNS lookup, time to first
byte, and so on.

The functionality of the bot is basic. But this example shows how you can:

- Start an HTTP service to listen on a bot endpoint
- Parse the JSON from the message request
- Respond to that request with some HTML-formatted text
