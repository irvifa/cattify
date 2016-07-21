require 'sinatra'

get '/' do
  logger.info "loading..."
  "the beauty of pure logic and math.. :cat:"
end
