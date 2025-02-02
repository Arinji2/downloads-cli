import { IconProps, IconSvg } from "./base";

export const SpeedIcon = (props: IconProps) => (
  <IconSvg viewBox="0 0 24 24" fill="none" stroke="currentColor" {...props}>
    <path
      fill="currentColor"
      d="M15 5H9v2H5v2H3v2H1v6h2v2h2v-2H3v-6h2V9h4V7h6zm8 6h-2v6h-2v2h2v-2h2zm-13 2h4v4h-4zm6-2h-2v2h2zm2-2v2h-2V9zm0 0V7h2v2z"
    />
  </IconSvg>
);
