require "bundler/setup"
require "sinatra"
require_relative "helpers"

# Healtcheck for Kamal
get("/up") { "✓" }

get "/" do
  "Hello world!"
end

post "/hello" do
  content = extract_json_from(request)
  "Hey #{content.dig("user", "name")}, <i>you said</i>: #{content.dig("message", "body", "plain")}"
end

post "/image" do
  content_type "image/png"
  send_file Pathname.new(__dir__).realpath.join("assets/images/campfire.png"), disposition: "inline"
end
