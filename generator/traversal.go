package generator

import (
	"google.golang.org/protobuf/proto"
)

// FindMessage finds a message by name in a file descriptor
func FindMessage(file proto.Message, name string) proto.Message {
	if file == nil || name == "" {
		return nil
	}

	fileDesc, ok := AsFileType(file)
	if !ok {
		return nil
	}

	for _, msg := range fileDesc.MessageType {
		if msg.GetName() == name {
			return msg
		}
	}

	return nil
}

// FindField finds a field by name in a message descriptor
func FindField(msg proto.Message, name string) proto.Message {
	if msg == nil || name == "" {
		return nil
	}

	msgDesc, ok := AsMessage(msg)
	if !ok {
		return nil
	}

	for _, field := range msgDesc.Field {
		if field.GetName() == name {
			return field
		}
	}

	return nil
}

// FindEnum finds an enum by name in a file descriptor
func FindEnum(file proto.Message, name string) proto.Message {
	if file == nil || name == "" {
		return nil
	}

	fileDesc, ok := AsFileType(file)
	if !ok {
		return nil
	}

	for _, enum := range fileDesc.EnumType {
		if enum.GetName() == name {
			return enum
		}
	}

	return nil
}

// FindService finds a service by name in a file descriptor
func FindService(file proto.Message, name string) proto.Message {
	if file == nil || name == "" {
		return nil
	}

	fileDesc, ok := AsFileType(file)
	if !ok {
		return nil
	}

	for _, svc := range fileDesc.Service {
		if svc.GetName() == name {
			return svc
		}
	}

	return nil
}

// FindMethod finds a method by name in a service descriptor
func FindMethod(service proto.Message, name string) proto.Message {
	if service == nil || name == "" {
		return nil
	}

	svcDesc, ok := AsServiceType(service)
	if !ok {
		return nil
	}

	for _, method := range svcDesc.Method {
		if method.GetName() == name {
			return method
		}
	}

	return nil
}

// ForEachField iterates over all fields in a message descriptor
// The callback returns true to continue iteration, false to stop
func ForEachField(msg proto.Message, fn func(proto.Message) bool) error {
	if msg == nil || fn == nil {
		return nil
	}

	msgDesc, ok := AsMessage(msg)
	if !ok {
		return nil
	}

	for _, field := range msgDesc.Field {
		if !fn(field) {
			break
		}
	}

	return nil
}

// ForEachMessage iterates over all messages in a file descriptor
// The callback returns true to continue iteration, false to stop
func ForEachMessage(file proto.Message, fn func(proto.Message) bool) error {
	if file == nil || fn == nil {
		return nil
	}

	fileDesc, ok := AsFileType(file)
	if !ok {
		return nil
	}

	for _, msg := range fileDesc.MessageType {
		if !fn(msg) {
			break
		}
	}

	return nil
}

// ForEachEnum iterates over all enums in a file descriptor
// The callback returns true to continue iteration, false to stop
func ForEachEnum(file proto.Message, fn func(proto.Message) bool) error {
	if file == nil || fn == nil {
		return nil
	}

	fileDesc, ok := AsFileType(file)
	if !ok {
		return nil
	}

	for _, enum := range fileDesc.EnumType {
		if !fn(enum) {
			break
		}
	}

	return nil
}

// ForEachService iterates over all services in a file descriptor
// The callback returns true to continue iteration, false to stop
func ForEachService(file proto.Message, fn func(proto.Message) bool) error {
	if file == nil || fn == nil {
		return nil
	}

	fileDesc, ok := AsFileType(file)
	if !ok {
		return nil
	}

	for _, svc := range fileDesc.Service {
		if !fn(svc) {
			break
		}
	}

	return nil
}

// ForEachMethod iterates over all methods in a service descriptor
// The callback returns true to continue iteration, false to stop
func ForEachMethod(service proto.Message, fn func(proto.Message) bool) error {
	if service == nil || fn == nil {
		return nil
	}

	svcDesc, ok := AsServiceType(service)
	if !ok {
		return nil
	}

	for _, method := range svcDesc.Method {
		if !fn(method) {
			break
		}
	}

	return nil
}

// ForEachNestedMessage iterates over all nested messages in a message descriptor
// The callback returns true to continue iteration, false to stop
func ForEachNestedMessage(msg proto.Message, fn func(proto.Message) bool) error {
	if msg == nil || fn == nil {
		return nil
	}

	msgDesc, ok := AsMessage(msg)
	if !ok {
		return nil
	}

	for _, nested := range msgDesc.NestedType {
		if !fn(nested) {
			break
		}
	}

	return nil
}

// ForEachNestedEnum iterates over all nested enums in a message descriptor
// The callback returns true to continue iteration, false to stop
func ForEachNestedEnum(msg proto.Message, fn func(proto.Message) bool) error {
	if msg == nil || fn == nil {
		return nil
	}

	msgDesc, ok := AsMessage(msg)
	if !ok {
		return nil
	}

	for _, enum := range msgDesc.EnumType {
		if !fn(enum) {
			break
		}
	}

	return nil
}
