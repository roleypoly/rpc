// package: com.roleypoly.platform
// file: platform/platform.proto

import * as platform_platform_pb from "../platform/platform_pb";
import * as discord_discord_pb from "../discord/discord_pb";
import {grpc} from "@improbable-eng/grpc-web";

type PlatformGetGuild = {
  readonly methodName: string;
  readonly service: typeof Platform;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof discord_discord_pb.IDQuery;
  readonly responseType: typeof discord_discord_pb.Guild;
};

export class Platform {
  static readonly serviceName: string;
  static readonly GetGuild: PlatformGetGuild;
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
  getGuild(
    requestMessage: discord_discord_pb.IDQuery,
    metadata?: grpc.Metadata,
  ): Promise<discord_discord_pb.Guild>;
}

