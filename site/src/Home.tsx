import { useEffect } from "react";

const URL_REGEX = /^[a-zA-Z0-9]+\.[a-zA-Z]+$/;
export default function Home() {
  const redirectFunc = () => {
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
  };
  useEffect(() => {
    redirectFunc();
  }, []);
  return (
    <div className="flex h-[100svh] w-full flex-col items-center justify-center gap-6 bg-slate-800 py-4">
      <h1 className="text-center text-4xl font-bold text-white md:text-6xl">
        "Downloads On Steroids"
      </h1>
      <p className="text-xl text-white/50 md:text-2xl">
        You will be redirected shortly.
      </p>
      <button
        onClick={redirectFunc}
        className="mt-20 rounded-md p-2 px-6 py-4 text-center text-sm text-white/50 shadow-md shadow-black transition-all duration-300 ease-in-out hover:shadow-sm md:text-base"
      >
        If you haven't been redirected yet, click here.
      </button>
    </div>
  );
}
