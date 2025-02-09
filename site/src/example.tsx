export type ConventionData = {
  name: string;
  description: string;
  items: ConventionItem[];
};

export type ConventionItem = {
  name: string;
  args: string[];
  description: string;
  demo: {
    terminal: string;
    explorer: string;
  };
};
export const ConventionsData = [
  {
    name: "MOVE",
    description: `Conventions to help you move files, from the downloads folder to anywhere on your computer.`,
    items: [
      {
        name: "Move Default",
        args: ["md", "pictures", "test.png"],
        description: `<span class="text-sm text-brand-offWhite">
          Move a file, using a location preset in your 
          <span class="text-brand-primaryLight">options file</span>
        </span>`,
        demo: {
          terminal: "https://cdn.arinji.com/u/P2qlH9.mp4",
          explorer: "https://cdn.arinji.com/u/U8mnff.mp4",
        },
      },
      {
        name: "Move Custom",
        args: ["mc", "~#pictures#test", "test.png"],
        description: `<span class="text-sm text-brand-offWhite"> 
          Move a file, using a custom location 
          <span class="text-brand-primaryLight">
            using # as separators
          </span>
        </span>`,
        demo: {
          terminal: "https://cdn.arinji.com/u/JuQIMk.mp4",
          explorer: "https://cdn.arinji.com/u/uBfSEp.mp4",
        },
      },
      {
        name: "Move Custom Default",
        args: ["mcd", "pictures#test", "test.png"],
        description: `<span class="text-sm text-brand-offWhite">
          Move a file, using a custom location 
          <span class="text-brand-primaryLight">
            using # as separators
          </span> based on a default location preset in your 
          <span class="text-brand-primaryLight">options file</span>
        </span>`,
        demo: {
          terminal: "https://cdn.arinji.com/u/i9m9RK.mp4",
          explorer: "https://cdn.arinji.com/u/xSjmvj.mp4",
        },
      },
    ],
  },
  {
    name: "LINK",
    description: `Conventions to help you convert your downloads, into shareable links.`,
    items: [
      {
        name: "Link Temporary",
        args: ["l", "t", "test.png"],
        description: `<span class="text-sm text-brand-offWhite">
          Make a temporary CDN link for a file, up to 150MB which will expire in an hour.
        </span>`,
        demo: {
          terminal: "https://cdn.arinji.com/u/ZAq5g9.mp4",
          explorer: "https://cdn.arinji.com/u/OOtbFf.mp4",
        },
      },
      {
        name: "Link Permanent",
        args: ["l", "p", "test.png"],
        description: `<span class="text-sm text-brand-offWhite">
          Make a permanent CDN link for a file, up to 100MB which will never expire.
        </span>`,
        demo: {
          terminal: "https://cdn.arinji.com/u/RK9q17.mp4",
          explorer: "https://cdn.arinji.com/u/sLYhRJ.mp4",
        },
      },
    ],
  },
  {
    name: "DELETE",
    description: `Conventions to help you make temporary files, and delete them after a set amount of time.`,
    items: [
      {
        name: "Delete",
        args: ["d", "2s", "test.png"],
        description: `<span class="text-sm text-brand-offWhite">
          Make a temporary file, and delete it after 2 seconds.
        </span>`,
        demo: {
          terminal: "https://cdn.arinji.com/u/rDPGlx.mp4",
          explorer: "https://cdn.arinji.com/u/z5vHCo.mp4",
        },
      },
    ],
  },
] as ConventionData[];
