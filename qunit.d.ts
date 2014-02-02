// Type definitions for QUnit 1.10
// Project: http://qunitjs.com/
// Definitions by: Diullei Gomes <https://github.com/diullei>
// DefinitelyTyped: https://github.com/borisyankov/DefinitelyTyped


interface DoneCallbackObject {
	failed: number;
	passed: number;
	total: number;
	runtime: number;
}

interface LogCallbackObject {
	result: boolean;
	actual: Object;
	expected: Object;
	message: string;
	source: string;
}

interface ModuleStartCallbackObject {
	name: string;
}

interface ModuleDoneCallbackObject {
	name: string;
	failed: number;
	passed: number;
	total: number;
}

interface TestDoneCallbackObject {
	name: string;
	module: string;
	failed: number;
	passed: number;
	total: number;
	duration: number;
}

interface TestStartCallbackObject {
	name: string;
	module: string;
}

interface Config {
	altertitle: boolean;
	autostart: boolean;
	current: Object;
	reorder: boolean;
	requireExpects: boolean;
	testTimeout: number;
	urlConfig: Array<URLConfigItem>;
	done: any;
}

interface URLConfigItem {
	id: string;
	label: string;
	tooltip: string;
}

interface LifecycleObject {
	setup?: () => any;
	teardown?: () => any;
}

interface QUnitAssert {
	/* ASSERT */
	assert: any;
	current_testEnvironment: any;
	jsDump: any;
	deepEqual(actual: any, expected: any, message?: string): any;
	equal(actual: any, expected: any, message?: string): any;
	notDeepEqual(actual: any, expected: any, message?: string): any;
	notEqual(actual: any, expected: any, message?: string): any;
	notPropEqual(actual: any, expected: any, message?: string): any;
	propEqual(actual: any, expected: any, message?: string): any;
	notStrictEqual(actual: any, expected: any, message?: string): any;
	ok(state: any, message?: string): any;
	strictEqual(actual: any, expected: any, message?: string): any;
	throws(block: () => any, expected: any, message?: string): any;
	throws(block: () => any, message?: string): any;
}

interface QUnitStatic extends QUnitAssert{	
	start(decrement?: number): any;
	stop(increment? : number): any;
	begin(callback: () => any): any;
	done(callback: (details: DoneCallbackObject) => any): any;
	log(callback: (details: LogCallbackObject) => any): any;
	moduleDone(callback: (details: ModuleDoneCallbackObject) => any): any;
	moduleStart(callback: (details: ModuleStartCallbackObject) => any): any;
	testDone(callback: (details: TestDoneCallbackObject) => any): any;
	testStart(callback: (details: TestStartCallbackObject) => any): any;
	config: Config;
	asyncTest(name: string, expected: number, test: () => any): any;
	asyncTest(name: string, test: () => any): any;
	expect(amount: number): any;
	module(name: string, lifecycle?: LifecycleObject): any;
	test(title: string, expected: number, test: (assert: QUnitAssert) => any): any;
	test(title: string, test: (assert: QUnitAssert) => any): any;
	equiv(a: any, b: any): any;
	raises: any;
	push(result: any, actual: any, expected: any, message: string): any;
	reset(): any;
}

declare function deepEqual(actual: any, expected: any, message?: string): any;
declare function equal(actual: any, expected: any, message?: string): any;
declare function notDeepEqual(actual: any, expected: any, message?: string): any;
declare function notEqual(actual: any, expected: any, message?: string): any;
declare function notStrictEqual(actual: any, expected: any, message?: string): any;
declare function ok(state: any, message?: string): any;
declare function strictEqual(actual: any, expected: any, message?: string): any;
declare function throws(block: () => any, expected: any, message?: string): any;
declare function throws(block: () => any, message?: string): any;
declare function start(decrement?: number): any;
declare function stop(increment? : number): any;
declare function begin(callback: () => any): any;
declare function done(callback: (details: DoneCallbackObject) => any): any;
declare function log(callback: (details: LogCallbackObject) => any): any;
declare function moduleDone(callback: (details: ModuleDoneCallbackObject) => any): any;
declare function moduleStart(callback: (name: string) => any): any;
declare function testDone(callback: (details: TestDoneCallbackObject) => any): any;
declare function testStart(callback: (details: TestStartCallbackObject) => any): any;
declare function asyncTest(name: string, expected?: any, test?: () => any): any;
declare function asyncTest(name: string, test: () => any): any;
declare function expect(amount: number): any;
declare function test(title: string, expected: number, test: (assert?: QUnitAssert) => any): any;
declare function test(title: string, test: (assert?: QUnitAssert) => any): any;
declare function notPropEqual(actual: any, expected: any, message?: string): any;
declare function propEqual(actual: any, expected: any, message?: string): any;
declare function equiv(a: any, b: any): any;
declare var raises: any;
declare var QUnit: QUnitStatic;