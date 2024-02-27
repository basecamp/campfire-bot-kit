require "bundler/setup"
require "sinatra"
require "uri"
require "net/https"
require_relative "rollbar_webhook"

configure do
  set :server_settings, max_threads: 2
end

helpers do
  def json_body(request)
    request.body.rewind
    JSON.parse(request.body.read)
  end
end

get("/up") { "✓" } # Healthcheck for Kamal

post "/rollbar" do
  webhook = RollbarWebhook.new(json_body(request))

  Net::HTTP.post(URI.parse(ENV["ROOM_URL"]), webhook.formatted_message)

  status 204
end

not_found do
  "This is nowhere to be found"
end

error do
  "Something went wrong"
end
