import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/app/")({
  component: RouteComponent,
});

function RouteComponent() {
  return (
    <div>
      <section>
        <h2>Communities</h2>
        <p>Popular chat rooms</p>
      </section>
    </div>
  );
}
