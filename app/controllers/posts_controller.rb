class PostsController < ApplicationController
  def show
    @post = Post.find(params[:id])
  end

  def index
    # Take the most recent 10 by default
    @posts = Post.all.reverse.take(10)
  end

  def new
  end
end
