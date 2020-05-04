// package: roleypoly.discord
// file: discord/discord.proto

import * as jspb from "google-protobuf";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";
import * as shared_shared_pb from "../shared/shared_pb";

export class GuildMembers extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  clearMembersList(): void;
  getMembersList(): Array<Member>;
  setMembersList(value: Array<Member>): void;
  addMembers(value?: Member, index?: number): Member;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GuildMembers.AsObject;
  static toObject(includeInstance: boolean, msg: GuildMembers): GuildMembers.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GuildMembers, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GuildMembers;
  static deserializeBinaryFromReader(message: GuildMembers, reader: jspb.BinaryReader): GuildMembers;
}

export namespace GuildMembers {
  export type AsObject = {
    id: string,
    membersList: Array<Member.AsObject>,
  }
}

export class Member extends jspb.Message {
  getGuildid(): string;
  setGuildid(value: string): void;

  clearRolesList(): void;
  getRolesList(): Array<string>;
  setRolesList(value: Array<string>): void;
  addRoles(value: string, index?: number): string;

  getNick(): string;
  setNick(value: string): void;

  hasUser(): boolean;
  clearUser(): void;
  getUser(): shared_shared_pb.DiscordUser | undefined;
  setUser(value?: shared_shared_pb.DiscordUser): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Member.AsObject;
  static toObject(includeInstance: boolean, msg: Member): Member.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Member, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Member;
  static deserializeBinaryFromReader(message: Member, reader: jspb.BinaryReader): Member;
}

export namespace Member {
  export type AsObject = {
    guildid: string,
    rolesList: Array<string>,
    nick: string,
    user?: shared_shared_pb.DiscordUser.AsObject,
  }
}

export class RoleTransaction extends jspb.Message {
  hasMember(): boolean;
  clearMember(): void;
  getMember(): shared_shared_pb.IDQuery | undefined;
  setMember(value?: shared_shared_pb.IDQuery): void;

  clearDeltaList(): void;
  getDeltaList(): Array<TxDelta>;
  setDeltaList(value: Array<TxDelta>): void;
  addDelta(value?: TxDelta, index?: number): TxDelta;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RoleTransaction.AsObject;
  static toObject(includeInstance: boolean, msg: RoleTransaction): RoleTransaction.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: RoleTransaction, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RoleTransaction;
  static deserializeBinaryFromReader(message: RoleTransaction, reader: jspb.BinaryReader): RoleTransaction;
}

export namespace RoleTransaction {
  export type AsObject = {
    member?: shared_shared_pb.IDQuery.AsObject,
    deltaList: Array<TxDelta.AsObject>,
  }
}

export class TxDelta extends jspb.Message {
  getRole(): string;
  setRole(value: string): void;

  getAction(): TxDelta.ActionMap[keyof TxDelta.ActionMap];
  setAction(value: TxDelta.ActionMap[keyof TxDelta.ActionMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TxDelta.AsObject;
  static toObject(includeInstance: boolean, msg: TxDelta): TxDelta.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: TxDelta, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TxDelta;
  static deserializeBinaryFromReader(message: TxDelta, reader: jspb.BinaryReader): TxDelta;
}

export namespace TxDelta {
  export type AsObject = {
    role: string,
    action: TxDelta.ActionMap[keyof TxDelta.ActionMap],
  }

  export interface ActionMap {
    UNKNOWN: 0;
    ADD: 1;
    REMOVE: 2;
  }

  export const Action: ActionMap;
}

export class RoleTransactionResult extends jspb.Message {
  hasMember(): boolean;
  clearMember(): void;
  getMember(): Member | undefined;
  setMember(value?: Member): void;

  getStatus(): RoleTransactionResult.StatusMap[keyof RoleTransactionResult.StatusMap];
  setStatus(value: RoleTransactionResult.StatusMap[keyof RoleTransactionResult.StatusMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RoleTransactionResult.AsObject;
  static toObject(includeInstance: boolean, msg: RoleTransactionResult): RoleTransactionResult.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: RoleTransactionResult, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RoleTransactionResult;
  static deserializeBinaryFromReader(message: RoleTransactionResult, reader: jspb.BinaryReader): RoleTransactionResult;
}

export namespace RoleTransactionResult {
  export type AsObject = {
    member?: Member.AsObject,
    status: RoleTransactionResult.StatusMap[keyof RoleTransactionResult.StatusMap],
  }

  export interface StatusMap {
    DONE: 0;
    QUEUED: 1;
    FAILED: 2;
  }

  export const Status: StatusMap;
}

export class ShardInfo extends jspb.Message {
  getShards(): number;
  setShards(value: number): void;

  getServers(): number;
  setServers(value: number): void;

  getUsers(): number;
  setUsers(value: number): void;

  getRoles(): number;
  setRoles(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ShardInfo.AsObject;
  static toObject(includeInstance: boolean, msg: ShardInfo): ShardInfo.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ShardInfo, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ShardInfo;
  static deserializeBinaryFromReader(message: ShardInfo, reader: jspb.BinaryReader): ShardInfo;
}

export namespace ShardInfo {
  export type AsObject = {
    shards: number,
    servers: number,
    users: number,
    roles: number,
  }
}

