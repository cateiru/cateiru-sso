import z from 'zod';

export const UserSchema = z.object({
  id: z.string(),
  user_name: z.string(),
  email: z.string().email(),
  family_name: z.string().nullable(),
  middle_name: z.string().nullable(),
  given_name: z.string().nullable(),
  gender: z.string(),
  birthdate: z.string().datetime({offset: true}).nullable(),
  avatar: z.string().nullable(),
  locale_id: z.string(),

  created_at: z.string().datetime({offset: true}),
  updated_at: z.string().datetime({offset: true}),
});
export type User = z.infer<typeof UserSchema>;

export const UserSettingSchema = z.object({
  user_id: z.string(),
  notice_email: z.boolean(),
  notice_webpush: z.boolean(),

  created_at: z.string().datetime({offset: true}),
  updated_at: z.string().datetime({offset: true}),
});
export type UserSetting = z.infer<typeof UserSettingSchema>;

export const UserMeSchema = z.object({
  user: UserSchema,
  setting: UserSettingSchema.optional(),
  is_staff: z.boolean(),
  joined_organization: z.boolean(),
});
export type UserMe = z.infer<typeof UserMeSchema>;

export const UserAvatarSchema = z.object({
  avatar: z.string(),
});
export type UserAvatar = z.infer<typeof UserAvatarSchema>;

export const UserUserName = z.object({
  user_name: z.string(),
  ok: z.boolean(),
  message: z.string(),
});
export type UserUserName = z.infer<typeof UserUserName>;

export const PublicUserSchema = z.object({
  id: z.string(),
  user_name: z.string(),
  avatar: z.string().nullable(),
});
export type PublicUser = z.infer<typeof PublicUserSchema>;
