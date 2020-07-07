// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: flyteidl/core/catalog.proto

#ifndef PROTOBUF_INCLUDED_flyteidl_2fcore_2fcatalog_2eproto
#define PROTOBUF_INCLUDED_flyteidl_2fcore_2fcatalog_2eproto

#include <limits>
#include <string>

#include <google/protobuf/port_def.inc>
#if PROTOBUF_VERSION < 3007000
#error This file was generated by a newer version of protoc which is
#error incompatible with your Protocol Buffer headers. Please update
#error your headers.
#endif
#if 3007000 < PROTOBUF_MIN_PROTOC_VERSION
#error This file was generated by an older version of protoc which is
#error incompatible with your Protocol Buffer headers. Please
#error regenerate this file with a newer version of protoc.
#endif

#include <google/protobuf/port_undef.inc>
#include <google/protobuf/io/coded_stream.h>
#include <google/protobuf/arena.h>
#include <google/protobuf/arenastring.h>
#include <google/protobuf/generated_message_table_driven.h>
#include <google/protobuf/generated_message_util.h>
#include <google/protobuf/inlined_string_field.h>
#include <google/protobuf/metadata.h>
#include <google/protobuf/message.h>
#include <google/protobuf/repeated_field.h>  // IWYU pragma: export
#include <google/protobuf/extension_set.h>  // IWYU pragma: export
#include <google/protobuf/generated_enum_reflection.h>
#include <google/protobuf/unknown_field_set.h>
#include "flyteidl/core/identifier.pb.h"
// @@protoc_insertion_point(includes)
#include <google/protobuf/port_def.inc>
#define PROTOBUF_INTERNAL_EXPORT_flyteidl_2fcore_2fcatalog_2eproto

// Internal implementation detail -- do not use these members.
struct TableStruct_flyteidl_2fcore_2fcatalog_2eproto {
  static const ::google::protobuf::internal::ParseTableField entries[]
    PROTOBUF_SECTION_VARIABLE(protodesc_cold);
  static const ::google::protobuf::internal::AuxillaryParseTableField aux[]
    PROTOBUF_SECTION_VARIABLE(protodesc_cold);
  static const ::google::protobuf::internal::ParseTable schema[2]
    PROTOBUF_SECTION_VARIABLE(protodesc_cold);
  static const ::google::protobuf::internal::FieldMetadata field_metadata[];
  static const ::google::protobuf::internal::SerializationTable serialization_table[];
  static const ::google::protobuf::uint32 offsets[];
};
void AddDescriptors_flyteidl_2fcore_2fcatalog_2eproto();
namespace flyteidl {
namespace core {
class CatalogArtifactTag;
class CatalogArtifactTagDefaultTypeInternal;
extern CatalogArtifactTagDefaultTypeInternal _CatalogArtifactTag_default_instance_;
class CatalogMetadata;
class CatalogMetadataDefaultTypeInternal;
extern CatalogMetadataDefaultTypeInternal _CatalogMetadata_default_instance_;
}  // namespace core
}  // namespace flyteidl
namespace google {
namespace protobuf {
template<> ::flyteidl::core::CatalogArtifactTag* Arena::CreateMaybeMessage<::flyteidl::core::CatalogArtifactTag>(Arena*);
template<> ::flyteidl::core::CatalogMetadata* Arena::CreateMaybeMessage<::flyteidl::core::CatalogMetadata>(Arena*);
}  // namespace protobuf
}  // namespace google
namespace flyteidl {
namespace core {

enum CatalogCacheStatus {
  CACHE_DISABLED = 0,
  CACHE_MISS = 1,
  CACHE_HIT = 2,
  CACHE_POPULATED = 3,
  CACHE_LOOKUP_FAILURE = 4,
  CACHE_PUT_FAILURE = 5,
  CatalogCacheStatus_INT_MIN_SENTINEL_DO_NOT_USE_ = std::numeric_limits<::google::protobuf::int32>::min(),
  CatalogCacheStatus_INT_MAX_SENTINEL_DO_NOT_USE_ = std::numeric_limits<::google::protobuf::int32>::max()
};
bool CatalogCacheStatus_IsValid(int value);
const CatalogCacheStatus CatalogCacheStatus_MIN = CACHE_DISABLED;
const CatalogCacheStatus CatalogCacheStatus_MAX = CACHE_PUT_FAILURE;
const int CatalogCacheStatus_ARRAYSIZE = CatalogCacheStatus_MAX + 1;

const ::google::protobuf::EnumDescriptor* CatalogCacheStatus_descriptor();
inline const ::std::string& CatalogCacheStatus_Name(CatalogCacheStatus value) {
  return ::google::protobuf::internal::NameOfEnum(
    CatalogCacheStatus_descriptor(), value);
}
inline bool CatalogCacheStatus_Parse(
    const ::std::string& name, CatalogCacheStatus* value) {
  return ::google::protobuf::internal::ParseNamedEnum<CatalogCacheStatus>(
    CatalogCacheStatus_descriptor(), name, value);
}
// ===================================================================

class CatalogArtifactTag final :
    public ::google::protobuf::Message /* @@protoc_insertion_point(class_definition:flyteidl.core.CatalogArtifactTag) */ {
 public:
  CatalogArtifactTag();
  virtual ~CatalogArtifactTag();

  CatalogArtifactTag(const CatalogArtifactTag& from);

  inline CatalogArtifactTag& operator=(const CatalogArtifactTag& from) {
    CopyFrom(from);
    return *this;
  }
  #if LANG_CXX11
  CatalogArtifactTag(CatalogArtifactTag&& from) noexcept
    : CatalogArtifactTag() {
    *this = ::std::move(from);
  }

  inline CatalogArtifactTag& operator=(CatalogArtifactTag&& from) noexcept {
    if (GetArenaNoVirtual() == from.GetArenaNoVirtual()) {
      if (this != &from) InternalSwap(&from);
    } else {
      CopyFrom(from);
    }
    return *this;
  }
  #endif
  static const ::google::protobuf::Descriptor* descriptor() {
    return default_instance().GetDescriptor();
  }
  static const CatalogArtifactTag& default_instance();

  static void InitAsDefaultInstance();  // FOR INTERNAL USE ONLY
  static inline const CatalogArtifactTag* internal_default_instance() {
    return reinterpret_cast<const CatalogArtifactTag*>(
               &_CatalogArtifactTag_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    0;

  void Swap(CatalogArtifactTag* other);
  friend void swap(CatalogArtifactTag& a, CatalogArtifactTag& b) {
    a.Swap(&b);
  }

  // implements Message ----------------------------------------------

  inline CatalogArtifactTag* New() const final {
    return CreateMaybeMessage<CatalogArtifactTag>(nullptr);
  }

  CatalogArtifactTag* New(::google::protobuf::Arena* arena) const final {
    return CreateMaybeMessage<CatalogArtifactTag>(arena);
  }
  void CopyFrom(const ::google::protobuf::Message& from) final;
  void MergeFrom(const ::google::protobuf::Message& from) final;
  void CopyFrom(const CatalogArtifactTag& from);
  void MergeFrom(const CatalogArtifactTag& from);
  PROTOBUF_ATTRIBUTE_REINITIALIZES void Clear() final;
  bool IsInitialized() const final;

  size_t ByteSizeLong() const final;
  #if GOOGLE_PROTOBUF_ENABLE_EXPERIMENTAL_PARSER
  static const char* _InternalParse(const char* begin, const char* end, void* object, ::google::protobuf::internal::ParseContext* ctx);
  ::google::protobuf::internal::ParseFunc _ParseFunc() const final { return _InternalParse; }
  #else
  bool MergePartialFromCodedStream(
      ::google::protobuf::io::CodedInputStream* input) final;
  #endif  // GOOGLE_PROTOBUF_ENABLE_EXPERIMENTAL_PARSER
  void SerializeWithCachedSizes(
      ::google::protobuf::io::CodedOutputStream* output) const final;
  ::google::protobuf::uint8* InternalSerializeWithCachedSizesToArray(
      ::google::protobuf::uint8* target) const final;
  int GetCachedSize() const final { return _cached_size_.Get(); }

  private:
  void SharedCtor();
  void SharedDtor();
  void SetCachedSize(int size) const final;
  void InternalSwap(CatalogArtifactTag* other);
  private:
  inline ::google::protobuf::Arena* GetArenaNoVirtual() const {
    return nullptr;
  }
  inline void* MaybeArenaPtr() const {
    return nullptr;
  }
  public:

  ::google::protobuf::Metadata GetMetadata() const final;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  // string artifact_id = 1;
  void clear_artifact_id();
  static const int kArtifactIdFieldNumber = 1;
  const ::std::string& artifact_id() const;
  void set_artifact_id(const ::std::string& value);
  #if LANG_CXX11
  void set_artifact_id(::std::string&& value);
  #endif
  void set_artifact_id(const char* value);
  void set_artifact_id(const char* value, size_t size);
  ::std::string* mutable_artifact_id();
  ::std::string* release_artifact_id();
  void set_allocated_artifact_id(::std::string* artifact_id);

  // string name = 2;
  void clear_name();
  static const int kNameFieldNumber = 2;
  const ::std::string& name() const;
  void set_name(const ::std::string& value);
  #if LANG_CXX11
  void set_name(::std::string&& value);
  #endif
  void set_name(const char* value);
  void set_name(const char* value, size_t size);
  ::std::string* mutable_name();
  ::std::string* release_name();
  void set_allocated_name(::std::string* name);

  // @@protoc_insertion_point(class_scope:flyteidl.core.CatalogArtifactTag)
 private:
  class HasBitSetters;

  ::google::protobuf::internal::InternalMetadataWithArena _internal_metadata_;
  ::google::protobuf::internal::ArenaStringPtr artifact_id_;
  ::google::protobuf::internal::ArenaStringPtr name_;
  mutable ::google::protobuf::internal::CachedSize _cached_size_;
  friend struct ::TableStruct_flyteidl_2fcore_2fcatalog_2eproto;
};
// -------------------------------------------------------------------

class CatalogMetadata final :
    public ::google::protobuf::Message /* @@protoc_insertion_point(class_definition:flyteidl.core.CatalogMetadata) */ {
 public:
  CatalogMetadata();
  virtual ~CatalogMetadata();

  CatalogMetadata(const CatalogMetadata& from);

  inline CatalogMetadata& operator=(const CatalogMetadata& from) {
    CopyFrom(from);
    return *this;
  }
  #if LANG_CXX11
  CatalogMetadata(CatalogMetadata&& from) noexcept
    : CatalogMetadata() {
    *this = ::std::move(from);
  }

  inline CatalogMetadata& operator=(CatalogMetadata&& from) noexcept {
    if (GetArenaNoVirtual() == from.GetArenaNoVirtual()) {
      if (this != &from) InternalSwap(&from);
    } else {
      CopyFrom(from);
    }
    return *this;
  }
  #endif
  static const ::google::protobuf::Descriptor* descriptor() {
    return default_instance().GetDescriptor();
  }
  static const CatalogMetadata& default_instance();

  enum SourceExecutionCase {
    kSourceTaskExecution = 3,
    SOURCE_EXECUTION_NOT_SET = 0,
  };

  static void InitAsDefaultInstance();  // FOR INTERNAL USE ONLY
  static inline const CatalogMetadata* internal_default_instance() {
    return reinterpret_cast<const CatalogMetadata*>(
               &_CatalogMetadata_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    1;

  void Swap(CatalogMetadata* other);
  friend void swap(CatalogMetadata& a, CatalogMetadata& b) {
    a.Swap(&b);
  }

  // implements Message ----------------------------------------------

  inline CatalogMetadata* New() const final {
    return CreateMaybeMessage<CatalogMetadata>(nullptr);
  }

  CatalogMetadata* New(::google::protobuf::Arena* arena) const final {
    return CreateMaybeMessage<CatalogMetadata>(arena);
  }
  void CopyFrom(const ::google::protobuf::Message& from) final;
  void MergeFrom(const ::google::protobuf::Message& from) final;
  void CopyFrom(const CatalogMetadata& from);
  void MergeFrom(const CatalogMetadata& from);
  PROTOBUF_ATTRIBUTE_REINITIALIZES void Clear() final;
  bool IsInitialized() const final;

  size_t ByteSizeLong() const final;
  #if GOOGLE_PROTOBUF_ENABLE_EXPERIMENTAL_PARSER
  static const char* _InternalParse(const char* begin, const char* end, void* object, ::google::protobuf::internal::ParseContext* ctx);
  ::google::protobuf::internal::ParseFunc _ParseFunc() const final { return _InternalParse; }
  #else
  bool MergePartialFromCodedStream(
      ::google::protobuf::io::CodedInputStream* input) final;
  #endif  // GOOGLE_PROTOBUF_ENABLE_EXPERIMENTAL_PARSER
  void SerializeWithCachedSizes(
      ::google::protobuf::io::CodedOutputStream* output) const final;
  ::google::protobuf::uint8* InternalSerializeWithCachedSizesToArray(
      ::google::protobuf::uint8* target) const final;
  int GetCachedSize() const final { return _cached_size_.Get(); }

  private:
  void SharedCtor();
  void SharedDtor();
  void SetCachedSize(int size) const final;
  void InternalSwap(CatalogMetadata* other);
  private:
  inline ::google::protobuf::Arena* GetArenaNoVirtual() const {
    return nullptr;
  }
  inline void* MaybeArenaPtr() const {
    return nullptr;
  }
  public:

  ::google::protobuf::Metadata GetMetadata() const final;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  // .flyteidl.core.Identifier dataset_id = 1;
  bool has_dataset_id() const;
  void clear_dataset_id();
  static const int kDatasetIdFieldNumber = 1;
  const ::flyteidl::core::Identifier& dataset_id() const;
  ::flyteidl::core::Identifier* release_dataset_id();
  ::flyteidl::core::Identifier* mutable_dataset_id();
  void set_allocated_dataset_id(::flyteidl::core::Identifier* dataset_id);

  // .flyteidl.core.CatalogArtifactTag artifact_tag = 2;
  bool has_artifact_tag() const;
  void clear_artifact_tag();
  static const int kArtifactTagFieldNumber = 2;
  const ::flyteidl::core::CatalogArtifactTag& artifact_tag() const;
  ::flyteidl::core::CatalogArtifactTag* release_artifact_tag();
  ::flyteidl::core::CatalogArtifactTag* mutable_artifact_tag();
  void set_allocated_artifact_tag(::flyteidl::core::CatalogArtifactTag* artifact_tag);

  // .flyteidl.core.TaskExecutionIdentifier source_task_execution = 3;
  bool has_source_task_execution() const;
  void clear_source_task_execution();
  static const int kSourceTaskExecutionFieldNumber = 3;
  const ::flyteidl::core::TaskExecutionIdentifier& source_task_execution() const;
  ::flyteidl::core::TaskExecutionIdentifier* release_source_task_execution();
  ::flyteidl::core::TaskExecutionIdentifier* mutable_source_task_execution();
  void set_allocated_source_task_execution(::flyteidl::core::TaskExecutionIdentifier* source_task_execution);

  void clear_source_execution();
  SourceExecutionCase source_execution_case() const;
  // @@protoc_insertion_point(class_scope:flyteidl.core.CatalogMetadata)
 private:
  class HasBitSetters;
  void set_has_source_task_execution();

  inline bool has_source_execution() const;
  inline void clear_has_source_execution();

  ::google::protobuf::internal::InternalMetadataWithArena _internal_metadata_;
  ::flyteidl::core::Identifier* dataset_id_;
  ::flyteidl::core::CatalogArtifactTag* artifact_tag_;
  union SourceExecutionUnion {
    SourceExecutionUnion() {}
    ::flyteidl::core::TaskExecutionIdentifier* source_task_execution_;
  } source_execution_;
  mutable ::google::protobuf::internal::CachedSize _cached_size_;
  ::google::protobuf::uint32 _oneof_case_[1];

  friend struct ::TableStruct_flyteidl_2fcore_2fcatalog_2eproto;
};
// ===================================================================


// ===================================================================

#ifdef __GNUC__
  #pragma GCC diagnostic push
  #pragma GCC diagnostic ignored "-Wstrict-aliasing"
#endif  // __GNUC__
// CatalogArtifactTag

// string artifact_id = 1;
inline void CatalogArtifactTag::clear_artifact_id() {
  artifact_id_.ClearToEmptyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline const ::std::string& CatalogArtifactTag::artifact_id() const {
  // @@protoc_insertion_point(field_get:flyteidl.core.CatalogArtifactTag.artifact_id)
  return artifact_id_.GetNoArena();
}
inline void CatalogArtifactTag::set_artifact_id(const ::std::string& value) {
  
  artifact_id_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), value);
  // @@protoc_insertion_point(field_set:flyteidl.core.CatalogArtifactTag.artifact_id)
}
#if LANG_CXX11
inline void CatalogArtifactTag::set_artifact_id(::std::string&& value) {
  
  artifact_id_.SetNoArena(
    &::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::move(value));
  // @@protoc_insertion_point(field_set_rvalue:flyteidl.core.CatalogArtifactTag.artifact_id)
}
#endif
inline void CatalogArtifactTag::set_artifact_id(const char* value) {
  GOOGLE_DCHECK(value != nullptr);
  
  artifact_id_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::string(value));
  // @@protoc_insertion_point(field_set_char:flyteidl.core.CatalogArtifactTag.artifact_id)
}
inline void CatalogArtifactTag::set_artifact_id(const char* value, size_t size) {
  
  artifact_id_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(),
      ::std::string(reinterpret_cast<const char*>(value), size));
  // @@protoc_insertion_point(field_set_pointer:flyteidl.core.CatalogArtifactTag.artifact_id)
}
inline ::std::string* CatalogArtifactTag::mutable_artifact_id() {
  
  // @@protoc_insertion_point(field_mutable:flyteidl.core.CatalogArtifactTag.artifact_id)
  return artifact_id_.MutableNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline ::std::string* CatalogArtifactTag::release_artifact_id() {
  // @@protoc_insertion_point(field_release:flyteidl.core.CatalogArtifactTag.artifact_id)
  
  return artifact_id_.ReleaseNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline void CatalogArtifactTag::set_allocated_artifact_id(::std::string* artifact_id) {
  if (artifact_id != nullptr) {
    
  } else {
    
  }
  artifact_id_.SetAllocatedNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), artifact_id);
  // @@protoc_insertion_point(field_set_allocated:flyteidl.core.CatalogArtifactTag.artifact_id)
}

// string name = 2;
inline void CatalogArtifactTag::clear_name() {
  name_.ClearToEmptyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline const ::std::string& CatalogArtifactTag::name() const {
  // @@protoc_insertion_point(field_get:flyteidl.core.CatalogArtifactTag.name)
  return name_.GetNoArena();
}
inline void CatalogArtifactTag::set_name(const ::std::string& value) {
  
  name_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), value);
  // @@protoc_insertion_point(field_set:flyteidl.core.CatalogArtifactTag.name)
}
#if LANG_CXX11
inline void CatalogArtifactTag::set_name(::std::string&& value) {
  
  name_.SetNoArena(
    &::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::move(value));
  // @@protoc_insertion_point(field_set_rvalue:flyteidl.core.CatalogArtifactTag.name)
}
#endif
inline void CatalogArtifactTag::set_name(const char* value) {
  GOOGLE_DCHECK(value != nullptr);
  
  name_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::string(value));
  // @@protoc_insertion_point(field_set_char:flyteidl.core.CatalogArtifactTag.name)
}
inline void CatalogArtifactTag::set_name(const char* value, size_t size) {
  
  name_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(),
      ::std::string(reinterpret_cast<const char*>(value), size));
  // @@protoc_insertion_point(field_set_pointer:flyteidl.core.CatalogArtifactTag.name)
}
inline ::std::string* CatalogArtifactTag::mutable_name() {
  
  // @@protoc_insertion_point(field_mutable:flyteidl.core.CatalogArtifactTag.name)
  return name_.MutableNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline ::std::string* CatalogArtifactTag::release_name() {
  // @@protoc_insertion_point(field_release:flyteidl.core.CatalogArtifactTag.name)
  
  return name_.ReleaseNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline void CatalogArtifactTag::set_allocated_name(::std::string* name) {
  if (name != nullptr) {
    
  } else {
    
  }
  name_.SetAllocatedNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), name);
  // @@protoc_insertion_point(field_set_allocated:flyteidl.core.CatalogArtifactTag.name)
}

// -------------------------------------------------------------------

// CatalogMetadata

// .flyteidl.core.Identifier dataset_id = 1;
inline bool CatalogMetadata::has_dataset_id() const {
  return this != internal_default_instance() && dataset_id_ != nullptr;
}
inline const ::flyteidl::core::Identifier& CatalogMetadata::dataset_id() const {
  const ::flyteidl::core::Identifier* p = dataset_id_;
  // @@protoc_insertion_point(field_get:flyteidl.core.CatalogMetadata.dataset_id)
  return p != nullptr ? *p : *reinterpret_cast<const ::flyteidl::core::Identifier*>(
      &::flyteidl::core::_Identifier_default_instance_);
}
inline ::flyteidl::core::Identifier* CatalogMetadata::release_dataset_id() {
  // @@protoc_insertion_point(field_release:flyteidl.core.CatalogMetadata.dataset_id)
  
  ::flyteidl::core::Identifier* temp = dataset_id_;
  dataset_id_ = nullptr;
  return temp;
}
inline ::flyteidl::core::Identifier* CatalogMetadata::mutable_dataset_id() {
  
  if (dataset_id_ == nullptr) {
    auto* p = CreateMaybeMessage<::flyteidl::core::Identifier>(GetArenaNoVirtual());
    dataset_id_ = p;
  }
  // @@protoc_insertion_point(field_mutable:flyteidl.core.CatalogMetadata.dataset_id)
  return dataset_id_;
}
inline void CatalogMetadata::set_allocated_dataset_id(::flyteidl::core::Identifier* dataset_id) {
  ::google::protobuf::Arena* message_arena = GetArenaNoVirtual();
  if (message_arena == nullptr) {
    delete reinterpret_cast< ::google::protobuf::MessageLite*>(dataset_id_);
  }
  if (dataset_id) {
    ::google::protobuf::Arena* submessage_arena = nullptr;
    if (message_arena != submessage_arena) {
      dataset_id = ::google::protobuf::internal::GetOwnedMessage(
          message_arena, dataset_id, submessage_arena);
    }
    
  } else {
    
  }
  dataset_id_ = dataset_id;
  // @@protoc_insertion_point(field_set_allocated:flyteidl.core.CatalogMetadata.dataset_id)
}

// .flyteidl.core.CatalogArtifactTag artifact_tag = 2;
inline bool CatalogMetadata::has_artifact_tag() const {
  return this != internal_default_instance() && artifact_tag_ != nullptr;
}
inline void CatalogMetadata::clear_artifact_tag() {
  if (GetArenaNoVirtual() == nullptr && artifact_tag_ != nullptr) {
    delete artifact_tag_;
  }
  artifact_tag_ = nullptr;
}
inline const ::flyteidl::core::CatalogArtifactTag& CatalogMetadata::artifact_tag() const {
  const ::flyteidl::core::CatalogArtifactTag* p = artifact_tag_;
  // @@protoc_insertion_point(field_get:flyteidl.core.CatalogMetadata.artifact_tag)
  return p != nullptr ? *p : *reinterpret_cast<const ::flyteidl::core::CatalogArtifactTag*>(
      &::flyteidl::core::_CatalogArtifactTag_default_instance_);
}
inline ::flyteidl::core::CatalogArtifactTag* CatalogMetadata::release_artifact_tag() {
  // @@protoc_insertion_point(field_release:flyteidl.core.CatalogMetadata.artifact_tag)
  
  ::flyteidl::core::CatalogArtifactTag* temp = artifact_tag_;
  artifact_tag_ = nullptr;
  return temp;
}
inline ::flyteidl::core::CatalogArtifactTag* CatalogMetadata::mutable_artifact_tag() {
  
  if (artifact_tag_ == nullptr) {
    auto* p = CreateMaybeMessage<::flyteidl::core::CatalogArtifactTag>(GetArenaNoVirtual());
    artifact_tag_ = p;
  }
  // @@protoc_insertion_point(field_mutable:flyteidl.core.CatalogMetadata.artifact_tag)
  return artifact_tag_;
}
inline void CatalogMetadata::set_allocated_artifact_tag(::flyteidl::core::CatalogArtifactTag* artifact_tag) {
  ::google::protobuf::Arena* message_arena = GetArenaNoVirtual();
  if (message_arena == nullptr) {
    delete artifact_tag_;
  }
  if (artifact_tag) {
    ::google::protobuf::Arena* submessage_arena = nullptr;
    if (message_arena != submessage_arena) {
      artifact_tag = ::google::protobuf::internal::GetOwnedMessage(
          message_arena, artifact_tag, submessage_arena);
    }
    
  } else {
    
  }
  artifact_tag_ = artifact_tag;
  // @@protoc_insertion_point(field_set_allocated:flyteidl.core.CatalogMetadata.artifact_tag)
}

// .flyteidl.core.TaskExecutionIdentifier source_task_execution = 3;
inline bool CatalogMetadata::has_source_task_execution() const {
  return source_execution_case() == kSourceTaskExecution;
}
inline void CatalogMetadata::set_has_source_task_execution() {
  _oneof_case_[0] = kSourceTaskExecution;
}
inline ::flyteidl::core::TaskExecutionIdentifier* CatalogMetadata::release_source_task_execution() {
  // @@protoc_insertion_point(field_release:flyteidl.core.CatalogMetadata.source_task_execution)
  if (has_source_task_execution()) {
    clear_has_source_execution();
      ::flyteidl::core::TaskExecutionIdentifier* temp = source_execution_.source_task_execution_;
    source_execution_.source_task_execution_ = nullptr;
    return temp;
  } else {
    return nullptr;
  }
}
inline const ::flyteidl::core::TaskExecutionIdentifier& CatalogMetadata::source_task_execution() const {
  // @@protoc_insertion_point(field_get:flyteidl.core.CatalogMetadata.source_task_execution)
  return has_source_task_execution()
      ? *source_execution_.source_task_execution_
      : *reinterpret_cast< ::flyteidl::core::TaskExecutionIdentifier*>(&::flyteidl::core::_TaskExecutionIdentifier_default_instance_);
}
inline ::flyteidl::core::TaskExecutionIdentifier* CatalogMetadata::mutable_source_task_execution() {
  if (!has_source_task_execution()) {
    clear_source_execution();
    set_has_source_task_execution();
    source_execution_.source_task_execution_ = CreateMaybeMessage< ::flyteidl::core::TaskExecutionIdentifier >(
        GetArenaNoVirtual());
  }
  // @@protoc_insertion_point(field_mutable:flyteidl.core.CatalogMetadata.source_task_execution)
  return source_execution_.source_task_execution_;
}

inline bool CatalogMetadata::has_source_execution() const {
  return source_execution_case() != SOURCE_EXECUTION_NOT_SET;
}
inline void CatalogMetadata::clear_has_source_execution() {
  _oneof_case_[0] = SOURCE_EXECUTION_NOT_SET;
}
inline CatalogMetadata::SourceExecutionCase CatalogMetadata::source_execution_case() const {
  return CatalogMetadata::SourceExecutionCase(_oneof_case_[0]);
}
#ifdef __GNUC__
  #pragma GCC diagnostic pop
#endif  // __GNUC__
// -------------------------------------------------------------------


// @@protoc_insertion_point(namespace_scope)

}  // namespace core
}  // namespace flyteidl

namespace google {
namespace protobuf {

template <> struct is_proto_enum< ::flyteidl::core::CatalogCacheStatus> : ::std::true_type {};
template <>
inline const EnumDescriptor* GetEnumDescriptor< ::flyteidl::core::CatalogCacheStatus>() {
  return ::flyteidl::core::CatalogCacheStatus_descriptor();
}

}  // namespace protobuf
}  // namespace google

// @@protoc_insertion_point(global_scope)

#include <google/protobuf/port_undef.inc>
#endif  // PROTOBUF_INCLUDED_flyteidl_2fcore_2fcatalog_2eproto
