import {z} from 'zod';
import {UserSchema} from './user';

export const StaffUsersSchema = z.array(UserSchema);
export type StaffUsers = z.infer<typeof StaffUsersSchema>;

export const StaffSchema = z.object({
  user_id: z.string(),
  memo: z.string().nullable(),

  created_at: z.string().datetime({offset: true}),
  updated_at: z.string().datetime({offset: true}),
});
export type Staff = z.infer<typeof StaffSchema>;

export const UserBrandSchema = z.object({
  id: z.string(),
  user_id: z.string(),
  brand_id: z.string(),

  created_at: z.string().datetime({offset: true}),
});
export type Brand = z.infer<typeof UserBrandSchema>;

export const StaffClientSchema = z.object({
  client_id: z.string(),
  name: z.string(),
  image: z.string().nullable(),
});

export const UserDetailSchema = z.object({
  user: UserSchema,

  staff: StaffSchema.nullable(),
  user_brands: z.array(UserBrandSchema),

  clients: z.array(StaffClientSchema),
});
export type UserDetail = z.infer<typeof UserDetailSchema>;
