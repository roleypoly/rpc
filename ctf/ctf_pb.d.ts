// package: roleypoly.ctf
// file: ctf/ctf.proto

import * as jspb from "google-protobuf";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";

export class Canary extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  getPercent(): number;
  setPercent(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Canary.AsObject;
  static toObject(includeInstance: boolean, msg: Canary): Canary.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Canary, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Canary;
  static deserializeBinaryFromReader(message: Canary, reader: jspb.BinaryReader): Canary;
}

export namespace Canary {
  export type AsObject = {
    name: string,
    percent: number,
  }
}

export class Canaries extends jspb.Message {
  clearCanariesList(): void;
  getCanariesList(): Array<Canary>;
  setCanariesList(value: Array<Canary>): void;
  addCanaries(value?: Canary, index?: number): Canary;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Canaries.AsObject;
  static toObject(includeInstance: boolean, msg: Canaries): Canaries.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Canaries, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Canaries;
  static deserializeBinaryFromReader(message: Canaries, reader: jspb.BinaryReader): Canaries;
}

export namespace Canaries {
  export type AsObject = {
    canariesList: Array<Canary.AsObject>,
  }
}

export class CanaryQuery extends jspb.Message {
  getThreshold(): number;
  setThreshold(value: number): void;

  hasCanary(): boolean;
  clearCanary(): void;
  getCanary(): Canary | undefined;
  setCanary(value?: Canary): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CanaryQuery.AsObject;
  static toObject(includeInstance: boolean, msg: CanaryQuery): CanaryQuery.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CanaryQuery, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CanaryQuery;
  static deserializeBinaryFromReader(message: CanaryQuery, reader: jspb.BinaryReader): CanaryQuery;
}

export namespace CanaryQuery {
  export type AsObject = {
    threshold: number,
    canary?: Canary.AsObject,
  }
}

