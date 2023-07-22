import {z} from 'zod';

export const ClientSchema = z.object({
  client_id: z.string(),

  name: z.string(),
  description: z.string().nullable(),
  image: z.string().url().nullable(),

  is_allow: z.boolean(),
  prompt: z.string().nullable(),

  org_member_only: z.boolean().optional(),

  created_at: z.string().datetime({offset: true}),
  updated_at: z.string().datetime({offset: true}),
});
export type Client = z.infer<typeof ClientSchema>;

export const ClientListSchema = z.array(ClientSchema);
export type ClientList = z.infer<typeof ClientListSchema>;

export const ClientDetailSchema = ClientSchema.extend({
  client_secret: z.string(),
  redirect_uris: z.array(z.string()),
  referrer_urls: z.array(z.string()),

  scopes: z.array(z.string()),

  org_member_only: z.boolean(),
});
export type ClientDetail = z.infer<typeof ClientDetailSchema>;

export const ClientDetailListSchema = z.array(ClientDetailSchema);
export type ClientDetailList = z.infer<typeof ClientDetailListSchema>;
