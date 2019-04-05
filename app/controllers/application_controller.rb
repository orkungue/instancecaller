class ApplicationController < ActionController::Base
  def build
    render file: "test"  
  end
end
