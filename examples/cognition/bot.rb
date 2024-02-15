require "bundler/setup"
require "sinatra"
require_relative "helpers"
require "cognition"
require_relative "basic_chatbot"

bot = Cognition::Bot.new
bot.register(BasicChatbot)

# Healthcheck for Kamal
get("/up") { "✓" }

# A bot endpoint for a chatbot
post "/chatbot" do
  content = extract_json_from(request)
  command = content.dig("message", "body", "plain")
  metadata = content.slice("user", "room")

  bot.process(command, metadata)
end

not_found do
  'This is nowhere to be found.'
end

error do
  'Something went wrong.'
end
