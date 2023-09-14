// @generated by protoc-gen-es v1.3.1
// @generated from file onehub/v1/messages.proto (package onehub.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { FieldMask, proto3 } from "@bufbuild/protobuf";
import { Message } from "./models_pb.js";

/**
 * *
 * Message creation request object
 *
 * @generated from message onehub.v1.CreateMessagesRequest
 */
export const CreateMessagesRequest = proto3.makeMessageType(
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
export const CreateMessagesResponse = proto3.makeMessageType(
  "onehub.v1.CreateMessagesResponse",
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
export const ListMessagesRequest = proto3.makeMessageType(
  "onehub.v1.ListMessagesRequest",
  () => [
    { no: 1, name: "page_key", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "page_size", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "topic_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * *
 * Response of a topic search/listing.
 *
 * @generated from message onehub.v1.ListMessagesResponse
 */
export const ListMessagesResponse = proto3.makeMessageType(
  "onehub.v1.ListMessagesResponse",
  () => [
    { no: 1, name: "messages", kind: "message", T: Message, repeated: true },
    { no: 2, name: "next_page_key", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * *
 * Request to get a single message.
 *
 * @generated from message onehub.v1.GetMessageRequest
 */
export const GetMessageRequest = proto3.makeMessageType(
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
export const GetMessageResponse = proto3.makeMessageType(
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
export const GetMessagesRequest = proto3.makeMessageType(
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
export const GetMessagesResponse = proto3.makeMessageType(
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
export const DeleteMessageRequest = proto3.makeMessageType(
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
export const DeleteMessageResponse = proto3.makeMessageType(
  "onehub.v1.DeleteMessageResponse",
  [],
);

/**
 * @generated from message onehub.v1.UpdateMessageRequest
 */
export const UpdateMessageRequest = proto3.makeMessageType(
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
export const UpdateMessageResponse = proto3.makeMessageType(
  "onehub.v1.UpdateMessageResponse",
  () => [
    { no: 1, name: "message", kind: "message", T: Message },
  ],
);

