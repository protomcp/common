package generator

import (
	"testing"

	"darvaza.org/core"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

// Compile-time verification that test case types implement TestCase interface
var _ core.TestCase = findTestCase{}
var _ core.TestCase = forEachTestCase{}

// Test case for Find functions
type findTestCase struct {
	// Larger fields first for field alignment
	findFunc func(proto.Message, string) proto.Message
	input    proto.Message
	name     string
	findName string
	fnName   string
	wantNil  bool
}

// Specific factory functions for each Find test type

func newFindMessageTestCase(name string, input proto.Message, findName string,
	wantNil bool) findTestCase {
	return findTestCase{
		name:     name,
		input:    input,
		findName: findName,
		findFunc: FindMessage,
		fnName:   "FindMessage",
		wantNil:  wantNil,
	}
}

func newFindFieldTestCase(name string, input proto.Message, findName string,
	wantNil bool) findTestCase {
	return findTestCase{
		name:     name,
		input:    input,
		findName: findName,
		findFunc: FindField,
		fnName:   "FindField",
		wantNil:  wantNil,
	}
}

func newFindEnumTestCase(name string, input proto.Message, findName string,
	wantNil bool) findTestCase {
	return findTestCase{
		name:     name,
		input:    input,
		findName: findName,
		findFunc: FindEnum,
		fnName:   "FindEnum",
		wantNil:  wantNil,
	}
}

func newFindServiceTestCase(name string, input proto.Message, findName string,
	wantNil bool) findTestCase {
	return findTestCase{
		name:     name,
		input:    input,
		findName: findName,
		findFunc: FindService,
		fnName:   "FindService",
		wantNil:  wantNil,
	}
}

func newFindMethodTestCase(name string, input proto.Message, findName string,
	wantNil bool) findTestCase {
	return findTestCase{
		name:     name,
		input:    input,
		findName: findName,
		findFunc: FindMethod,
		fnName:   "FindMethod",
		wantNil:  wantNil,
	}
}

func (tc findTestCase) Name() string {
	return tc.name
}

func (tc findTestCase) Test(t *testing.T) {
	t.Helper()
	result := tc.findFunc(tc.input, tc.findName)
	if tc.wantNil {
		core.AssertNil(t, result, tc.fnName+" result")
	} else {
		core.AssertNotNil(t, result, tc.fnName+" result")
	}
}

// Test case for ForEach functions
type forEachTestCase struct {
	// Larger fields first for field alignment
	forEachFunc func(proto.Message, func(proto.Message) bool) error
	input       proto.Message
	name        string
	fnName      string
	stopAt      int // Stop iteration at this index (-1 for no stop)
	wantCount   int
	wantErr     bool
}

//revive:disable-next-line:argument-limit,flag-parameter
func newForEachTestCase(name string, input proto.Message,
	forEachFunc func(proto.Message, func(proto.Message) bool) error,
	fnName string, stopAt int, wantCount int, wantErr bool) forEachTestCase {
	return forEachTestCase{
		name:        name,
		input:       input,
		forEachFunc: forEachFunc,
		fnName:      fnName,
		stopAt:      stopAt,
		wantCount:   wantCount,
		wantErr:     wantErr,
	}
}

func (tc forEachTestCase) Name() string {
	return tc.name
}

func (tc forEachTestCase) Test(t *testing.T) {
	t.Helper()
	count := 0
	err := tc.forEachFunc(tc.input, func(_ proto.Message) bool {
		count++
		if tc.stopAt >= 0 && count >= tc.stopAt {
			return false // Stop iteration
		}
		return true // Continue
	})

	if tc.wantErr {
		core.AssertError(t, err, tc.fnName+" error")
	} else {
		core.AssertNoError(t, err, tc.fnName)
	}
	core.AssertEqual(t, tc.wantCount, count, tc.fnName+" count")
}

// Helper factory that creates both error simulation function and error message
func makeForEachErrorTest(delegateFn func(proto.Message, func(proto.Message) bool) error) (
	func(proto.Message, func(proto.Message) bool) error,
	*descriptorpb.DescriptorProto) {
	const errorName = "ErrorCase"

	errorFn := func(msg proto.Message, fn func(proto.Message) bool) error {
		// Simulate an error condition for testing
		if msg != nil {
			if msgDesc, ok := AsMessage(msg); ok && msgDesc.GetName() == errorName {
				return core.ErrInvalid
			}
		}
		// Otherwise delegate to the real function
		return delegateFn(msg, fn)
	}

	errorMsg := &descriptorpb.DescriptorProto{
		Name: proto.String(errorName),
		Field: []*descriptorpb.FieldDescriptorProto{
			NewField("test", 1, TypeString),
		},
	}

	return errorFn, errorMsg
}

// Helper function to create test file descriptor
func createTestFileDescriptor() *descriptorpb.FileDescriptorProto {
	stringType := descriptorpb.FieldDescriptorProto_TYPE_STRING
	int32Type := descriptorpb.FieldDescriptorProto_TYPE_INT32
	boolType := descriptorpb.FieldDescriptorProto_TYPE_BOOL
	messageType := descriptorpb.FieldDescriptorProto_TYPE_MESSAGE
	enumType := descriptorpb.FieldDescriptorProto_TYPE_ENUM

	optionalLabel := descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL
	repeatedLabel := descriptorpb.FieldDescriptorProto_LABEL_REPEATED

	fileName := "test.proto"
	packageName := "test.package"

	// Create enum
	enumName := "Status"
	enum := &descriptorpb.EnumDescriptorProto{
		Name: &enumName,
		Value: []*descriptorpb.EnumValueDescriptorProto{
			{
				Name:   proto.String("UNKNOWN"),
				Number: proto.Int32(0),
			},
			{
				Name:   proto.String("ACTIVE"),
				Number: proto.Int32(1),
			},
			{
				Name:   proto.String("INACTIVE"),
				Number: proto.Int32(2),
			},
		},
	}

	// Create nested message
	nestedMsgName := "Nested"
	nestedMsg := &descriptorpb.DescriptorProto{
		Name: &nestedMsgName,
		Field: []*descriptorpb.FieldDescriptorProto{
			{
				Name:   proto.String("nested_field"),
				Number: proto.Int32(1),
				Label:  &optionalLabel,
				Type:   &stringType,
			},
		},
	}

	// Create main message with nested types
	msgName := "TestMessage"
	msg := &descriptorpb.DescriptorProto{
		Name: &msgName,
		Field: []*descriptorpb.FieldDescriptorProto{
			{
				Name:   proto.String("name"),
				Number: proto.Int32(1),
				Label:  &optionalLabel,
				Type:   &stringType,
			},
			{
				Name:   proto.String("id"),
				Number: proto.Int32(2),
				Label:  &optionalLabel,
				Type:   &int32Type,
			},
			{
				Name:   proto.String("active"),
				Number: proto.Int32(3),
				Label:  &optionalLabel,
				Type:   &boolType,
			},
			{
				Name:     proto.String("nested"),
				Number:   proto.Int32(4),
				Label:    &optionalLabel,
				Type:     &messageType,
				TypeName: proto.String(".test.package.TestMessage.Nested"),
			},
			{
				Name:     proto.String("status"),
				Number:   proto.Int32(5),
				Label:    &optionalLabel,
				Type:     &enumType,
				TypeName: proto.String(".test.package.Status"),
			},
			{
				Name:   proto.String("tags"),
				Number: proto.Int32(6),
				Label:  &repeatedLabel,
				Type:   &stringType,
			},
		},
		NestedType: []*descriptorpb.DescriptorProto{nestedMsg},
		EnumType: []*descriptorpb.EnumDescriptorProto{
			{
				Name: proto.String("InnerEnum"),
				Value: []*descriptorpb.EnumValueDescriptorProto{
					{
						Name:   proto.String("INNER_UNKNOWN"),
						Number: proto.Int32(0),
					},
				},
			},
		},
	}

	// Create another message
	anotherMsgName := "AnotherMessage"
	anotherMsg := &descriptorpb.DescriptorProto{
		Name: &anotherMsgName,
		Field: []*descriptorpb.FieldDescriptorProto{
			{
				Name:   proto.String("value"),
				Number: proto.Int32(1),
				Label:  &optionalLabel,
				Type:   &stringType,
			},
		},
	}

	// Create service with methods
	serviceName := "TestService"
	service := &descriptorpb.ServiceDescriptorProto{
		Name: &serviceName,
		Method: []*descriptorpb.MethodDescriptorProto{
			{
				Name:       proto.String("GetTest"),
				InputType:  proto.String(".test.package.TestMessage"),
				OutputType: proto.String(".test.package.TestMessage"),
			},
			{
				Name:       proto.String("ListTests"),
				InputType:  proto.String(".test.package.TestMessage"),
				OutputType: proto.String(".test.package.TestMessage"),
			},
			{
				Name:       proto.String("UpdateTest"),
				InputType:  proto.String(".test.package.TestMessage"),
				OutputType: proto.String(".test.package.TestMessage"),
			},
		},
	}

	// Create another service
	anotherServiceName := "AnotherService"
	anotherService := &descriptorpb.ServiceDescriptorProto{
		Name: &anotherServiceName,
		Method: []*descriptorpb.MethodDescriptorProto{
			{
				Name:       proto.String("DoSomething"),
				InputType:  proto.String(".test.package.AnotherMessage"),
				OutputType: proto.String(".test.package.AnotherMessage"),
			},
		},
	}

	return &descriptorpb.FileDescriptorProto{
		Name:        &fileName,
		Package:     &packageName,
		MessageType: []*descriptorpb.DescriptorProto{msg, anotherMsg},
		EnumType:    []*descriptorpb.EnumDescriptorProto{enum},
		Service:     []*descriptorpb.ServiceDescriptorProto{service, anotherService},
	}
}

func TestFindMessage(t *testing.T) {
	file := createTestFileDescriptor()

	testCases := []findTestCase{
		newFindMessageTestCase("find existing message", file, "TestMessage", false),
		newFindMessageTestCase("find another message", file, "AnotherMessage", false),
		newFindMessageTestCase("find non-existent message", file, "NonExistent", true),
		newFindMessageTestCase("find with empty name", file, "", true),
		newFindMessageTestCase("find with nil input", nil, "TestMessage", true),
	}

	core.RunTestCases(t, testCases)
}

func TestFindField(t *testing.T) {
	file := createTestFileDescriptor()
	msg := file.MessageType[0] // TestMessage

	testCases := []findTestCase{
		newFindFieldTestCase("find existing field", msg, "name", false),
		newFindFieldTestCase("find another field", msg, "id", false),
		newFindFieldTestCase("find nested field", msg, "nested", false),
		newFindFieldTestCase("find repeated field", msg, "tags", false),
		newFindFieldTestCase("find non-existent field", msg, "missing", true),
		newFindFieldTestCase("find with empty name", msg, "", true),
		newFindFieldTestCase("find with nil input", nil, "name", true),
	}

	core.RunTestCases(t, testCases)
}

func TestFindEnum(t *testing.T) {
	file := createTestFileDescriptor()

	testCases := []findTestCase{
		newFindEnumTestCase("find existing enum", file, "Status", false),
		newFindEnumTestCase("find non-existent enum", file, "NonExistent", true),
		newFindEnumTestCase("find with empty name", file, "", true),
		newFindEnumTestCase("find with nil input", nil, "Status", true),
	}

	core.RunTestCases(t, testCases)
}

func TestFindService(t *testing.T) {
	file := createTestFileDescriptor()

	testCases := []findTestCase{
		newFindServiceTestCase("find existing service", file, "TestService", false),
		newFindServiceTestCase("find another service", file, "AnotherService", false),
		newFindServiceTestCase("find non-existent service", file, "NonExistent", true),
		newFindServiceTestCase("find with empty name", file, "", true),
		newFindServiceTestCase("find with nil input", nil, "TestService", true),
	}

	core.RunTestCases(t, testCases)
}

func TestFindMethod(t *testing.T) {
	file := createTestFileDescriptor()
	service := file.Service[0] // TestService

	testCases := []findTestCase{
		newFindMethodTestCase("find existing method", service, "GetTest", false),
		newFindMethodTestCase("find another method", service, "ListTests", false),
		newFindMethodTestCase("find third method", service, "UpdateTest", false),
		newFindMethodTestCase("find non-existent method", service, "NonExistent", true),
		newFindMethodTestCase("find with empty name", service, "", true),
		newFindMethodTestCase("find with nil input", nil, "GetTest", true),
	}

	core.RunTestCases(t, testCases)
}

func forEachFieldTestCases() []forEachTestCase {
	file := createTestFileDescriptor()
	msg := file.MessageType[0] // TestMessage with 6 fields
	emptyMsg := &descriptorpb.DescriptorProto{
		Name: proto.String("Empty"),
	}

	errorFn, errorMsg := makeForEachErrorTest(ForEachField)

	return []forEachTestCase{
		newForEachTestCase("iterate all fields", msg,
			ForEachField, "ForEachField", -1, 6, false),
		newForEachTestCase("early termination", msg,
			ForEachField, "ForEachField", 3, 3, false),
		newForEachTestCase("empty message", emptyMsg,
			ForEachField, "ForEachField", -1, 0, false),
		newForEachTestCase("nil input", nil,
			ForEachField, "ForEachField", -1, 0, false),
		newForEachTestCase("error case", errorMsg,
			errorFn, "ForEachField", -1, 0, true),
	}
}

func TestForEachField(t *testing.T) {
	core.RunTestCases(t, forEachFieldTestCases())
}

func TestForEachMessage(t *testing.T) {
	file := createTestFileDescriptor() // Has 2 messages
	emptyFile := &descriptorpb.FileDescriptorProto{
		Name: proto.String("empty.proto"),
	}

	testCases := []forEachTestCase{
		newForEachTestCase("iterate all messages", file,
			ForEachMessage, "ForEachMessage", -1, 2, false),
		newForEachTestCase("early termination", file,
			ForEachMessage, "ForEachMessage", 1, 1, false),
		newForEachTestCase("empty file", emptyFile,
			ForEachMessage, "ForEachMessage", -1, 0, false),
		newForEachTestCase("nil input", nil,
			ForEachMessage, "ForEachMessage", -1, 0, false),
	}

	core.RunTestCases(t, testCases)
}

func TestForEachEnum(t *testing.T) {
	file := createTestFileDescriptor() // Has 1 enum at file level
	emptyFile := &descriptorpb.FileDescriptorProto{
		Name: proto.String("empty.proto"),
	}

	testCases := []forEachTestCase{
		newForEachTestCase("iterate all enums", file,
			ForEachEnum, "ForEachEnum", -1, 1, false),
		newForEachTestCase("early termination", file,
			ForEachEnum, "ForEachEnum", 1, 1, false),
		newForEachTestCase("empty file", emptyFile,
			ForEachEnum, "ForEachEnum", -1, 0, false),
		newForEachTestCase("nil input", nil,
			ForEachEnum, "ForEachEnum", -1, 0, false),
	}

	core.RunTestCases(t, testCases)
}

func TestForEachService(t *testing.T) {
	file := createTestFileDescriptor() // Has 2 services
	emptyFile := &descriptorpb.FileDescriptorProto{
		Name: proto.String("empty.proto"),
	}

	testCases := []forEachTestCase{
		newForEachTestCase("iterate all services", file,
			ForEachService, "ForEachService", -1, 2, false),
		newForEachTestCase("early termination", file,
			ForEachService, "ForEachService", 1, 1, false),
		newForEachTestCase("empty file", emptyFile,
			ForEachService, "ForEachService", -1, 0, false),
		newForEachTestCase("nil input", nil,
			ForEachService, "ForEachService", -1, 0, false),
	}

	core.RunTestCases(t, testCases)
}

func TestForEachMethod(t *testing.T) {
	file := createTestFileDescriptor()
	service := file.Service[0] // TestService with 3 methods
	emptyService := &descriptorpb.ServiceDescriptorProto{
		Name: proto.String("EmptyService"),
	}

	testCases := []forEachTestCase{
		newForEachTestCase("iterate all methods", service,
			ForEachMethod, "ForEachMethod", -1, 3, false),
		newForEachTestCase("early termination", service,
			ForEachMethod, "ForEachMethod", 2, 2, false),
		newForEachTestCase("empty service", emptyService,
			ForEachMethod, "ForEachMethod", -1, 0, false),
		newForEachTestCase("nil input", nil,
			ForEachMethod, "ForEachMethod", -1, 0, false),
	}

	core.RunTestCases(t, testCases)
}

func TestForEachNestedMessage(t *testing.T) {
	file := createTestFileDescriptor()
	msg := file.MessageType[0] // TestMessage with 1 nested message
	emptyMsg := &descriptorpb.DescriptorProto{
		Name: proto.String("Empty"),
	}

	testCases := []forEachTestCase{
		newForEachTestCase("iterate nested messages", msg,
			ForEachNestedMessage, "ForEachNestedMessage", -1, 1, false),
		newForEachTestCase("early termination", msg,
			ForEachNestedMessage, "ForEachNestedMessage", 1, 1, false),
		newForEachTestCase("no nested messages", emptyMsg,
			ForEachNestedMessage, "ForEachNestedMessage", -1, 0, false),
		newForEachTestCase("nil input", nil,
			ForEachNestedMessage, "ForEachNestedMessage", -1, 0, false),
	}

	core.RunTestCases(t, testCases)
}

func TestForEachNestedEnum(t *testing.T) {
	file := createTestFileDescriptor()
	msg := file.MessageType[0] // TestMessage with 1 nested enum
	emptyMsg := &descriptorpb.DescriptorProto{
		Name: proto.String("Empty"),
	}

	testCases := []forEachTestCase{
		newForEachTestCase("iterate nested enums", msg,
			ForEachNestedEnum, "ForEachNestedEnum", -1, 1, false),
		newForEachTestCase("early termination", msg,
			ForEachNestedEnum, "ForEachNestedEnum", 1, 1, false),
		newForEachTestCase("no nested enums", emptyMsg,
			ForEachNestedEnum, "ForEachNestedEnum", -1, 0, false),
		newForEachTestCase("nil input", nil,
			ForEachNestedEnum, "ForEachNestedEnum", -1, 0, false),
	}

	core.RunTestCases(t, testCases)
}
