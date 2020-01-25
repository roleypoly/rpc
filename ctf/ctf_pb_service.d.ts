// package: roleypoly.ctf
// file: ctf/ctf.proto

import * as ctf_ctf_pb from "../ctf/ctf_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";
import {grpc} from "@improbable-eng/grpc-web";

type CTFGetRingFlags = {
  readonly methodName: string;
  readonly service: typeof CTF;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof ctf_ctf_pb.Ring;
  readonly responseType: typeof ctf_ctf_pb.Flags;
};

type CTFCreateFlag = {
  readonly methodName: string;
  readonly service: typeof CTF;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof ctf_ctf_pb.Flag;
  readonly responseType: typeof ctf_ctf_pb.Flag;
};

type CTFPromoteFlag = {
  readonly methodName: string;
  readonly service: typeof CTF;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof ctf_ctf_pb.Flag;
  readonly responseType: typeof ctf_ctf_pb.Flag;
};

type CTFRemoveFlag = {
  readonly methodName: string;
  readonly service: typeof CTF;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof ctf_ctf_pb.Flag;
  readonly responseType: typeof google_protobuf_empty_pb.Empty;
};

export class CTF {
  static readonly serviceName: string;
  static readonly GetRingFlags: CTFGetRingFlags;
  static readonly CreateFlag: CTFCreateFlag;
  static readonly PromoteFlag: CTFPromoteFlag;
  static readonly RemoveFlag: CTFRemoveFlag;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }

interface UnaryResponse {
  cancel(): void;
}
interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: (status?: Status) => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}
interface RequestStream<T> {
  write(message: T): RequestStream<T>;
  end(): void;
  cancel(): void;
  on(type: 'end', handler: (status?: Status) => void): RequestStream<T>;
  on(type: 'status', handler: (status: Status) => void): RequestStream<T>;
}
interface BidirectionalStream<ReqT, ResT> {
  write(message: ReqT): BidirectionalStream<ReqT, ResT>;
  end(): void;
  cancel(): void;
  on(type: 'data', handler: (message: ResT) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'end', handler: (status?: Status) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'status', handler: (status: Status) => void): BidirectionalStream<ReqT, ResT>;
}

export class CTFClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  getRingFlags(
    requestMessage: ctf_ctf_pb.Ring,
    metadata?: grpc.Metadata,
  ): Promise<ctf_ctf_pb.Flags>;
  createFlag(
    requestMessage: ctf_ctf_pb.Flag,
    metadata?: grpc.Metadata,
  ): Promise<ctf_ctf_pb.Flag>;
  promoteFlag(
    requestMessage: ctf_ctf_pb.Flag,
    metadata?: grpc.Metadata,
  ): Promise<ctf_ctf_pb.Flag>;
  removeFlag(
    requestMessage: ctf_ctf_pb.Flag,
    metadata?: grpc.Metadata,
  ): Promise<google_protobuf_empty_pb.Empty>;
}

