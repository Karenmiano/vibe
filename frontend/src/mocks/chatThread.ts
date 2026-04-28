export const recipient = {
  avatar: "/images/frankenstein.jpeg",
  name: "So I watched Frankenstein, and I cried",
};

/** Mirrors API shape: join users for group threads (Frankenstein discussion). */
export type Message = {
  content: string;
  senderId: string;
  createdAt: string;
  senderName: string;
  senderAvatar: string;
};

export const conversation: Message[] = [
  {
    senderId: "pixel-panda",
    senderName: "pixel_panda",
    senderAvatar: "https://i.pravatar.cc/300?img=1",
    content:
      "I went in expecting horror and came out questioning what it means to be human. Anyone else?",
    createdAt: "10:31 AM",
  },
  {
    senderId: "code-wizard",
    senderName: "codewizard99",
    senderAvatar: "https://i.pravatar.cc/300?img=2",
    content:
      "YES. I literally sat in silence for like 10 minutes after it ended.",
    createdAt: "10:31 AM",
  },
  {
    senderId: "pixel-panda",
    senderName: "pixel_panda",
    senderAvatar: "https://i.pravatar.cc/300?img=1",
    content: "The creature never asked to exist. That hit different.",
    createdAt: "10:32 AM",
  },
  {
    senderId: "luna-sky",
    senderName: "luna.sky",
    senderAvatar: "https://i.pravatar.cc/300?img=3",
    content:
      "And Frankenstein just... abandoned him the moment he opened his eyes. Like, you made a person and then noped out.",
    createdAt: "10:32 AM",
  },
  {
    senderId: "pixel-panda",
    senderName: "pixel_panda",
    senderAvatar: "https://i.pravatar.cc/300?img=1",
    content: "The real monster was the scientist all along lol",
    createdAt: "10:33 AM",
  },
  {
    senderId: "shadow-fox",
    senderName: "shadowfox_",
    senderAvatar: "https://i.pravatar.cc/300?img=4",
    content: "Classic 'we live in a society' moment from 1818",
    createdAt: "10:33 AM",
  },
  {
    senderId: "pixel-panda",
    senderName: "pixel_panda",
    senderAvatar: "https://i.pravatar.cc/300?img=1",
    content:
      "What got me was how desperately the creature wanted connection. He just wanted ONE friend.",
    createdAt: "10:34 AM",
  },
  {
    senderId: "echo-wave",
    senderName: "echo.wave",
    senderAvatar: "https://i.pravatar.cc/300?img=5",
    content:
      "The part where he watches the De Lacey family for months just to learn how to speak... he was so hopeful",
    createdAt: "10:34 AM",
  },
  {
    senderId: "pixel-panda",
    senderName: "pixel_panda",
    senderAvatar: "https://i.pravatar.cc/300?img=1",
    content: "And then they screamed and ran away. That was brutal.",
    createdAt: "10:35 AM",
  },
  {
    senderId: "nova-blast",
    senderName: "novablast",
    senderAvatar: "https://i.pravatar.cc/300?img=6",
    content:
      "You could feel him shutting down emotionally after that. Like something just died in him.",
    createdAt: "10:35 AM",
  },
  {
    senderId: "pixel-panda",
    senderName: "pixel_panda",
    senderAvatar: "https://i.pravatar.cc/300?img=1",
    content:
      "Which is what makes the rage make sense. It's not random evil. It's grief.",
    createdAt: "10:36 AM",
  },
  {
    senderId: "glitch-master",
    senderName: "glitch_master",
    senderAvatar: "https://i.pravatar.cc/300?img=7",
    content:
      "Exactly. Every act of destruction was a response to rejection. He was literally taught to hate himself by the way people treated him.",
    createdAt: "10:36 AM",
  },
  {
    senderId: "pixel-panda",
    senderName: "pixel_panda",
    senderAvatar: "https://i.pravatar.cc/300?img=1",
    content:
      "The blind man scene wrecked me. The one person who couldn't see his appearance actually talked to him like a human.",
    createdAt: "10:37 AM",
  },
  {
    senderId: "solar-flare",
    senderName: "solar.flare",
    senderAvatar: "https://i.pravatar.cc/300?img=8",
    content:
      "That scene is so quietly devastating. And then the son comes back and ruins everything.",
    createdAt: "10:37 AM",
  },
  {
    senderId: "pixel-panda",
    senderName: "pixel_panda",
    senderAvatar: "https://i.pravatar.cc/300?img=1",
    content:
      "I kept thinking about what would've happened if the old man had just... believed him.",
    createdAt: "10:38 AM",
  },
  {
    senderId: "neon-tiger",
    senderName: "neontiger",
    senderAvatar: "https://i.pravatar.cc/300?img=9",
    content:
      "The whole story pivots on that one moment of rejection. It's wild how much weight it carries.",
    createdAt: "10:38 AM",
  },
  {
    senderId: "pixel-panda",
    senderName: "pixel_panda",
    senderAvatar: "https://i.pravatar.cc/300?img=1",
    content:
      "Do you think Frankenstein was supposed to be a cautionary tale about science or about responsibility?",
    createdAt: "10:39 AM",
  },
  {
    senderId: "crimson-raven",
    senderName: "crimson_raven",
    senderAvatar: "https://i.pravatar.cc/300?img=10",
    content:
      "Both honestly. Shelley wrote it during the industrial revolution. People were terrified of what progress without ethics looks like.",
    createdAt: "10:39 AM",
  },
  {
    senderId: "pixel-panda",
    senderName: "pixel_panda",
    senderAvatar: "https://i.pravatar.cc/300?img=1",
    content:
      "And here we are in 2024 having the same conversation about AI lol",
    createdAt: "10:40 AM",
  },
  {
    senderId: "frost-byte",
    senderName: "frostbyte",
    senderAvatar: "https://i.pravatar.cc/300?img=11",
    content:
      "The book basically predicted every ethics debate we're having now. Create something sentient, ignore the consequences, act surprised when it goes wrong.",
    createdAt: "10:40 AM",
  },
  {
    senderId: "pixel-panda",
    senderName: "pixel_panda",
    senderAvatar: "https://i.pravatar.cc/300?img=1",
    content:
      "The creature asking for a companion hit me especially hard with that framing.",
    createdAt: "10:41 AM",
  },
  {
    senderId: "iron-lotus",
    senderName: "iron.lotus",
    senderAvatar: "https://i.pravatar.cc/300?img=12",
    content:
      "He wasn't asking for power or revenge at that point. Just someone who understood him. That's so human it hurts.",
    createdAt: "10:41 AM",
  },
  {
    senderId: "pixel-panda",
    senderName: "pixel_panda",
    senderAvatar: "https://i.pravatar.cc/300?img=1",
    content:
      "And Frankenstein destroys the female creature out of fear of what they might create together. Still punishing him for existing.",
    createdAt: "10:42 AM",
  },
  {
    senderId: "code-wizard",
    senderName: "codewizard99",
    senderAvatar: "https://i.pravatar.cc/300?img=2",
    content:
      "That scene made me genuinely angry. The creature had ONE ask and it got ripped away.",
    createdAt: "10:42 AM",
  },
  {
    senderId: "pixel-panda",
    senderName: "pixel_panda",
    senderAvatar: "https://i.pravatar.cc/300?img=1",
    content:
      "The ending is so bleak too. Both of them destroyed by the whole thing. No winners.",
    createdAt: "10:43 AM",
  },
  {
    senderId: "luna-sky",
    senderName: "luna.sky",
    senderAvatar: "https://i.pravatar.cc/300?img=3",
    content:
      "That's what makes it literature and not just horror. It doesn't let anyone off the hook.",
    createdAt: "10:43 AM",
  },
  {
    senderId: "pixel-panda",
    senderName: "pixel_panda",
    senderAvatar: "https://i.pravatar.cc/300?img=1",
    content:
      "I feel like I need to read it again now. I definitely missed things.",
    createdAt: "10:44 AM",
  },
  {
    senderId: "shadow-fox",
    senderName: "shadowfox_",
    senderAvatar: "https://i.pravatar.cc/300?img=4",
    content:
      "The letters framing it as a story within a story also adds so much. Walton's obsession mirrors Frankenstein's exactly.",
    createdAt: "10:44 AM",
  },
  {
    senderId: "pixel-panda",
    senderName: "pixel_panda",
    senderAvatar: "https://i.pravatar.cc/300?img=1",
    content:
      "Oh wow I didn't even clock that parallel properly. Same hubris, different scale.",
    createdAt: "10:45 AM",
  },
  {
    senderId: "echo-wave",
    senderName: "echo.wave",
    senderAvatar: "https://i.pravatar.cc/300?img=5",
    content:
      "Shelley was 18 when she wrote this. I need to just lie down and think about my life choices.",
    createdAt: "10:45 AM",
  },
];
