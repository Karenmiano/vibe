import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/app/chat/$chatId")({
  component: ChatThread,
});

function ChatThread() {
  return (
    <main>
      <header>
        <button>Back</button>
        <div>
          <img src="" alt="" />
          <h2>Chat Name</h2>
        </div>
      </header>
      <ul role="log" aria-live="polite"></ul>
      <form>
        <input type="text" placeholder="Type a message" name="message" />
      </form>
    </main>
  );
}
