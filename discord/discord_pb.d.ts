// package: 
// file: discord.proto

import * as jspb from "google-protobuf";

export class Empty extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Empty.AsObject;
  static toObject(includeInstance: boolean, msg: Empty): Empty.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Empty, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Empty;
  static deserializeBinaryFromReader(message: Empty, reader: jspb.BinaryReader): Empty;
}

export namespace Empty {
  export type AsObject = {
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

  getBot(): string;
  setBot(value: string): void;

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
    bot: string,
  }
}

