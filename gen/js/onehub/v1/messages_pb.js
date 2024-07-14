// @generated by protoc-gen-es v1.10.0
// @generated from file onehub/v1/messages.proto (package onehub.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { FieldMask, proto3 } from "@bufbuild/protobuf";
import { Message, Pagination, PaginationResponse } from "./models_pb.js";

/**
 * *
 * Message creation request object
 *
 * @generated from message onehub.v1.CreateMessagesRequest
 */
export const CreateMessagesRequest = /*@__PURE__*/ proto3.makeMessageType(
  "onehub.v1.CreateMessagesRequest",
  () => [
    { no: 1, name: "topic_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "messages", kind: "message", T: Message, repeated: true },
    { no: 3, name: "allow_userids", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ],
);

/**
 * *
 * Response of an message creation.
 *
 * @generated from message onehub.v1.CreateMessagesResponse
 */
export const CreateMessagesResponse = /*@__PURE__*/ proto3.makeMessageType(
  "onehub.v1.CreateMessagesResponse",
  () => [
    { no: 1, name: "messages", kind: "message", T: Message, repeated: true },
  ],
);

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
export const ImportMessagesRequest = /*@__PURE__*/ proto3.makeMessageType(
  "onehub.v1.ImportMessagesRequest",
  () => [
    { no: 2, name: "messages", kind: "message", T: Message, repeated: true },
  ],
);

/**
 * *
 * Response of an message import.
 *
 * @generated from message onehub.v1.ImportMessagesResponse
 */
export const ImportMessagesResponse = /*@__PURE__*/ proto3.makeMessageType(
  "onehub.v1.ImportMessagesResponse",
  () => [
    { no: 1, name: "messages", kind: "message", T: Message, repeated: true },
  ],
);

/**
 * *
 * A message listing request.  For now only paginations params are provided.
 *
 * @generated from message onehub.v1.ListMessagesRequest
 */
export const ListMessagesRequest = /*@__PURE__*/ proto3.makeMessageType(
  "onehub.v1.ListMessagesRequest",
  () => [
    { no: 1, name: "topic_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "pagination", kind: "message", T: Pagination },
  ],
);

/**
 * *
 * Response of a topic search/listing.
 *
 * @generated from message onehub.v1.ListMessagesResponse
 */
export const ListMessagesResponse = /*@__PURE__*/ proto3.makeMessageType(
  "onehub.v1.ListMessagesResponse",
  () => [
    { no: 1, name: "messages", kind: "message", T: Message, repeated: true },
    { no: 2, name: "pagination", kind: "message", T: PaginationResponse },
  ],
);

/**
 * *
 * Unified API for message search.  Searches across all messages with a bunch of filters.
 *
 * @generated from message onehub.v1.SearchMessagesRequest
 */
export const SearchMessagesRequest = /*@__PURE__*/ proto3.makeMessageType(
  "onehub.v1.SearchMessagesRequest",
  () => [
    { no: 1, name: "search_phrase", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "sender_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "topic_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "order_by", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ],
);

/**
 * *
 * Response of an message import.
 *
 * @generated from message onehub.v1.SearchMessagesResponse
 */
export const SearchMessagesResponse = /*@__PURE__*/ proto3.makeMessageType(
  "onehub.v1.SearchMessagesResponse",
  () => [
    { no: 1, name: "messages", kind: "message", T: Message, repeated: true },
    { no: 2, name: "pagination", kind: "message", T: PaginationResponse },
  ],
);

/**
 * *
 * Request to get a single message.
 *
 * @generated from message onehub.v1.GetMessageRequest
 */
export const GetMessageRequest = /*@__PURE__*/ proto3.makeMessageType(
  "onehub.v1.GetMessageRequest",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * *
 * Message get response
 *
 * @generated from message onehub.v1.GetMessageResponse
 */
export const GetMessageResponse = /*@__PURE__*/ proto3.makeMessageType(
  "onehub.v1.GetMessageResponse",
  () => [
    { no: 1, name: "message", kind: "message", T: Message },
  ],
);

/**
 * *
 * Request to batch get messages
 *
 * @generated from message onehub.v1.GetMessagesRequest
 */
export const GetMessagesRequest = /*@__PURE__*/ proto3.makeMessageType(
  "onehub.v1.GetMessagesRequest",
  () => [
    { no: 1, name: "ids", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ],
);

/**
 * *
 * Message batch-get response
 *
 * @generated from message onehub.v1.GetMessagesResponse
 */
export const GetMessagesResponse = /*@__PURE__*/ proto3.makeMessageType(
  "onehub.v1.GetMessagesResponse",
  () => [
    { no: 1, name: "messages", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "message", T: Message} },
  ],
);

/**
 * *
 * Request to delete an message
 *
 * @generated from message onehub.v1.DeleteMessageRequest
 */
export const DeleteMessageRequest = /*@__PURE__*/ proto3.makeMessageType(
  "onehub.v1.DeleteMessageRequest",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * *
 * Message deletion response
 *
 * @generated from message onehub.v1.DeleteMessageResponse
 */
export const DeleteMessageResponse = /*@__PURE__*/ proto3.makeMessageType(
  "onehub.v1.DeleteMessageResponse",
  [],
);

/**
 * @generated from message onehub.v1.UpdateMessageRequest
 */
export const UpdateMessageRequest = /*@__PURE__*/ proto3.makeMessageType(
  "onehub.v1.UpdateMessageRequest",
  () => [
    { no: 1, name: "message", kind: "message", T: Message },
    { no: 3, name: "update_mask", kind: "message", T: FieldMask },
    { no: 4, name: "append_mask", kind: "message", T: FieldMask },
  ],
);

/**
 * @generated from message onehub.v1.UpdateMessageResponse
 */
export const UpdateMessageResponse = /*@__PURE__*/ proto3.makeMessageType(
  "onehub.v1.UpdateMessageResponse",
  () => [
    { no: 1, name: "message", kind: "message", T: Message },
  ],
);

