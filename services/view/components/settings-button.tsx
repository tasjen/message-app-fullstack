import { Settings } from "lucide-react";
import { Button } from "@/components/ui/button";
import Link from "next/link";

export default function SettingsButton() {
  return (
    <Link href="/settings">
      <Button variant="secondary">
        <Settings />
      </Button>
    </Link>
  );
}
