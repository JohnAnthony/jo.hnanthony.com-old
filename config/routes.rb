Rails.application.routes.draw do
  get 'posts/new'

  resources :posts
  get 'welcome/index'
  root 'welcome#index'
end
