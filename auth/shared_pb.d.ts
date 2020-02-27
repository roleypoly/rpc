// package: roleypoly.auth
// file: auth/shared.proto

import * as jspb from "google-protobuf";

export class Token extends jspb.Message {
  getToken(): string;
  setToken(value: string): void;

  getType(): Token.TypeMap[keyof Token.TypeMap];
  setType(value: Token.TypeMap[keyof Token.TypeMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Token.AsObject;
  static toObject(includeInstance: boolean, msg: Token): Token.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Token, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Token;
  static deserializeBinaryFromReader(message: Token, reader: jspb.BinaryReader): Token;
}

export namespace Token {
  export type AsObject = {
    token: string,
    type: Token.TypeMap[keyof Token.TypeMap],
  }

  export interface TypeMap {
    UNKNOWN: 0;
    SESSIONKEY: 1;
    CLIENTTOKEN: 2;
  }

  export const Type: TypeMap;
}

export class AuthChallenge extends jspb.Message {
  getUserid(): string;
  setUserid(value: string): void;

  getMagicurl(): string;
  setMagicurl(value: string): void;

  getMagicwords(): string;
  setMagicwords(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AuthChallenge.AsObject;
  static toObject(includeInstance: boolean, msg: AuthChallenge): AuthChallenge.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: AuthChallenge, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AuthChallenge;
  static deserializeBinaryFromReader(message: AuthChallenge, reader: jspb.BinaryReader): AuthChallenge;
}

export namespace AuthChallenge {
  export type AsObject = {
    userid: string,
    magicurl: string,
    magicwords: string,
  }
}

