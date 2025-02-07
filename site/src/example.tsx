export const ConventionsData = [
  {
    name: "MOVE",
    description: `Conventions to help you move files, from the downloads folder to anywhere on your computer.`,
    items: [
      {
        name: "Move Default",
        args: ["md", "pictures", "test.png"],
        infoLink: "/move?selected=default",
        description: `<span class="text-sm text-brand-offWhite">
  Move a file, using a location preset in your 
  <span class="text-brand-primaryLight">options file</span>
</span>`,
      },
      {
        name: "Move Custom",
        args: ["mc", "~#pictures#test", "test.png"],
        infoLink: "/move?selected=custom",
        description: `<span class="text-sm text-brand-offWhite"> 
  Move a file, using a custom location 
  <span class="text-brand-primaryLight">
    using # as seperators
  </span>
</span>`,
      },

      {
        name: "Move Custom Default",
        args: ["mcd", "pictures#test", "test.png"],
        infoLink: "/move?selected=customdefault",
        description: `<span class="text-sm text-brand-offWhite">
  Move a file, using a custom location 
  <span class="text-brand-primaryLight">
    using # as seperators
  </span> based on a default location preset in your 
  <span class="text-brand-primaryLight">options file</span>
      </span>`,
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
        infoLink: "/link?selected=temporary",
        description: `<span class="text-sm text-brand-offWhite">
   Make a temporary cdn link for a file, upto 150mb which will expire in a hour.
  </span>`,
      },
      {
        name: "Link Permanent",
        args: ["l", "p", "test.png"],
        infoLink: "/link?selected=permanent",
        description: `<span class="text-sm text-brand-offWhite">
   Make a permanent cdn link for a file, upto 100mb which will never expire.
  </span>`,
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
        infoLink: "/delete",
        description: `<span class="text-sm text-brand-offWhite">
    Make a temporary file, and delete it after 2 seconds
  </span>`,
      },
    ],
  },
];
