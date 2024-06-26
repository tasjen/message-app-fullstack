import ChatBoard from "@/components/chat/chat-board";
import ChatForm from "@/components/chat/chat-form";
import { fetchLatestMessages } from "@/lib/data";

type Props = {
  params: { groupId: string };
};

export default async function ChatPage({ params: { groupId } }: Props) {
  const messages = await fetchLatestMessages(groupId);
  return (
    <div className="flex flex-col h-full justify-between gap-4">
      <ChatBoard messages={messages} />
      <ChatForm />
    </div>
  );
}
