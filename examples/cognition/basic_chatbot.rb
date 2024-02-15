class BasicChatbot < Cognition::Plugins::Base
  # Matchers are evaluated in order. First one wins.

  match /.*/, :echo, help: {
    '<anything>' => 'Echoes back the message'
  }

  def echo(msg, match_data = nil)
    pp msg
    msg.command
  end
end
