// @generated by protoc-gen-es v1.3.1
// @generated from file onehub/v1/messages.proto (package onehub.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, FieldMask, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { Message as Message$1 } from "./models_pb.js";

/**
 * *
 * Message creation request object
 *
 * @generated from message onehub.v1.CreateMessageRequest
 */
export declare class CreateMessageRequest extends Message<CreateMessageRequest> {
  /**
   * *
   * Message being updated
   *
   * @generated from field: repeated onehub.v1.Message message = 1;
   */
  message: Message$1[];

  constructor(data?: PartialMessage<CreateMessageRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "onehub.v1.CreateMessageRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateMessageRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateMessageRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateMessageRequest;

  static equals(a: CreateMessageRequest | PlainMessage<CreateMessageRequest> | undefined, b: CreateMessageRequest | PlainMessage<CreateMessageRequest> | undefined): boolean;
}

/**
 * *
 * Response of an message creation.
 *
 * @generated from message onehub.v1.CreateMessageResponse
 */
export declare class CreateMessageResponse extends Message<CreateMessageResponse> {
  /**
   * *
   * Message being created
   *
   * @generated from field: repeated onehub.v1.Message message = 1;
   */
  message: Message$1[];

  constructor(data?: PartialMessage<CreateMessageResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "onehub.v1.CreateMessageResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateMessageResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateMessageResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateMessageResponse;

  static equals(a: CreateMessageResponse | PlainMessage<CreateMessageResponse> | undefined, b: CreateMessageResponse | PlainMessage<CreateMessageResponse> | undefined): boolean;
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
   * Instead of an offset an abstract  "page" key is provided that offers
   * an opaque "pointer" into some offset in a result set.
   *
   * @generated from field: string page_key = 1;
   */
  pageKey: string;

  /**
   * *
   * Number of results to return.
   *
   * @generated from field: int32 page_size = 2;
   */
  pageSize: number;

  /**
   * *
   * Topic in which messages are to be listed.  Required.
   *
   * @generated from field: string topic_id = 3;
   */
  topicId: string;

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
   * The key/pointer string that subsequent List requests should pass to
   * continue the pagination.
   *
   * @generated from field: string next_page_key = 2;
   */
  nextPageKey: string;

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

