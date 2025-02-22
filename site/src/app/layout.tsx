import type { Metadata } from "next";
import { JetBrains_Mono } from "next/font/google";
import "./globals.css";

const jetBrainsMono = JetBrains_Mono({
  subsets: ["latin"],
});

export const metadata: Metadata = {
  title: "Create Next App",
  description: "Generated by create next app",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <head>
        <link
          rel="icon"
          type="image/png"
          href="https://cdn.arinji.com/u/eqdgi9.png"
          sizes="48×48"
        />
        <meta name="viewport" content="width=device-width, initial-scale=1.0" />
        <meta name="theme-color" content="#1D6704" />

        <link rel="canonical" href="https://dos.arinji.com/" />
        <meta
          name="description"
          content="Downloads on Steroids (DOS) is a tool that makes downloading files from the web easy and fast. It works using just filenames, making it cross platform, simple and easy to use. Delete, Move and Link files with ease, only on DOS"
        />

        <meta
          name="keywords"
          content="downloads, termina, file explorer, linux, windows"
        />
        <meta name="author" content="Arinji" />
        {/* Open Graph / Facebook */}
        <meta property="og:title" content="Downloads on Steroids" />
        <meta
          property="og:description"
          content="Downloads on Steroids (DOS) is a tool that makes downloading files from the web easy and fast. It works using just filenames, making it cross platform, simple and easy to use. Delete, Move and Link files with ease, only on DOS"
        />
        <meta
          property="og:image"
          content="https://cdn.arinji.com/u/2LEQrt.png"
        />
        <meta property="og:url" content="https://dos.arinji.com/" />
        <meta property="og:type" content="website" />

        {/* Twitter */}
        <meta name="twitter:card" content="summary_large_image" />
        <meta name="twitter:title" content="Downloads On Steroids" />
        <meta
          name="twitter:description"
          content="Downloads on Steroids (DOS) is a tool that makes downloading files from the web easy and fast. It works using just filenames, making it cross platform, simple and easy to use. Delete, Move and Link files with ease, only on DOS"
        />
        <meta
          name="twitter:image"
          content="https://cdn.arinji.com/u/2LEQrt.png"
        />
        <title>Downloads on Steroids</title>
      </head>
      <body
        className={`${jetBrainsMono.className} bg-brand-background py-20 antialiased w-full h-full flex flex-col items-center justify-center`}
      >
        <div className="w-full  px-3 flex flex-col items-center justify-start h-fit max-w-screen">
          {children}
        </div>
      </body>
    </html>
  );
}
