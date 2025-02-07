import RedirectButton from "@/components/redirect-button";
import { Suspense } from "react";
import Example from "./(example)/example.client";
import Feature from "./feature";
import Hero from "./hero";
import Support from "./support";

export default async function Home() {
  return (
    <div className="w-full h-full flex flex-col items-center justify-start gap-20">
      <Hero />
      <Support />
      <Feature />
      <Suspense fallback={<></>}>
        <Example />
      </Suspense>
      <RedirectButton />
    </div>
  );
}
