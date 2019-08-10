// package: 
// file: discord.proto

import * as discord_pb from "./discord_pb";
import {grpc} from "@improbable-eng/grpc-web";

type DiscordListServers = {
  readonly methodName: string;
  readonly service: typeof Discord;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof discord_pb.Empty;
  readonly responseType: typeof discord_pb.GuildList;
};

export class Discord {
  static readonly serviceName: string;
  static readonly ListServers: DiscordListServers;
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

export class DiscordClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  listServers(
    requestMessage: discord_pb.Empty,
    metadata?: grpc.Metadata,
  ): Promise<discord_pb.GuildList>;
}

