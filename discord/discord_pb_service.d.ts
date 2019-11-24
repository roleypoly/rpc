// package: roleypoly.discord
// file: discord/discord.proto

import * as discord_discord_pb from "../discord/discord_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";
import * as shared_shared_pb from "../shared/shared_pb";
import {grpc} from "@improbable-eng/grpc-web";

type DiscordListGuilds = {
  readonly methodName: string;
  readonly service: typeof Discord;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof google_protobuf_empty_pb.Empty;
  readonly responseType: typeof shared_shared_pb.GuildList;
};

type DiscordGetGuild = {
  readonly methodName: string;
  readonly service: typeof Discord;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof shared_shared_pb.IDQuery;
  readonly responseType: typeof shared_shared_pb.Guild;
};

type DiscordGetGuildRoles = {
  readonly methodName: string;
  readonly service: typeof Discord;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof shared_shared_pb.IDQuery;
  readonly responseType: typeof shared_shared_pb.GuildRoles;
};

type DiscordGetGuildsByMember = {
  readonly methodName: string;
  readonly service: typeof Discord;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof shared_shared_pb.IDQuery;
  readonly responseType: typeof shared_shared_pb.GuildList;
};

type DiscordGetMember = {
  readonly methodName: string;
  readonly service: typeof Discord;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof shared_shared_pb.IDQuery;
  readonly responseType: typeof discord_discord_pb.Member;
};

type DiscordGetUser = {
  readonly methodName: string;
  readonly service: typeof Discord;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof shared_shared_pb.IDQuery;
  readonly responseType: typeof discord_discord_pb.User;
};

type DiscordUpdateMember = {
  readonly methodName: string;
  readonly service: typeof Discord;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof discord_discord_pb.Member;
  readonly responseType: typeof discord_discord_pb.Member;
};

type DiscordOwnUser = {
  readonly methodName: string;
  readonly service: typeof Discord;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof google_protobuf_empty_pb.Empty;
  readonly responseType: typeof discord_discord_pb.User;
};

export class Discord {
  static readonly serviceName: string;
  static readonly ListGuilds: DiscordListGuilds;
  static readonly GetGuild: DiscordGetGuild;
  static readonly GetGuildRoles: DiscordGetGuildRoles;
  static readonly GetGuildsByMember: DiscordGetGuildsByMember;
  static readonly GetMember: DiscordGetMember;
  static readonly GetUser: DiscordGetUser;
  static readonly UpdateMember: DiscordUpdateMember;
  static readonly OwnUser: DiscordOwnUser;
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
  listGuilds(
    requestMessage: google_protobuf_empty_pb.Empty,
    metadata?: grpc.Metadata,
  ): Promise<shared_shared_pb.GuildList>;
  getGuild(
    requestMessage: shared_shared_pb.IDQuery,
    metadata?: grpc.Metadata,
  ): Promise<shared_shared_pb.Guild>;
  getGuildRoles(
    requestMessage: shared_shared_pb.IDQuery,
    metadata?: grpc.Metadata,
  ): Promise<shared_shared_pb.GuildRoles>;
  getGuildsByMember(
    requestMessage: shared_shared_pb.IDQuery,
    metadata?: grpc.Metadata,
  ): Promise<shared_shared_pb.GuildList>;
  getMember(
    requestMessage: shared_shared_pb.IDQuery,
    metadata?: grpc.Metadata,
  ): Promise<discord_discord_pb.Member>;
  getUser(
    requestMessage: shared_shared_pb.IDQuery,
    metadata?: grpc.Metadata,
  ): Promise<discord_discord_pb.User>;
  updateMember(
    requestMessage: discord_discord_pb.Member,
    metadata?: grpc.Metadata,
  ): Promise<discord_discord_pb.Member>;
  ownUser(
    requestMessage: google_protobuf_empty_pb.Empty,
    metadata?: grpc.Metadata,
  ): Promise<discord_discord_pb.User>;
}

