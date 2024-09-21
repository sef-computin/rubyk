a = 0
b = 1

i = 2
n = 10
tmp = 0
while i < n
  tmp = b
  b = a + b
  a = tmp
  i = i + 1  
end

return b
