require 'fileutils'

folder_path = './_posts'  # 替换为你的文件夹路径

Dir.glob(File.join(folder_path, '**', '*')).select { |file| File.file?(file) }.each do |file|
  content = File.read(file)
  content.gsub!('{% endhighlight %}', '')

end
