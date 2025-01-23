import { useEffect } from "react";

const URL_REGEX = /^[a-zA-Z0-9]+\.[a-zA-Z]+$/;
export default function Home() {
  useEffect(() => {
    const hash = window.location.hash.slice(1);
    const params = new URLSearchParams(hash);

    const urlID = params.get("urlID");
    const type = params.get("type");
    if (urlID && type) {
      if (URL_REGEX.test(urlID)) {
        if (type === "t") {
          window.location.href = `https://litter.catbox.moe/${urlID}`;
        } else if (type === "p") {
          window.location.href = `https://files.catbox.moe/${urlID}`;
        }
      }
    }
  }, []);
  return <div className="w-full h-full"></div>;
}
