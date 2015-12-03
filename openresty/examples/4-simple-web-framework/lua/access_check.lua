local param = require("common.param")
local args = ngx.req.get_uri_args()

ngx.log(ngx.DEBUG, "args.a = ", args.a, ", args.b = ", args.b)

if args.a == nil or args.b == nil then
    ngx.exit(ngx.HTTP_BAD_REQUEST)
    return
end

if not param.is_number(args.a, args.b) then
    ngx.exit(ngx.HTTP_BAD_REQUEST)
    return
end
