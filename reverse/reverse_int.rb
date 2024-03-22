# 反转整数，但是正负号不变
def reverse_int x
  sign = x > 0 ? 1 : -1
  # Abs是absolute,也就是绝对值
  x = x.abs
  # 反转整数
  reversed = 0
  while x >0
      reversed = reversed * 10 + x % 10
      x /= 10
  end
  reversed *= sign
  return 0 if reversed < -2 ** 31 || reversed > 2 ** 31 -1
  reversed
end
x = 1342
y = -1342
puts reverse_int x
puts reverse_int y
