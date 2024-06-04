"use client";

import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { Button } from "@/components/ui/button";
import { FriendContact, User } from "@/lib/schema";
import { useStore } from "@/lib/stores/client-store";
import Link from "next/link";
import { UserCheck } from "lucide-react";
import AddFriendButton from "./add-friend-button";
import { useSearchParams } from "next/navigation";

type Props = {
  user: User;
};

export default function FriendSearchResultClient({ user }: Props) {
  const store = useStore();

  // for preventing the previous search result from showing up even if there is no `q` searchParam
  const searchParams = useSearchParams();
  if (!searchParams.get("q")) {
    return <></>;
  }

  const contact = store.contacts.find(
    (contact) => (contact as FriendContact).userId === user.id
  );

  return (
    <div className="mt-4 flex flex-col items-center">
      <Avatar className="self-center h-36 w-36 mb-4">
        <AvatarImage src={user.pfp} alt={`${user.pfp}'s pfp`} />
        <AvatarFallback>{user.username[0].toUpperCase()}</AvatarFallback>
      </Avatar>
      <div className="flex text-xl mb-2 gap-2.5">
        {user.username} {contact && <UserCheck />}
      </div>
      {contact ? (
        <Link href={`/chat/${contact.groupId}`}>
          <Button>Chat</Button>
        </Link>
      ) : (
        <AddFriendButton user={user} />
      )}
    </div>
  );
}
