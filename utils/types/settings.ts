import {z} from 'zod';

export const UserUpdateEmailScheme = z.object({
  session: z.string(),
});
export type UserUpdateEmail = z.infer<typeof UserUpdateEmailScheme>;

export const UserUpdateEmailRegisterScheme = z.object({
  email: z.string().email(),
});
export type UserUpdateEmailRegister = z.infer<
  typeof UserUpdateEmailRegisterScheme
>;
