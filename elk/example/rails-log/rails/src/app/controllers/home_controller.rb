require 'logstash-logger'

class HomeController < ApplicationController

	def index
		logger = LogStashLogger.new(uri: ENV['LOGSTASH_URI'])

		@greeting = "i'm nyan cat"
		logger.info("inside the index")
		render json: @greeting
	end	
end
