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

type DiscordGetGuild = {
  readonly methodName: string;
  readonly service: typeof Discord;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof discord_pb.IDQuery;
  readonly responseType: typeof discord_pb.Guild;
};

type DiscordGetMember = {
  readonly methodName: string;
  readonly service: typeof Discord;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof discord_pb.IDQuery;
  readonly responseType: typeof discord_pb.Member;
};

type DiscordUpdateMember = {
  readonly methodName: string;
  readonly service: typeof Discord;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof discord_pb.Member;
  readonly responseType: typeof discord_pb.Member;
};

export class Discord {
  static readonly serviceName: string;
  static readonly ListServers: DiscordListServers;
  static readonly GetGuild: DiscordGetGuild;
  static readonly GetMember: DiscordGetMember;
  static readonly UpdateMember: DiscordUpdateMember;
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
    metadata: grpc.Metadata,
  ): Promise<UnaryResponse | null>;
  listServers(
    requestMessage: discord_pb.Empty,
  ): Promise<UnaryResponse | null>;
  getGuild(
    requestMessage: discord_pb.IDQuery,
    metadata: grpc.Metadata,
  ): Promise<UnaryResponse | null>;
  getGuild(
    requestMessage: discord_pb.IDQuery,
  ): Promise<UnaryResponse | null>;
  getMember(
    requestMessage: discord_pb.IDQuery,
    metadata: grpc.Metadata,
  ): Promise<UnaryResponse | null>;
  getMember(
    requestMessage: discord_pb.IDQuery,
  ): Promise<UnaryResponse | null>;
  updateMember(
    requestMessage: discord_pb.Member,
    metadata: grpc.Metadata,
  ): Promise<UnaryResponse | null>;
  updateMember(
    requestMessage: discord_pb.Member,
  ): Promise<UnaryResponse | null>;
}

