// Package generator provides utilities for protoc plugin development.
//
// This package includes:
//
// Descriptor type checking utilities:
//   - AsMessage, IsMessage - for DescriptorProto
//   - AsMessageWithName, IsMessageWithName - for DescriptorProto with name check
//   - AsFieldType, IsFieldType - for FieldDescriptorProto
//   - AsEnumType, IsEnumType - for EnumDescriptorProto
//   - AsServiceType, IsServiceType - for ServiceDescriptorProto
//   - AsMethodType, IsMethodType - for MethodDescriptorProto
//   - AsFileType, IsFileType - for FileDescriptorProto
//
// Field classification utilities:
//   - AsScalarField, IsScalarField - for scalar type fields
//   - AsMessageField, IsMessageField - for message type fields (TYPE_MESSAGE only)
//   - AsGroupField, IsGroupField - for group type fields (TYPE_GROUP, deprecated)
//   - AsEnumField, IsEnumField - for enum type fields
//   - AsRepeatedField, IsRepeatedField - for repeated fields
//   - AsMapField, IsMapField - for map fields (heuristic check)
//   - AsMapFieldWithMessage, IsMapFieldWithMessage - for map fields (definitive check)
//   - AsOneOfField, IsOneOfField - for oneof fields
//   - AsOptionalField, IsOptionalField - for optional fields
//   - AsRequiredField, IsRequiredField - for required fields
//
// Test utilities for creating descriptor objects:
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
//   - Descriptor traversal utilities.
//   - Path construction and naming helpers.
//   - Visitor patterns for walking descriptor trees.
//   - Code generation output management.
//   - Context management for build environments.
package generator
