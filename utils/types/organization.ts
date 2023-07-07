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
