"use client";
import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from "@/components/accordion";
import Link from "next/link";
export default function Faq() {
  return (
    <div className="w-full h-fit md:w-align flex flex-col items-center justify-start gap-7">
      <div className="items-center justify-start gap-7  w-full h-fit flex flex-col text-center">
        <h3 className="text-3xl font-bold tracking-tight text-white">FAQ</h3>
        <p className="text-sm text-brand-offWhite md:max-w-[80%]">
          Frequently asked questions
        </p>
      </div>
      <div className="md:w-[70%] w-full h-fit flex flex-col items-start justify-start gap-5">
        <Accordion type="single" className="w-full" collapsible>
          <AccordionItem value="item-1">
            <AccordionTrigger>
              Where can I find the source code?
            </AccordionTrigger>
            <AccordionContent>
              The source code is available on{" "}
              <Link
                href="https://github.com/Arinji2/downloads-cli"
                target="_blank"
                className="text-brand-primaryLight"
              >
                Github
              </Link>
              .
            </AccordionContent>
          </AccordionItem>
          <AccordionItem value="item-2">
            <AccordionTrigger>What is the options.json file?</AccordionTrigger>
            <AccordionContent>
              The configuration file for DOS, created on first startup. DOS
              guides you through its setup when it is created.
            </AccordionContent>
          </AccordionItem>
          <AccordionItem value="item-3">
            <AccordionTrigger>
              How do i edit the options.json file?
            </AccordionTrigger>
            <AccordionContent>
              You can edit the file in any text editor. If DOS is running, it
              will automaticaly restart, make sure to check the app logs if
              there is any errors in it.
            </AccordionContent>
          </AccordionItem>
          <AccordionItem value="item-4">
            <AccordionTrigger>How do i stop DOS?</AccordionTrigger>
            <AccordionContent>
              To delete DOS, just delete the file called{" "}
              <span className="text-brand-primaryLight">status</span> in the
              same directory as the dos file. This will instantly stop DOS.
            </AccordionContent>
          </AccordionItem>
          <AccordionItem value="item-5">
            <AccordionTrigger>How do i check logs?</AccordionTrigger>
            <AccordionContent>
              If you face any issues, make sure to check the file called{" "}
              <span className="text-brand-primaryLight">app.log</span> in the
              same directory as the dos file. This will contain all the app
              logs. If you dont know whats wrong, make a issue in github and
              send your app log there.
            </AccordionContent>
          </AccordionItem>
        </Accordion>
      </div>
    </div>
  );
}
