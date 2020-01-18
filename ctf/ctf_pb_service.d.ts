// package: roleypoly.ctf
// file: ctf/ctf.proto

import * as ctf_ctf_pb from "../ctf/ctf_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";
import {grpc} from "@improbable-eng/grpc-web";

type CTFGetCanaries = {
  readonly methodName: string;
  readonly service: typeof CTF;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof ctf_ctf_pb.CanaryQuery;
  readonly responseType: typeof ctf_ctf_pb.Canaries;
};

type CTFSetSingleCanaryRollout = {
  readonly methodName: string;
  readonly service: typeof CTF;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof ctf_ctf_pb.CanaryQuery;
  readonly responseType: typeof google_protobuf_empty_pb.Empty;
};

type CTFSetGlobalCanaryRollout = {
  readonly methodName: string;
  readonly service: typeof CTF;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof ctf_ctf_pb.CanaryQuery;
  readonly responseType: typeof google_protobuf_empty_pb.Empty;
};

type CTFCreateCanary = {
  readonly methodName: string;
  readonly service: typeof CTF;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof ctf_ctf_pb.Canary;
  readonly responseType: typeof ctf_ctf_pb.Canary;
};

type CTFDeleteCanary = {
  readonly methodName: string;
  readonly service: typeof CTF;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof ctf_ctf_pb.CanaryQuery;
  readonly responseType: typeof google_protobuf_empty_pb.Empty;
};

export class CTF {
  static readonly serviceName: string;
  static readonly GetCanaries: CTFGetCanaries;
  static readonly SetSingleCanaryRollout: CTFSetSingleCanaryRollout;
  static readonly SetGlobalCanaryRollout: CTFSetGlobalCanaryRollout;
  static readonly CreateCanary: CTFCreateCanary;
  static readonly DeleteCanary: CTFDeleteCanary;
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
  getCanaries(
    requestMessage: ctf_ctf_pb.CanaryQuery,
    metadata?: grpc.Metadata,
  ): Promise<ctf_ctf_pb.Canaries>;
  setSingleCanaryRollout(
    requestMessage: ctf_ctf_pb.CanaryQuery,
    metadata?: grpc.Metadata,
  ): Promise<google_protobuf_empty_pb.Empty>;
  setGlobalCanaryRollout(
    requestMessage: ctf_ctf_pb.CanaryQuery,
    metadata?: grpc.Metadata,
  ): Promise<google_protobuf_empty_pb.Empty>;
  createCanary(
    requestMessage: ctf_ctf_pb.Canary,
    metadata?: grpc.Metadata,
  ): Promise<ctf_ctf_pb.Canary>;
  deleteCanary(
    requestMessage: ctf_ctf_pb.CanaryQuery,
    metadata?: grpc.Metadata,
  ): Promise<google_protobuf_empty_pb.Empty>;
}

