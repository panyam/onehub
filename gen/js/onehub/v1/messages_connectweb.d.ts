// @generated by protoc-gen-connect-web v0.8.6
// @generated from file onehub/v1/messages.proto (package onehub.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreateMessageRequest, CreateMessageResponse, DeleteMessageRequest, DeleteMessageResponse, GetMessageRequest, GetMessageResponse, GetMessagesRequest, GetMessagesResponse, ListMessagesRequest, ListMessagesResponse, UpdateMessageRequest, UpdateMessageResponse } from "./messages_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * *
 * Service for operating on messages
 *
 * @generated from service onehub.v1.MessageService
 */
export declare const MessageService: {
  readonly typeName: "onehub.v1.MessageService",
  readonly methods: {
    /**
     * *
     * Create a single message or messages in batch
     *
     * @generated from rpc onehub.v1.MessageService.CreateMessage
     */
    readonly createMessage: {
      readonly name: "CreateMessage",
      readonly I: typeof CreateMessageRequest,
      readonly O: typeof CreateMessageResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * *
     * List all messages in a topic
     *
     * @generated from rpc onehub.v1.MessageService.ListMessages
     */
    readonly listMessages: {
      readonly name: "ListMessages",
      readonly I: typeof ListMessagesRequest,
      readonly O: typeof ListMessagesResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * *
     * Get a particular message
     *
     * @generated from rpc onehub.v1.MessageService.GetMessage
     */
    readonly getMessage: {
      readonly name: "GetMessage",
      readonly I: typeof GetMessageRequest,
      readonly O: typeof GetMessageResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * *
     * Batch get multiple messages by IDs
     *
     * @generated from rpc onehub.v1.MessageService.GetMessages
     */
    readonly getMessages: {
      readonly name: "GetMessages",
      readonly I: typeof GetMessagesRequest,
      readonly O: typeof GetMessagesResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * *
     * Delete a particular message
     *
     * @generated from rpc onehub.v1.MessageService.DeleteMessage
     */
    readonly deleteMessage: {
      readonly name: "DeleteMessage",
      readonly I: typeof DeleteMessageRequest,
      readonly O: typeof DeleteMessageResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * *
     * Update a message within a topic.
     *
     * @generated from rpc onehub.v1.MessageService.UpdateMessage
     */
    readonly updateMessage: {
      readonly name: "UpdateMessage",
      readonly I: typeof UpdateMessageRequest,
      readonly O: typeof UpdateMessageResponse,
      readonly kind: MethodKind.Unary,
    },
  }
};

