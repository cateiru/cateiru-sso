import {z} from 'zod';
import {PublicUserSchema} from './user';

export const OrganizationUserSchema = z.object({
  id: z.number(),
  user: PublicUserSchema,
  role: z.string(),

  created_at: z.string().datetime({offset: true}),
  updated_at: z.string().datetime({offset: true}),
});
export type OrganizationUser = z.infer<typeof OrganizationUserSchema>;

export const PublicOrganizationSchema = z.object({
  id: z.string(),
  name: z.string(),
  image: z.string().nullable(),
  link: z.string().nullable(),
  role: z.string(),
  join_date: z.string().datetime({offset: true}),
});
export type PublicOrganization = z.infer<typeof PublicOrganizationSchema>;

export const PublicOrganizationListSchema = z.array(PublicOrganizationSchema);
export type PublicOrganizationList = z.infer<
  typeof PublicOrganizationListSchema
>;

export const PublicOrganizationDetailSchema = PublicOrganizationSchema.extend({
  created_at: z.string().datetime({offset: true}),
});
export type PublicOrganizationDetail = z.infer<
  typeof PublicOrganizationDetailSchema
>;
