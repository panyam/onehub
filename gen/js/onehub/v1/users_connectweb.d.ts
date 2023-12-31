// @generated by protoc-gen-connect-web v0.8.6
// @generated from file onehub/v1/users.proto (package onehub.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreateUserRequest, CreateUserResponse, DeleteUserRequest, DeleteUserResponse, GetUserRequest, GetUserResponse, GetUsersRequest, GetUsersResponse, ListUsersRequest, ListUsersResponse, UpdateUserRequest, UpdateUserResponse } from "./users_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * *
 * Service for operating on users
 *
 * @generated from service onehub.v1.UserService
 */
export declare const UserService: {
  readonly typeName: "onehub.v1.UserService",
  readonly methods: {
    /**
     * *
     * Create a new sesssion
     *
     * @generated from rpc onehub.v1.UserService.CreateUser
     */
    readonly createUser: {
      readonly name: "CreateUser",
      readonly I: typeof CreateUserRequest,
      readonly O: typeof CreateUserResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * *
     * List all users from a user.
     *
     * @generated from rpc onehub.v1.UserService.ListUsers
     */
    readonly listUsers: {
      readonly name: "ListUsers",
      readonly I: typeof ListUsersRequest,
      readonly O: typeof ListUsersResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * *
     * Get a particular user
     *
     * @generated from rpc onehub.v1.UserService.GetUser
     */
    readonly getUser: {
      readonly name: "GetUser",
      readonly I: typeof GetUserRequest,
      readonly O: typeof GetUserResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * *
     * Batch get multiple users by ID
     *
     * @generated from rpc onehub.v1.UserService.GetUsers
     */
    readonly getUsers: {
      readonly name: "GetUsers",
      readonly I: typeof GetUsersRequest,
      readonly O: typeof GetUsersResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * *
     * Delete a particular user
     *
     * @generated from rpc onehub.v1.UserService.DeleteUser
     */
    readonly deleteUser: {
      readonly name: "DeleteUser",
      readonly I: typeof DeleteUserRequest,
      readonly O: typeof DeleteUserResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * *
     * Updates specific fields of a user
     *
     * @generated from rpc onehub.v1.UserService.UpdateUser
     */
    readonly updateUser: {
      readonly name: "UpdateUser",
      readonly I: typeof UpdateUserRequest,
      readonly O: typeof UpdateUserResponse,
      readonly kind: MethodKind.Unary,
    },
  }
};

