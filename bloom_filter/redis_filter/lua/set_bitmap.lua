local bloomKeyArray = KEYS
local sign = 1
local delimiter = ","
-- 使用gsub函数进行分割
local n = 1
for substr in string.gmatch(ARGV[1], "[^" .. delimiter .. "]+") do
    --print(bloomKeyArray[n],tonumber(substr),sign)
    redis.call('SETBIT', bloomKeyArray[n], tonumber(substr), sign)
    n = n+1
end


return "true"
