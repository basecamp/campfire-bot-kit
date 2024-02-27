# Campfire Bot Kit

Creating a Campfire bot is as simple as replying to a request with either text
or an image. This reply will be posted directly into the room from which the bot
was either mentioned or pinged. Alternatively, you can have your bot post to a
room on its own accord by using the bot-specific URLs that are listed in
Campfire in the bot section.

This repository contains a simple Ruby bot that can be used as a starting point
for your own bot.

## Developing with Ruby

You'll need a working Ruby environment to build the example bot. We recommend
using [rbenv](https://github.com/rbenv/rbenv) to install and manage your Ruby
versions.

If you're using macOS, you can install rbenv using Homebrew:

```sh
brew install rbenv
```

Installation instructions for other platforms are available on the [rbenv GitHub
page](https://github.com/rbenv/rbenv).

Once you have rbenv installed, you can install the required Ruby version by
running:

```sh
rbenv install
```

## The sample bot

Our sample actually implements two bots, so we can demonstrate both text and
image responses. We've defined two endpoints, one for each bot:

- `/hello` responds to its messages with a text greeting
- `/cat` responds to any messages it receives with a random cat image

The `/hello` endpoint also shows how to access the content of the message that
was sent to the bot.

## How Campfire bots work

Campfire bots can receive and respond to messages sent by other users, and they
can also post messages to rooms on their own accord.

### Receiving and responding to messages

When a user mentions or pings a bot, Campfire sends a POST request to the bot's
webhook URL. That request includes a JSON payload with the details of the
message content, the room it was sent from, and the user who sent it.

Here's an example of a valid message payload:

```json
{
  "room": {
    "id": 23,
    "name": "All Talk"
  },
  "user": {
    "id": 42,
    "name": "Kevin"
  },
  "message": {
    "id": 100,
    "body": {
      "html": "<p>hello</p>",
      "plain": "hello"
    }
  }
}
```

The bot should return a successful HTTP status code to acknowledge that it
received the message. It can also include a response, which will be posted to
the room the message was sent from.

To respond with a text message, simply return the text. This is what the
`/hello` endpoint does in our example bot.

To respond with an attachment, like an image, return its content, and make sure
to set the `Content-Type` header to the appropriate MIME type. This is what our
`/cat` endpoint does.

### Posting messages to rooms

To send messages to a room that aren't in response to a message, make an HTTP
request to the bot's room-specific URL.  Each bot gets unique URLs for every
room that it's a member of. The authentication token is included in the URL, so
all you need to do is make the request.

To send a text message, make a POST request to the room's URL with the message
in the request body. To send an attachment, use a `multipart/form-data` request
with the attachment as the `attachment` field.

In the bot section of the Campfire UI you'll see example `curl` commands for
sending to each room in both text and attachment format.

## Deploying your bot

You can deploy the sample bot to try it out. We've added some
[Kamal](https://kamal-deploy.org) configuration to the repo, so all that's left
to do is provide the details of your server and Docker registry credentials. You
can use an existing server to deploy to, or you can create one with a cloud
provider like Digital Ocean or Hetzner.

Most of the deployment details will go into the `config/deploy.yml` file. The
registry credentials should be kept separate though, to make sure they don't end
up in the repository; those can be added to a `.env` file instead.

For a typical setup, the steps you'll need to do are:

- Add your server's IP address or DNS name to the `servers` section of `config/deploy.yml`
- Put your Docker username in the `registry` section. If you're using a registry
  other than Docker Hub, you'll need to add the registry's URL as well.
- Create a `.env` file to hold your registry password. You can use an access
  token here if you have one. Your `.env` file should look something like this:

```
KAMAL_REGISTRY_PASSWORD=yourpassword
```

You can now run `kamal setup` to set up your server. Each time you want to
deploy your bot, just run `kamal deploy`.

Depending on your particular setup you might need some other configuration. For
example, to use a different port, or to use a user other than `root` to connect.
You can find all the available options in the
[Kamal documentation](https://kamal-deploy.org/docs/configuration).

Once your bot has been deployed, add it to your Campfire instance. In the bot
section of the Campfire UI you can give your bot a name, and set its webhook URL
to be an endpoint of your bot (like `http://mybot.example.com/cats`). Then try
sending a message to your bot to see it respond.

## Exploring other examples

Check the [examples](./examples) directory for some more examples that use
different languages, libraries or frameworks.

## Share your bots!

If you create a bot that you think others might find useful, we'd love to see
it! Feel free to open a pull request to add it to a `community` directory in
this repo so we can include it for others to use.

## License

The Campfire Bot Kit is licensed under the MIT license.

Campfire itself uses a commercial license: https://once.com/license
