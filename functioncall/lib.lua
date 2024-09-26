function Fibonacci(n)
  local result = {}
  for i = 1, n do
    local a = result[i - 2]
    local b = result[i - 1]
    if a == nil or b == nil then
      result[i] = 1
    else
      result[i] = a + b
    end
  end
  return result
end