// @generated by protoc-gen-es v1.10.0
// @generated from file onehub/v1/search.proto (package onehub.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message onehub.v1.SearchTopicsRequest
 */
export declare class SearchTopicsRequest extends Message<SearchTopicsRequest> {
  constructor(data?: PartialMessage<SearchTopicsRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "onehub.v1.SearchTopicsRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SearchTopicsRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SearchTopicsRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SearchTopicsRequest;

  static equals(a: SearchTopicsRequest | PlainMessage<SearchTopicsRequest> | undefined, b: SearchTopicsRequest | PlainMessage<SearchTopicsRequest> | undefined): boolean;
}

