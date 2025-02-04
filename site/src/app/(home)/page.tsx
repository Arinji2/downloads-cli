import RedirectButton from "@/components/redirect-button";
import Example from "./example.client";
import Feature from "./feature";
import Hero from "./hero";
import Support from "./support";

export default async function Home() {
  return (
    <div className="w-full h-full flex flex-col items-center justify-start gap-20">
      <Hero />
      <Support />
      <Feature />
      <Example />
      <RedirectButton />
    </div>
  );
}
