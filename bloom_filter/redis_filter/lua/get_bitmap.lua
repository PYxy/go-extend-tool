--local bloomKeyArray = KEYS
--local sign = tonumber(ARGV[1])
--local bitMapArray = ARGV[2]
--for i, v in pairs(bloomKeyArray) do
--    print("操作", v, bitMapArray[i],sign)
--    -- 0 或者 1
--    local val = redis.call('GETBIT', v, bitMapArray[i])
--    if sign ~= val then
--        return "false"
--    end
--
--end
--
--
--return "true"


local bloomKeyArray = KEYS
local sign = 1
local delimiter = ","
-- 使用gsub函数进行分割
local n = 1
for substr in string.gmatch(ARGV[1], "[^" .. delimiter .. "]+") do
    --print(bloomKeyArray[n],tonumber(substr),sign)
    --redis.call('SETBIT', bloomKeyArray[n], tonumber(substr), sign)
    local val = redis.call('GETBIT',  bloomKeyArray[n], tonumber(substr))
    if val ~= sign then
        return "false"
    end
    n = n+1
end

return "true"
