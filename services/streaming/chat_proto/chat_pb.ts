// @generated by protoc-gen-es v1.10.0 with parameter "target=ts"
// @generated from file chat.proto (syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message GetGroupIdsReq
 */
export class GetGroupIdsReq extends Message<GetGroupIdsReq> {
  /**
   * @generated from field: bytes user_id = 1;
   */
  userId = new Uint8Array(0);

  constructor(data?: PartialMessage<GetGroupIdsReq>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "GetGroupIdsReq";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "user_id", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetGroupIdsReq {
    return new GetGroupIdsReq().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetGroupIdsReq {
    return new GetGroupIdsReq().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetGroupIdsReq {
    return new GetGroupIdsReq().fromJsonString(jsonString, options);
  }

  static equals(a: GetGroupIdsReq | PlainMessage<GetGroupIdsReq> | undefined, b: GetGroupIdsReq | PlainMessage<GetGroupIdsReq> | undefined): boolean {
    return proto3.util.equals(GetGroupIdsReq, a, b);
  }
}

/**
 * @generated from message GetGroupIdsRes
 */
export class GetGroupIdsRes extends Message<GetGroupIdsRes> {
  /**
   * @generated from field: repeated bytes group_ids = 1;
   */
  groupIds: Uint8Array[] = [];

  constructor(data?: PartialMessage<GetGroupIdsRes>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "GetGroupIdsRes";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "group_ids", kind: "scalar", T: 12 /* ScalarType.BYTES */, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetGroupIdsRes {
    return new GetGroupIdsRes().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetGroupIdsRes {
    return new GetGroupIdsRes().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetGroupIdsRes {
    return new GetGroupIdsRes().fromJsonString(jsonString, options);
  }

  static equals(a: GetGroupIdsRes | PlainMessage<GetGroupIdsRes> | undefined, b: GetGroupIdsRes | PlainMessage<GetGroupIdsRes> | undefined): boolean {
    return proto3.util.equals(GetGroupIdsRes, a, b);
  }
}

