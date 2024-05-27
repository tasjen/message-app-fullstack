// Original file: lib/chat_proto/chat.proto

import type { Timestamp as _google_protobuf_Timestamp, Timestamp__Output as _google_protobuf_Timestamp__Output } from './google/protobuf/Timestamp';

export interface Message {
  'id'?: (number);
  'fromUsername'?: (string);
  'fromPfp'?: (string);
  'content'?: (string);
  'sentAt'?: (_google_protobuf_Timestamp | null);
}

export interface Message__Output {
  'id'?: (number);
  'fromUsername'?: (string);
  'fromPfp'?: (string);
  'content'?: (string);
  'sentAt'?: (_google_protobuf_Timestamp__Output);
}
