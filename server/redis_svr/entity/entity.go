package entity

// 存放实体

const (
	GetAndSetLua = `
    local key = KEYS[1]
    local value = ARGV[1]
    local exists = redis.call('EXISTS', key)
    if exists == 0 then
        redis.call('SET', key, value)
        return true
    else
        return false
    end
	`
)
