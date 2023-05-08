import {z} from 'zod';

export const AccountUserSchema = z.object({
  user_name: z.string(),
  id: z.string(),
  avatar: z.string().optional(),
});
export type AccountUser = z.infer<typeof AccountUserSchema>;

export const AccountUserListSchema = z.array(AccountUserSchema);
export type AccountUserList = z.infer<typeof AccountUserListSchema>;
