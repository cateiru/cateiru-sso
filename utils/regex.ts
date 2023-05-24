export const emailRegex = /[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,}/i;
export const userIDRegex = /[A-Z0-9_]{3,15}/i;
export const userIdEmailRegex = new RegExp(
  `^(${userIDRegex.source})|(${emailRegex.source})$`,
  'i'
);
