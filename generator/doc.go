// Package generator provides utilities for protoc plugin development.
//
// Currently provides test utilities for creating descriptor objects:
//   - NewField - create optional field with scalar type.
//   - NewRepeatedField - create repeated field.
//   - NewRequiredField - create required field (proto2).
//   - NewMessageField - create message type field.
//   - NewEnumField - create enum type field.
//   - NewMapField - create map field.
//   - NewOneOfField - create oneof field.
//   - NewFieldWithType - create minimal field with only type set.
//   - NewFieldWithLabel - create minimal field with only label set.
//   - NewRepeatedMessageField - create repeated message field.
//   - NewMessage - create message descriptor.
//   - NewMessageWithNested - create message with nested types.
//   - NewEnum - create enum descriptor.
//   - NewEnumValue - create enum value descriptor.
//   - NewService - create service descriptor.
//   - NewMethod - create method descriptor.
//   - NewFile - create file descriptor.
//   - NewFileWithTypes - create file with messages, enums, and services.
//   - NewOneOf - create oneof descriptor.
//   - Type constants (TypeString, TypeInt32, etc.) for field types.
//
// Future releases will add:
//   - Descriptor type checking and classification utilities.
//   - Path construction and naming helpers.
//   - Visitor patterns for walking descriptor trees.
//   - Code generation output management.
//   - Context management for build environments.
package generator
