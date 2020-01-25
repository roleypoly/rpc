// package: roleypoly.ctf
// file: ctf/ctf.proto

import * as jspb from "google-protobuf";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";

export class Flags extends jspb.Message {
  clearFlagsList(): void;
  getFlagsList(): Array<Flag>;
  setFlagsList(value: Array<Flag>): void;
  addFlags(value?: Flag, index?: number): Flag;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Flags.AsObject;
  static toObject(includeInstance: boolean, msg: Flags): Flags.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Flags, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Flags;
  static deserializeBinaryFromReader(message: Flags, reader: jspb.BinaryReader): Flags;
}

export namespace Flags {
  export type AsObject = {
    flagsList: Array<Flag.AsObject>,
  }
}

export class Flag extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  getRing(): number;
  setRing(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Flag.AsObject;
  static toObject(includeInstance: boolean, msg: Flag): Flag.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Flag, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Flag;
  static deserializeBinaryFromReader(message: Flag, reader: jspb.BinaryReader): Flag;
}

export namespace Flag {
  export type AsObject = {
    name: string,
    ring: number,
  }
}

export class Ring extends jspb.Message {
  getRing(): number;
  setRing(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Ring.AsObject;
  static toObject(includeInstance: boolean, msg: Ring): Ring.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Ring, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Ring;
  static deserializeBinaryFromReader(message: Ring, reader: jspb.BinaryReader): Ring;
}

export namespace Ring {
  export type AsObject = {
    ring: number,
  }
}

