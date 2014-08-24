module ApplicationHelper
  # Returns the full title of the page
  def full_title(page_title)
    ret = "John Anthony"
    if not page_title.empty?
      ret = ret + " | " + page_title
    end
    return ret
  end 
end
