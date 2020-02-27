// package: roleypoly.auth.client
// file: auth/client/auth-client.proto

import * as auth_client_auth_client_pb from "../../auth/client/auth-client_pb";
import * as shared_internal_pb from "../../shared/internal_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";
import * as auth_shared_pb from "../../auth/shared_pb";
import {grpc} from "@improbable-eng/grpc-web";

type AuthClientGetClientToken = {
  readonly methodName: string;
  readonly service: typeof AuthClient;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof google_protobuf_empty_pb.Empty;
  readonly responseType: typeof auth_shared_pb.Token;
};

type AuthClientGetUserSession = {
  readonly methodName: string;
  readonly service: typeof AuthClient;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof google_protobuf_empty_pb.Empty;
  readonly responseType: typeof shared_internal_pb.RoleypolySession;
};

type AuthClientResolveSessionKey = {
  readonly methodName: string;
  readonly service: typeof AuthClient;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof auth_shared_pb.Token;
  readonly responseType: typeof auth_shared_pb.Token;
};

type AuthClientAuthorizeChallenge = {
  readonly methodName: string;
  readonly service: typeof AuthClient;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof auth_shared_pb.AuthChallenge;
  readonly responseType: typeof auth_shared_pb.Token;
};

export class AuthClient {
  static readonly serviceName: string;
  static readonly GetClientToken: AuthClientGetClientToken;
  static readonly GetUserSession: AuthClientGetUserSession;
  static readonly ResolveSessionKey: AuthClientResolveSessionKey;
  static readonly AuthorizeChallenge: AuthClientAuthorizeChallenge;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }

interface UnaryResponse {
  cancel(): void;
}
interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: (status?: Status) => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}
interface RequestStream<T> {
  write(message: T): RequestStream<T>;
  end(): void;
  cancel(): void;
  on(type: 'end', handler: (status?: Status) => void): RequestStream<T>;
  on(type: 'status', handler: (status: Status) => void): RequestStream<T>;
}
interface BidirectionalStream<ReqT, ResT> {
  write(message: ReqT): BidirectionalStream<ReqT, ResT>;
  end(): void;
  cancel(): void;
  on(type: 'data', handler: (message: ResT) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'end', handler: (status?: Status) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'status', handler: (status: Status) => void): BidirectionalStream<ReqT, ResT>;
}

export class AuthClientClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  getClientToken(
    requestMessage: google_protobuf_empty_pb.Empty,
    metadata?: grpc.Metadata,
  ): Promise<auth_shared_pb.Token>;
  getUserSession(
    requestMessage: google_protobuf_empty_pb.Empty,
    metadata?: grpc.Metadata,
  ): Promise<shared_internal_pb.RoleypolySession>;
  resolveSessionKey(
    requestMessage: auth_shared_pb.Token,
    metadata?: grpc.Metadata,
  ): Promise<auth_shared_pb.Token>;
  authorizeChallenge(
    requestMessage: auth_shared_pb.AuthChallenge,
    metadata?: grpc.Metadata,
  ): Promise<auth_shared_pb.Token>;
}

