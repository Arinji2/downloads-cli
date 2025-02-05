import { IconProps, IconSvg } from "./base";

export const CloseIcon = (props: IconProps) => (
  <IconSvg viewBox="0 0 24 24" fill="none" stroke="currentColor" {...props}>
    <path
      fill="currentColor"
      d="M5 3H3v18h18V3zm14 2v14H5V5zm-8 4H9V7H7v2h2v2h2v2H9v2H7v2h2v-2h2v-2h2v2h2v2h2v-2h-2v-2h-2v-2h2V9h2V7h-2v2h-2v2h-2z"
    />
  </IconSvg>
);
