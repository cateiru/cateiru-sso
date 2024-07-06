import {z} from 'zod';
import {UserMeSchema} from './user';

export const LoginUserSchema = z.object({
  avatar: z.string().nullable(),
});
export type LoginUser = z.infer<typeof LoginUserSchema>;

export const LoginResponseSchema = z.object({
  user: UserMeSchema.optional(),
  otp: z
    .object({
      token: z.string(),
      login_user: LoginUserSchema,
    })
    .optional(),
});
export type LoginResponse = z.infer<typeof LoginResponseSchema>;

export const AccountReRegisterPasswordIsSessionSchema = z.object({
  active: z.boolean(),
});
export type AccountReRegisterPasswordIsSession = z.infer<
  typeof AccountReRegisterPasswordIsSessionSchema
>;
