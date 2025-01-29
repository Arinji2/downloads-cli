import RedirectButton from "@/components/redirect-button";

export default function Home() {
  return (
    <div className="flex h-[100svh] w-full flex-col items-center justify-center gap-6 bg-slate-800 py-4">
      <h1 className="text-4xl font-bold text-white">Welcome to Next.js!</h1>
      <RedirectButton />
    </div>
  );
}
