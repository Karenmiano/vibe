import { communities } from "@/mocks/communities";

import CommunityCard from "./CommunityCard";
import styles from "./CommunityList.module.css";

export default function CommunityList() {
  return (
    <section className={styles.communities}>
      <header>
        <h2>Communities</h2>
        <p>Popular chat rooms</p>
      </header>

      <ul className={styles.communitiesList}>
        {communities.map((community) => (
          <li>
            <CommunityCard community={community} />
          </li>
        ))}
      </ul>
    </section>
  );
}
