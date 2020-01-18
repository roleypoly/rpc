// package: roleypoly.auth.client
// file: auth/client/auth-client.proto

import * as jspb from "google-protobuf";
import * as shared_internal_pb from "../../shared/internal_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";

export class ClientToken extends jspb.Message {
  getToken(): string;
  setToken(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ClientToken.AsObject;
  static toObject(includeInstance: boolean, msg: ClientToken): ClientToken.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ClientToken, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ClientToken;
  static deserializeBinaryFromReader(message: ClientToken, reader: jspb.BinaryReader): ClientToken;
}

export namespace ClientToken {
  export type AsObject = {
    token: string,
  }
}

