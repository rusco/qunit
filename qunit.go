package qunit

//GopherJS Bindings for qunitjs.com
import "github.com/neelance/gopherjs/js"

type QUnitAssert struct {
	js.Object
	//assert                  js.Object//2do
	//current_testEnvironment js.Object
	//jsDump                  js.Object
}

type DoneCallbackObject struct {
	js.Object
	Failed  int `js:"failed"`
	Passed  int `js:"passed"`
	Total   int `js:"total"`
	Runtime int `js:"runtime"`
}

type LogCallbackObject struct {
	js.Object
	result   bool      `js:"result"`
	actual   js.Object `js:"actual"`
	expected js.Object `js:"expected"`
	message  string    `js:"message"`
	source   string    `js:"source"`
}

type ModuleStartCallbackObject struct {
	js.Object
	name string `js:"name"`
}

type ModuleDoneCallbackObject struct {
	js.Object
	name   string `js:"name"`
	failed int    `js:"failed"`
	passed int    `js:"passed"`
	total  int    `js:"total"`
}

type TestDoneCallbackObject struct {
	js.Object
	name     string `js:"name"`
	module   string `js:"module"`
	failed   int    `js:"failed"`
	passed   int    `js:"passed"`
	total    int    `js:"total"`
	duration int    `js:"duration"`
}

type TestStartCallbackObject struct {
	js.Object
	name   string `js:"name"`
	module string `js:"module"`
}

func (qa QUnitAssert) DeepEqual(actual interface{}, expected interface{}, message string) interface{} {
	return qa.Call("deepEqual", actual, expected, message)
}

func (qa QUnitAssert) Equal(actual interface{}, expected interface{}, message string) interface{} {
	return qa.Call("equal", actual, expected, message)
}

func (qa QUnitAssert) NotDeepEqual(actual interface{}, expected interface{}, message string) interface{} {
	return qa.Call("notDeepEqual", actual, expected, message)
}

func (qa QUnitAssert) NotEqual(actual interface{}, expected interface{}, message string) interface{} {
	return qa.Call("notEqual", actual, expected, message)
}

func (qa QUnitAssert) NotPropEqual(actual interface{}, expected interface{}, message string) interface{} {
	return qa.Call("notPropEqual", actual, expected, message)
}

func (qa QUnitAssert) PropEqual(actual interface{}, expected interface{}, message string) interface{} {
	return qa.Call("propEqual", actual, expected, message)
}

func (qa QUnitAssert) NotStrictEqual(actual interface{}, expected interface{}, message string) interface{} {
	return qa.Call("notStrictEqual", actual, expected, message)
}

func (qa QUnitAssert) Ok(state interface{}, message string) interface{} {
	return qa.Call("ok", state, message)
}

func (qa QUnitAssert) StrictEqual(actual interface{}, expected interface{}, message string) interface{} {
	return qa.Call("strictEqual", actual, expected, message)
}

func (qa QUnitAssert) ThrowsExpected(block func() interface{}, expected interface{}, message string) interface{} {
	return qa.Call("throwsExpected", block, expected, message)
}

func (qa QUnitAssert) Throws(block func() interface{}, message string) interface{} {
	return qa.Call("throws", block, message)
}

//start QUnit static methods
func Test(name string, testFn func(QUnitAssert)) {

	js.Global("QUnit").Call("test", name, func(e js.Object) {
		testFn(QUnitAssert{Object: e})
	})
}

func TestExpected(title string, expected int, testFn func(assert QUnitAssert) interface{}) interface{} {
	t := js.Global("QUnit").Call("test", title, expected, func(e js.Object) {
		testFn(QUnitAssert{Object: e})
	})
	return t
}

func Start() interface{} {
	return js.Global("QUnit").Call("start")
}
func StartDecrement(decrement int) interface{} {
	return js.Global("QUnit").Call("start", decrement)
}
func Stop() interface{} {
	return js.Global("QUnit").Call("stop")
}
func StopIncrement(increment int) interface{} {
	return js.Global("QUnit").Call("stop", increment)
}

func Begin(callbackFn func() interface{}) interface{} {
	t := js.Global("QUnit").Call("begin", func() {
		callbackFn()
	})
	return t
}
func Done(callbackFn func(details DoneCallbackObject) interface{}) interface{} {
	t := js.Global("QUnit").Call("done", func(e js.Object) {
		callbackFn(DoneCallbackObject{Object: e})
	})
	return t
}
func Log(callbackFn func(details LogCallbackObject) interface{}) interface{} {
	t := js.Global("QUnit").Call("log", func(e js.Object) {
		callbackFn(LogCallbackObject{Object: e})
	})
	return t
}
func ModuleDone(callbackFn func(details ModuleDoneCallbackObject) interface{}) interface{} {
	t := js.Global("QUnit").Call("moduleDone", func(e js.Object) {
		callbackFn(ModuleDoneCallbackObject{Object: e})
	})
	return t
}
func ModuleStart(callbackFn func(name string) interface{}) interface{} {
	t := js.Global("QUnit").Call("moduleStart", func(e js.Object) {
		callbackFn(e.String())
	})
	return t
}
func TestDone(callbackFn func(details TestDoneCallbackObject) interface{}) interface{} {
	t := js.Global("QUnit").Call("testDone", func(e js.Object) {
		callbackFn(TestDoneCallbackObject{Object: e})
	})
	return t
}
func TestStart(callbackFn func(details TestStartCallbackObject) interface{}) interface{} {
	t := js.Global("QUnit").Call("testStart", func(e js.Object) {
		callbackFn(TestStartCallbackObject{Object: e})
	})
	return t
}
func AsyncTestExpected(name string, expected interface{}, testFn func() interface{}) interface{} {
	t := js.Global("QUnit").Call("asyncTestExpected", name, expected, func() {
		testFn()
	})
	return t
}
func AsyncTest(name string, testFn func() interface{}) interface{} {
	t := js.Global("QUnit").Call("asyncTest", name, func() {
		testFn()
	})
	return t
}
func Expect(amount int) interface{} {
	return js.Global("QUnit").Call("expect", amount)
}

func Equiv(a interface{}, b interface{}) interface{} {
	return js.Global("QUnit").Call("equip", a, b)
}

func Module(name string) interface{} {
	return js.Global("QUnit").Call("module", name)
}

/*
2do: implement this:
func ModuleLifecycle(name string, lifecycle LifecycleObject) interface{} {
	return js.Global("QUnit").Call("module", name, lifecycle)
}
*/

type Raises struct {
	js.Object
	Raises js.Object `js:"raises"`
}

func Push(result interface{}, actual interface{}, expected interface{}, message string) interface{} {
	return js.Global("QUnit").Call("push", result, actual, expected, message)
}

func Reset() interface{} {
	return js.Global("QUnit").Call("reset")
}
