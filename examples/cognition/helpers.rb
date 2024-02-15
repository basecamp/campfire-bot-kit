require "json"
require "pathname"

def extract_json_from(request)
  request.body.rewind
  JSON.parse(request.body.read)
end
