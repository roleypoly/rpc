// package: roleypoly.platform
// file: platform/platform.proto

import * as jspb from "google-protobuf";
import * as shared_shared_pb from "../shared/shared_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";

export class GuildData extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getMessage(): string;
  setMessage(value: string): void;

  clearCategoriesList(): void;
  getCategoriesList(): Array<Category>;
  setCategoriesList(value: Array<Category>): void;
  addCategories(value?: Category, index?: number): Category;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GuildData.AsObject;
  static toObject(includeInstance: boolean, msg: GuildData): GuildData.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GuildData, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GuildData;
  static deserializeBinaryFromReader(message: GuildData, reader: jspb.BinaryReader): GuildData;
}

export namespace GuildData {
  export type AsObject = {
    id: string,
    message: string,
    categoriesList: Array<Category.AsObject>,
  }
}

export class Roles extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Roles.AsObject;
  static toObject(includeInstance: boolean, msg: Roles): Roles.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Roles, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Roles;
  static deserializeBinaryFromReader(message: Roles, reader: jspb.BinaryReader): Roles;
}

export namespace Roles {
  export type AsObject = {
  }
}

export class Category extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Category.AsObject;
  static toObject(includeInstance: boolean, msg: Category): Category.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Category, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Category;
  static deserializeBinaryFromReader(message: Category, reader: jspb.BinaryReader): Category;
}

export namespace Category {
  export type AsObject = {
  }
}

