// Package generator provides common utilities for protoc plugin development.
//
// This package includes:
//
// Descriptor type checking utilities:
//   - AsFileType, IsFileType - for FileDescriptorProto
//   - AsMessage, IsMessage - for DescriptorProto without name check
//   - AsMessageType, IsMessageType - for DescriptorProto with name check
//   - AsFieldType, IsFieldType - for FieldDescriptorProto
//   - AsEnumType, IsEnumType - for EnumDescriptorProto
//   - AsServiceType, IsServiceType - for ServiceDescriptorProto
//   - AsMethodType, IsMethodType - for MethodDescriptorProto
//
// Field classification utilities:
//   - AsScalarField, IsScalarField - for scalar type fields
//   - AsMessageField, IsMessageField - for message type fields
//   - AsEnumField, IsEnumField - for enum type fields
//   - AsRepeatedField, IsRepeatedField - for repeated fields
//   - AsMapField, IsMapField - for map fields
//   - AsOneOfField, IsOneOfField - for oneof fields
//   - AsOptionalField, IsOptionalField - for optional fields
//   - AsRequiredField, IsRequiredField - for required fields
//
// Future releases will add:
//   - Descriptor traversal utilities
//   - Path construction and naming helpers
//   - Visitor patterns for walking descriptor trees
//   - Code generation output management
//   - Context management for build environments
package generator
