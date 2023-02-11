#!lua name=mac_laren

local function lpop_order(keys)
    return redis.call('lpop', keys[1])
end

redis.register_function('lpop_order', lpop_order)
