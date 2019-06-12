
function add(str)
     return "test add" .. str
end

function concat(str1,str2)
     return str1 .. str2
end

function urlEncode(s)  
    s = string.gsub(s, "([^%w%.%- ])", function(c) return string.format("%%%02X", string.byte(c)) end)  
    return string.gsub(s, " ", "+")  
end  

function urlDecode(s)  
    s = string.gsub(s, '%%(%x%x)', function(h) return string.char(tonumber(h, 16)) end)  
    return s  
end 

function shell(cmd)
    local handle = io.popen(cmd)
    local result = handle:read("*a")
    handle:close()
    return result
end

function get(url)
    local http = require("http")
    response, error_message = http.request("GET", url, {
	query="page=1",
	headers={
	    Accept="*/*"
	}
    })
    return response["body"]
end
