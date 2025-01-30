export default function Terminal() {
  return (
    <div className="w-fit h-fit z-10 flex flex-col items-start justify-start gap-3 py-6 px-5 bg-[#1D1D1D] shadow-terminal-shadow">
      <TerminalLine
        updatedName="md-pictures-test.png"
        response="Moved text.png to /pictures folder!"
      />
      <TerminalLine
        updatedName="mc-pictures#dos-test.png"
        response="Moved text.png to /pictures/dos folder!"
      />
      <TerminalLine
        updatedName="l-p-test.png"
        response="Made a Permanent Shareable URL for test.png!"
      />
      <TerminalLine
        updatedName="d-2s-test.png"
        response="Deleting test.png in 2 seconds!"
      />
    </div>
  );
}
function TerminalLine({
  updatedName,
  response,
}: {
  updatedName: string;
  response: string;
}) {
  return (
    <div className="font-medium text-sm h-fit w-fit gap-2 flex flex-col items-start justify-start">
      <p>
        <span className="text-brand-offWhite">~/Downloads$</span>{" "}
        <span className="text-brand-darkWhite">mv</span>{" "}
        <span className="text-brand-darkBlue">test.png</span>{" "}
        <span className="text-brand-darkYellow">{updatedName}</span>
      </p>
      <p className="text-brand-primary">
        <span className="text-brand-offWhite">~/Downloads$</span> [DOS]
        {response}
      </p>
    </div>
  );
}
