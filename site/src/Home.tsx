import { useEffect } from "react";

const URL_REGEX = /^[a-zA-Z0-9]+\.[a-zA-Z]+$/;
export default function Home() {
  const redirectFunc = () => {
    const urlID = window.location.hash.slice(1);
    const type = window.location.hash.slice(1);
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
    <div className="py-4 w-full h-[100svh] gap-6 bg-slate-800 flex flex-col items-center justify-center">
      <h1 className="font-bold md:text-6xl text-4xl text-white text-center">
        "Downloads On Steroids"
      </h1>
      <p className="text-white/50 text-xl md:text-2xl">
        You will be redirected shortly.
      </p>
      <button
        onClick={redirectFunc}
        className=" text-white/50 shadow-black shadow-md mt-20 text-center hover:shadow-sm transition-all ease-in-out duration-300 px-6 py-4 text-sm md:text-base p-2 rounded-md"
      >
        If you haven't been redirected yet, click here.
      </button>
    </div>
  );
}
