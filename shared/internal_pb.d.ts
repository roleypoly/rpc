// package: roleypoly
// file: shared/internal.proto

import * as jspb from "google-protobuf";
import * as shared_shared_pb from "../shared/shared_pb";

export class RoleypolyUser extends jspb.Message {
  hasDiscorduser(): boolean;
  clearDiscorduser(): void;
  getDiscorduser(): shared_shared_pb.DiscordUser | undefined;
  setDiscorduser(value?: shared_shared_pb.DiscordUser): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RoleypolyUser.AsObject;
  static toObject(includeInstance: boolean, msg: RoleypolyUser): RoleypolyUser.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: RoleypolyUser, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RoleypolyUser;
  static deserializeBinaryFromReader(message: RoleypolyUser, reader: jspb.BinaryReader): RoleypolyUser;
}

export namespace RoleypolyUser {
  export type AsObject = {
    discorduser?: shared_shared_pb.DiscordUser.AsObject,
  }
}

export class RoleypolySession extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  hasUser(): boolean;
  clearUser(): void;
  getUser(): RoleypolyUser | undefined;
  setUser(value?: RoleypolyUser): void;

  getSource(): RoleypolySession.SessionSourceMap[keyof RoleypolySession.SessionSourceMap];
  setSource(value: RoleypolySession.SessionSourceMap[keyof RoleypolySession.SessionSourceMap]): void;

  getCreatedAt(): number;
  setCreatedAt(value: number): void;

  getExpiresIn(): number;
  setExpiresIn(value: number): void;

  getExtraMap(): jspb.Map<string, string>;
  clearExtraMap(): void;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RoleypolySession.AsObject;
  static toObject(includeInstance: boolean, msg: RoleypolySession): RoleypolySession.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: RoleypolySession, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RoleypolySession;
  static deserializeBinaryFromReader(message: RoleypolySession, reader: jspb.BinaryReader): RoleypolySession;
}

export namespace RoleypolySession {
  export type AsObject = {
    id: string,
    user?: RoleypolyUser.AsObject,
    source: RoleypolySession.SessionSourceMap[keyof RoleypolySession.SessionSourceMap],
    createdAt: number,
    expiresIn: number,
    extraMap: Array<[string, string]>,
  }

  export interface SessionSourceMap {
    OAUTH: 0;
    DM: 1;
  }

  export const SessionSource: SessionSourceMap;
}

