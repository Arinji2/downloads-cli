import RedirectButton from "@/components/redirect-button";
import { headers } from "next/headers";
import { Suspense } from "react";
import Example from "./(example)/example.client";
import Demo from "./demo.client";
import Feature from "./feature";
import Footer from "./footer";
import Hero from "./hero";
import Install from "./install.client";
import Support from "./support";

export default async function Home() {
  const headersList = await headers();
  const userAgent = (headersList.get("user-agent") || "Unknown").toLowerCase();
  const isWindowsOS =
    userAgent.includes("win") || userAgent.includes("windows");
  return (
    <div className="w-full h-full flex flex-col items-center justify-start gap-20">
      <Hero isWindowsOS={isWindowsOS} />
      <Support />
      <Feature />
      <Suspense fallback={<></>}>
        <Example />
      </Suspense>
      <Suspense fallback={<></>}>
        <Install isWindowsOS={isWindowsOS} />
        <Demo />
      </Suspense>
      <Footer />
      <RedirectButton />
    </div>
  );
}
