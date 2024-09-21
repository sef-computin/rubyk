n = 11
i = 2

fibo = []

fibo[0] = 0
fibo[1] = 1

while i < n 
  fibo[i] = fibo[i-1]+fibo[i-2]
  i = i + 1
end

return fibo[n-1]
