// @generated by protoc-gen-es v1.10.0
// @generated from file onehub/v1/messages.proto (package onehub.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, FieldMask, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { Message as Message$1, Pagination, PaginationResponse } from "./models_pb.js";

/**
 * *
 * Message creation request object
 *
 * @generated from message onehub.v1.CreateMessagesRequest
 */
export declare class CreateMessagesRequest extends Message<CreateMessagesRequest> {
  /**
   * *
   * Topic where messages are being created
   *
   * @generated from field: string topic_id = 1;
   */
  topicId: string;

  /**
   * *
   * Message being updated
   *
   * @generated from field: repeated onehub.v1.Message messages = 2;
   */
  messages: Message$1[];

  /**
   * *
   * Whether to allow custom user IDs or whether to
   * force user IDs to be overridden to the logged in user.
   * In batch mode we want the option to have diff user IDs
   * In prod - we want to ensure that only Admins can provide
   * this option.
   *
   * @generated from field: bool allow_userids = 3;
   */
  allowUserids: boolean;

  constructor(data?: PartialMessage<CreateMessagesRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "onehub.v1.CreateMessagesRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateMessagesRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateMessagesRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateMessagesRequest;

  static equals(a: CreateMessagesRequest | PlainMessage<CreateMessagesRequest> | undefined, b: CreateMessagesRequest | PlainMessage<CreateMessagesRequest> | undefined): boolean;
}

/**
 * *
 * Response of an message creation.
 *
 * @generated from message onehub.v1.CreateMessagesResponse
 */
export declare class CreateMessagesResponse extends Message<CreateMessagesResponse> {
  /**
   * *
   * Message being created
   *
   * @generated from field: repeated onehub.v1.Message messages = 1;
   */
  messages: Message$1[];

  constructor(data?: PartialMessage<CreateMessagesResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "onehub.v1.CreateMessagesResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateMessagesResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateMessagesResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateMessagesResponse;

  static equals(a: CreateMessagesResponse | PlainMessage<CreateMessagesResponse> | undefined, b: CreateMessagesResponse | PlainMessage<CreateMessagesResponse> | undefined): boolean;
}

/**
 * *
 * Bulk importing of messages with very minimal checks.
 * Here no validation is performed on the messages (ie checking topic IDS)
 * setting current user id, setting created/updated time stamps etc.
 *
 * Use this either for recovery (typically you should do DR on the DB) or
 * or for testing.
 *
 * @generated from message onehub.v1.ImportMessagesRequest
 */
export declare class ImportMessagesRequest extends Message<ImportMessagesRequest> {
  /**
   * *
   * Message being updated
   *
   * @generated from field: repeated onehub.v1.Message messages = 2;
   */
  messages: Message$1[];

  constructor(data?: PartialMessage<ImportMessagesRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "onehub.v1.ImportMessagesRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ImportMessagesRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ImportMessagesRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ImportMessagesRequest;

  static equals(a: ImportMessagesRequest | PlainMessage<ImportMessagesRequest> | undefined, b: ImportMessagesRequest | PlainMessage<ImportMessagesRequest> | undefined): boolean;
}

/**
 * *
 * Response of an message import.
 *
 * @generated from message onehub.v1.ImportMessagesResponse
 */
export declare class ImportMessagesResponse extends Message<ImportMessagesResponse> {
  /**
   * *
   * Message being created
   *
   * @generated from field: repeated onehub.v1.Message messages = 1;
   */
  messages: Message$1[];

  constructor(data?: PartialMessage<ImportMessagesResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "onehub.v1.ImportMessagesResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ImportMessagesResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ImportMessagesResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ImportMessagesResponse;

  static equals(a: ImportMessagesResponse | PlainMessage<ImportMessagesResponse> | undefined, b: ImportMessagesResponse | PlainMessage<ImportMessagesResponse> | undefined): boolean;
}

/**
 * *
 * A message listing request.  For now only paginations params are provided.
 *
 * @generated from message onehub.v1.ListMessagesRequest
 */
export declare class ListMessagesRequest extends Message<ListMessagesRequest> {
  /**
   * *
   * Topic in which messages are to be listed.  Required.
   *
   * @generated from field: string topic_id = 1;
   */
  topicId: string;

  /**
   * *
   * Pagination prameters.
   *
   * @generated from field: onehub.v1.Pagination pagination = 2;
   */
  pagination?: Pagination;

  constructor(data?: PartialMessage<ListMessagesRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "onehub.v1.ListMessagesRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListMessagesRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListMessagesRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListMessagesRequest;

  static equals(a: ListMessagesRequest | PlainMessage<ListMessagesRequest> | undefined, b: ListMessagesRequest | PlainMessage<ListMessagesRequest> | undefined): boolean;
}

/**
 * *
 * Response of a topic search/listing.
 *
 * @generated from message onehub.v1.ListMessagesResponse
 */
export declare class ListMessagesResponse extends Message<ListMessagesResponse> {
  /**
   * *
   * The list of topics found as part of this response.
   *
   * @generated from field: repeated onehub.v1.Message messages = 1;
   */
  messages: Message$1[];

  /**
   * *
   * Pagination response info
   *
   * @generated from field: onehub.v1.PaginationResponse pagination = 2;
   */
  pagination?: PaginationResponse;

  constructor(data?: PartialMessage<ListMessagesResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "onehub.v1.ListMessagesResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListMessagesResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListMessagesResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListMessagesResponse;

  static equals(a: ListMessagesResponse | PlainMessage<ListMessagesResponse> | undefined, b: ListMessagesResponse | PlainMessage<ListMessagesResponse> | undefined): boolean;
}

/**
 * *
 * Request to get a single message.
 *
 * @generated from message onehub.v1.GetMessageRequest
 */
export declare class GetMessageRequest extends Message<GetMessageRequest> {
  /**
   * *
   * ID of the topic to be fetched
   *
   * @generated from field: string id = 1;
   */
  id: string;

  constructor(data?: PartialMessage<GetMessageRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "onehub.v1.GetMessageRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetMessageRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetMessageRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetMessageRequest;

  static equals(a: GetMessageRequest | PlainMessage<GetMessageRequest> | undefined, b: GetMessageRequest | PlainMessage<GetMessageRequest> | undefined): boolean;
}

/**
 * *
 * Message get response
 *
 * @generated from message onehub.v1.GetMessageResponse
 */
export declare class GetMessageResponse extends Message<GetMessageResponse> {
  /**
   * @generated from field: onehub.v1.Message message = 1;
   */
  message?: Message$1;

  constructor(data?: PartialMessage<GetMessageResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "onehub.v1.GetMessageResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetMessageResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetMessageResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetMessageResponse;

  static equals(a: GetMessageResponse | PlainMessage<GetMessageResponse> | undefined, b: GetMessageResponse | PlainMessage<GetMessageResponse> | undefined): boolean;
}

/**
 * *
 * Request to batch get messages
 *
 * @generated from message onehub.v1.GetMessagesRequest
 */
export declare class GetMessagesRequest extends Message<GetMessagesRequest> {
  /**
   * *
   * IDs of the messages to be fetched
   *
   * @generated from field: repeated string ids = 1;
   */
  ids: string[];

  constructor(data?: PartialMessage<GetMessagesRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "onehub.v1.GetMessagesRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetMessagesRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetMessagesRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetMessagesRequest;

  static equals(a: GetMessagesRequest | PlainMessage<GetMessagesRequest> | undefined, b: GetMessagesRequest | PlainMessage<GetMessagesRequest> | undefined): boolean;
}

/**
 * *
 * Message batch-get response
 *
 * @generated from message onehub.v1.GetMessagesResponse
 */
export declare class GetMessagesResponse extends Message<GetMessagesResponse> {
  /**
   * @generated from field: map<string, onehub.v1.Message> messages = 1;
   */
  messages: { [key: string]: Message$1 };

  constructor(data?: PartialMessage<GetMessagesResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "onehub.v1.GetMessagesResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetMessagesResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetMessagesResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetMessagesResponse;

  static equals(a: GetMessagesResponse | PlainMessage<GetMessagesResponse> | undefined, b: GetMessagesResponse | PlainMessage<GetMessagesResponse> | undefined): boolean;
}

/**
 * *
 * Request to delete an message
 *
 * @generated from message onehub.v1.DeleteMessageRequest
 */
export declare class DeleteMessageRequest extends Message<DeleteMessageRequest> {
  /**
   * *
   * ID of the message to be deleted.
   *
   * @generated from field: string id = 1;
   */
  id: string;

  constructor(data?: PartialMessage<DeleteMessageRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "onehub.v1.DeleteMessageRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteMessageRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteMessageRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteMessageRequest;

  static equals(a: DeleteMessageRequest | PlainMessage<DeleteMessageRequest> | undefined, b: DeleteMessageRequest | PlainMessage<DeleteMessageRequest> | undefined): boolean;
}

/**
 * *
 * Message deletion response
 *
 * @generated from message onehub.v1.DeleteMessageResponse
 */
export declare class DeleteMessageResponse extends Message<DeleteMessageResponse> {
  constructor(data?: PartialMessage<DeleteMessageResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "onehub.v1.DeleteMessageResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteMessageResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteMessageResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteMessageResponse;

  static equals(a: DeleteMessageResponse | PlainMessage<DeleteMessageResponse> | undefined, b: DeleteMessageResponse | PlainMessage<DeleteMessageResponse> | undefined): boolean;
}

/**
 * @generated from message onehub.v1.UpdateMessageRequest
 */
export declare class UpdateMessageRequest extends Message<UpdateMessageRequest> {
  /**
   * The message being updated.  The topic ID AND message ID fields *must*
   * be specified in this message object.  How other fields are used is
   * determined by the update_mask parameter enabling partial updates
   *
   * @generated from field: onehub.v1.Message message = 1;
   */
  message?: Message$1;

  /**
   * Indicates which fields are being updated
   * If the field_mask is *not* provided then we reject
   * a replace (as required by the standard convention) to prevent
   * full replace in error.  Instead an update_mask of "*" must be passed.
   *
   * @generated from field: google.protobuf.FieldMask update_mask = 3;
   */
  updateMask?: FieldMask;

  /**
   * Any fields specified here will be "appended" to instead of being
   * replaced
   *
   * @generated from field: google.protobuf.FieldMask append_mask = 4;
   */
  appendMask?: FieldMask;

  constructor(data?: PartialMessage<UpdateMessageRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "onehub.v1.UpdateMessageRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateMessageRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateMessageRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateMessageRequest;

  static equals(a: UpdateMessageRequest | PlainMessage<UpdateMessageRequest> | undefined, b: UpdateMessageRequest | PlainMessage<UpdateMessageRequest> | undefined): boolean;
}

/**
 * @generated from message onehub.v1.UpdateMessageResponse
 */
export declare class UpdateMessageResponse extends Message<UpdateMessageResponse> {
  /**
   * The updated message
   *
   * @generated from field: onehub.v1.Message message = 1;
   */
  message?: Message$1;

  constructor(data?: PartialMessage<UpdateMessageResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "onehub.v1.UpdateMessageResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateMessageResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateMessageResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateMessageResponse;

  static equals(a: UpdateMessageResponse | PlainMessage<UpdateMessageResponse> | undefined, b: UpdateMessageResponse | PlainMessage<UpdateMessageResponse> | undefined): boolean;
}

