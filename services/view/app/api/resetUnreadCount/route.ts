import { chatClient } from "@/lib/grpc-clients/node";
import { parse as uuidParse } from "uuid";
import { handleNodeError } from "@/lib/util";
import { cookies } from "next/headers";

export async function POST(req: Request) {
  try {
    const groupId = new URL(req.url).searchParams.get("groupId");
    if (!groupId) throw new Error("no 'groupId' searchParam");

    await chatClient.resetUnreadCount(
      {
        groupId: uuidParse(groupId),
      },
      {
        headers: { cookie: cookies().toString() },
      }
    );
  } catch (err) {
    handleNodeError(err);
  }
  return new Response();
}
