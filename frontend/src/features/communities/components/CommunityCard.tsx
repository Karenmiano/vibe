import { UsersIcon } from "@phosphor-icons/react";
import styles from "./CommunityCard.module.css";

type CommunityCardProps = {
  community: {
    topic: string;
    thumbnail: string;
    leadText: string;
    followersCount: number | string;
    theme: string;
  };
};

export default function CommunityCard({ community }: CommunityCardProps) {
  return (
    <article
      className={styles.communityCard}
      style={{ backgroundColor: community.theme }}
    >
      <div className={styles.communityHeader}>
        <img src={community.thumbnail} alt="" className={styles.thumbnail} />
        <div>
          <h3 className={styles.topic}>{community.topic}</h3>
          <div className={styles.followersCount}>
            <UsersIcon />
            <p>{community.followersCount} Followers</p>
          </div>
        </div>
      </div>
      <p className={styles.leadText}>{community.leadText}</p>
      <button className={styles.joinButton}>Join</button>
    </article>
  );
}
