# neurons-platform script

Used to write simple scripts, For neurons-platform



# Grammar 

## variabal
```
a = 1
b = 2
c = `hello `
d = `world`
print a + b
print c + d
```

## if
```
if a eq 1
then
   print b + ` ` + c
fi
```

## while
```
a = 10
while a > 0 
do
  print a
  a = a - 1
  _lua("sleep","1")
done

```

## call lua function
```
// you can call lua function 
r = _lua("hello", "hello", "world")
print r
```

# install
```
# install goyacc
git clone https://github.com/golang/tools.git
cd tools/cmd/goyacc/
go build
cp goyacc /usr/local/go/bin/

./build.sh
```

# usage
```
./dist/neurons_script -lua fun.lua test/test_lua_json.ns
```

