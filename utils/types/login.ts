import {z} from 'zod';
import {UserSchema} from './user';

export const LoginUserSchema = z.object({
  avatar: z.string().nullable(),
  user_name: z.string(),
  available_passkey: z.boolean(),
  available_password: z.boolean(),
});
export type LoginUser = z.infer<typeof LoginUserSchema>;

export const LoginResponseSchema = z.object({
  user: UserSchema.optional(),
  otp: z.string().optional(),
});
export type LoginResponse = z.infer<typeof LoginResponseSchema>;
