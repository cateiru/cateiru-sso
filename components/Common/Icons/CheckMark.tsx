import {Box, keyframes, usePrefersReducedMotion} from '@chakra-ui/react';
import React from 'react';

export interface CheckmarkProps {
  size: number;
  bgColor: string;
  color: string;
}

export const CheckMark: React.FC<CheckmarkProps> = props => {
  const prefersReducedMotion = usePrefersReducedMotion();

  const stroke = keyframes`
100% {
  stroke-dashoffset: 0;
}
`;

  const scale = keyframes`
0%,
100% {
  transform: none;
}
50% {
  transform: scale3d(1.1, 1.1, 1);
}
`;

  const fill = keyframes`
100% {
  box-shadow: inset 0px 0px 0px ${props.size / 2}px ${props.bgColor};
}
`;

  return (
    <Box>
      <Box
        as="svg"
        xmlns="http://www.w3.org/2000/svg"
        viewBox="0 0 52 52"
        w={`${props.size}px`}
        h={`${props.size}px`}
        borderRadius="50%"
        display="block"
        strokeWidth="5"
        stroke={props.color}
        strokeMiterlimit="10"
        boxShadow={`inset 0px 0px 0px ${props.bgColor}`}
        animation={
          prefersReducedMotion
            ? undefined
            : `${fill} .4s ease-in-out .4s forwards, ${scale} .3s ease-in-out .9s both`
        }
      >
        <Box
          as="circle"
          cx="26"
          cy="26"
          r="25"
          fill="none"
          strokeDasharray="166"
          strokeDashoffset="166"
          strokeWidth="2"
          strokeMiterlimit="10"
          stroke={props.bgColor}
          animation={
            prefersReducedMotion
              ? undefined
              : `${stroke} 0.6s cubic-bezier(0.65, 0, 0.45, 1) forwards`
          }
        />
        <Box
          as="path"
          fill="none"
          d="M14.1 27.2l7.1 7.2 16.7-16.8"
          transformOrigin="50% 50%"
          strokeDasharray="48"
          strokeDashoffset="48"
          animation={
            prefersReducedMotion
              ? undefined
              : `${stroke} 0.3s cubic-bezier(0.65, 0, 0.45, 1) 0.8s forwards`
          }
        />
      </Box>
    </Box>
  );
};
