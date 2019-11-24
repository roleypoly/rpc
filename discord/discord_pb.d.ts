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
  getUser(): User | undefined;
  setUser(value?: User): void;

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
    user?: User.AsObject,
  }
}

export class User extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getUsername(): string;
  setUsername(value: string): void;

  getDiscriminator(): string;
  setDiscriminator(value: string): void;

  getAvatar(): string;
  setAvatar(value: string): void;

  getBot(): boolean;
  setBot(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): User.AsObject;
  static toObject(includeInstance: boolean, msg: User): User.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: User, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): User;
  static deserializeBinaryFromReader(message: User, reader: jspb.BinaryReader): User;
}

export namespace User {
  export type AsObject = {
    id: string,
    username: string,
    discriminator: string,
    avatar: string,
    bot: boolean,
  }
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

