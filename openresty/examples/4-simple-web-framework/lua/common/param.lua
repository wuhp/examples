local _M = {}

function _M.is_number(...)
    args = {...}
    for _, v in ipairs(args) do
        if nil == tonumber(v) then
            return false
        end
    end

    return true
end

return _M
