require 'fileutils'

folder_path = './_posts'  # 替换为实际的文件夹路径

# 获取文件夹中的所有文件
file_list = Dir.glob(File.join(folder_path, '*'))

# 遍历每个文件
file_list.each do |file_path|
  # 仅处理 Markdown 文件
  if File.extname(file_path) == '.markdown' || File.extname(file_path) == '.md'
    # 读取文件内容
    content = File.read(file_path)

    # 清除无效字符，包括回车符
    content = content.scrub.gsub(/\{% highlight html %\}[\r\n]*\{% highlight html %\}/m, '{% highlight html %}')

    # 将修改后的内容写回文件
    File.write(file_path, content)
  end
end

#require 'fileutils'

#folder_path = './_posts'  # 替换为实际的文件夹路径

## 获取文件夹中的所有文件
#file_list = Dir.glob(File.join(folder_path, '*'))

## 遍历每个文件
#file_list.each do |file_path|
#  # 仅处理 Markdown 文件
#  if File.extname(file_path) == '.markdown' || File.extname(file_path) == '.md'
#    # 读取文件内容
#    content = File.read(file_path)

#    # 去除空行和每行开头的空格
#    content.gsub!(/^\s+/, '')
#    content.gsub!(/\n\s*\n/, "\n")

#    # 修改 highlight 标签的语法
#    content.gsub!("{% highlight %}", "{% highlight html %}")

#    # 将修改后的内容写回文件
#    File.write(file_path, content)
#  end
#end
