// @generated by protoc-gen-es v1.10.0
// @generated from file onehub/v1/topics.proto (package onehub.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { FieldMask, proto3 } from "@bufbuild/protobuf";
import { Pagination, PaginationResponse, Topic } from "./models_pb.js";

/**
 * *
 * Topic creation request object
 *
 * @generated from message onehub.v1.CreateTopicRequest
 */
export const CreateTopicRequest = /*@__PURE__*/ proto3.makeMessageType(
  "onehub.v1.CreateTopicRequest",
  () => [
    { no: 1, name: "topic", kind: "message", T: Topic },
  ],
);

/**
 * *
 * Response of an topic creation.
 *
 * @generated from message onehub.v1.CreateTopicResponse
 */
export const CreateTopicResponse = /*@__PURE__*/ proto3.makeMessageType(
  "onehub.v1.CreateTopicResponse",
  () => [
    { no: 1, name: "topic", kind: "message", T: Topic },
  ],
);

/**
 * *
 * An topic search request.  For now only paginations params are provided.
 *
 * @generated from message onehub.v1.ListTopicsRequest
 */
export const ListTopicsRequest = /*@__PURE__*/ proto3.makeMessageType(
  "onehub.v1.ListTopicsRequest",
  () => [
    { no: 1, name: "pagination", kind: "message", T: Pagination },
  ],
);

/**
 * *
 * Response of a topic search/listing.
 *
 * @generated from message onehub.v1.ListTopicsResponse
 */
export const ListTopicsResponse = /*@__PURE__*/ proto3.makeMessageType(
  "onehub.v1.ListTopicsResponse",
  () => [
    { no: 1, name: "topics", kind: "message", T: Topic, repeated: true },
    { no: 2, name: "pagination", kind: "message", T: PaginationResponse },
  ],
);

/**
 * *
 * Request to get an topic.
 *
 * @generated from message onehub.v1.GetTopicRequest
 */
export const GetTopicRequest = /*@__PURE__*/ proto3.makeMessageType(
  "onehub.v1.GetTopicRequest",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * *
 * Topic get response
 *
 * @generated from message onehub.v1.GetTopicResponse
 */
export const GetTopicResponse = /*@__PURE__*/ proto3.makeMessageType(
  "onehub.v1.GetTopicResponse",
  () => [
    { no: 1, name: "topic", kind: "message", T: Topic },
  ],
);

/**
 * *
 * Request to batch get topics
 *
 * @generated from message onehub.v1.GetTopicsRequest
 */
export const GetTopicsRequest = /*@__PURE__*/ proto3.makeMessageType(
  "onehub.v1.GetTopicsRequest",
  () => [
    { no: 1, name: "ids", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ],
);

/**
 * *
 * Topic batch-get response
 *
 * @generated from message onehub.v1.GetTopicsResponse
 */
export const GetTopicsResponse = /*@__PURE__*/ proto3.makeMessageType(
  "onehub.v1.GetTopicsResponse",
  () => [
    { no: 1, name: "topics", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "message", T: Topic} },
  ],
);

/**
 * *
 * Request to delete an topic.
 *
 * @generated from message onehub.v1.DeleteTopicRequest
 */
export const DeleteTopicRequest = /*@__PURE__*/ proto3.makeMessageType(
  "onehub.v1.DeleteTopicRequest",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * *
 * Topic deletion response
 *
 * @generated from message onehub.v1.DeleteTopicResponse
 */
export const DeleteTopicResponse = /*@__PURE__*/ proto3.makeMessageType(
  "onehub.v1.DeleteTopicResponse",
  [],
);

/**
 * *
 * The request for (partially) updating an Topic.
 *
 * @generated from message onehub.v1.UpdateTopicRequest
 */
export const UpdateTopicRequest = /*@__PURE__*/ proto3.makeMessageType(
  "onehub.v1.UpdateTopicRequest",
  () => [
    { no: 1, name: "topic", kind: "message", T: Topic },
    { no: 2, name: "update_mask", kind: "message", T: FieldMask },
    { no: 3, name: "add_users", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 4, name: "remove_users", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ],
);

/**
 * *
 * The request for (partially) updating an Topic.
 *
 * @generated from message onehub.v1.UpdateTopicResponse
 */
export const UpdateTopicResponse = /*@__PURE__*/ proto3.makeMessageType(
  "onehub.v1.UpdateTopicResponse",
  () => [
    { no: 1, name: "topic", kind: "message", T: Topic },
  ],
);

