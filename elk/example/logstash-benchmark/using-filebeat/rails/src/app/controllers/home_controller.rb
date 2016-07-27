class HomeController < ApplicationController

	def index

		@greeting = "i'm nyan cat"
		render json: @greeting
	end	
end
