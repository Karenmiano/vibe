import { createFileRoute } from "@tanstack/react-router";

import CommunityList from "@/features/communities/components/CommunityList";

export const Route = createFileRoute("/app/")({
  component: Home,
});

function Home() {
  return (
    <main>
      <CommunityList />
    </main>
  );
}
