// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"
	"time"

	"github.com/hyperledger-labs/fabric-smart-client/platform/fabric/core/generic/msp"
)

type ConfigProvider struct {
	ConfigFileUsedStub        func() string
	configFileUsedMutex       sync.RWMutex
	configFileUsedArgsForCall []struct {
	}
	configFileUsedReturns struct {
		result1 string
	}
	configFileUsedReturnsOnCall map[int]struct {
		result1 string
	}
	GetBoolStub        func(string) bool
	getBoolMutex       sync.RWMutex
	getBoolArgsForCall []struct {
		arg1 string
	}
	getBoolReturns struct {
		result1 bool
	}
	getBoolReturnsOnCall map[int]struct {
		result1 bool
	}
	GetDurationStub        func(string) time.Duration
	getDurationMutex       sync.RWMutex
	getDurationArgsForCall []struct {
		arg1 string
	}
	getDurationReturns struct {
		result1 time.Duration
	}
	getDurationReturnsOnCall map[int]struct {
		result1 time.Duration
	}
	GetIntStub        func(string) int
	getIntMutex       sync.RWMutex
	getIntArgsForCall []struct {
		arg1 string
	}
	getIntReturns struct {
		result1 int
	}
	getIntReturnsOnCall map[int]struct {
		result1 int
	}
	GetPathStub        func(string) string
	getPathMutex       sync.RWMutex
	getPathArgsForCall []struct {
		arg1 string
	}
	getPathReturns struct {
		result1 string
	}
	getPathReturnsOnCall map[int]struct {
		result1 string
	}
	GetStringStub        func(string) string
	getStringMutex       sync.RWMutex
	getStringArgsForCall []struct {
		arg1 string
	}
	getStringReturns struct {
		result1 string
	}
	getStringReturnsOnCall map[int]struct {
		result1 string
	}
	GetStringSliceStub        func(string) []string
	getStringSliceMutex       sync.RWMutex
	getStringSliceArgsForCall []struct {
		arg1 string
	}
	getStringSliceReturns struct {
		result1 []string
	}
	getStringSliceReturnsOnCall map[int]struct {
		result1 []string
	}
	IsSetStub        func(string) bool
	isSetMutex       sync.RWMutex
	isSetArgsForCall []struct {
		arg1 string
	}
	isSetReturns struct {
		result1 bool
	}
	isSetReturnsOnCall map[int]struct {
		result1 bool
	}
	TranslatePathStub        func(string) string
	translatePathMutex       sync.RWMutex
	translatePathArgsForCall []struct {
		arg1 string
	}
	translatePathReturns struct {
		result1 string
	}
	translatePathReturnsOnCall map[int]struct {
		result1 string
	}
	UnmarshalKeyStub        func(string, interface{}) error
	unmarshalKeyMutex       sync.RWMutex
	unmarshalKeyArgsForCall []struct {
		arg1 string
		arg2 interface{}
	}
	unmarshalKeyReturns struct {
		result1 error
	}
	unmarshalKeyReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ConfigProvider) ConfigFileUsed() string {
	fake.configFileUsedMutex.Lock()
	ret, specificReturn := fake.configFileUsedReturnsOnCall[len(fake.configFileUsedArgsForCall)]
	fake.configFileUsedArgsForCall = append(fake.configFileUsedArgsForCall, struct {
	}{})
	stub := fake.ConfigFileUsedStub
	fakeReturns := fake.configFileUsedReturns
	fake.recordInvocation("ConfigFileUsed", []interface{}{})
	fake.configFileUsedMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *ConfigProvider) ConfigFileUsedCallCount() int {
	fake.configFileUsedMutex.RLock()
	defer fake.configFileUsedMutex.RUnlock()
	return len(fake.configFileUsedArgsForCall)
}

func (fake *ConfigProvider) ConfigFileUsedCalls(stub func() string) {
	fake.configFileUsedMutex.Lock()
	defer fake.configFileUsedMutex.Unlock()
	fake.ConfigFileUsedStub = stub
}

func (fake *ConfigProvider) ConfigFileUsedReturns(result1 string) {
	fake.configFileUsedMutex.Lock()
	defer fake.configFileUsedMutex.Unlock()
	fake.ConfigFileUsedStub = nil
	fake.configFileUsedReturns = struct {
		result1 string
	}{result1}
}

func (fake *ConfigProvider) ConfigFileUsedReturnsOnCall(i int, result1 string) {
	fake.configFileUsedMutex.Lock()
	defer fake.configFileUsedMutex.Unlock()
	fake.ConfigFileUsedStub = nil
	if fake.configFileUsedReturnsOnCall == nil {
		fake.configFileUsedReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.configFileUsedReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *ConfigProvider) GetBool(arg1 string) bool {
	fake.getBoolMutex.Lock()
	ret, specificReturn := fake.getBoolReturnsOnCall[len(fake.getBoolArgsForCall)]
	fake.getBoolArgsForCall = append(fake.getBoolArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetBoolStub
	fakeReturns := fake.getBoolReturns
	fake.recordInvocation("GetBool", []interface{}{arg1})
	fake.getBoolMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *ConfigProvider) GetBoolCallCount() int {
	fake.getBoolMutex.RLock()
	defer fake.getBoolMutex.RUnlock()
	return len(fake.getBoolArgsForCall)
}

func (fake *ConfigProvider) GetBoolCalls(stub func(string) bool) {
	fake.getBoolMutex.Lock()
	defer fake.getBoolMutex.Unlock()
	fake.GetBoolStub = stub
}

func (fake *ConfigProvider) GetBoolArgsForCall(i int) string {
	fake.getBoolMutex.RLock()
	defer fake.getBoolMutex.RUnlock()
	argsForCall := fake.getBoolArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ConfigProvider) GetBoolReturns(result1 bool) {
	fake.getBoolMutex.Lock()
	defer fake.getBoolMutex.Unlock()
	fake.GetBoolStub = nil
	fake.getBoolReturns = struct {
		result1 bool
	}{result1}
}

func (fake *ConfigProvider) GetBoolReturnsOnCall(i int, result1 bool) {
	fake.getBoolMutex.Lock()
	defer fake.getBoolMutex.Unlock()
	fake.GetBoolStub = nil
	if fake.getBoolReturnsOnCall == nil {
		fake.getBoolReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.getBoolReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *ConfigProvider) GetDuration(arg1 string) time.Duration {
	fake.getDurationMutex.Lock()
	ret, specificReturn := fake.getDurationReturnsOnCall[len(fake.getDurationArgsForCall)]
	fake.getDurationArgsForCall = append(fake.getDurationArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetDurationStub
	fakeReturns := fake.getDurationReturns
	fake.recordInvocation("GetDuration", []interface{}{arg1})
	fake.getDurationMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *ConfigProvider) GetDurationCallCount() int {
	fake.getDurationMutex.RLock()
	defer fake.getDurationMutex.RUnlock()
	return len(fake.getDurationArgsForCall)
}

func (fake *ConfigProvider) GetDurationCalls(stub func(string) time.Duration) {
	fake.getDurationMutex.Lock()
	defer fake.getDurationMutex.Unlock()
	fake.GetDurationStub = stub
}

func (fake *ConfigProvider) GetDurationArgsForCall(i int) string {
	fake.getDurationMutex.RLock()
	defer fake.getDurationMutex.RUnlock()
	argsForCall := fake.getDurationArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ConfigProvider) GetDurationReturns(result1 time.Duration) {
	fake.getDurationMutex.Lock()
	defer fake.getDurationMutex.Unlock()
	fake.GetDurationStub = nil
	fake.getDurationReturns = struct {
		result1 time.Duration
	}{result1}
}

func (fake *ConfigProvider) GetDurationReturnsOnCall(i int, result1 time.Duration) {
	fake.getDurationMutex.Lock()
	defer fake.getDurationMutex.Unlock()
	fake.GetDurationStub = nil
	if fake.getDurationReturnsOnCall == nil {
		fake.getDurationReturnsOnCall = make(map[int]struct {
			result1 time.Duration
		})
	}
	fake.getDurationReturnsOnCall[i] = struct {
		result1 time.Duration
	}{result1}
}

func (fake *ConfigProvider) GetInt(arg1 string) int {
	fake.getIntMutex.Lock()
	ret, specificReturn := fake.getIntReturnsOnCall[len(fake.getIntArgsForCall)]
	fake.getIntArgsForCall = append(fake.getIntArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetIntStub
	fakeReturns := fake.getIntReturns
	fake.recordInvocation("GetInt", []interface{}{arg1})
	fake.getIntMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *ConfigProvider) GetIntCallCount() int {
	fake.getIntMutex.RLock()
	defer fake.getIntMutex.RUnlock()
	return len(fake.getIntArgsForCall)
}

func (fake *ConfigProvider) GetIntCalls(stub func(string) int) {
	fake.getIntMutex.Lock()
	defer fake.getIntMutex.Unlock()
	fake.GetIntStub = stub
}

func (fake *ConfigProvider) GetIntArgsForCall(i int) string {
	fake.getIntMutex.RLock()
	defer fake.getIntMutex.RUnlock()
	argsForCall := fake.getIntArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ConfigProvider) GetIntReturns(result1 int) {
	fake.getIntMutex.Lock()
	defer fake.getIntMutex.Unlock()
	fake.GetIntStub = nil
	fake.getIntReturns = struct {
		result1 int
	}{result1}
}

func (fake *ConfigProvider) GetIntReturnsOnCall(i int, result1 int) {
	fake.getIntMutex.Lock()
	defer fake.getIntMutex.Unlock()
	fake.GetIntStub = nil
	if fake.getIntReturnsOnCall == nil {
		fake.getIntReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.getIntReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *ConfigProvider) GetPath(arg1 string) string {
	fake.getPathMutex.Lock()
	ret, specificReturn := fake.getPathReturnsOnCall[len(fake.getPathArgsForCall)]
	fake.getPathArgsForCall = append(fake.getPathArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetPathStub
	fakeReturns := fake.getPathReturns
	fake.recordInvocation("GetPath", []interface{}{arg1})
	fake.getPathMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *ConfigProvider) GetPathCallCount() int {
	fake.getPathMutex.RLock()
	defer fake.getPathMutex.RUnlock()
	return len(fake.getPathArgsForCall)
}

func (fake *ConfigProvider) GetPathCalls(stub func(string) string) {
	fake.getPathMutex.Lock()
	defer fake.getPathMutex.Unlock()
	fake.GetPathStub = stub
}

func (fake *ConfigProvider) GetPathArgsForCall(i int) string {
	fake.getPathMutex.RLock()
	defer fake.getPathMutex.RUnlock()
	argsForCall := fake.getPathArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ConfigProvider) GetPathReturns(result1 string) {
	fake.getPathMutex.Lock()
	defer fake.getPathMutex.Unlock()
	fake.GetPathStub = nil
	fake.getPathReturns = struct {
		result1 string
	}{result1}
}

func (fake *ConfigProvider) GetPathReturnsOnCall(i int, result1 string) {
	fake.getPathMutex.Lock()
	defer fake.getPathMutex.Unlock()
	fake.GetPathStub = nil
	if fake.getPathReturnsOnCall == nil {
		fake.getPathReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.getPathReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *ConfigProvider) GetString(arg1 string) string {
	fake.getStringMutex.Lock()
	ret, specificReturn := fake.getStringReturnsOnCall[len(fake.getStringArgsForCall)]
	fake.getStringArgsForCall = append(fake.getStringArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetStringStub
	fakeReturns := fake.getStringReturns
	fake.recordInvocation("GetString", []interface{}{arg1})
	fake.getStringMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *ConfigProvider) GetStringCallCount() int {
	fake.getStringMutex.RLock()
	defer fake.getStringMutex.RUnlock()
	return len(fake.getStringArgsForCall)
}

func (fake *ConfigProvider) GetStringCalls(stub func(string) string) {
	fake.getStringMutex.Lock()
	defer fake.getStringMutex.Unlock()
	fake.GetStringStub = stub
}

func (fake *ConfigProvider) GetStringArgsForCall(i int) string {
	fake.getStringMutex.RLock()
	defer fake.getStringMutex.RUnlock()
	argsForCall := fake.getStringArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ConfigProvider) GetStringReturns(result1 string) {
	fake.getStringMutex.Lock()
	defer fake.getStringMutex.Unlock()
	fake.GetStringStub = nil
	fake.getStringReturns = struct {
		result1 string
	}{result1}
}

func (fake *ConfigProvider) GetStringReturnsOnCall(i int, result1 string) {
	fake.getStringMutex.Lock()
	defer fake.getStringMutex.Unlock()
	fake.GetStringStub = nil
	if fake.getStringReturnsOnCall == nil {
		fake.getStringReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.getStringReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *ConfigProvider) GetStringSlice(arg1 string) []string {
	fake.getStringSliceMutex.Lock()
	ret, specificReturn := fake.getStringSliceReturnsOnCall[len(fake.getStringSliceArgsForCall)]
	fake.getStringSliceArgsForCall = append(fake.getStringSliceArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetStringSliceStub
	fakeReturns := fake.getStringSliceReturns
	fake.recordInvocation("GetStringSlice", []interface{}{arg1})
	fake.getStringSliceMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *ConfigProvider) GetStringSliceCallCount() int {
	fake.getStringSliceMutex.RLock()
	defer fake.getStringSliceMutex.RUnlock()
	return len(fake.getStringSliceArgsForCall)
}

func (fake *ConfigProvider) GetStringSliceCalls(stub func(string) []string) {
	fake.getStringSliceMutex.Lock()
	defer fake.getStringSliceMutex.Unlock()
	fake.GetStringSliceStub = stub
}

func (fake *ConfigProvider) GetStringSliceArgsForCall(i int) string {
	fake.getStringSliceMutex.RLock()
	defer fake.getStringSliceMutex.RUnlock()
	argsForCall := fake.getStringSliceArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ConfigProvider) GetStringSliceReturns(result1 []string) {
	fake.getStringSliceMutex.Lock()
	defer fake.getStringSliceMutex.Unlock()
	fake.GetStringSliceStub = nil
	fake.getStringSliceReturns = struct {
		result1 []string
	}{result1}
}

func (fake *ConfigProvider) GetStringSliceReturnsOnCall(i int, result1 []string) {
	fake.getStringSliceMutex.Lock()
	defer fake.getStringSliceMutex.Unlock()
	fake.GetStringSliceStub = nil
	if fake.getStringSliceReturnsOnCall == nil {
		fake.getStringSliceReturnsOnCall = make(map[int]struct {
			result1 []string
		})
	}
	fake.getStringSliceReturnsOnCall[i] = struct {
		result1 []string
	}{result1}
}

func (fake *ConfigProvider) IsSet(arg1 string) bool {
	fake.isSetMutex.Lock()
	ret, specificReturn := fake.isSetReturnsOnCall[len(fake.isSetArgsForCall)]
	fake.isSetArgsForCall = append(fake.isSetArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.IsSetStub
	fakeReturns := fake.isSetReturns
	fake.recordInvocation("IsSet", []interface{}{arg1})
	fake.isSetMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *ConfigProvider) IsSetCallCount() int {
	fake.isSetMutex.RLock()
	defer fake.isSetMutex.RUnlock()
	return len(fake.isSetArgsForCall)
}

func (fake *ConfigProvider) IsSetCalls(stub func(string) bool) {
	fake.isSetMutex.Lock()
	defer fake.isSetMutex.Unlock()
	fake.IsSetStub = stub
}

func (fake *ConfigProvider) IsSetArgsForCall(i int) string {
	fake.isSetMutex.RLock()
	defer fake.isSetMutex.RUnlock()
	argsForCall := fake.isSetArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ConfigProvider) IsSetReturns(result1 bool) {
	fake.isSetMutex.Lock()
	defer fake.isSetMutex.Unlock()
	fake.IsSetStub = nil
	fake.isSetReturns = struct {
		result1 bool
	}{result1}
}

func (fake *ConfigProvider) IsSetReturnsOnCall(i int, result1 bool) {
	fake.isSetMutex.Lock()
	defer fake.isSetMutex.Unlock()
	fake.IsSetStub = nil
	if fake.isSetReturnsOnCall == nil {
		fake.isSetReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isSetReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *ConfigProvider) TranslatePath(arg1 string) string {
	fake.translatePathMutex.Lock()
	ret, specificReturn := fake.translatePathReturnsOnCall[len(fake.translatePathArgsForCall)]
	fake.translatePathArgsForCall = append(fake.translatePathArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.TranslatePathStub
	fakeReturns := fake.translatePathReturns
	fake.recordInvocation("TranslatePath", []interface{}{arg1})
	fake.translatePathMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *ConfigProvider) TranslatePathCallCount() int {
	fake.translatePathMutex.RLock()
	defer fake.translatePathMutex.RUnlock()
	return len(fake.translatePathArgsForCall)
}

func (fake *ConfigProvider) TranslatePathCalls(stub func(string) string) {
	fake.translatePathMutex.Lock()
	defer fake.translatePathMutex.Unlock()
	fake.TranslatePathStub = stub
}

func (fake *ConfigProvider) TranslatePathArgsForCall(i int) string {
	fake.translatePathMutex.RLock()
	defer fake.translatePathMutex.RUnlock()
	argsForCall := fake.translatePathArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ConfigProvider) TranslatePathReturns(result1 string) {
	fake.translatePathMutex.Lock()
	defer fake.translatePathMutex.Unlock()
	fake.TranslatePathStub = nil
	fake.translatePathReturns = struct {
		result1 string
	}{result1}
}

func (fake *ConfigProvider) TranslatePathReturnsOnCall(i int, result1 string) {
	fake.translatePathMutex.Lock()
	defer fake.translatePathMutex.Unlock()
	fake.TranslatePathStub = nil
	if fake.translatePathReturnsOnCall == nil {
		fake.translatePathReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.translatePathReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *ConfigProvider) UnmarshalKey(arg1 string, arg2 interface{}) error {
	fake.unmarshalKeyMutex.Lock()
	ret, specificReturn := fake.unmarshalKeyReturnsOnCall[len(fake.unmarshalKeyArgsForCall)]
	fake.unmarshalKeyArgsForCall = append(fake.unmarshalKeyArgsForCall, struct {
		arg1 string
		arg2 interface{}
	}{arg1, arg2})
	stub := fake.UnmarshalKeyStub
	fakeReturns := fake.unmarshalKeyReturns
	fake.recordInvocation("UnmarshalKey", []interface{}{arg1, arg2})
	fake.unmarshalKeyMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *ConfigProvider) UnmarshalKeyCallCount() int {
	fake.unmarshalKeyMutex.RLock()
	defer fake.unmarshalKeyMutex.RUnlock()
	return len(fake.unmarshalKeyArgsForCall)
}

func (fake *ConfigProvider) UnmarshalKeyCalls(stub func(string, interface{}) error) {
	fake.unmarshalKeyMutex.Lock()
	defer fake.unmarshalKeyMutex.Unlock()
	fake.UnmarshalKeyStub = stub
}

func (fake *ConfigProvider) UnmarshalKeyArgsForCall(i int) (string, interface{}) {
	fake.unmarshalKeyMutex.RLock()
	defer fake.unmarshalKeyMutex.RUnlock()
	argsForCall := fake.unmarshalKeyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *ConfigProvider) UnmarshalKeyReturns(result1 error) {
	fake.unmarshalKeyMutex.Lock()
	defer fake.unmarshalKeyMutex.Unlock()
	fake.UnmarshalKeyStub = nil
	fake.unmarshalKeyReturns = struct {
		result1 error
	}{result1}
}

func (fake *ConfigProvider) UnmarshalKeyReturnsOnCall(i int, result1 error) {
	fake.unmarshalKeyMutex.Lock()
	defer fake.unmarshalKeyMutex.Unlock()
	fake.UnmarshalKeyStub = nil
	if fake.unmarshalKeyReturnsOnCall == nil {
		fake.unmarshalKeyReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.unmarshalKeyReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ConfigProvider) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.configFileUsedMutex.RLock()
	defer fake.configFileUsedMutex.RUnlock()
	fake.getBoolMutex.RLock()
	defer fake.getBoolMutex.RUnlock()
	fake.getDurationMutex.RLock()
	defer fake.getDurationMutex.RUnlock()
	fake.getIntMutex.RLock()
	defer fake.getIntMutex.RUnlock()
	fake.getPathMutex.RLock()
	defer fake.getPathMutex.RUnlock()
	fake.getStringMutex.RLock()
	defer fake.getStringMutex.RUnlock()
	fake.getStringSliceMutex.RLock()
	defer fake.getStringSliceMutex.RUnlock()
	fake.isSetMutex.RLock()
	defer fake.isSetMutex.RUnlock()
	fake.translatePathMutex.RLock()
	defer fake.translatePathMutex.RUnlock()
	fake.unmarshalKeyMutex.RLock()
	defer fake.unmarshalKeyMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ConfigProvider) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ msp.ConfigProvider = new(ConfigProvider)
