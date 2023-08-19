import {z} from 'zod';
import {ClientAllowUserSchema} from './client';
import {OrganizationUserSchema} from './organization';
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
  id: z.number(),
  brand_id: z.string(),
  brand_name: z.string(),

  created_at: z.string().datetime({offset: true}),
});
export type UserBrand = z.infer<typeof UserBrandSchema>;

export const StaffClientSchema = z.object({
  client_id: z.string(),
  name: z.string(),
  image: z.string().nullable(),
});
export type StaffClient = z.infer<typeof StaffClientSchema>;

export const StaffClientsSchema = z.array(StaffClientSchema);
export type StaffClients = z.infer<typeof StaffClientsSchema>;

export const UserDetailSchema = z.object({
  user: UserSchema,

  staff: StaffSchema.optional(),
  user_brands: z.array(UserBrandSchema),

  clients: z.array(StaffClientSchema),
});
export type UserDetail = z.infer<typeof UserDetailSchema>;

export const BrandSchema = z.object({
  id: z.string(),
  name: z.string(),
  description: z.string().nullable(),

  created_at: z.string().datetime({offset: true}),
  updated_at: z.string().datetime({offset: true}),
});
export type Brand = z.infer<typeof BrandSchema>;

export const BrandsSchema = z.array(BrandSchema);
export type Brands = z.infer<typeof BrandsSchema>;

export const OrganizationSchema = z.object({
  id: z.string(),
  name: z.string(),
  image: z.string().nullable(),
  link: z.string().nullable(),

  created_at: z.string().datetime({offset: true}),
  updated_at: z.string().datetime({offset: true}),
});
export type Organization = z.infer<typeof OrganizationSchema>;

export const OrganizationsSchema = z.array(OrganizationSchema);
export type Organizations = z.infer<typeof OrganizationsSchema>;

export const OrganizationDetailSchema = z.object({
  org: OrganizationSchema,
  users: z.array(OrganizationUserSchema),
  clients: z.array(StaffClientSchema),
});
export type OrganizationDetail = z.infer<typeof OrganizationDetailSchema>;

export const ClientDetailSchema = z.object({
  client: z.object({
    client_id: z.string(),
    name: z.string(),
    description: z.string().nullable(),
    image: z.string().nullable(),
    org_id: z.string().nullable(),
    org_member_only: z.boolean(),
    is_allow: z.boolean(),
    prompt: z.string().nullable(),
    owner_user_id: z.string(),
    client_secret: z.string(),

    created_at: z.string().datetime({offset: true}),
    updated_at: z.string().datetime({offset: true}),
  }),

  redirect_urls: z.array(z.string()),
  referrer_urls: z.array(z.string()),

  scopes: z.array(z.string()),

  allow_rules: z.array(ClientAllowUserSchema),
});
export type ClientDetail = z.infer<typeof ClientDetailSchema>;

export const RegisterSessionSchema = z.object({
  email: z.string(),
  email_verified: z.boolean(),
  send_count: z.number(),
  retry_count: z.number(),
  org_id: z.string().nullable(),

  period: z.string().datetime({offset: true}),

  created_at: z.string().datetime({offset: true}),
  updated_at: z.string().datetime({offset: true}),
});
export type RegisterSession = z.infer<typeof RegisterSessionSchema>;

export const RegisterSessionsSchema = z.array(RegisterSessionSchema);
export type RegisterSessions = z.infer<typeof RegisterSessionsSchema>;
