-- file: scripts/limiter.lua

-- KEYS[1]: 限流的 key, 比如 ip:127.0.0.1
-- ARGV[1]: 过期时间（秒）
-- ARGV[2]: 限制次数

local current = redis.call('get', KEYS[1])
if current and tonumber(current) >= tonumber(ARGV[2]) then
    return 0
end

local count = redis.call('incr', KEYS[1])

if count == 1 then
    redis.call('expire', KEYS[1], ARGV[1])
end

return 1
