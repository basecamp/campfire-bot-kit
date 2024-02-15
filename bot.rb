require "bundler/setup"
require "sinatra"
require_relative "helpers"

# Healthcheck for Kamal
get("/up") { "✓" }

# A bot endpoint that responds with a text greeting
post "/hello" do
  content = extract_json_from(request)
  "Hey #{content.dig("user", "name")}, <i>you said</i>: #{content.dig("message", "body", "plain")}"
end

# A bot endpoint that responds with a random cat image
post "/cat" do
  files = Dir.glob("./assets/images/*.jpg")

  content_type "image/jpg"
  send_file files.sample, disposition: "inline"
end
