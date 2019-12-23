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

  clearEntitlementsList(): void;
  getEntitlementsList(): Array<string>;
  setEntitlementsList(value: Array<string>): void;
  addEntitlements(value: string, index?: number): string;

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
    entitlementsList: Array<string>,
  }
}

export class Roles extends jspb.Message {
  clearRolesList(): void;
  getRolesList(): Array<string>;
  setRolesList(value: Array<string>): void;
  addRoles(value: string, index?: number): string;

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
    rolesList: Array<string>,
  }
}

export class Category extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getName(): string;
  setName(value: string): void;

  clearRolesList(): void;
  getRolesList(): Array<string>;
  setRolesList(value: Array<string>): void;
  addRoles(value: string, index?: number): string;

  getHidden(): boolean;
  setHidden(value: boolean): void;

  getType(): Category.CategoryTypeMap[keyof Category.CategoryTypeMap];
  setType(value: Category.CategoryTypeMap[keyof Category.CategoryTypeMap]): void;

  getPosition(): number;
  setPosition(value: number): void;

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
    id: string,
    name: string,
    rolesList: Array<string>,
    hidden: boolean,
    type: Category.CategoryTypeMap[keyof Category.CategoryTypeMap],
    position: number,
  }

  export interface CategoryTypeMap {
    MULTI: 0;
    SINGLE: 1;
  }

  export const CategoryType: CategoryTypeMap;
}

export class UpdateEntitlement extends jspb.Message {
  hasQuery(): boolean;
  clearQuery(): void;
  getQuery(): shared_shared_pb.IDQuery | undefined;
  setQuery(value?: shared_shared_pb.IDQuery): void;

  getName(): string;
  setName(value: string): void;

  getState(): boolean;
  setState(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateEntitlement.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateEntitlement): UpdateEntitlement.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateEntitlement, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateEntitlement;
  static deserializeBinaryFromReader(message: UpdateEntitlement, reader: jspb.BinaryReader): UpdateEntitlement;
}

export namespace UpdateEntitlement {
  export type AsObject = {
    query?: shared_shared_pb.IDQuery.AsObject,
    name: string,
    state: boolean,
  }
}

