// @generated by protoc-gen-es v1.10.0 with parameter "target=ts"
// @generated from file auth/auth.proto (package messenjo, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message messenjo.VerifyTokenReq
 */
export class VerifyTokenReq extends Message<VerifyTokenReq> {
  /**
   * @generated from field: string token = 1;
   */
  token = "";

  constructor(data?: PartialMessage<VerifyTokenReq>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "messenjo.VerifyTokenReq";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): VerifyTokenReq {
    return new VerifyTokenReq().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): VerifyTokenReq {
    return new VerifyTokenReq().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): VerifyTokenReq {
    return new VerifyTokenReq().fromJsonString(jsonString, options);
  }

  static equals(a: VerifyTokenReq | PlainMessage<VerifyTokenReq> | undefined, b: VerifyTokenReq | PlainMessage<VerifyTokenReq> | undefined): boolean {
    return proto3.util.equals(VerifyTokenReq, a, b);
  }
}

/**
 * @generated from message messenjo.VerifyTokenRes
 */
export class VerifyTokenRes extends Message<VerifyTokenRes> {
  /**
   * @generated from field: bytes user_id = 1;
   */
  userId = new Uint8Array(0);

  constructor(data?: PartialMessage<VerifyTokenRes>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "messenjo.VerifyTokenRes";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "user_id", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): VerifyTokenRes {
    return new VerifyTokenRes().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): VerifyTokenRes {
    return new VerifyTokenRes().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): VerifyTokenRes {
    return new VerifyTokenRes().fromJsonString(jsonString, options);
  }

  static equals(a: VerifyTokenRes | PlainMessage<VerifyTokenRes> | undefined, b: VerifyTokenRes | PlainMessage<VerifyTokenRes> | undefined): boolean {
    return proto3.util.equals(VerifyTokenRes, a, b);
  }
}

