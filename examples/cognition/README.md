# A cognition-based chatbot

[Cognition](https://rubygems.org/gems/cognition) is a chatbot framework
developed at 37signals over the years. It provides a very simple method to
match text and run commands.

Most of this can be found in the [gem's README.md](https://github.com/basecamp/cognition/blob/master/README.md), but
we'll hit the high points.

First, let's add `cognition` to our app:

```shell
bundle add cognition
```

Next, we create a bot plugin. It's just a ruby class that inherits from `Cognition::Plugins::Base`. In it, we'll define a `match` rule, and the method that rule will call:

```ruby
class BasicChatbot < Cognition::Plugins::Base
  # Matchers are evaluated in order. First one wins.

  match /.*/, :echo, help: {
    '<anything>' => 'Echoes back the message'
  }

  def echo(msg, match_data = nil)
    msg.command
  end
end
```

Here, we have a single rule that matches everything via the regular expression
`/.*/`. It then sends that data to the `echo` method, which just replies with
the command text. Stick this into a file called `basic_chatbot.rb`, and let's
wire it into our app:

```ruby
require "cognition"
require_relative "basic_chatbot"

bot = Cognition::Bot.new
bot.register(BasicChatbot)
```

Once we have the chatbot in our app, we'll use it inside a sinatra endpoint:

```ruby
# A bot endpoint for a chatbot
post "/chatbot" do
  content = extract_json_from(request)
  command = content.dig("message", "body", "plain")
  metadata = content.slice("user", "room")

  bot.process(command, metadata)
end
```

This extracts the body content, and puts the plaintext rendering of the chat
message into the `command` variable, and the `user` and `room` data into a
`metadata` hash that we can optionally use inside our chatbot.

The last thing that happens is that we tell the bot to process the command,
and send the response back up to Campfire.
