import { communities } from "@/mocks/communities";
import { UsersIcon } from "@phosphor-icons/react";
import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/app/")({
  component: RouteComponent,
});

function RouteComponent() {
  return (
    <div>
      <section className="communities">
        <header>
          <h2>Communities</h2>
          <p>Popular chat rooms</p>
        </header>

        <ul className="communities-list">
          {communities.map((community) => (
            <li style={{ backgroundColor: community.theme }}>
              <div className="community-header">
                <img src={community.thumbnail} alt="" className="thumbnail" />
                <div>
                  <h3 className="topic">{community.topic}</h3>
                  <div className="followers-count">
                    <UsersIcon />
                    <p>{community.followersCount} Followers</p>
                  </div>
                </div>
              </div>
              <p className="lead-text">{community.leadText}</p>
              <button className="join-button">Join</button>
            </li>
          ))}
        </ul>
      </section>
    </div>
  );
}
