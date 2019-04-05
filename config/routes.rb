Rails.application.routes.draw do
  get '/', to: 'build#index'
  get '/build/build_yml', to: 'build#build_yml'
end
