import z from 'zod';

export const UserSchema = z.object({
  id: z.string(),
  user_name: z.string(),
  email: z.string(),
  family_name: z.string().optional(),
  middle_name: z.string().optional(),
  given_name: z.string().optional(),
  gender: z.string(),
  birthdate: z.string().datetime().optional(),
  avatar: z.string().optional(),
  locale_id: z.string(),

  created: z.string().datetime(),
  modified: z.string().datetime(),
});
export type User = z.infer<typeof UserSchema>;

export const UserSettingSchema = z.object({
  user_id: z.string(),
  notice_email: z.boolean(),
  notice_webpush: z.boolean(),

  created: z.string().datetime(),
  modified: z.string().datetime(),
});
export type UserSetting = z.infer<typeof UserSettingSchema>;

export const UserMeSchema = z.object({
  user: UserSchema,
  setting: UserSettingSchema,
});
export type UserMe = z.infer<typeof UserMeSchema>;
