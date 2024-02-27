require "cgi"

# https://docs.rollbar.com/docs/webhooks
class RollbarWebhook
  # Static mapping so we don't need to make a request to rollbar API just for the name
  PROJECTS = {
    111 => "Your project name"
  }

  def initialize(content)
    @content = content
  end

  def formatted_message
    case event_name
    when "new_item", "exp_repeat_item", "item_velocity", "occurrence", "reactivated_item", "reopened_item", "resolved_item"
      %Q[<a href="#{url}" target="_blank">#{title}</a><br>🐞 <em>#{humanized_event_name} on</em> #{project_name} #{environment}]
    when "deploy"
      username = @content.dig("data", "deploy", "local_username")
      revision = @content.dig("data", "deploy", "revision")[...7]
      "🚀 #{username} deployed <strong>#{project_name}</strong> to <strong>#{environment}</strong> (rev: #{revision})"
    when "test"
      @content.dig("data", "message")
    end
  end

  private

  def event_name = @content.fetch("event_name")
  def url = @content.dig("data", "url")
  def environment = @content.dig("data", "item", "environment") || @content.dig("data", "deploy", "environment")
  def title = CGI.escapeHTML(@content.dig("data", "item", "title"))
  def occurrences = @content.dig("data", "item", "occurrences")
  def project_id = @content.dig("data", "item", "project_id") || @content.dig("data", "deploy", "project_id")
  def project_name = PROJECTS[project_id] || "Unknown: #{project_id}"

  def humanized_event_name
    case event_name
    when "new_item" then "New error"
    when "reactivated_item" then "Reactivated error"
    when "reopened_item" then "Reopened error"
    when "resolved_item" then "Resolved error"
    when "exp_repeat_item" then "#{occurrences}th error"
    when "item_velocity", "occurrence" then event_name
    end
  end
end
