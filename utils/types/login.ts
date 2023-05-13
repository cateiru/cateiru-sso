import {z} from 'zod';
import {UserSchema} from './user';

export const LoginUserSchema = z.object({
  avatar: z.string().nullable(),
});
export type LoginUser = z.infer<typeof LoginUserSchema>;

export const LoginResponseSchema = z.object({
  user: UserSchema.optional(),
  otp: z
    .object({
      token: z.string(),
      login_user: LoginUserSchema,
    })
    .optional(),
});
export type LoginResponse = z.infer<typeof LoginResponseSchema>;
