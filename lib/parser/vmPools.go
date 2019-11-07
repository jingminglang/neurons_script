package parser


import (
	"github.com/yuin/gopher-lua"
	"github.com/BixData/gluabit32"
	"github.com/BixData/gluasocket"
	"github.com/cjoudrey/gluahttp"
	"github.com/zhu327/gluadb"
	LuaJson "layeh.com/gopher-json"
	"net/http"
	"sync"
	// "fmt"
)


type LuaPool struct {
	m     sync.Mutex
	saved []*lua.LState
}

func (pl *LuaPool) Get() *lua.LState {
	pl.m.Lock()
	defer pl.m.Unlock()
	n := len(pl.saved)
	if n == 0 {
		// TODO: 这里的初始化方式有问题
		return pl.New()
	}
	x := pl.saved[n-1]
	pl.saved = pl.saved[0 : n-1]
	return x
}

func (pl *LuaPool) CallFunction(fname string, fargs ...string) string {
	//先获取lua中定义的函数
	ls := []lua.LValue{}
	for _, fg := range fargs {
		ls = append(ls, lua.LString(fg))
		// ls = append(ls,fg)
	}
	L := pl.Get()
	defer pl.Put(L)
	fn := L.GetGlobal(fname)
	if err := L.CallByParam(lua.P{
		Fn:      fn,
		NRet:    1,
		Protect: true,
	}, ls...); err != nil {
		panic(err)
	}
	//这里获取函数返回值
	ret := L.Get(-1)
	return ret.String()

}

func (pl *LuaPool) New() *lua.LState {
	L := lua.NewState()
	L.PreloadModule("http", gluahttp.NewHttpModule(&http.Client{}).Loader)
	LuaJson.Preload(L)
	gluasocket.Preload(L)
	gluabit32.Preload(L)
	gluadb.Preload(L)
	return L
}


func (pl *LuaPool) Put(L *lua.LState) {
	pl.m.Lock()
	defer pl.m.Unlock()
	pl.saved = append(pl.saved, L)
}


func (pl *LuaPool) DoString(str string) {
	// fmt.Print("len: ",len(pl.saved))
	for _, L := range pl.saved {
		err := L.DoString(str)
		if err != nil {
			panic(err)
		}
	}
}

func (pl *LuaPool) PreloadModule(name string,loader lua.LGFunction) {
	for _, L := range pl.saved {
		L.PreloadModule(name, loader)
	}
}


func (pl *LuaPool) Shutdown() {
	for _, L := range pl.saved {
		L.Close()
	}
}

var LuaPools = &LuaPool{
	saved: make([]*lua.LState,100),
}

func init() {
	for i,_  := range LuaPools.saved {
		LuaPools.saved[i] = LuaPools.New()
	}
}