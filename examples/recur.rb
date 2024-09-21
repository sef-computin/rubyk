x = 2
y = 4
n = 0

def pow(a, i)
  while n < i
    a = a * a
    n = n + 1
  end

  return n
end



return pow(x, y)
