import { PlusIcon } from "@phosphor-icons/react";

import ProfileMenu from "@/features/user/components/ProfileMenu";
import Logo from "@/ui/Logo";
import styles from "./AppBar.module.css";

export default function AppBar() {
  return (
    <div className={styles.appBar}>
      <Logo />
      <div>
        <button
          aria-label="Create new room"
          title="Create new room"
          className={styles.createRoomBtn}
        >
          <PlusIcon size={18} />
        </button>
        <ProfileMenu />
      </div>
    </div>
  );
}
