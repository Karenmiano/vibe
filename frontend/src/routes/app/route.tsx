import { chats } from "@/mocks/chats";
import { createFileRoute, Outlet } from "@tanstack/react-router";

export const Route = createFileRoute("/app")({
  component: RouteComponent,
});

function RouteComponent() {
  return (
    <div className="container">
      <section className="chats">
        <header>
          <h2>Chats</h2>
        </header>

        <ul className="chat-filters">
          <li className="active">All</li>
          <li>Personal</li>
          <li>Groups</li>
        </ul>

        <ul className="chat-list">
          {chats.map((chat) => (
            <li>
              <img
                src={chat.avatar}
                alt={`${chat.name}'s avatar`}
                className="avatar"
              />
              <div className="chat-content">
                <div>
                  <p>{chat.name}</p>
                  <p className="timestamp">{chat.timestamp}</p>
                </div>
                <div>
                  <p className="recent-message">{chat.recentMessage}</p>
                  <div className="unread-count">{chat.unreadCount}</div>
                </div>
              </div>
            </li>
          ))}
        </ul>
      </section>
      <Outlet />
    </div>
  );
}
