// @generated by protoc-gen-es v1.10.0
// @generated from file onehub/v1/models.proto (package onehub.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3, Struct, Timestamp } from "@bufbuild/protobuf";

/**
 * @generated from message onehub.v1.User
 */
export const User = /*@__PURE__*/ proto3.makeMessageType(
  "onehub.v1.User",
  () => [
    { no: 1, name: "created_at", kind: "message", T: Timestamp },
    { no: 2, name: "updated_at", kind: "message", T: Timestamp },
    { no: 3, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "avatar", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "profile_data", kind: "message", T: Struct, opt: true },
  ],
);

/**
 * @generated from message onehub.v1.MessageBase
 */
export const MessageBase = /*@__PURE__*/ proto3.makeMessageType(
  "onehub.v1.MessageBase",
  () => [
    { no: 1, name: "created_at", kind: "message", T: Timestamp },
    { no: 2, name: "updated_at", kind: "message", T: Timestamp },
    { no: 3, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "creator_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * Artists perform/play/sing songs
 *
 * @generated from message onehub.v1.Topic
 */
export const Topic = /*@__PURE__*/ proto3.makeMessageType(
  "onehub.v1.Topic",
  () => [
    { no: 1, name: "base", kind: "message", T: MessageBase },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "users", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "scalar", T: 8 /* ScalarType.BOOL */} },
  ],
);

/**
 * *
 * Base message type of entities that have custom "content" in them.
 *
 * @generated from message onehub.v1.ContentBase
 */
export const ContentBase = /*@__PURE__*/ proto3.makeMessageType(
  "onehub.v1.ContentBase",
  () => [
    { no: 1, name: "content_type", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "content_text", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "content_data", kind: "message", T: Struct },
  ],
);

/**
 * *
 * An individual message in a topic
 *
 * @generated from message onehub.v1.Message
 */
export const Message = /*@__PURE__*/ proto3.makeMessageType(
  "onehub.v1.Message",
  () => [
    { no: 1, name: "base", kind: "message", T: MessageBase },
    { no: 2, name: "content_base", kind: "message", T: ContentBase },
    { no: 3, name: "topic_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "parent_message_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * *
 * Nodes are blocks inside a graph.
 *
 * @generated from message onehub.v1.Node
 */
export const Node = /*@__PURE__*/ proto3.makeMessageType(
  "onehub.v1.Node",
  () => [
    { no: 1, name: "base", kind: "message", T: MessageBase },
    { no: 2, name: "content_base", kind: "message", T: ContentBase },
    { no: 3, name: "topic_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * *
 * Edges between two nodes in a graph.
 *
 * @generated from message onehub.v1.Edge
 */
export const Edge = /*@__PURE__*/ proto3.makeMessageType(
  "onehub.v1.Edge",
  () => [
    { no: 1, name: "base", kind: "message", T: MessageBase },
    { no: 2, name: "content_base", kind: "message", T: ContentBase },
    { no: 3, name: "source_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "target_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "undirected", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * *
 * General way to handle pagination in all listing resources.
 *
 * @generated from message onehub.v1.Pagination
 */
export const Pagination = /*@__PURE__*/ proto3.makeMessageType(
  "onehub.v1.Pagination",
  () => [
    { no: 1, name: "page_key", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "page_size", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ],
);

/**
 * *
 * Standard way to pass pagination related responses, eg the next page key
 * that can be passed on a paginated request to get the "next page" of results.
 *
 * @generated from message onehub.v1.PaginationResponse
 */
export const PaginationResponse = /*@__PURE__*/ proto3.makeMessageType(
  "onehub.v1.PaginationResponse",
  () => [
    { no: 1, name: "next_page_key", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "has_more_results", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 3, name: "total_num_results", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ],
);

