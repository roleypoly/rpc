// package: roleypoly.platform
// file: platform/platform.proto

import * as platform_platform_pb from "../platform/platform_pb";
import * as shared_shared_pb from "../shared/shared_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";
import {grpc} from "@improbable-eng/grpc-web";

type PlatformEnumerateMyGuilds = {
  readonly methodName: string;
  readonly service: typeof Platform;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof google_protobuf_empty_pb.Empty;
  readonly responseType: typeof platform_platform_pb.GuildEnumeration;
};

type PlatformGetGuildSlug = {
  readonly methodName: string;
  readonly service: typeof Platform;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof shared_shared_pb.IDQuery;
  readonly responseType: typeof shared_shared_pb.Guild;
};

type PlatformGetGuild = {
  readonly methodName: string;
  readonly service: typeof Platform;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof shared_shared_pb.IDQuery;
  readonly responseType: typeof platform_platform_pb.PresentableGuild;
};

type PlatformUpdateMyRoles = {
  readonly methodName: string;
  readonly service: typeof Platform;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof platform_platform_pb.UpdateRoles;
  readonly responseType: typeof google_protobuf_empty_pb.Empty;
};

type PlatformUpdateGuildData = {
  readonly methodName: string;
  readonly service: typeof Platform;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof platform_platform_pb.GuildData;
  readonly responseType: typeof google_protobuf_empty_pb.Empty;
};

export class Platform {
  static readonly serviceName: string;
  static readonly EnumerateMyGuilds: PlatformEnumerateMyGuilds;
  static readonly GetGuildSlug: PlatformGetGuildSlug;
  static readonly GetGuild: PlatformGetGuild;
  static readonly UpdateMyRoles: PlatformUpdateMyRoles;
  static readonly UpdateGuildData: PlatformUpdateGuildData;
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

export class PlatformClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  enumerateMyGuilds(
    requestMessage: google_protobuf_empty_pb.Empty,
    metadata?: grpc.Metadata,
  ): Promise<platform_platform_pb.GuildEnumeration>;
  getGuildSlug(
    requestMessage: shared_shared_pb.IDQuery,
    metadata?: grpc.Metadata,
  ): Promise<shared_shared_pb.Guild>;
  getGuild(
    requestMessage: shared_shared_pb.IDQuery,
    metadata?: grpc.Metadata,
  ): Promise<platform_platform_pb.PresentableGuild>;
  updateMyRoles(
    requestMessage: platform_platform_pb.UpdateRoles,
    metadata?: grpc.Metadata,
  ): Promise<google_protobuf_empty_pb.Empty>;
  updateGuildData(
    requestMessage: platform_platform_pb.GuildData,
    metadata?: grpc.Metadata,
  ): Promise<google_protobuf_empty_pb.Empty>;
}

