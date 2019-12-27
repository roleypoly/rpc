// package: roleypoly
// file: shared/shared.proto

import * as jspb from "google-protobuf";

export class IDQuery extends jspb.Message {
  getMemberid(): string;
  setMemberid(value: string): void;

  getGuildid(): string;
  setGuildid(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): IDQuery.AsObject;
  static toObject(includeInstance: boolean, msg: IDQuery): IDQuery.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: IDQuery, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): IDQuery;
  static deserializeBinaryFromReader(message: IDQuery, reader: jspb.BinaryReader): IDQuery;
}

export namespace IDQuery {
  export type AsObject = {
    memberid: string,
    guildid: string,
  }
}

export class GuildList extends jspb.Message {
  clearGuildsList(): void;
  getGuildsList(): Array<Guild>;
  setGuildsList(value: Array<Guild>): void;
  addGuilds(value?: Guild, index?: number): Guild;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GuildList.AsObject;
  static toObject(includeInstance: boolean, msg: GuildList): GuildList.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GuildList, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GuildList;
  static deserializeBinaryFromReader(message: GuildList, reader: jspb.BinaryReader): GuildList;
}

export namespace GuildList {
  export type AsObject = {
    guildsList: Array<Guild.AsObject>,
  }
}

export class Guild extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getName(): string;
  setName(value: string): void;

  getIcon(): string;
  setIcon(value: string): void;

  getOwnerid(): string;
  setOwnerid(value: string): void;

  getMembercount(): number;
  setMembercount(value: number): void;

  getSplash(): string;
  setSplash(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Guild.AsObject;
  static toObject(includeInstance: boolean, msg: Guild): Guild.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Guild, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Guild;
  static deserializeBinaryFromReader(message: Guild, reader: jspb.BinaryReader): Guild;
}

export namespace Guild {
  export type AsObject = {
    id: string,
    name: string,
    icon: string,
    ownerid: string,
    membercount: number,
    splash: string,
  }
}

export class GuildRoles extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  clearRolesList(): void;
  getRolesList(): Array<Role>;
  setRolesList(value: Array<Role>): void;
  addRoles(value?: Role, index?: number): Role;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GuildRoles.AsObject;
  static toObject(includeInstance: boolean, msg: GuildRoles): GuildRoles.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GuildRoles, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GuildRoles;
  static deserializeBinaryFromReader(message: GuildRoles, reader: jspb.BinaryReader): GuildRoles;
}

export namespace GuildRoles {
  export type AsObject = {
    id: string,
    rolesList: Array<Role.AsObject>,
  }
}

export class Role extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getName(): string;
  setName(value: string): void;

  getPermissions(): number;
  setPermissions(value: number): void;

  getColor(): number;
  setColor(value: number): void;

  getPosition(): number;
  setPosition(value: number): void;

  getManaged(): boolean;
  setManaged(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Role.AsObject;
  static toObject(includeInstance: boolean, msg: Role): Role.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Role, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Role;
  static deserializeBinaryFromReader(message: Role, reader: jspb.BinaryReader): Role;
}

export namespace Role {
  export type AsObject = {
    id: string,
    name: string,
    permissions: number,
    color: number,
    position: number,
    managed: boolean,
  }
}

export class DiscordUser extends jspb.Message {
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
  toObject(includeInstance?: boolean): DiscordUser.AsObject;
  static toObject(includeInstance: boolean, msg: DiscordUser): DiscordUser.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DiscordUser, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DiscordUser;
  static deserializeBinaryFromReader(message: DiscordUser, reader: jspb.BinaryReader): DiscordUser;
}

export namespace DiscordUser {
  export type AsObject = {
    id: string,
    username: string,
    discriminator: string,
    avatar: string,
    bot: boolean,
  }
}

